package models

type Episodes struct {
	E_name     string `json:"name"`
	E_date     string `json:"date"`
	E_id       int    `json:"id"`
	E_idSeason int    `json:"idSeason"`
	E_num      int    `json:"num"`
	E_duration string `json:"duration"`
}

func (e *Episodes) Validate() bool {
	if e.E_idSeason <= 0 || e.E_name == "" || e.E_num < 0 || e.E_duration == "" || e.E_date == "" {
		return false
	}
	return true
}

func (e *Episodes) GetId() int {
	return e.E_id
}

func (e *Episodes) GetIdSeason() int {
	return e.E_idSeason
}

func (e *Episodes) GetName() string {
	return e.E_name
}

func (e *Episodes) GetNum() int {
	return e.E_num
}

func (e *Episodes) GetDuration() string {
	return e.E_duration
}

func (e *Episodes) GetDate() string {
	return e.E_date
}

func (e *Episodes) SetId(id int) {
	e.E_id = id
}

func (e *Episodes) SetIdSeason(idSeason int) {
	e.E_idSeason = idSeason
}

func (e *Episodes) SetName(name string) {
	e.E_name = name
}

func (e *Episodes) SetNum(num int) {
	e.E_num = num
}

func (e *Episodes) SetDuration(duration string) {
	e.E_duration = duration
}

func (e *Episodes) SetDate(date string) {
	e.E_date = date
}
