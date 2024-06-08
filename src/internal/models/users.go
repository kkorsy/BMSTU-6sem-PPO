package models

import (
	"time"
)

type Users struct {
	U_login        string `json:"login"`
	U_password     string `json:"password"`
	U_role         string `json:"role"`
	U_name         string `json:"name"`
	U_surname      string `json:"surname"`
	U_gender       string `json:"gender"`
	U_bdate        string `json:"bdate"`
	U_id           int    `json:"id"`
	U_idFavourites int    `json:"idFavourites"`
}

func (u *Users) Validate() bool {
	if u.U_idFavourites <= 0 || u.U_login == "" ||
		u.U_password == "" || u.U_role == "" || u.U_name == "" ||
		u.U_surname == "" || u.U_gender == "" || u.U_bdate == "" {
		return false
	}
	return true
}

func (u *Users) GetAge() int {
	bdate, err := time.Parse("2006-01-02", u.U_bdate)
	if err != nil {
		bdate, _ = time.Parse("02.01.2006", u.U_bdate)
	}
	now := time.Now()
	age := now.Year() - bdate.Year()
	if now.Month() < bdate.Month() || (now.Month() == bdate.Month() && now.Day() < bdate.Day()) {
		age--
	}
	return age
}

func (u *Users) GetId() int {
	return u.U_id
}

func (u *Users) GetIdFavourites() int {
	return u.U_idFavourites
}

func (u *Users) GetLogin() string {
	return u.U_login
}

func (u *Users) GetPassword() string {
	return u.U_password
}

func (u *Users) GetRole() string {
	return u.U_role
}

func (u *Users) GetName() string {
	return u.U_name
}

func (u *Users) GetSurname() string {
	return u.U_surname
}

func (u *Users) GetGender() string {
	return u.U_gender
}

func (u *Users) GetBdate() string {
	return u.U_bdate
}

func (u *Users) SetId(id int) {
	u.U_id = id
}

func (u *Users) SetIdFavourites(idFavourites int) {
	u.U_idFavourites = idFavourites
}

func (u *Users) SetLogin(login string) {
	u.U_login = login
}

func (u *Users) SetPassword(password string) {
	u.U_password = password
}

func (u *Users) SetRole(role string) {
	u.U_role = role
}

func (u *Users) SetName(name string) {
	u.U_name = name
}

func (u *Users) SetSurname(surname string) {
	u.U_surname = surname
}

func (u *Users) SetGender(gender string) {
	u.U_gender = gender
}

func (u *Users) SetBdate(bdate string) {
	u.U_bdate = bdate
}
