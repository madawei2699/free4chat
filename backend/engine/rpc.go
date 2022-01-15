package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/MixinNetwork/mixin/logger"
	"github.com/dimfeld/httptreemux"
	"github.com/gorilla/handlers"
	"github.com/pion/webrtc/v3"
	"github.com/unrolled/render"
)

type R struct {
	router *Router
	conf   *Configuration
}

type Call struct {
	Id     string        `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

type Render struct {
	w       http.ResponseWriter
	impl    *render.Render
	id      string
	startAt time.Time
}

func NewRender(w http.ResponseWriter, id string) *Render {
	r := &Render{
		w:       w,
		id:      id,
		impl:    render.New(),
		startAt: time.Now(),
	}
	return r
}

func (r *Render) RenderData(data interface{}) {
	body := map[string]interface{}{"data": data}
	if r.id != "" {
		body["id"] = r.id
	}
	r.impl.JSON(r.w, http.StatusOK, body)
	logger.Printf("RPC.handle(id: %s, time: %fs) OK\n", r.id, time.Now().Sub(r.startAt).Seconds())
}

func (r *Render) RenderError(err error) {
	body := map[string]interface{}{"error": err}
	if r.id != "" {
		body["id"] = r.id
	}
	r.impl.JSON(r.w, http.StatusOK, body)
	logger.Printf("RPC.handle(id: %s, time: %fs) ERROR %s\n", r.id, time.Now().Sub(r.startAt).Seconds(), err.Error())
}

func (impl *R) handle(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	var call Call
	d := json.NewDecoder(r.Body)
	d.UseNumber()
	if err := d.Decode(&call); err != nil {
		render.New().JSON(w, http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}
	renderer := NewRender(w, call.Id)
	logger.Printf("RPC.handle(id: %s, method: %s, params: %v)\n", call.Id, call.Method, call.Params)
	switch call.Method {
	case "turn":
		servers, err := impl.turn(call.Params)
		if err != nil {
			renderer.RenderError(err)
		} else {
			renderer.RenderData(servers)
		}
	case "info":
		info, err := impl.info(call.Params)
		if err != nil {
			renderer.RenderError(err)
		} else {
			renderer.RenderData(info)
		}
	case "list":
		peers, err := impl.list(call.Params)
		if err != nil {
			renderer.RenderError(err)
		} else {
			renderer.RenderData(map[string]interface{}{"peers": peers})
		}
	case "publish":
		cid, answer, err := impl.publish(call.Params)
		if err != nil {
			renderer.RenderError(err)
		} else {
			jsep, _ := json.Marshal(answer)
			renderer.RenderData(map[string]interface{}{"track": cid, "sdp": answer, "jsep": string(jsep)})
		}
	case "restart":
		answer, err := impl.restart(call.Params)
		if err != nil {
			renderer.RenderError(err)
		} else {
			jsep, _ := json.Marshal(answer)
			renderer.RenderData(map[string]interface{}{"jsep": string(jsep)})
		}
	case "end":
		err := impl.end(call.Params)
		if err != nil {
			renderer.RenderError(err)
		} else {
			renderer.RenderData(map[string]string{})
		}
	case "trickle":
		err := impl.trickle(call.Params)
		if err != nil {
			renderer.RenderError(err)
		} else {
			renderer.RenderData(map[string]string{})
		}
	case "subscribe":
		offer, err := impl.subscribe(call.Params)
		if err != nil {
			renderer.RenderError(err)
		} else {
			jsep, _ := json.Marshal(offer)
			renderer.RenderData(map[string]interface{}{"type": offer.Type, "sdp": offer.SDP, "jsep": string(jsep)})
		}
	case "answer":
		err := impl.answer(call.Params)
		if err != nil {
			renderer.RenderError(err)
		} else {
			renderer.RenderData(map[string]string{})
		}
	default:
		renderer.RenderError(fmt.Errorf("invalid method %s", call.Method))
	}
}

func (r *R) turn(params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid params count %d", len(params)))
	}
	uid, ok := params[0].(string)
	if !ok {
		return nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid uid type %s", params[0]))
	}
	return turn(r.conf, uid)
}

func (r *R) info(params []interface{}) (interface{}, error) {
	if len(params) != 0 {
		return nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid params count %d", len(params)))
	}
	return r.router.info()
}

func (r *R) list(params []interface{}) ([]map[string]interface{}, error) {
	if len(params) != 1 {
		return nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid params count %d", len(params)))
	}
	rid, ok := params[0].(string)
	if !ok {
		return nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid rid type %s", params[0]))
	}
	return r.router.list(rid)
}

func (r *R) publish(params []interface{}) (string, *webrtc.SessionDescription, error) {
	if len(params) < 3 {
		return "", nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid params count %d", len(params)))
	}
	rid, ok := params[0].(string)
	if !ok {
		return "", nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid rid type %v", params[0]))
	}
	uid, ok := params[1].(string)
	if !ok {
		return "", nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid uid type %v", params[1]))
	}
	sdp, ok := params[2].(string)
	if !ok {
		return "", nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid sdp type %v", params[2]))
	}
	var limit int
	var callback string
	if len(params) == 5 {
		i, err := strconv.ParseInt(fmt.Sprint(params[3]), 10, 64)
		if err != nil {
			return "", nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid limit type %v %v", params[3], err))
		}
		limit = int(i)
		cbk, ok := params[4].(string)
		if !ok {
			return "", nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid callback type %v", params[4]))
		}
		if !strings.HasPrefix(cbk, "https://") {
			return "", nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid callback value %s", cbk))
		}
		callback = cbk
	}
	return r.router.publish(rid, uid, sdp, limit, callback)
}

func (r *R) restart(params []interface{}) (*webrtc.SessionDescription, error) {
	if len(params) != 4 {
		return nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid params count %d", len(params)))
	}
	ids, err := r.parseId(params)
	if err != nil {
		return nil, buildError(ErrorInvalidParams, err)
	}
	jsep, ok := params[3].(string)
	if !ok {
		return nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid jsep type %s", params[3]))
	}
	return r.router.restart(ids[0], ids[1], ids[2], jsep)
}

func (r *R) end(params []interface{}) error {
	if len(params) != 3 {
		return buildError(ErrorInvalidParams, fmt.Errorf("invalid params count %d", len(params)))
	}
	ids, err := r.parseId(params)
	if err != nil {
		return buildError(ErrorInvalidParams, err)
	}
	return r.router.end(ids[0], ids[1], ids[2])
}

func (r *R) trickle(params []interface{}) error {
	if len(params) != 4 {
		return buildError(ErrorInvalidParams, fmt.Errorf("invalid params count %d", len(params)))
	}
	ids, err := r.parseId(params)
	if err != nil {
		return buildError(ErrorInvalidParams, err)
	}
	candi, ok := params[3].(string)
	if !ok {
		return buildError(ErrorInvalidParams, fmt.Errorf("invalid candi type %s", params[3]))
	}
	return r.router.trickle(ids[0], ids[1], ids[2], candi)
}

func (r *R) subscribe(params []interface{}) (*webrtc.SessionDescription, error) {
	if len(params) != 3 {
		return nil, buildError(ErrorInvalidParams, fmt.Errorf("invalid params count %d", len(params)))
	}
	ids, err := r.parseId(params)
	if err != nil {
		return nil, buildError(ErrorInvalidParams, err)
	}
	return r.router.subscribe(ids[0], ids[1], ids[2])
}

func (r *R) answer(params []interface{}) error {
	if len(params) != 4 {
		return buildError(ErrorInvalidParams, fmt.Errorf("invalid params count %d", len(params)))
	}
	ids, err := r.parseId(params)
	if err != nil {
		return buildError(ErrorInvalidParams, err)
	}
	sdp, ok := params[3].(string)
	if !ok {
		return buildError(ErrorInvalidParams, fmt.Errorf("invalid sdp type %s", params[3]))
	}
	return r.router.answer(ids[0], ids[1], ids[2], sdp)
}

func (r *R) parseId(params []interface{}) ([]string, error) {
	rid, ok := params[0].(string)
	if !ok {
		return nil, fmt.Errorf("invalid rid type %s", params[0])
	}
	uid, ok := params[1].(string)
	if !ok {
		return nil, fmt.Errorf("invalid uid type %s", params[1])
	}
	cid, ok := params[2].(string)
	if !ok {
		return nil, fmt.Errorf("invalid cid type %s", params[2])
	}
	return []string{rid, uid, cid}, nil
}

func registerHandlers(router *httptreemux.TreeMux) {
	router.MethodNotAllowedHandler = func(w http.ResponseWriter, r *http.Request, _ map[string]httptreemux.HandlerFunc) {
		render.New().JSON(w, http.StatusNotFound, map[string]interface{}{"error": "not found"})
	}
	router.NotFoundHandler = func(w http.ResponseWriter, r *http.Request) {
		render.New().JSON(w, http.StatusNotFound, map[string]interface{}{"error": "not found"})
	}
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, rcv interface{}) {
		logger.Println(rcv)
		render.New().JSON(w, http.StatusInternalServerError, map[string]interface{}{"error": "server error"})
	}
}

func handleCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			handler.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Authorization,Mixin-Conversation-ID")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET,POST,DELETE")
		w.Header().Set("Access-Control-Max-Age", "600")
		if r.Method == "OPTIONS" {
			render.New().JSON(w, http.StatusOK, map[string]interface{}{})
		} else {
			handler.ServeHTTP(w, r)
		}
	})
}

func ServeRPC(engine *Engine, conf *Configuration) error {
	logger.Printf("ServeRPC(:%d)\n", conf.RPC.Port)
	impl := &R{router: NewRouter(engine), conf: conf}
	router := httptreemux.New()
	router.POST("/", impl.handle)
	registerHandlers(router)
	handler := handleCORS(router)
	handler = handlers.ProxyHeaders(handler)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", conf.RPC.Port),
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	return server.ListenAndServe()
}
