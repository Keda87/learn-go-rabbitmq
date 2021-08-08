package models

import "time"

type Health struct {
	Status string        `json:"status"`
	Time   time.Time     `json:"time"`
	Uptime time.Duration `json:"uptime"`
}
