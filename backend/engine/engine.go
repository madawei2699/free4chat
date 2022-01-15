package engine

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/MixinNetwork/mixin/logger"
)

const (
	engineStateLoopPeriod = 60 * time.Second
)

type State struct {
	UpdatedAt   time.Time `json:"updated_at"`
	ActivePeers int       `json:"active_peers"`
	ClosedPeers int       `json:"closed_peers"`
	ActiveRooms int       `json:"active_rooms"`
	ClosedRooms int       `json:"closed_rooms"`
}

type Engine struct {
	IP        string
	Interface string
	PortMin   uint16
	PortMax   uint16

	State State
	rooms *rmap
}

func BuildEngine(conf *Configuration) (*Engine, error) {
	ip, err := getIPFromInterface(conf.Engine.Interface, conf.Engine.Address)
	if err != nil {
		return nil, err
	}
	engine := &Engine{
		IP:        ip,
		Interface: conf.Engine.Interface,
		PortMin:   conf.Engine.PortMin,
		PortMax:   conf.Engine.PortMax,
		rooms:     rmapAllocate(),
	}
	logger.Printf("BuildEngine(IP: %s, Interface: %s, Ports: %d-%d)\n", engine.IP, engine.Interface, engine.PortMin, engine.PortMax)
	return engine, nil
}

func (engine *Engine) Loop() {
	for {
		engine.rooms.RLock()

		engine.State.UpdatedAt = time.Now()
		engine.State.ActiveRooms = 0
		engine.State.ClosedRooms = 0
		engine.State.ActivePeers = 0
		engine.State.ClosedPeers = 0

		for _, pm := range engine.rooms.m {
			pm.RLock()
			ap, cp := 0, 0
			for _, p := range pm.m {
				if p.cid == peerTrackClosedId {
					cp += 1
				} else {
					ap += 1
				}
			}
			engine.State.ActivePeers += ap
			engine.State.ClosedPeers += cp
			if ap > 0 {
				engine.State.ActiveRooms += 1
			} else {
				engine.State.ClosedRooms += 1
			}
			pm.RUnlock()
		}

		engine.rooms.RUnlock()
		time.Sleep(engineStateLoopPeriod)
	}
}

func getIPFromInterface(iname string, addr string) (string, error) {
	if addr != "" {
		return addr, nil
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, i := range ifaces {
		if i.Name != iname {
			continue
		}
		addrs, err := i.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				return v.IP.String(), nil
			case *net.IPAddr:
				return v.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no address for interface %s", iname)
}

type pmap struct {
	sync.RWMutex
	id string
	m  map[string]*Peer
}

func pmapAllocate(id string) *pmap {
	pm := new(pmap)
	pm.id = id
	pm.m = make(map[string]*Peer)
	return pm
}

type rmap struct {
	sync.RWMutex
	m map[string]*pmap
}

func rmapAllocate() *rmap {
	rm := new(rmap)
	rm.m = make(map[string]*pmap)
	return rm
}

func (engine *Engine) getRoom(rid string) *pmap {
	rm := engine.rooms
	rm.RLock()
	defer rm.RUnlock()

	return rm.m[rid]
}

func (engine *Engine) GetRoom(rid string) *pmap {
	pm := engine.getRoom(rid)
	if pm != nil {
		return pm
	}

	rm := engine.rooms
	rm.Lock()
	defer rm.Unlock()
	if rm.m[rid] == nil {
		rm.m[rid] = pmapAllocate(rid)
	}
	return rm.m[rid]
}

func (room *pmap) get(uid, cid string) (*Peer, error) {
	peer := room.m[uid]
	if peer == nil {
		return nil, buildError(ErrorPeerNotFound, fmt.Errorf("peer %s not found in %s", uid, room.id))
	}
	if peer.cid == peerTrackClosedId {
		return nil, buildError(ErrorPeerClosed, fmt.Errorf("peer %s closed in %s", uid, room.id))
	}
	if peer.cid != cid {
		return nil, buildError(ErrorTrackNotFound, fmt.Errorf("peer %s track not match %s %s in %s", uid, cid, peer.cid, room.id))
	}
	return peer, nil
}
