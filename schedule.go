package rachio

import "fmt"

type ScheduleRule struct {
	client *Rachio
	ID     string `json:"id"`
	Zones  []struct {
		ZoneID    string `json:"zoneId"`
		Duration  int    `json:"duration"`
		SortOrder int    `json:"sortOrder"`
	} `json:"zones"`
	ScheduleJobTypes []string `json:"scheduleJobTypes"`
	StartHour        int      `json:"startHour,omitempty"`
	StartMinute      int      `json:"startMinute,omitempty"`
	Operator         string   `json:"operator"`
	Summary          string   `json:"summary"`
	CycleSoakStatus  string   `json:"cycleSoakStatus"`
	StartDate        int64    `json:"startDate"`
	Name             string   `json:"name"`
	Enabled          bool     `json:"enabled"`
	StartDay         int      `json:"startDay"`
	StartMonth       int      `json:"startMonth"`
	StartYear        int      `json:"startYear"`
	TotalDuration    int      `json:"totalDuration"`
	EtSkip           bool     `json:"etSkip"`
	ExternalName     string   `json:"externalName"`
	CycleSoak        bool     `json:"cycleSoak"`
}

func (s *ScheduleRule) Start() error {
	msg := struct {
		ID string `json:"id"`
	}{
		ID: s.ID,
	}

	err := s.client.do("PUT", "/1/public/schedulerule/start", &msg, nil)
	return err
}

func (s *ScheduleRule) Skip() error {
	msg := struct {
		ID string `json:"id"`
	}{
		ID: s.ID,
	}

	err := s.client.do("PUT", "/1/public/schedulerule/skip", &msg, nil)
	return err
}

func (s *ScheduleRule) SeasonalAdjustment(adjustment float64) error {
	msg := struct {
		ID         string  `json:"id"`
		Adjustment float64 `json:"adjustment"`
	}{
		ID:         s.ID,
		Adjustment: adjustment,
	}

	err := s.client.do("PUT", "/1/public/schedulerule/seasonal_adjustment", &msg, nil)
	return err
}

func (s *ScheduleRule) Get() error {
	err := s.client.do("GET", fmt.Sprintf("/1/public/schedulerule/%s", s.ID), nil, s)
	if err != nil {
		return err
	}

	return err
}

// FlexScheduleRule describes a flex schedule rule
type FlexScheduleRule struct {
	client *Rachio
	ID     string `json:"id"`
	Zones  []struct {
		ZoneID     string `json:"zoneId"`
		ZoneNumber int    `json:"zoneNumber"`
		SortOrder  int    `json:"sortOrder"`
	} `json:"zones"`
	ScheduleJobTypes     []string `json:"scheduleJobTypes"`
	Summary              string   `json:"summary"`
	CycleSoak            bool     `json:"cycleSoak"`
	StartDate            int64    `json:"startDate"`
	Name                 string   `json:"name"`
	Enabled              bool     `json:"enabled"`
	TotalDuration        int      `json:"totalDuration"`
	TotalDurationNoCycle int      `json:"totalDurationNoCycle"`
	ExternalName         string   `json:"externalName"`
	Type                 string   `json:"type"`
}

func (s *FlexScheduleRule) Get() error {
	err := s.client.do("GET", fmt.Sprintf("/1/public/flexschedulerule/%s", s.ID), nil, s)
	if err != nil {
		return err
	}

	return err
}
