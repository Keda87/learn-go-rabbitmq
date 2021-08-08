package publisher

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/Keda87/learn-go-rabbitmq/models"
)

type publisherController struct {
	validator *validator.Validate
}

func NewController(validator *validator.Validate) *publisherController {
	return &publisherController{
		validator: validator,
	}
}

func (c *publisherController) HandlerPublishMessage(w http.ResponseWriter, r *http.Request) {
	var data models.PublishEventPayload

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.validator.Struct(data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
