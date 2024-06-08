package models

type Seasons struct {
	Ss_name        string `json:"name"`
	Ss_date        string `json:"date"`
	Ss_id          int    `json:"id"`
	Ss_idSerial    int    `json:"idSerial"`
	Ss_num         int    `json:"num"`
	Ss_cntEpisodes int    `json:"cntEpisodes"`
}

func (s *Seasons) Validate() bool {
	if s.Ss_idSerial <= 0 || s.Ss_name == "" || s.Ss_num < 0 || s.Ss_cntEpisodes < 0 || s.Ss_date == "" {
		return false
	}
	return true
}

func (s *Seasons) GetId() int {
	return s.Ss_id
}

func (s *Seasons) GetIdSerial() int {
	return s.Ss_idSerial
}

func (s *Seasons) GetName() string {
	return s.Ss_name
}

func (s *Seasons) GetNum() int {
	return s.Ss_num
}

func (s *Seasons) GetCntEpisodes() int {
	return s.Ss_cntEpisodes
}

func (s *Seasons) GetDate() string {
	return s.Ss_date
}

func (s *Seasons) SetId(id int) {
	s.Ss_id = id
}

func (s *Seasons) SetIdSerial(idSerial int) {
	s.Ss_idSerial = idSerial
}

func (s *Seasons) SetName(name string) {
	s.Ss_name = name
}

func (s *Seasons) SetNum(num int) {
	s.Ss_num = num
}

func (s *Seasons) SetCntEpisodes(cntEpisodes int) {
	s.Ss_cntEpisodes = cntEpisodes
}

func (s *Seasons) SetDate(date string) {
	s.Ss_date = date
}
