package server

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/cffmnk/yametrics/internal/metrics"
	"github.com/cffmnk/yametrics/internal/storage"
)

type Server struct {
	storage storage.Storage
}

func NewServer() *Server {
	return &Server{
		storage: storage.NewMemStorage(),
	}
}

func (s *Server) HandleUpdateMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}
	path := strings.Split(r.URL.Path[8:], "/")
	if len(path) < 1 || path[0] == "" {
		http.Error(w, "No metric type", http.StatusBadRequest)
		return
	} else if len(path) < 2 || path[1] == "" {
		http.Error(w, "No metric name", http.StatusNotFound)
		return
	} else if len(path) < 3 || path[2] == "" {
		http.Error(w, "No metric value", http.StatusBadRequest)
		return
	}
	metricType := path[0]
	metricName := path[1]
	metricValue := path[2]
	switch metricType {
	case metrics.Gauge:
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			http.Error(w, "Invalid gauge value", http.StatusBadRequest)
			return
		}
		s.storage.UpdateGauge(metricName, value)
	case metrics.Counter:
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			http.Error(w, "Invalid counter value", http.StatusBadRequest)
			return
		}
		s.storage.UpdateCounter(metricName, value)
	default:
		http.Error(w, "Invalid metric type", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
