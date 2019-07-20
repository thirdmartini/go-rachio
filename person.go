package rachio

import "fmt"

type Person struct {
	client     *Rachio
	ID         string   `json:"id"`
	CreateDate int64    `json:"createDate"`
	Username   string   `json:"username"`
	FullName   string   `json:"fullName"`
	Email      string   `json:"email"`
	Devices    []Device `json:"devices"`
	Deleted    bool     `json:"deleted"`
}

func (p *Person) resolve() {
	for idx := range p.Devices {
		dev := &p.Devices[idx]
		dev.client = p.client
		dev.resolve()
	}
}

func (p *Person) Get() error {
	err := p.client.do("GET", fmt.Sprintf("/1/public/person/%s", p.ID), nil, p)
	if err != nil {
		return err
	}

	p.resolve()
	return nil
}
