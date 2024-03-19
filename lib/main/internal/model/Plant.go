package model

type Plant struct {
	ID             int         `json:"id"`
	Latin          string      `json:"latin"`
	Family         string      `json:"family"`
	Common         []string    `json:"common"`
	Category       Category    `json:"category"`
	Origin         string      `json:"origin"`
	Climate        string      `json:"climate"`
	TempMax        Temperature `json:"tempmax"`
	TempMin        Temperature `json:"tempmin"`
	IdealLight     string      `json:"ideallight"`
	ToleratedLight string      `json:"toleratedlight"`
	Watering       string      `json:"watering"`
	Insects        []string    `json:"insects"`
	Diseases       string      `json:"diseases"`
	Use            []string    `json:"use"`
}
