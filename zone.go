package rachio

import (
	"fmt"
	"time"
)

type Zone struct {
	client       *Rachio
	ID           string `json:"id"`
	ZoneNumber   int    `json:"zoneNumber"`
	Name         string `json:"name"`
	Enabled      bool   `json:"enabled"`
	CustomNozzle struct {
		Name          string  `json:"name"`
		InchesPerHour float64 `json:"inchesPerHour"`
	} `json:"customNozzle"`
	CustomSoil struct {
		Name string `json:"name"`
	} `json:"customSoil"`
	CustomSlope struct {
		Name      string `json:"name"`
		SortOrder int    `json:"sortOrder"`
	} `json:"customSlope"`
	CustomCrop struct {
		Name        string  `json:"name"`
		Coefficient float64 `json:"coefficient"`
	} `json:"customCrop"`
	CustomShade struct {
		Name string `json:"name"`
	} `json:"customShade"`
	AvailableWater             float64 `json:"availableWater"`
	RootZoneDepth              float64 `json:"rootZoneDepth"`
	ManagementAllowedDepletion float64 `json:"managementAllowedDepletion"`
	Efficiency                 float64 `json:"efficiency"`
	YardAreaSquareFeet         int     `json:"yardAreaSquareFeet"`
	ImageURL                   string  `json:"imageUrl"`
	ScheduleDataModified       bool    `json:"scheduleDataModified"`
	FixedRuntime               int     `json:"fixedRuntime"`
	RuntimeNoMultiplier        int     `json:"runtimeNoMultiplier"`
	WateringAdjustmentRuntimes struct {
		Num1 int `json:"1"`
		Num2 int `json:"2"`
		Num3 int `json:"3"`
		Num4 int `json:"4"`
		Num5 int `json:"5"`
	} `json:"wateringAdjustmentRuntimes"`
	SaturatedDepthOfWater float64 `json:"saturatedDepthOfWater"`
	DepthOfWater          float64 `json:"depthOfWater"`
	MaxRuntime            int     `json:"maxRuntime"`
	Runtime               int     `json:"runtime"`
	LastWateredDuration   int     `json:"lastWateredDuration,omitempty"`
	LastWateredDate       int64   `json:"lastWateredDate,omitempty"`
}

func (z *Zone) Start(duration time.Duration) error {
	msg := struct {
		ID       string `json:"id"`
		Duration uint64 `json:"duration"`
	}{
		ID:       z.ID,
		Duration: uint64(duration.Seconds()),
	}

	err := z.client.do("PUT", "/1/public/device/rain_delay", &msg, nil)
	return err
}

func (z *Zone) Get() error {
	err := z.client.do("GET", fmt.Sprintf("/1/public/zone/%s", z.ID), nil, z)
	if err != nil {
		return err
	}

	return nil
}
