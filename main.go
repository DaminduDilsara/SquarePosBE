package main

import (
	"github.com/Square-POC/SquarePosBE/configurations"
	"github.com/Square-POC/SquarePosBE/internal/clients"
	"github.com/Square-POC/SquarePosBE/internal/controllers"
	"github.com/Square-POC/SquarePosBE/internal/services"
	"github.com/Square-POC/SquarePosBE/internal/transport/http"
	"log"
	"os"
)

func main() {
	log.Println("starting userservice")

	sig := make(chan os.Signal, 0)

	conf := configurations.LoadConfigurations()

	loyaltyClient := clients.NewLoyaltyClient(conf.SquareConfig)

	servicesCollection := services.InitServices(conf, loyaltyClient)

	controller := controllers.NewControllerV1(servicesCollection)

	http.InitServer(conf.AppConfig, controller)

	select {
	case <-sig:
		log.Println("Application is shutting down..")

		http.Shutdown()
		os.Exit(0)
	}
}
