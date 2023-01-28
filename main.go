package main

import (
	"fmt"

	"github.com/luuisavelino/short-circuit-analysis-critical-time/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequests()
}
