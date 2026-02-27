package database

import (
	"log"
	"time"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectaComBancoDeDados() {
	stringDeConexao := "host=postgres user=root password=root dbname=root port=5432 sslmode=disable"

	var err error

	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
		if err == nil {
			log.Println("Conectado ao banco com sucesso!")
			break
		}

		log.Println("Banco ainda não está pronto, tentando novamente...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados após várias tentativas")
	}

	DB.AutoMigrate(&models.Aluno{})
}