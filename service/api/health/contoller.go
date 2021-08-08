package health

import (
	"encoding/json"
	"github.com/Keda87/learn-go-rabbitmq/models"
	"net/http"
	"time"
)

type healthController struct{}

func NewController() *healthController {
	return &healthController{}
}

func (c *healthController) HandlerHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	healthData := &models.Health{}
	healthData.Status = "OK"
	healthData.Time = time.Now()
	healthData.Uptime = time.Since(healthData.Time)

	_ = json.NewEncoder(w).Encode(healthData)
}
