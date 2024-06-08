package models

type Serial struct {
	S_name        string  `json:"name"`
	S_description string  `json:"description"`
	S_genre       string  `json:"genre"`
	S_state       string  `json:"state"`
	S_id          int     `json:"id"`
	S_idProducer  int     `json:"idProducer"`
	S_year        int     `json:"year"`
	S_seasons     int     `json:"seasons"`
	S_rating      float32 `json:"rating"`
	S_img         string  `json:"img"`
	S_duration    string  `json:"duration"`
}

func (s *Serial) Validate() bool {
	if s.S_idProducer <= 0 || s.S_name == "" || s.S_description == "" || s.S_year <= 0 || s.S_genre == "" || s.S_rating < 0 || s.S_seasons < 0 || s.S_state == "" || s.S_img == "" || s.S_duration == "" {
		return false
	}
	return true
}

func (s *Serial) GetId() int {
	return s.S_id
}

func (s *Serial) GetIdProducer() int {
	return s.S_idProducer
}

func (s *Serial) GetName() string {
	return s.S_name
}

func (s *Serial) GetDescription() string {
	return s.S_description
}

func (s *Serial) GetYear() int {
	return s.S_year
}

func (s *Serial) GetGenre() string {
	return s.S_genre
}

func (s *Serial) GetRating() float32 {
	return s.S_rating
}

func (s *Serial) GetSeasons() int {
	return s.S_seasons
}

func (s *Serial) GetState() string {
	return s.S_state
}

func (s *Serial) SetId(id int) {
	s.S_id = id
}

func (s *Serial) SetIdProducer(idProducer int) {
	s.S_idProducer = idProducer
}

func (s *Serial) SetName(name string) {
	s.S_name = name
}

func (s *Serial) SetDescription(description string) {
	s.S_description = description
}

func (s *Serial) SetYear(year int) {
	s.S_year = year
}

func (s *Serial) SetGenre(genre string) {
	s.S_genre = genre
}

func (s *Serial) SetRating(rating float32) {
	s.S_rating = rating
}

func (s *Serial) SetSeasons(seasons int) {
	s.S_seasons = seasons
}

func (s *Serial) SetState(state string) {
	s.S_state = state
}
