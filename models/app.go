package models

import "time"

type Status struct {
	Ok   bool      `json:"ok"`
	Time time.Time `json:"time"`
}

func NewStatus(ok bool) *Status {
	return &Status{
		Ok:   ok,
		Time: time.Now().UTC(),
	}
}
