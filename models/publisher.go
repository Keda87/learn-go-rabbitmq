package models

type PublishEventPayload struct {
	Event   string `json:"event" validate:"required"`
	Payload struct {
		Name string `json:"name" validate:"required"`
	} `json:"payload" validate:"required"`
}
