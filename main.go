package main

import (
	"log"

	"github.com/diego-dm-morais/order-manager/framework_drivers"
)

func main() {
	log.Println("Iniciando a aplicação Order Manager")
	new(framework.Application).Init()
}
