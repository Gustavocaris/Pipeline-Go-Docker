package database

import (
	"log"
	"os"
	"time"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectaComBancoDeDados() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	stringDeConexao := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + "stgres user=root password=root dbname=root port=5432 sslmode=disable"

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

	_ = DB.AutoMigrate(&models.Aluno{})
}