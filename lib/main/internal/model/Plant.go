package model

type Plant struct {
	Id int `json:"id"`
	//Latin          string      `json:"latin"`
	Family string `json:"family"`
	Common string `json:"common"`
	//Category       Category    `json:"category"`
	//Origin         string      `json:"origin"`
	//Climate        string      `json:"climate"`
	//TempMax        Temperature `json:"tempmax"`
	//TempMin        Temperature `json:"tempmin"`
	//IdealLight     string      `json:"ideallight"`
	//ToleratedLight string      `json:"toleratedlight"`
	//Watering       string      `json:"watering"`
	//Insects        []string    `json:"insects"`
	//Diseases       string      `json:"diseases"`
	//Use            []string    `json:"use"`
}

type Plants struct {
	Plants []Plant
}

func PlantExists(plants Plants, name string) bool {
	for _, p := range plants.Plants {
		if p.Common == name {
			return true
		}
	}
	return false
}

func (p Plant) NewPlant(common string, family string, id int) Plant {
	return Plant{
		Common: common,
		Family: family,
		Id:     id,
	}
}
