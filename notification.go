package rachio

import (
	"errors"
	"fmt"
)

type Webhook struct {
	client     *Rachio
	ID         string `json:"id"`
	URL        string `json:"url"`
	ExternalID string `json:"externalId"`
	EventTypes []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"eventTypes"`
}

// Get refresh the contents of this webhook
func (w *Webhook) Get() error {
	err := w.client.do("GET", fmt.Sprintf("/1/public/notification/webhook/%s", w.ID), nil, w)
	if err != nil {
		return err
	}
	return err
}

func (w *Webhook) Update() error {
	err := w.client.do("PUT", fmt.Sprintf("/1/public/notification/webhook", w.ID), w, w)
	if err != nil {
		return err
	}
	return err
}

func (w *Webhook) Delete() error {
	err := w.client.do("DELETE", fmt.Sprintf("/1/public/notification/webhook/%s", w.ID), nil, nil)
	if err != nil {
		return err
	}
	return err
}

func (w *Webhook) Create() error {
	return errors.New("not implemented")
}
