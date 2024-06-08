package models

type Producers struct {
	P_name    string `json:"name"`
	P_surname string `json:"surname"`
	P_id      int    `json:"id"`
}

func (p *Producers) Validate() bool {
	if p.P_name == "" || p.P_surname == "" {
		return false
	}
	return true
}

func (p *Producers) GetId() int {
	return p.P_id
}

func (p *Producers) GetName() string {
	return p.P_name
}

func (p *Producers) GetSurname() string {
	return p.P_surname
}

func (p *Producers) SetId(id int) {
	p.P_id = id
}

func (p *Producers) SetName(name string) {
	p.P_name = name
}

func (p *Producers) SetSurname(surname string) {
	p.P_surname = surname
}
