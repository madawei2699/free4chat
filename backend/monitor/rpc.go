package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dimfeld/httptreemux"
	"github.com/gorilla/handlers"
	"github.com/unrolled/render"
)

type R struct {
	monitor *Monitor
}

type Call struct {
	Id     string        `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

type Render struct {
	w    http.ResponseWriter
	impl *render.Render
	id   string
}

func (r *Render) RenderData(data interface{}) {
	body := map[string]interface{}{"data": data}
	if r.id != "" {
		body["id"] = r.id
	}
	r.impl.JSON(r.w, http.StatusOK, body)
}

func (r *Render) RenderError(err error) {
	body := map[string]interface{}{"error": err.Error()}
	if r.id != "" {
		body["id"] = r.id
	}
	r.impl.JSON(r.w, http.StatusOK, body)
}

func (impl *R) handle(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	var call Call
	d := json.NewDecoder(r.Body)
	d.UseNumber()
	if err := d.Decode(&call); err != nil {
		render.New().JSON(w, http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}
	renderer := &Render{w: w, impl: render.New(), id: call.Id}
	switch call.Method {
	default:
		renderer.RenderError(fmt.Errorf("invalid method %s", call.Method))
	}
}

func registerHanders(router *httptreemux.TreeMux) {
	router.MethodNotAllowedHandler = func(w http.ResponseWriter, r *http.Request, _ map[string]httptreemux.HandlerFunc) {
		render.New().JSON(w, http.StatusNotFound, map[string]interface{}{"error": "not found"})
	}
	router.NotFoundHandler = func(w http.ResponseWriter, r *http.Request) {
		render.New().JSON(w, http.StatusNotFound, map[string]interface{}{"error": "not found"})
	}
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, rcv interface{}) {
		render.New().JSON(w, http.StatusInternalServerError, map[string]interface{}{"error": "server error"})
	}
}

func ServeRPC(monitor *Monitor, conf *Configuration) error {
	impl := &R{monitor: monitor}
	router := httptreemux.New()
	router.POST("/", impl.handle)
	registerHanders(router)
	handler := handlers.ProxyHeaders(router)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", conf.RPC.Port),
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	return server.ListenAndServe()
}
