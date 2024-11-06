package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// example metrics
	// Task metrics
	TasksTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "semaphore_tasks_total",
			Help: "Total number of tasks created",
		},
	)

	TasksRunning = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "semaphore_tasks_running",
			Help: "Number of currently running tasks",
		},
	)
)
