package router

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yalagtyarzh/aggregator/pkg/http/middleware"
	"net/http"
	"net/http/pprof"
)

type IAPI interface {
	Router(r *mux.Router, c *middleware.Middleware)
}

func Router(mw *middleware.Middleware, apis ...IAPI) http.Handler {
	root := mux.NewRouter().StrictSlash(true).PathPrefix("/").Subrouter()

	//initPprof(root)

	root.Use(mw.RecoverMiddleware, mw.PrometheusMiddleware, mw.ReqIDMiddleware, mw.EnableCORS)

	for _, v := range apis {
		v.Router(root, mw)
	}

	return root
}

func Metrics() http.Handler {
	root := mux.NewRouter().StrictSlash(true)
	root.Handle("/metrics", promhttp.Handler())

	return root
}

func Probe() http.Handler {
	root := mux.NewRouter().StrictSlash(true)

	root.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	root.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return root
}

func initPprof(r *mux.Router) {
	debug := r.PathPrefix("/debug/pprof").Subrouter()
	debug.HandleFunc("/", pprof.Index)
	debug.HandleFunc("/cmdline", pprof.Cmdline)
	debug.HandleFunc("/symbol", pprof.Symbol)
	debug.HandleFunc("/trace", pprof.Trace)
	profile := debug.PathPrefix("/profile").Subrouter()
	profile.HandleFunc("", pprof.Profile)
	profile.Handle("/goroutine", pprof.Handler("goroutine"))
	profile.Handle("/threadcreate", pprof.Handler("threadcreate"))
	profile.Handle("/heap", pprof.Handler("heap"))
	profile.Handle("/block", pprof.Handler("block"))
	profile.Handle("/mutex", pprof.Handler("mutex"))
}
