package engine

import (
	"encoding/json"
	"net/http"
)

const (
	ErrorInvalidParams           = 5001000
	ErrorInvalidSDP              = 5001001
	ErrorInvalidCandidate        = 5001002
	ErrorRoomFull                = 5002000
	ErrorPeerNotFound            = 5002001
	ErrorPeerClosed              = 5002002
	ErrorTrackNotFound           = 5002003
	ErrorServerNewPeerConnection = 5003000
	ErrorServerCreateOffer       = 5003001
	ErrorServerSetLocalOffer     = 5003002
	ErrorServerNewTrack          = 5003003
	ErrorServerAddTransceiver    = 5003004
	ErrorServerSetRemoteOffer    = 5003005
	ErrorServerCreateAnswer      = 5003006
	ErrorServerSetLocalAnswer    = 5003007
	ErrorServerSetRemoteAnswer   = 5003008
	ErrorServerTimeout           = 5003999
)

type Error struct {
	Status      int    `json:"status"`
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func (e Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func buildError(code int, err error) error {
	status := http.StatusAccepted
	if code >= ErrorServerNewPeerConnection && code <= ErrorServerTimeout {
		status = http.StatusInternalServerError
	}
	return Error{
		Status:      status,
		Code:        code,
		Description: err.Error(),
	}
}
