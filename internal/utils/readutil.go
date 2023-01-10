package utils

import (
	"encoding/json"
	"go-gin-demo/internal/models"
	"io/ioutil"
)

func ReadMockPositionData() models.Positions {
	file, _ := ioutil.ReadFile("internal/mockdata/mockpositions.json")
	data := models.Positions{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

func ReadMockPriceData() models.Prices {
	file, _ := ioutil.ReadFile("internal/mockdata/mockprices.json")
	data := models.Prices{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}
