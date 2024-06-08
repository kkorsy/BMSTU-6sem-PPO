package models

type Actors struct {
	A_name    string `json:"name"`
	A_surname string `json:"surname"`
	A_gender  string `json:"gender"`
	A_bdate   string `json:"bdate"`
	A_id      int    `json:"id"`
}

func (a *Actors) Validate() bool {
	if a.A_name == "" || a.A_surname == "" || a.A_gender == "" || a.A_bdate == "" {
		return false
	}
	return true
}

func (a *Actors) GetId() int {
	return a.A_id
}

func (a *Actors) GetName() string {
	return a.A_name
}

func (a *Actors) GetSurname() string {
	return a.A_surname
}

func (a *Actors) GetGender() string {
	return a.A_gender
}

func (a *Actors) GetBdate() string {
	return a.A_bdate
}

func (a *Actors) SetId(id int) {
	a.A_id = id
}

func (a *Actors) SetName(name string) {
	a.A_name = name
}

func (a *Actors) SetSurname(surname string) {
	a.A_surname = surname
}

func (a *Actors) SetGender(gender string) {
	a.A_gender = gender
}

func (a *Actors) SetBdate(bdate string) {
	a.A_bdate = bdate
}
