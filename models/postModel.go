package models

import "gorm.io/gorm"

type Paciente struct { //Variavéis do paciente
	gorm.Model
	Nome  string
	Idade int
	Tipo  string
}

type Medico struct { //Variavéis do médico
	gorm.Model
	Nome  string
	Idade int
	Tipo  string
}

func CreatePac() {
	initializers.LoadVariables()
	paciente := Paciente{Nome: "Jinzhu", Idade: 18, Tipo: "Paciente"}
	criar := db.Create(&Paciente)
	user.ID      
	result.Error
		if err != nil {
			fmt.Println("erro: ", err.Error()) 

	
	db.Select([]string{"nome", "idade", "tipo"}).Find(&Medico{Nome}; {Idade}; {Tipo})
}

func ReadPac() {

}

func UpdatePac() {
	initializers.LoadVariables()
	db.First(&Paciente)
	user.Name = "jinzhu 2" //manter user no banco de dados 
	user.Idade = 100
	user.Tipo = "paciente"

	db.Save(&Paciente)
	db.Save(&Paciente{Name: "jinzhu", Idade: 100 , Tipo: "Paciente"})
}

func DeletePac() {
	initializers.LoadVariables()
}


func CreateMed() {
	initializers.LoadVariables()
	medico := Medico{Name: "Jinzhu", Idade: 18, Tipo: "Medico"}
	criar := db.Create(&Medico)
	user.ID      
	result.Error
		if err != nil {
			fmt.Println("erro: ", err.Error()) 

	
	db.Select([]string{"nome"}).Find(&Medico{Nome})
}

func ReadMed() {

}

func UpdateMed() {
	initializers.LoadVariables()
	db.First(&Medico)
	user.Name = "jinzhu 2" //manter user no banco de dados 
	user.Idade = 100
	user.Tipo = "medico"

	db.Save(&Medico)
	db.Save(&Medico{Nome: "jinzhu", Idade: 100, Tipo: "Medico"})
}

func DeleteMed() {
	initializers.LoadVariables()
}