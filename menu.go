package main

import (
	"database/sql"
	"fmt"
	_ "html/template"
	"net/http"
	_ "net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:601220@tcp(localhost:3306)/microservice")
	if err != nil {
		fmt.Println("Erro na validação do sql.Open")
		panic(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/productsearch", productSearchHandler)
	http.HandleFunc("/", homePageHandler)
	http.ListenAndServe("localhost:8080", nil)


	func productSearchHandler(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			tp1.ExecuteTemplate(w, "productsearch.html", nil)
			return
		}
		r.ParseForm()
		var P Product
		name := r.FormValue("productName")
		fmt.Println("name", name)

		stmt := "SELECT * FROM products WHERE name = ?;"

		row := db.QueryRow(stmt, name)

		err := row.Scan(&P.ID, &P.Name, &P.Price, &P.Description)
		if err != nil {
			panic(err)
		}
		tp1.ExecuteTemplate(w, "productsearch.html", P)
	}
	
	func homePageHandler2(w http.ResponseWriter, r *http.Request) {

	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Erro na conexão com Ping")
		panic(err.Error())
	}
	//insert, err := db.Query("INSERT INTO 'microservice' , 'pacientes' ('codigo', 'nome', 'idade') VALUES ('2', 'Cristhian', '20');")
	//if err != nil {
	//panic(err.Error())
	//}

	//defer insert.Close()

	fmt.Println("Sucesso na conexão com o banco de dados ")
}
