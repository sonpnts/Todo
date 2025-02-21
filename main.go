package main

import (
	"fmt"
	"github.com/sonpnts/todo-list/config"
)

func main() {
	r := config.SetupRouter()
	fmt.Println("Server is running on port 8080...")
	r.Run(":8080")
}
