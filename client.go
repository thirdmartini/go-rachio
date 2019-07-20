package rachio

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	rachioApiServer     = "https://api.rach.io"
	rachioApiPersonInfo = "/1/public/person/info"
)

type Rachio struct {
	token string
}

func (c *Rachio) rawRequest(path string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", rachioApiServer, path)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

func (c *Rachio) do(kind string, path string, in interface{}, out interface{}) error {
	url := fmt.Sprintf("%s%s", rachioApiServer, path)

	b := new(bytes.Buffer)
	if in != nil {
		json.NewEncoder(b).Encode(in)
	}

	req, err := http.NewRequest(kind, url, b)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		if out != nil {
			err = json.Unmarshal(body, out)
			if err != nil {
				return err
			}
		}
		return nil
	}

	return errors.New(string(body))
}

// Self Retrieve the person entity for the currently logged in user
func (c *Rachio) Self() (*Person, error) {
	msg := struct {
		ID string `json:"id"`
	}{}

	err := c.do("GET", rachioApiPersonInfo, nil, &msg)
	if err != nil {
		return nil, err
	}

	return c.Person(msg.ID)
}

// Person retrieves the information for a person entity
func (c *Rachio) Person(id string) (*Person, error) {
	person := &Person{
		client: c,
		ID:     id,
	}

	err := person.Get()
	if err != nil {
		return nil, err
	}
	return person, nil
}

// Device retrieves the information for a device entity
func (c *Rachio) Device(id string) (*Device, error) {
	device := &Device{
		client: c,
		ID:     id,
	}

	err := device.Get()
	if err != nil {
		return nil, err
	}
	return device, err
}

func (c *Rachio) DeviceList() ([]*Device, error) {
	var devices []*Device

	err := c.do("GET", "/1/public/person/", nil, devices)
	if err != nil {
		return nil, err
	}

	return devices, nil
}

// Zone retrieves the information for a zone entity
func (c *Rachio) Zone(id string) (*Zone, error) {
	zone := &Zone{
		client: c,
		ID:     id,
	}

	err := zone.Get()
	if err != nil {
		return nil, err
	}
	return zone, err
}

func (c *Rachio) ScheduleRule(id string) (*ScheduleRule, error) {
	sched := &ScheduleRule{
		client: c,
		ID:     id,
	}

	err := sched.Get()
	if err != nil {
		return nil, err
	}
	return sched, err
}

func (c *Rachio) FlexScheduleRule(id string) (*FlexScheduleRule, error) {
	sched := &FlexScheduleRule{
		client: c,
		ID:     id,
	}

	err := sched.Get()
	if err != nil {
		return nil, err
	}
	return sched, err
}

func (c *Rachio) Webhook(id string) (*Webhook, error) {
	hook := &Webhook{
		client: c,
		ID:     id,
	}

	err := hook.Get()
	if err != nil {
		return nil, err
	}
	return hook, err
}

// NewClient creates a new rachio sprinkler system API client
func NewClient(token string) *Rachio {
	c := Rachio{
		token: token,
	}
	return &c
}
