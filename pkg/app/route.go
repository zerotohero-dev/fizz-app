/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package app

import (
	"github.com/gorilla/mux"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"io"
	"net/http"
	"path"
)


func Handle404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	log.Info("404:url: '%s'", r.URL.String())

	_, _ = io.WriteString(w, ":(")
}

func RouteHealthEndpoints(r *mux.Router) {
	r.Methods("GET").Path("/readyz").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, `{"ready": true}`)
	})

	r.Methods("GET").Path("/healthz").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, `{"alive": true}`)
	})

	r.NotFoundHandler = http.HandlerFunc(Handle404)
}

func Route(router *mux.Router, handler http.Handler, method string, path string) {
	router.Methods(method).Path(path).Handler(handler)
}

func RoutePaths(handler http.Handler, router *mux.Router, method string, paths []string) {
	for _, p := range paths {
		Route(
			router, handler,
			method, p,
		)
	}
}

// RoutePrefixedPath can be used to normalize ALB Ingress Controller behavior.
func RoutePrefixedPath(
	handler http.Handler, router *mux.Router, method string,
	prefix string, pth string,
) {
	RoutePaths(handler, router, method, []string{
		pth,
		path.Join("/", prefix, pth),
	})
}