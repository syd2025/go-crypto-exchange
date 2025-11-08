package types

type RateRequest struct {
	Unit string `path:"unit" json:"unit"`
	Ip   string `json:"ip,omitempty"`
}

type RateResponse struct {
	Rate float64 `json:"rate"`
}
