package irishrail

import (
	"net/url"
	"sort"
)

const stationDataAPI = "http://api.irishrail.ie/realtime/realtime.asmx/getStationDataByNameXML"

const (
	DirectionNorthbound = "Northbound"
	DirectionSouthbound = "Southbound"
)

type StationData struct {
	Elements []StationDataElement `xml:"objStationData"`
}

type StationDataElement struct {
	ExpectedArrivalTime string `xml:"Exparrival"`
	Direction           string `xml:"Direction"`
	DueInMinutes        int    `xml:"Duein"`
}

func GetStationData(station string) (StationData, error) {

	stationData := StationData{}
	params := url.Values{"StationDesc": []string{station}}
	err := SendHTTPGetRequestXML(stationDataAPI, params, &stationData)
	if err != nil {
		return StationData{}, err
	}

	dueTrains := stationData.Elements

	sort.Slice(dueTrains, func(i, j int) bool {
		return dueTrains[i].DueInMinutes < dueTrains[j].DueInMinutes
	})

	stationData.Elements = dueTrains
	return stationData, nil

}
