package main

import (
	"fmt"

	lg "github.com/bahodurnazarov/middleTask/utils"
)

func main() {
	fmt.Println("Hello")
	lg.Errl.Println("Hello Error ")
	lg.Server.Println("Server Hello")

}
