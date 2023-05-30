package main

import (
	"fmt"
	"log"
	"os"

    "github.com/joho/godotenv"
	d "github.com/bahodurnazarov/middleTask/db"
	lg "github.com/bahodurnazarov/middleTask/utils"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dfs := os.Getenv("DB_PORT")
	log.Println(dfs)

	db := d.Conn()

	fmt.Println(db)
	lg.Errl.Println("Hello Error ")
	lg.Server.Println("Server Hello")

}
