package models

type SerialsActors struct {
	Sa_id       int `json:"id"`
	Sa_idSerial int `json:"idSerial"`
	Sa_idActor  int `json:"idActor"`
}

func (sa *SerialsActors) Validate() bool {
	if sa.Sa_idSerial <= 0 || sa.Sa_idActor <= 0 {
		return false
	}
	return true
}

func (sa *SerialsActors) GetId() int {
	return sa.Sa_id
}

func (sa *SerialsActors) GetIdSerial() int {
	return sa.Sa_idSerial
}

func (sa *SerialsActors) GetIdActor() int {
	return sa.Sa_idActor
}

func (sa *SerialsActors) SetId(id int) {
	sa.Sa_id = id
}

func (sa *SerialsActors) SetIdSerial(idSerial int) {
	sa.Sa_idSerial = idSerial
}

func (sa *SerialsActors) SetIdActor(idActor int) {
	sa.Sa_idActor = idActor
}
