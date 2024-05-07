package main

import (
	"fmt"
	"net/http"

	"os"

	"github.com/gorilla/mux"
)

func loadData() []byte {

	jsonFile, err := os.Open("banco.json")
	if err != nil {
		fmt.Println("erro: ", err.Error())

	}
	defer jsonFile.Close()

	data, err := os.ReadFile("banco.json")
	return data
}
func ListUsers(w http.ResponseWriter, r *http.Request) {

	banco := loadData()

	w.Write([]byte(banco))
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", ListUsers)

	http.ListenAndServe(":8081", r)
}
