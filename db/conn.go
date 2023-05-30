package conn

import (
	"database/sql"
	"log"

	lg "github.com/bahodurnazarov/middleTask/utils"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // add this
)

func Conn() *sql.DB {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	psqlInfo := "postgresql://postgres:postgres@localhost/AlifTask?sslmode=disable"

	// Подключение к базе данных PostgreSQL
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		lg.Errl.Fatal("Failed to connect to database. \n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		lg.Errl.Fatal("Panic . \n", err)
	}

	lg.Server.Println("Successfully connected!")
	return db
}
