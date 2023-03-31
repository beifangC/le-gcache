package models

type Pers struct {
	Id int
	K  string
	V  string
}

func (Pers) TableNmae() string {
	return "pers"
}
