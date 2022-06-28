package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andrersp/controle-financeiro/src/core"
	"github.com/andrersp/controle-financeiro/src/database"
	"github.com/andrersp/controle-financeiro/src/routers"
)

func main() {

	core.LoadConfig()

	if err := database.SetupAPP(); err != nil {
		log.Fatal(err)

	}

	r := routers.Load()

	fmt.Printf("Open on port %d \n", core.API_PORT)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", core.API_PORT), r))

}
