package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/kaiaverkvist/luftmon/airquality"
)

func main() {
	currentTimeString := time.Now().Format("2006-01-02")

	previousDayString := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	log.Println(currentTimeString, previousDayString)

	response, err := http.Get("https://api.nilu.no/obs/historical/" + previousDayString + "/" + currentTimeString + "/Sofienbergparken")

	if err != nil {
		log.Println(err.Error())
	}

	defer response.Body.Close()

	bodyStr, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println(err.Error())
	}

	rawJSON := string(bodyStr)

	aq := airquality.MarshalIntoResponse(rawJSON)

	graph, lastMeasurement := airquality.GetPlottedComponent(aq, "PM2.5", 4)

	lastMeasurementStr := fmt.Sprintf("%f", lastMeasurement)

	fmt.Println("\n\n" + graph + "\n" + "Last value: " + lastMeasurementStr)
}
