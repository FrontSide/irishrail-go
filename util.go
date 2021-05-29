package irishrail

import "fmt"

func FilterDirection(stationData StationData, direction string) (StationData, error) {

	if direction != DirectionNorthbound && direction != DirectionSouthbound {
		return StationData{}, fmt.Errorf("invalid direction: %s", direction)
	}

	filteredTrains := []StationDataElement{}
	for _, dueTrain := range stationData.Elements {
		if dueTrain.Direction == direction {
			filteredTrains = append(filteredTrains, dueTrain)
		}
	}
	stationData.Elements = filteredTrains
	return stationData, nil
}
