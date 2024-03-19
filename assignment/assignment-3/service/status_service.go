package service

import (
	"hacktiv-digitalent-golang/assignment/assignment-3/model"
	"math/rand"
	"time"
)

func GenerateRandomStatusValue() model.Status {
	source := rand.NewSource(time.Now().UnixNano())
	randomValue := rand.New(source)

	water := randomValue.Intn(100) + 1
	wind := randomValue.Intn(100) + 1

	statusWater := ""
	statusWind := ""

	if water < 5 {
		statusWater = "aman"
	} else if water >= 6 && water <= 8 {
		statusWater = "siaga"
	} else {
		statusWater = "bahaya"
	}

	if wind < 6 {
		statusWind = "aman"
	} else if wind >= 7 && wind <= 15 {
		statusWind = "siaga"
	} else {
		statusWind = "bahaya"
	}

	return model.Status {
		Water: water,
		Wind: wind,
		StatusWater: statusWater,
		StatusWind: statusWind,
	}
}