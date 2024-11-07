package monitor

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/semaphoreui/semaphore/util"
)

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if util.Config.Metrics.Enabled {
			w.Header().Set("Content-Type", "application-json")
			next.ServeHTTP(w, r)
		} else {
			w.Write([]byte(`{"metrics":"disabled"}`))
		}
	})
}

var (
	// examples
	TasksTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "semaphore_tasks_total",
			Help: "total number of tasks created",
		},
	)

	TasksRunning = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "semaphroe_tasks_running",
			Help: "number of currently running tasks",
		},
	)
)
