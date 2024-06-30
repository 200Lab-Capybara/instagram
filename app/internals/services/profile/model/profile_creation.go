package model

type ProfileCreation struct {
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Avatar      string `json:"avatar"`
}

func (p *ProfileCreation) Validate() error {
	return nil
}
