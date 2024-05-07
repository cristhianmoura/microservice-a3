package main

import "menu/initializers"

func init() {
	initializers.LoadVariables()
	//initializers.ConnectDB()
}

func main() {
	//initializers.DB.AutoMigrate(&model.Post{}})
}
