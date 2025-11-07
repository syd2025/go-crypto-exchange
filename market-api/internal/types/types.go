package types

type RateRequest struct {
	Unit string `json:"unit"`
	Ip   string `json:"ip"`
}

type RateResponse struct {
	Rate float64 `json:"rate"`
}
