package models

type Comments struct {
	C_text     string `json:"text"`
	C_date     string `json:"date"`
	C_id       int    `json:"id"`
	C_idUser   int    `json:"idUser"`
	C_idSerial int    `json:"idSerial"`
}

func (c *Comments) Validate() bool {
	if c.C_idUser <= 0 || c.C_text == "" || c.C_date == "" || c.C_idSerial <= 0 {
		return false
	}
	return true
}

func (c *Comments) GetId() int {
	return c.C_id
}

func (c *Comments) GetIdUser() int {
	return c.C_idUser
}

func (c *Comments) GetText() string {
	return c.C_text
}

func (c *Comments) GetDate() string {
	return c.C_date
}

func (c *Comments) GetIdSerial() int {
	return c.C_idSerial
}

func (c *Comments) SetId(id int) {
	c.C_id = id
}

func (c *Comments) SetIdUser(idUser int) {
	c.C_idUser = idUser
}

func (c *Comments) SetText(text string) {
	c.C_text = text
}

func (c *Comments) SetDate(date string) {
	c.C_date = date
}

func (c *Comments) SetIdSerial(idSerial int) {
	c.C_idSerial = idSerial
}
