package entity

import "time"

type Tenor struct {
	TenorId       string    `json:"tenor_id"`
	UserId        string    `json:"user_id"`
	Tenor         int       `json:"tenor"`
	TenorMaxLimit int       `json:"tenor_max_limit"`
	TenorInterest int       `json:"tenor_interest"`
	TenorCreateAt time.Time `json:"-"`
	TenorUpdateAt time.Time `json:"-"`
}

type TenorResponse struct {
	TenorId       string    `json:"tenor_id"`
	UserId        string    `json:"user_id"`
	Tenor         int       `json:"tenor"`
	TenorMaxLimit int       `json:"tenor_max_limit"`
	TenorInterest int       `json:"tenor_interest"`
	TenorCreateAt time.Time `json:"-"`
	TenorUpdateAt time.Time `json:"-"`
}
