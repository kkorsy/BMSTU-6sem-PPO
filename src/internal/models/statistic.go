package models

type Statistic struct {
	St_id            int
	St_gender_male   int
	St_gender_female int
	St_role_user     int
	St_role_admin    int
	St_age_0_18      int
	St_age_19_30     int
	St_age_31_50     int
	St_age_51_100    int
}

func (s *Statistic) Validate() bool {
	if s.St_id <= 0 || s.St_gender_male < 0 || s.St_gender_female < 0 || s.St_role_user < 0 || s.St_role_admin < 0 || s.St_age_0_18 < 0 || s.St_age_19_30 < 0 || s.St_age_31_50 < 0 || s.St_age_51_100 < 0 {
		return false
	}
	return true
}

func (s *Statistic) GetId() int {
	return s.St_id
}

func (s *Statistic) GetGenderMale() int {
	return s.St_gender_male
}

func (s *Statistic) GetGenderFemale() int {
	return s.St_gender_female
}

func (s *Statistic) GetRoleUser() int {
	return s.St_role_user
}

func (s *Statistic) GetRoleAdmin() int {
	return s.St_role_admin
}

func (s *Statistic) GetAge0_18() int {
	return s.St_age_0_18
}

func (s *Statistic) GetAge19_30() int {
	return s.St_age_19_30
}

func (s *Statistic) GetAge31_50() int {
	return s.St_age_31_50
}

func (s *Statistic) GetAge51_100() int {
	return s.St_age_51_100
}

func (s *Statistic) SetId(id int) {
	s.St_id = id
}

func (s *Statistic) IncreaseGenderMale() {
	s.St_gender_male++
}

func (s *Statistic) IncreaseGenderFemale() {
	s.St_gender_female++
}

func (s *Statistic) IncreaseRoleUser() {
	s.St_role_user++
}

func (s *Statistic) IncreaseRoleAdmin() {
	s.St_role_admin++
}

func (s *Statistic) IncreaseAge0_18() {
	s.St_age_0_18++
}

func (s *Statistic) IncreaseAge19_30() {
	s.St_age_19_30++
}

func (s *Statistic) IncreaseAge31_50() {
	s.St_age_31_50++
}

func (s *Statistic) IncreaseAge51_100() {
	s.St_age_51_100++
}

func (s *Statistic) DecreaseGenderMale() {
	s.St_gender_male--
}

func (s *Statistic) DecreaseGenderFemale() {
	s.St_gender_female--
}

func (s *Statistic) DecreaseRoleUser() {
	s.St_role_user--
}

func (s *Statistic) DecreaseRoleAdmin() {
	s.St_role_admin--
}

func (s *Statistic) DecreaseAge0_18() {
	s.St_age_0_18--
}

func (s *Statistic) DecreaseAge19_30() {
	s.St_age_19_30--
}

func (s *Statistic) DecreaseAge31_50() {
	s.St_age_31_50--
}

func (s *Statistic) DecreaseAge51_100() {
	s.St_age_51_100--
}
