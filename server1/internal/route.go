package api

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Router struct {
	mux     *http.ServeMux
	handler Handler
}

func NewRouter(handler Handler) *Router {
	router := &Router{
		mux:     http.NewServeMux(),
		handler: handler,
	}
	router.initRouter()
	return router
}

func (r *Router) initRouter() {
	r.mux.HandleFunc("/server1/test", r.handler.HandleTest)
	r.mux.HandleFunc("/", r.handler.HandleBase)
	r.mux.HandleFunc("/server1", r.handler.HandleBaseServer)
	r.mux.HandleFunc("/test", r.handler.HandleBaseTest)
	r.mux.Handle("/metrics", promhttp.Handler())
}

func (r *Router) GetMux() *http.ServeMux {
	return r.mux
}
