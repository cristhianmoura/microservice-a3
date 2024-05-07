package main

import (
	"fmt"
	"menu/initializers"
	"net/http"

	"os"

	"github.com/gorilla/mux"
)

func init() {
	initializers.LoadVariables() //carregando as variáveis do banco de dados assim que o serviço for iniciado
}

func loadData() []byte { //iniciando o microserviço

	jsonFile, err := os.Open("banco.json")
	if err != nil {
		fmt.Println("erro: ", err.Error()) //retorna erro caso algo dê errado
	}
	defer jsonFile.Close()

	data, err := os.ReadFile("banco.json")
	return data //processa os dados do banco(que por enquanto está sendo um arquivo .json)
}
func ListUsers(w http.ResponseWriter, r *http.Request) { //lista e imprime na tela as informações desse banco temporário
	banco := loadData()

	w.Write([]byte(banco))
}
func main() { //essa função é um direcionamento para localizar o serviço
	r := mux.NewRouter()
	r.HandleFunc("/users", ListUsers)

	http.ListenAndServe(":8081", r)
}
