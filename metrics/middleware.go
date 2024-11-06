package metrics

import (
	"net/http"

	"github.com/semaphoreui/semaphore/util"
)

func PrometheusMiddleware(next http.Handler) http.Handler {
	if util.Config.Metrics.Enabled {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("enable metrics in config"))
	})
}
