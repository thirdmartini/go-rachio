package rachio

import (
	"fmt"
	"time"
)

type Units string

const (
	US     Units = "US"
	METRIC Units = "Metric"
)

// Device describes a Rachio sprinkler controller
type Device struct {
	client            *Rachio
	CreateDate        int64              `json:"createDate"`
	ID                string             `json:"id"`
	Status            string             `json:"status"`
	Zones             []Zone             `json:"zones"`
	TimeZone          string             `json:"timeZone"`
	Latitude          float64            `json:"latitude"`
	Longitude         float64            `json:"longitude"`
	Name              string             `json:"name"`
	ScheduleRules     []ScheduleRule     `json:"scheduleRules"`
	SerialNumber      string             `json:"serialNumber"`
	MacAddress        string             `json:"macAddress"`
	IsOn              bool               `json:"on"`
	FlexScheduleRules []FlexScheduleRule `json:"flexScheduleRules"`
	Model             string             `json:"model"`
	ScheduleModeType  string             `json:"scheduleModeType"`
	Deleted           bool               `json:"deleted"`
	HomeKitCompatible bool               `json:"homeKitCompatible"`
	UtcOffset         int                `json:"utcOffset"`
}

// Event describes a sprinkler event
type Event struct {
	Id        string
	DeviceId  string
	Category  string
	Type      string
	SubType   string
	EventDate int64
	Topic     string
	Summary   string
	Hidden    bool
}

// fixup our zones so we can interact with them directly
func (d *Device) resolve() {
	for idx := range d.Zones {
		zone := &d.Zones[idx]
		zone.client = d.client
	}

	for idx := range d.ScheduleRules {
		sched := &d.ScheduleRules[idx]
		sched.client = d.client
	}

	for idx := range d.FlexScheduleRules {
		sched := &d.FlexScheduleRules[idx]
		sched.client = d.client
	}
}

// Events returns a list of events in the given time window
//
// Note: the rachio API restricts this to about a month of events
func (d *Device) Events(start, end time.Time) ([]Event, error) {
	startV := start.Unix() * 1000
	endV := end.Unix() * 1000

	var events []Event

	err := d.client.do("GET", fmt.Sprintf("/1/public/device/%s/event?startTime=%d&endTime=%d", d.ID, startV, endV), nil, &events)
	if err != nil {
		return nil, err
	}

	return events, nil
}

type CurrentSchedule struct {
	client          *Rachio
	DeviceID        string `json:"deviceId"`
	ScheduleID      string `json:"scheduleId"`
	Type            string `json:"type"`
	Status          string `json:"status"`
	StartDate       int64  `json:"startDate"`
	Duration        int    `json:"duration"`
	ZoneID          string `json:"zoneId"`
	ZoneStartDate   int64  `json:"zoneStartDate"`
	ZoneDuration    int    `json:"zoneDuration"`
	CycleCount      int    `json:"cycleCount"`
	TotalCycleCount int    `json:"totalCycleCount"`
	Cycling         bool   `json:"cycling"`
	DurationNoCycle int    `json:"durationNoCycle"`
}

func (d *Device) CurrentSchedule() (*CurrentSchedule, error) {
	sched := &CurrentSchedule{
		client: d.client,
	}

	err := d.client.do("GET", fmt.Sprintf("/1/public/device/%s/current_schedule", d.ID), nil, sched)
	if err != nil {
		return nil, err
	}

	return sched, nil
}

type Conditions struct {
	Time               int       `json:"time"`
	PrecipIntensity    int       `json:"precipIntensity"`
	PrecipProbability  int       `json:"precipProbability"`
	WindSpeed          int       `json:"windSpeed"`
	Humidity           float64   `json:"humidity"`
	CloudCover         int       `json:"cloudCover"`
	DewPoint           int       `json:"dewPoint"`
	WeatherType        string    `json:"weatherType"`
	UnitType           string    `json:"unitType"`
	CurrentTemperature int       `json:"currentTemperature"`
	WeatherSummary     string    `json:"weatherSummary"`
	DailyWeatherType   string    `json:"dailyWeatherType"`
	PrettyTime         time.Time `json:"prettyTime"`
}

type Forecast struct {
	Current  Conditions   `json:"current"`
	Forecast []Conditions `json:"forecast"`
}

// Forecast returns the weather forecast for this device
//
// Units can be US|METRIC
func (d *Device) Forecast(units Units) (*Forecast, error) {
	forecast := &Forecast{}

	err := d.client.do("GET",
		fmt.Sprintf("/1/public/device/%s/forcast?units=%s", d.ID, units),
		nil, forecast)
	if err != nil {
		return nil, err
	}
	return forecast, nil
}

// StopWater stops all watering on device
func (d *Device) StopWater() error {
	msg := struct {
		ID string `json:"id"`
	}{
		ID: d.ID,
	}

	err := d.client.do("PUT", "/1/public/device/stop_water", &msg, nil)
	return err
}

// RainDelay set rain delay on the device
//
// Duration max is 7 days
func (d *Device) RainDelay(duration time.Duration) error {
	msg := struct {
		ID       string `json:"id"`
		Duration uint64 `json:"duration"`
	}{
		ID:       d.ID,
		Duration: uint64(duration.Seconds()),
	}

	err := d.client.do("PUT", "/1/public/device/rain_delay", &msg, nil)
	return err
}

// On turns ON all features of the device (schedules, weather intelligence, water budget, etc.)
func (d *Device) On() error {
	msg := struct {
		ID string `json:"id"`
	}{
		ID: d.ID,
	}

	err := d.client.do("PUT", "/1/public/device/on", &msg, nil)
	return err
}

// Off turns OFF all features of the device (schedules, weather intelligence, water budget, etc.)
func (d *Device) Off() error {
	msg := struct {
		ID string `json:"id"`
	}{
		ID: d.ID,
	}

	err := d.client.do("PUT", "/1/public/device/off", &msg, nil)
	return err
}

// PauseZoneRun pauses running zones on the device
//  Duration max is 1 hour
func (d *Device) PauseZoneRun(duration time.Duration) error {
	msg := struct {
		ID       string `json:"id"`
		Duration uint64 `json:"duration"`
	}{
		ID:       d.ID,
		Duration: uint64(duration.Seconds()),
	}

	err := d.client.do("PUT", "/1/public/device/pause_zone_run", &msg, nil)
	return err
}

// ResumeZoneRun resumes zone runs on the device
func (d *Device) ResumeZoneRun() error {
	msg := struct {
		ID string `json:"id"`
	}{
		ID: d.ID,
	}

	err := d.client.do("PUT", "/1/public/device/resume_zone_run", &msg, nil)
	return err
}

// Get updates this device with any new information
func (d *Device) Get() error {
	err := d.client.do("GET", fmt.Sprintf("/1/public/device/%s", d.ID), nil, d)
	if err != nil {
		return err
	}

	d.resolve()
	return err
}

func (d *Device) Webhook() ([]Webhook, error) {
	var webhooks []Webhook

	err := d.client.do("GET", fmt.Sprintf("/1/public/notification/%s/webhook", d.ID), nil, webhooks)
	if err != nil {
		return nil, err
	}

	for idx := range webhooks {
		hook := &webhooks[idx]
		hook.client = d.client
	}

	return webhooks, nil
}
