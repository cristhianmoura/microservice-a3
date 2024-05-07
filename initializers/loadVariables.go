package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro no carregamento") //carrega as variaveis do banco e caso dê algo errado informa.
	}
}
