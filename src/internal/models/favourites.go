package models

type Favourites struct {
	F_id         int `json:"id"`
	F_cntSerials int `json:"cntSerials"`
}

func (f *Favourites) Validate() bool {
	return f.F_cntSerials >= 0
}

func (f *Favourites) GetId() int {
	return f.F_id
}

func (f *Favourites) GetCntSerials() int {
	return f.F_cntSerials
}

func (f *Favourites) SetId(id int) {
	f.F_id = id
}

func (f *Favourites) SetCntSerials(cntSerials int) {
	f.F_cntSerials = cntSerials
}
