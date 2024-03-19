package model

type HTTPResponse struct {
	Ok     bool `json:"ok"`
	Status int  `json:"status"`
}
