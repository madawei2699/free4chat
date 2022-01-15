package engine

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/MixinNetwork/mixin/logger"
	"github.com/gofrs/uuid"
	"github.com/pion/interceptor"
	"github.com/pion/sdp/v2"
	"github.com/pion/webrtc/v3"
)

type Router struct {
	engine *Engine
}

func NewRouter(engine *Engine) *Router {
	return &Router{engine: engine}
}

func (r *Router) info() (interface{}, error) {
	rm := r.engine.rooms
	rm.RLock()
	defer rm.RUnlock()

	return r.engine.State, nil
}

func (r *Router) list(rid string) ([]map[string]interface{}, error) {
	room := r.engine.GetRoom(rid)
	room.RLock()
	defer room.RUnlock()
	peers := make([]map[string]interface{}, 0)
	for _, p := range room.m {
		cid := uuid.FromStringOrNil(p.cid)
		if cid.String() == uuid.Nil.String() {
			continue
		}
		peers = append(peers, map[string]interface{}{
			"id":    p.uid,
			"track": cid.String(),
		})
	}
	return peers, nil
}

func (r *Router) create(rid, uid, callback string, offer webrtc.SessionDescription) (*Peer, error) {
	se := webrtc.SettingEngine{}
	se.SetLite(true)
	se.SetInterfaceFilter(func(in string) bool { return in == r.engine.Interface })
	se.SetNAT1To1IPs([]string{r.engine.IP}, webrtc.ICECandidateTypeHost)
	se.SetICETimeouts(10*time.Second, 30*time.Second, 2*time.Second)
	se.SetEphemeralUDPPortRange(r.engine.PortMin, r.engine.PortMax)

	me := &webrtc.MediaEngine{}
	opusChrome := webrtc.RTPCodecParameters{
		RTPCodecCapability: webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeOpus, ClockRate: 48000, Channels: 2, SDPFmtpLine: "minptime=10;useinbandfec=1", RTCPFeedback: nil},
		PayloadType:        111,
	}
	opusFirefox := webrtc.RTPCodecParameters{
		RTPCodecCapability: webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeOpus, ClockRate: 48000, Channels: 2, SDPFmtpLine: "minptime=10;useinbandfec=1", RTCPFeedback: nil},
		PayloadType:        109,
	}
	me.RegisterCodec(opusChrome, webrtc.RTPCodecTypeAudio)
	me.RegisterCodec(opusFirefox, webrtc.RTPCodecTypeAudio)

	ir := &interceptor.Registry{}
	err := webrtc.RegisterDefaultInterceptors(me, ir)
	if err != nil {
		panic(err)
	}

	api := webrtc.NewAPI(webrtc.WithMediaEngine(me), webrtc.WithSettingEngine(se), webrtc.WithInterceptorRegistry(ir))

	pcConfig := webrtc.Configuration{
		BundlePolicy:  webrtc.BundlePolicyMaxBundle,
		RTCPMuxPolicy: webrtc.RTCPMuxPolicyRequire,
	}
	pc, err := api.NewPeerConnection(pcConfig)
	if err != nil {
		return nil, buildError(ErrorServerNewPeerConnection, err)
	}

	err = pc.SetRemoteDescription(offer)
	if err != nil {
		pc.Close()
		return nil, buildError(ErrorServerSetRemoteOffer, err)
	}
	answer, err := pc.CreateAnswer(nil)
	if err != nil {
		pc.Close()
		return nil, buildError(ErrorServerCreateAnswer, err)
	}
	gatherComplete := webrtc.GatheringCompletePromise(pc)
	err = pc.SetLocalDescription(answer)
	if err != nil {
		pc.Close()
		return nil, buildError(ErrorServerSetLocalAnswer, err)
	}
	<-gatherComplete

	peer := BuildPeer(rid, uid, pc, callback)
	return peer, nil
}

func (r *Router) publish(rid, uid string, jsep string, limit int, callback string) (string, *webrtc.SessionDescription, error) {
	if err := validateId(rid); err != nil {
		return "", nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid rid format %s %s", rid, err.Error()))
	}
	if err := validateId(uid); err != nil {
		return "", nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid uid format %s %s", uid, err.Error()))
	}
	var offer webrtc.SessionDescription
	err := json.Unmarshal([]byte(jsep), &offer)
	if err != nil {
		return "", nil, buildError(ErrorInvalidSDP, err)
	}
	if offer.Type != webrtc.SDPTypeOffer {
		return "", nil, buildError(ErrorInvalidSDP, fmt.Errorf("invalid sdp type %s", offer.Type))
	}

	parser := sdp.SessionDescription{}
	err = parser.Unmarshal([]byte(offer.SDP))
	if err != nil {
		return "", nil, buildError(ErrorInvalidSDP, err)
	}

	room := r.engine.GetRoom(rid)
	room.Lock()
	defer room.Unlock()

	if limit > 0 {
		for i, p := range room.m {
			cid := uuid.FromStringOrNil(p.cid)
			if cid.String() == uuid.Nil.String() || uid == i {
				continue
			}
			limit--
		}
		if limit <= 0 {
			return "", nil, buildError(ErrorRoomFull, fmt.Errorf("room full %d", limit))
		}
	}

	timer := time.NewTimer(peerTrackConnectionTimeout)
	defer timer.Stop()

	pc := make(chan *Peer)
	ec := make(chan error)
	go func() {
		peer, err := r.create(rid, uid, callback, offer)
		if err != nil {
			ec <- err
		} else {
			pc <- peer
		}
	}()
	select {
	case err := <-ec:
		return "", nil, err
	case peer := <-pc:
		old := room.m[peer.uid]
		if old != nil {
			old.Close()
		}
		room.m[peer.uid] = peer
		return peer.cid, peer.pc.LocalDescription(), nil
	case <-timer.C:
		err := fmt.Errorf("publish(%s,%s) timeout", rid, uid)
		return "", nil, buildError(ErrorServerTimeout, err)
	}
}

func (r *Router) restart(rid, uid, cid string, jsep string) (*webrtc.SessionDescription, error) {
	room := r.engine.GetRoom(rid)
	room.Lock()
	peer, err := room.get(uid, cid)
	room.Unlock()

	if err != nil {
		return nil, err
	}
	peer.Lock()
	defer peer.Unlock()

	var offer webrtc.SessionDescription
	err = json.Unmarshal([]byte(jsep), &offer)
	if err != nil {
		return nil, buildError(ErrorInvalidSDP, err)
	}
	if offer.Type != webrtc.SDPTypeOffer {
		return nil, buildError(ErrorInvalidSDP, fmt.Errorf("invalid sdp type %s", offer.Type))
	}

	parser := sdp.SessionDescription{}
	err = parser.Unmarshal([]byte(offer.SDP))
	if err != nil {
		return nil, buildError(ErrorInvalidSDP, err)
	}

	err = peer.pc.SetRemoteDescription(offer)
	if err != nil {
		peer.pc.Close()
		return nil, buildError(ErrorServerSetRemoteOffer, err)
	}
	answer, err := peer.pc.CreateAnswer(nil)
	if err != nil {
		peer.pc.Close()
		return nil, buildError(ErrorServerCreateAnswer, err)
	}
	gatherComplete := webrtc.GatheringCompletePromise(peer.pc)
	err = peer.pc.SetLocalDescription(answer)
	if err != nil {
		peer.pc.Close()
		return nil, buildError(ErrorServerSetLocalAnswer, err)
	}
	<-gatherComplete
	return peer.pc.LocalDescription(), nil
}

func (r *Router) end(rid, uid, cid string) error {
	room := r.engine.GetRoom(rid)
	room.Lock()
	peer, err := room.get(uid, cid)
	room.Unlock()

	if err != nil {
		return err
	}
	return peer.Close()
}

func (r *Router) trickle(rid, uid, cid string, candi string) error {
	var ici webrtc.ICECandidateInit
	err := json.Unmarshal([]byte(candi), &ici)
	if err != nil {
		return buildError(ErrorInvalidCandidate, err)
	}
	if ici.Candidate == "" {
		return nil
	}

	room := r.engine.GetRoom(rid)
	room.Lock()
	peer, err := room.get(uid, cid)
	room.Unlock()
	if err != nil {
		return err
	}
	peer.Lock()
	defer peer.Unlock()

	return peer.pc.AddICECandidate(ici)
}

func (r *Router) subscribe(rid, uid, cid string) (*webrtc.SessionDescription, error) {
	room := r.engine.GetRoom(rid)
	room.Lock()
	defer room.Unlock()

	peer, err := room.get(uid, cid)
	if err != nil {
		return nil, err
	}

	timer := time.NewTimer(peerTrackConnectionTimeout)
	defer timer.Stop()

	ec := make(chan error)
	gc := make(chan struct{})
	go func() {
		peer.Lock()
		defer peer.Unlock()

		var renegotiate bool
		for _, p := range room.m {
			if p.uid == peer.uid {
				continue
			}
			p.Lock()
			old := peer.publishers[p.uid]

			if old != nil && (p.track == nil || old.id != p.cid) {
				err := peer.pc.RemoveTrack(old.rtp)
				if err != nil {
					logger.Printf("failed to remove sender %s from peer %s with error %s\n", p.id(), peer.id(), err.Error())
				} else {
					delete(peer.publishers, p.uid)
					delete(p.subscribers, peer.uid)
					renegotiate = true
				}
			}
			if p.track != nil && (old == nil || old.id != p.cid) {
				sender, err := peer.pc.AddTrack(p.track)
				if err != nil {
					logger.Printf("failed to add sender %s to peer %s with error %s\n", p.id(), peer.id(), err.Error())
				} else if id := sender.Track().ID(); id != p.cid {
					panic(fmt.Errorf("malformed peer and track id %s %s", p.cid, id))
				} else {
					peer.publishers[p.uid] = &Sender{id: p.cid, rtp: sender}
					p.subscribers[peer.uid] = &Sender{id: peer.cid, rtp: sender}
					renegotiate = true
				}
			}

			p.Unlock()
		}
		if !renegotiate {
			ec <- nil
			return
		}

		offer, err := peer.pc.CreateOffer(nil)
		if err != nil {
			ec <- buildError(ErrorServerCreateOffer, err)
			return
		}
		gatherComplete := webrtc.GatheringCompletePromise(peer.pc)
		err = peer.pc.SetLocalDescription(offer)
		if err != nil {
			ec <- buildError(ErrorServerSetLocalOffer, err)
			return
		}
		c := <-gatherComplete
		gc <- c
	}()

	select {
	case err := <-ec:
		return &webrtc.SessionDescription{}, err
	case <-gc:
		return peer.pc.LocalDescription(), nil
	case <-timer.C:
		err := fmt.Errorf("subscribe(%s,%s,%s) timeout", rid, uid, cid)
		return nil, buildError(ErrorServerTimeout, err)
	}
}

func (r *Router) answer(rid, uid, cid string, jsep string) error {
	var answer webrtc.SessionDescription
	err := json.Unmarshal([]byte(jsep), &answer)
	if err != nil {
		return buildError(ErrorInvalidSDP, err)
	}
	if answer.Type != webrtc.SDPTypeAnswer {
		return buildError(ErrorInvalidSDP, fmt.Errorf("invalid sdp type %s", answer.Type))
	}

	parser := sdp.SessionDescription{}
	err = parser.Unmarshal([]byte(answer.SDP))
	if err != nil {
		return buildError(ErrorInvalidSDP, err)
	}

	room := r.engine.GetRoom(rid)
	room.Lock()
	peer, err := room.get(uid, cid)
	room.Unlock()
	if err != nil {
		return err
	}
	peer.Lock()
	defer peer.Unlock()

	timer := time.NewTimer(peerTrackConnectionTimeout)
	defer timer.Stop()

	renegotiated := make(chan error)
	go func() {
		err := peer.pc.SetRemoteDescription(answer)
		logger.Printf("answer(%s,%s,%s) SetRemoteDescription with %v\n", rid, uid, cid, err)
		renegotiated <- err
	}()
	select {
	case err := <-renegotiated:
		if err != nil {
			return buildError(ErrorServerSetRemoteAnswer, err)
		}
	case <-timer.C:
		err := fmt.Errorf("answer(%s,%s,%s) timeout", rid, uid, cid)
		return buildError(ErrorServerTimeout, err)
	}

	return nil
}

func validateId(id string) error {
	if len(id) > 256 {
		return fmt.Errorf("id %s too long, the maximum is %d", id, 256)
	}
	uid, err := url.QueryUnescape(id)
	if err != nil {
		return err
	}
	if eid := url.QueryEscape(uid); eid != id {
		return fmt.Errorf("unmatch %s %s", id, eid)
	}
	return nil
}
