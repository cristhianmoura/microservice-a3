package model

import "gorm.io/gorm"

type Paciente struct { //Variavéis do paciente
	gorm.Model
	Name  string
	Idade int
	Tipo  string
}

type Medico struct { //Variavéis do médico
	gorm.Model
	Name  string
	Idade int
	Tipo  string
}
