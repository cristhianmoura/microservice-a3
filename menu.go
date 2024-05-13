package main

import (
	"database/sql"
	"encoding/json"
	_ "fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Paciente struct { //estruturação dos pacientes
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:601220@tcp(localhost:3306)/microservice") //conexão com o banco de dados "microservice"
	if err != nil {
		log.Fatal("Erro na conexão com o banco de dados:", err)
	}
	defer db.Close()

	r := mux.NewRouter() //CRUD de pacientes
	r.HandleFunc("/pacientes", getPacientesHandler).Methods("GET")
	r.HandleFunc("/pacientes/{id}", getPacientesHandler).Methods("GET")
	//r.HandleFunc("/pacientes", createPacientesHandler).Methods("POST")
	//r.HandleFunc("/pacientes/{id}", updatePacientesHandler).Methods("PUT")
	//r.HandleFunc("/pacientes/{id}", deletePacientesHandler).Methods("DELETE")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8081", nil)) //iniciando um servidor na porta 8081
}

func getPacientesHandler(w http.ResponseWriter, r *http.Request) { //GET pacientes
	rows, err := db.Query("SELECT * FROM pacientes")
	if err != nil {
		http.Error(w, "Erro ao buscar pacientes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pacientes []Paciente
	for rows.Next() {
		var p Paciente
		err := rows.Scan(&p.ID, &p.Nome, &p.Idade)
		if err != nil {
			http.Error(w, "Erro ao ler pacientes", http.StatusInternalServerError)
			return
		}
		pacientes = append(pacientes, p)
	}

	json.NewEncoder(w).Encode(pacientes)
}

func getPacienteHandler(w http.ResponseWriter, r *http.Request) { //GET pacientes através de um "ID" de controle de usuários
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM pacientes WHERE id = ?", id)
	var p Paciente
	err = row.Scan(&p.ID, &p.Nome, &p.Idade)
	if err != nil {
		http.Error(w, "Paciente não encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(p)
}

/*func createPacienteHandler(w http.ResponseWriter, r *http.Request) {
	var p Paciente
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := db.Exec("INSERT INTO pacientes (nome, idade) VALUES (?, ?)", p.Nome, p.Idade)
	if err != nil {
		http.Error(w, "Erro ao inserir o paciente", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	p.ID = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func updatePacienteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var p Paciente
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_, err = db.Exec("UPDATE pacientes SET nome=?, idade=? WHERE id=?", p.Nome, p.Idade, id)
	if err != nil {
		http.Error(w, "Erro ao atualizar o paciente", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deletePacienteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM pacientes WHERE id=?", id)
	if err != nil {
		http.Error(w, "Erro ao deletar o paciente", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
*/
