package model

import "time"

type ProfileCreation struct {
	DateOfBirth time.Time `json:"date_of_birth"`
	Gender      string    `json:"gender"`
	Avatar      string    `json:"avatar"`
}

func (p *ProfileCreation) Validate() error {
	return nil
}
