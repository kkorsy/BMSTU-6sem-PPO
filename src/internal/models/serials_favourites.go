package models

type SerialsFavourites struct {
	Sf_id          int `json:"id"`
	Sf_idSerial    int `json:"idSerial"`
	Sf_idFavourite int `json:"idFavourite"`
}

func (sf *SerialsFavourites) Validate() bool {
	if sf.Sf_idSerial <= 0 || sf.Sf_idFavourite <= 0 {
		return false
	}
	return true
}

func (sf *SerialsFavourites) GetId() int {
	return sf.Sf_id
}

func (sf *SerialsFavourites) GetIdSerial() int {
	return sf.Sf_idSerial
}

func (sf *SerialsFavourites) GetIdFavourite() int {
	return sf.Sf_idFavourite
}

func (sf *SerialsFavourites) SetId(id int) {
	sf.Sf_id = id
}

func (sf *SerialsFavourites) SetIdSerial(idSerial int) {
	sf.Sf_idSerial = idSerial
}

func (sf *SerialsFavourites) SetIdFavourite(idFavourite int) {
	sf.Sf_idFavourite = idFavourite
}
