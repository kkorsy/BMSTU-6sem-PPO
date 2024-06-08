package models

type SerialsUsers struct {
	Su_lastSeen string `json:"lastSeen"`
	Su_id       int    `json:"id"`
	Su_idSerial int    `json:"idSerial"`
	Su_idUser   int    `json:"idUser"`
}

func (su *SerialsUsers) Validate() bool {
	if su.Su_idSerial <= 0 || su.Su_idUser <= 0 || su.Su_lastSeen == "" {
		return false
	}
	return true
}

func (su *SerialsUsers) GetId() int {
	return su.Su_id
}

func (su *SerialsUsers) GetIdSerial() int {
	return su.Su_idSerial
}

func (su *SerialsUsers) GetIdUser() int {
	return su.Su_idUser
}

func (su *SerialsUsers) GetLastSeen() string {
	return su.Su_lastSeen
}

func (su *SerialsUsers) SetId(id int) {
	su.Su_id = id
}

func (su *SerialsUsers) SetIdSerial(idSerial int) {
	su.Su_idSerial = idSerial
}

func (su *SerialsUsers) SetIdUser(idUser int) {
	su.Su_idUser = idUser
}

func (su *SerialsUsers) SetLastSeen(lastSeen string) {
	su.Su_lastSeen = lastSeen
}
