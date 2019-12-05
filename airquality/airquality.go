package airquality

import (
	"encoding/json"
	"time"

	"github.com/guptarohit/asciigraph"
)

type AirQualityJson []struct {
	Zone         string  `json:"zone"`
	Municipality string  `json:"municipality"`
	Area         string  `json:"area"`
	Station      string  `json:"station"`
	Eoi          string  `json:"eoi"`
	Component    string  `json:"component"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Timestep     int     `json:"timestep"`
	Unit         string  `json:"unit"`
	Values       []struct {
		FromTime          time.Time `json:"fromTime"`
		ToTime            time.Time `json:"toTime"`
		Value             float64   `json:"value"`
		QualityControlled bool      `json:"qualityControlled"`
	} `json:"values"`
}

func MarshalIntoResponse(rawJson string) AirQualityJson {
	var response AirQualityJson
	json.Unmarshal([]byte(rawJson), &response)

	return response
}

func GetPlottedComponent(aq AirQualityJson, componentFilter string, graphHeight int) (string, float64) {

	for _, stationData := range aq {

		if componentFilter == stationData.Component {
			var data []float64

			for i := 0; i < len(stationData.Values); i++ {
				data = append(data, stationData.Values[i].Value)
			}

			graph := asciigraph.Plot(data, asciigraph.Height(graphHeight), asciigraph.Caption(stationData.Component+" in ("+stationData.Unit+")"))

			lastValue := data[len(data)-1] // C

			return graph, lastValue
		}
	}

	return "", 99.99
}
