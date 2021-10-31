package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v2"
)

var (
	// only support unified plan
	cfg = webrtc.Configuration{
		SDPSemantics: webrtc.SDPSemanticsUnifiedPlanWithFallback,
	}

	setting webrtc.SettingEngine

	errChanClosed    = errors.New("channel closed")
	errInvalidTrack  = errors.New("track is nil")
	errInvalidPacket = errors.New("packet is nil")
	// errInvalidPC      = errors.New("pc is nil")
	// errInvalidOptions = errors.New("invalid options")
	errNotImplemented = errors.New("not implemented")
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 51200
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// User is a middleman between the websocket connection and the hub.
type User struct {
	ID            string
	room          *Room
	conn          *websocket.Conn          // The websocket connection.
	send          chan []byte              // Buffered channel of outbound messages.
	pc            *webrtc.PeerConnection   // WebRTC Peer Connection
	inTracks      map[uint32]*webrtc.Track // Microphone
	inTracksLock  sync.RWMutex
	outTracks     map[uint32]*webrtc.Track // Rest of the room's tracks
	outTracksLock sync.RWMutex

	rtpCh chan *rtp.Packet

	stop bool

	info UserInfo
}

// UserInfo contains some user data
type UserInfo struct {
	Emoji string `json:"emoji"` // emoji-face like on clients (for test)
	Mute  bool   `json:"mute"`
}

// UserWrap represents user object sent to client
type UserWrap struct {
	ID string `json:"id"`
	UserInfo
}

// Wrap wraps user
func (u *User) Wrap() *UserWrap {
	return &UserWrap{
		ID:       u.ID,
		UserInfo: u.info,
	}
}

// readPump pumps messages from the websocket connection to the hub.
func (u *User) readPump() {
	defer func() {
		u.stop = true
		u.pc.Close()
		u.room.Leave(u)
		u.conn.Close()
	}()
	u.conn.SetReadLimit(maxMessageSize)
	u.conn.SetReadDeadline(time.Now().Add(pongWait))
	u.conn.SetPongHandler(func(string) error { u.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := u.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
				log.Println(err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		go func() {
			err := u.HandleEvent(message)
			if err != nil {
				log.Println(err)
				u.SendErr(err)
			}
		}()
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (u *User) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		u.stop = true
		u.conn.Close()
	}()
	for {
		select {
		case message, ok := <-u.send:
			u.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				u.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := u.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			u.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := u.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// Event represents web socket user event
type Event struct {
	Type string `json:"type"`

	Offer     *webrtc.SessionDescription `json:"offer,omitempty"`
	Answer    *webrtc.SessionDescription `json:"answer,omitempty"`
	Candidate *webrtc.ICECandidateInit   `json:"candidate,omitempty"`
	User      *UserWrap                  `json:"user,omitempty"`
	Room      *RoomWrap                  `json:"room,omitempty"`
	Desc      string                     `json:"desc,omitempty"`
}

// SendEvent sends json body to web socket
func (u *User) SendEvent(event Event) error {
	json, err := json.Marshal(event)
	if err != nil {
		return err
	}
	u.send <- json
	return nil
}

// SendEventUser sends user to client to identify himself
func (u *User) SendEventUser() error {
	return u.SendEvent(Event{Type: "user", User: u.Wrap()})
}

// SendEventRoom sends room to client with users except me
func (u *User) SendEventRoom() error {
	return u.SendEvent(Event{Type: "room", Room: u.room.Wrap(u)})
}

// BroadcastEvent sends json body to everyone in the room except this user
func (u *User) BroadcastEvent(event Event) error {
	json, err := json.Marshal(event)
	if err != nil {
		return err
	}
	u.room.Broadcast(json, u)
	return nil
}

// BroadcastEventJoin sends user_join event
func (u *User) BroadcastEventJoin() error {
	return u.BroadcastEvent(Event{Type: "user_join", User: u.Wrap()})
}

// BroadcastEventLeave sends user_leave event
func (u *User) BroadcastEventLeave() error {
	return u.BroadcastEvent(Event{Type: "user_leave", User: u.Wrap()})
}

// BroadcastEventMute sends microphone mute event to everyone
func (u *User) BroadcastEventMute() error {
	return u.BroadcastEvent(Event{Type: "mute", User: u.Wrap()})
}

// BroadcastEventUnmute sends microphone unmute event to everyone
func (u *User) BroadcastEventUnmute() error {
	return u.BroadcastEvent(Event{Type: "unmute", User: u.Wrap()})
}

// SendErr sends error in json format to web socket
func (u *User) SendErr(err error) error {
	return u.SendEvent(Event{Type: "error", Desc: fmt.Sprint(err)})
}

func (u *User) log(msg ...interface{}) {
	log.Println(
		fmt.Sprintf("user %s:", u.ID),
		fmt.Sprint(msg...),
	)
}

// HandleEvent handles user event
func (u *User) HandleEvent(eventRaw []byte) error {
	var event *Event
	err := json.Unmarshal(eventRaw, &event)
	if err != nil {
		return err
	}
	u.log("handle event", event.Type)
	if event.Type == "offer" {
		if event.Offer == nil {
			return u.SendErr(errors.New("empty offer"))
		}
		err := u.HandleOffer(*event.Offer)
		if err != nil {
			return err
		}
		return nil
	} else if event.Type == "answer" {
		if event.Answer == nil {
			return u.SendErr(errors.New("empty answer"))
		}
		u.pc.SetRemoteDescription(*event.Answer)
		return nil
	} else if event.Type == "candidate" {
		if event.Candidate == nil {
			return u.SendErr(errors.New("empty candidate"))
		}
		u.log("adding candidate")
		u.pc.AddICECandidate(*event.Candidate)
		return nil
	} else if event.Type == "mute" {
		u.info.Mute = true
		u.BroadcastEventMute()
		return nil
	} else if event.Type == "unmute" {
		u.info.Mute = false
		u.BroadcastEventUnmute()
		return nil
	}

	return u.SendErr(errNotImplemented)
}

// GetRoomTracks returns list of room incoming tracks
func (u *User) GetRoomTracks() []*webrtc.Track {
	tracks := []*webrtc.Track{}
	for _, user := range u.room.GetUsers() {
		for _, track := range user.inTracks {
			tracks = append(tracks, track)
		}
	}
	return tracks
}

func (u *User) supportOpus(offer webrtc.SessionDescription) bool {
	mediaEngine := webrtc.MediaEngine{}
	mediaEngine.PopulateFromSDP(offer)
	var payloadType uint8
	// Search for Payload type. If the offer doesn't support codec exit since
	// since they won't be able to decode anything we send them
	for _, audioCodec := range mediaEngine.GetCodecsByKind(webrtc.RTPCodecTypeAudio) {
		if audioCodec.Name == "OPUS" {
			payloadType = audioCodec.PayloadType
			break
		}
	}
	if payloadType == 0 {
		return false
	}
	return true
}

// HandleOffer handles webrtc offer
func (u *User) HandleOffer(offer webrtc.SessionDescription) error {
	if ok := u.supportOpus(offer); !ok {
		return errors.New("remote peer does not support opus codec")
	}

	if len(u.pc.GetTransceivers()) == 0 {
		// add receive only transciever to get user microphone audio
		_, err := u.pc.AddTransceiver(webrtc.RTPCodecTypeAudio, webrtc.RtpTransceiverInit{
			Direction: webrtc.RTPTransceiverDirectionRecvonly,
		})
		if err != nil {
			return err
		}
	}

	// Set the remote SessionDescription
	if err := u.pc.SetRemoteDescription(offer); err != nil {
		return err
	}

	err := u.SendAnswer()
	if err != nil {
		return err
	}

	return nil
}

// Offer return a offer
func (u *User) Offer() (webrtc.SessionDescription, error) {
	offer, err := u.pc.CreateOffer(nil)
	if err != nil {
		return webrtc.SessionDescription{}, err
	}
	err = u.pc.SetLocalDescription(offer)
	if err != nil {
		return webrtc.SessionDescription{}, err
	}
	return offer, nil
}

// SendOffer creates webrtc offer
func (u *User) SendOffer() error {
	offer, err := u.Offer()
	err = u.SendEvent(Event{Type: "offer", Offer: &offer})
	if err != nil {
		panic(err)
	}
	return nil
}

// SendCandidate sends ice candidate to peer
func (u *User) SendCandidate(iceCandidate *webrtc.ICECandidate) error {
	if iceCandidate == nil {
		return errors.New("nil ice candidate")
	}
	iceCandidateInit := iceCandidate.ToJSON()
	err := u.SendEvent(Event{Type: "candidate", Candidate: &iceCandidateInit})
	if err != nil {
		return err
	}
	return nil
}

// Answer creates webrtc answer
func (u *User) Answer() (webrtc.SessionDescription, error) {
	answer, err := u.pc.CreateAnswer(nil)
	if err != nil {
		return webrtc.SessionDescription{}, err
	}
	// Sets the LocalDescription, and starts our UDP listeners
	if err = u.pc.SetLocalDescription(answer); err != nil {
		return webrtc.SessionDescription{}, err
	}
	return answer, nil
}

// SendAnswer creates answer and send it via websocket
func (u *User) SendAnswer() error {
	answer, err := u.Answer()
	if err != nil {
		return err
	}
	err = u.SendEvent(Event{Type: "answer", Answer: &answer})
	return nil
}

// receiveInTrackRTP receive all incoming tracks' rtp and sent to one channel
func (u *User) receiveInTrackRTP(remoteTrack *webrtc.Track) {
	for {
		if u.stop {
			return
		}
		rtp, err := remoteTrack.ReadRTP()
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatalf("rtp err => %v", err)
		}
		u.rtpCh <- rtp
	}
}

// ReadRTP read rtp packet
func (u *User) ReadRTP() (*rtp.Packet, error) {
	rtp, ok := <-u.rtpCh
	if !ok {
		return nil, errChanClosed
	}
	return rtp, nil
}

// WriteRTP send rtp packet to user outgoing tracks
func (u *User) WriteRTP(pkt *rtp.Packet) error {
	if pkt == nil {
		return errInvalidPacket
	}
	u.outTracksLock.RLock()
	track := u.outTracks[pkt.SSRC]
	u.outTracksLock.RUnlock()

	if track == nil {
		log.Printf("WebRTCTransport.WriteRTP track==nil pkt.SSRC=%d", pkt.SSRC)
		return errInvalidTrack
	}

	// log.Debugf("WebRTCTransport.WriteRTP pkt=%v", pkt)
	err := track.WriteRTP(pkt)
	if err != nil {
		// log.Errorf(err.Error())
		// u.writeErrCnt++
		return err
	}
	return nil
}

func (u *User) broadcastIncomingRTP() {
	for {
		rtp, err := u.ReadRTP()
		if err != nil {
			panic(err)
		}
		for _, user := range u.room.GetOtherUsers(u) {
			err := user.WriteRTP(rtp)
			if err != nil {
				// panic(err)
				fmt.Println(err)
			}
		}
	}
}

// GetInTracks return incoming tracks
func (u *User) GetInTracks() map[uint32]*webrtc.Track {
	u.inTracksLock.RLock()
	defer u.inTracksLock.RUnlock()
	return u.inTracks

}

// GetOutTracks return outgoing tracks
func (u *User) GetOutTracks() map[uint32]*webrtc.Track {
	u.outTracksLock.RLock()
	defer u.outTracksLock.RUnlock()
	return u.outTracks
}

// AddTrack adds track to peer connection
func (u *User) AddTrack(ssrc uint32) error {
	track, err := u.pc.NewTrack(webrtc.DefaultPayloadTypeOpus, ssrc, string(ssrc), string(ssrc))
	if err != nil {
		return err
	}
	if _, err := u.pc.AddTrack(track); err != nil {
		log.Println("ERROR Add remote track as peerConnection local track", err)
		return err
	}

	u.outTracksLock.Lock()
	u.outTracks[track.SSRC()] = track
	u.outTracksLock.Unlock()
	return nil
}

// Watch for debug
func (u *User) Watch() {
	ticker := time.NewTicker(time.Second * 5)
	for range ticker.C {
		if u.stop {
			ticker.Stop()
			return
		}
		fmt.Println("ID:", u.ID, "out: ", u.GetOutTracks())
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(rooms *Rooms, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	mediaEngine := webrtc.MediaEngine{}
	mediaEngine.RegisterCodec(webrtc.NewRTPOpusCodec(webrtc.DefaultPayloadTypeOpus, 48000))

	api := webrtc.NewAPI(webrtc.WithMediaEngine(mediaEngine))
	peerConnection, err := api.NewPeerConnection(peerConnectionConfig)

	roomID := strings.ReplaceAll(r.URL.Path, "/", "")
	room := rooms.GetOrCreate(roomID)

	log.Println("ws connection to room:", roomID, len(room.GetUsers()), "users")

	emojis := []string{
		"😎", "🧐", "🤡", "👻", "😷", "🤗", "😏",
		"👽", "👨‍🚀", "🐺", "🐯", "🦁", "🐶", "🐼", "🙈",
	}

	user := &User{
		ID:        strconv.FormatInt(time.Now().UnixNano(), 10), // generate random id based on timestamp
		room:      room,
		conn:      conn,
		send:      make(chan []byte, 256),
		pc:        peerConnection,
		inTracks:  make(map[uint32]*webrtc.Track),
		outTracks: make(map[uint32]*webrtc.Track),
		rtpCh:     make(chan *rtp.Packet, 100),

		info: UserInfo{
			Emoji: emojis[rand.Intn(len(emojis))],
			Mute:  true, // user is muted by default
		},
	}

	user.pc.OnICECandidate(func(iceCandidate *webrtc.ICECandidate) {
		if iceCandidate != nil {
			err := user.SendCandidate(iceCandidate)
			if err != nil {
				log.Println("fail send candidate", err)
			}
		}
	})

	user.pc.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		log.Printf("Connection State has changed %s \n", connectionState.String())
		if connectionState == webrtc.ICEConnectionStateConnected {
			log.Println("user joined")
			tracks := user.GetRoomTracks()
			fmt.Println("attach ", len(tracks), "tracks to new user")
			user.log("new user add tracks", len(tracks))
			for _, track := range tracks {
				err := user.AddTrack(track.SSRC())
				if err != nil {
					log.Println("ERROR Add remote track as peerConnection local track", err)
					panic(err)
				}
			}
			err = user.SendOffer()
			if err != nil {
				panic(err)
			}
		} else if connectionState == webrtc.ICEConnectionStateDisconnected ||
			connectionState == webrtc.ICEConnectionStateFailed ||
			connectionState == webrtc.ICEConnectionStateClosed {

			user.stop = true
			senders := user.pc.GetSenders()
			for _, roomUser := range user.room.GetOtherUsers(user) {
				user.log("removing tracks from user")
				for _, sender := range senders {
					ssrc := sender.Track().SSRC()

					roomUserSenders := roomUser.pc.GetSenders()
					for _, roomUserSender := range roomUserSenders {
						if roomUserSender.Track().SSRC() == ssrc {
							err := roomUser.pc.RemoveTrack(roomUserSender)
							if err != nil {
								panic(err)
							}
						}
					}
				}
			}

		}
	})

	user.pc.OnTrack(func(remoteTrack *webrtc.Track, receiver *webrtc.RTPReceiver) {
		user.log(
			"peerConnection.OnTrack",
			fmt.Sprintf("track has started, of type %d: %s, ssrc: %d \n", remoteTrack.PayloadType(), remoteTrack.Codec().Name, remoteTrack.SSRC()),
		)
		if _, alreadyAdded := user.inTracks[remoteTrack.SSRC()]; alreadyAdded {
			user.log("user.inTrack != nil", "already handled")
			return
		}

		user.inTracks[remoteTrack.SSRC()] = remoteTrack
		for _, roomUser := range user.room.GetOtherUsers(user) {
			log.Println("add remote track", fmt.Sprintf("(user: %s)", user.ID), "track to user ", roomUser.ID)
			if err := roomUser.AddTrack(remoteTrack.SSRC()); err != nil {
				log.Println(err)
				continue
			}
			err := roomUser.SendOffer()
			if err != nil {
				panic(err)
			}
		}
		go user.receiveInTrackRTP(remoteTrack)
		go user.broadcastIncomingRTP()
	})

	user.room.Join(user)

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go user.writePump()
	go user.readPump()
	go user.Watch()

	user.SendEventUser()
	user.SendEventRoom()
}
