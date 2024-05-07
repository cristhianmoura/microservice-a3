package main

import "menu/initializers"

func init() {
	initializers.LoadVariables()
	//initializers.ConnectDB()
}

func main() { //função para fazer a conexão do banco de dados com o serviço
	//initializers.DB.AutoMigrate(&model.Post{}})
}
