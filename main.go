package main

import "github.com/bahodurnazarov/middleTask/routes"

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// dfs := os.Getenv("DB_PORT")
	// log.Println(dfs)

	// db := d.Conn()

	// fmt.Println(db)
	// lg.Errl.Println("Hello Error ")
	// lg.Server.Println("Server Hello")
	routes.Listening()
}
