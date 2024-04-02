package main

import (
	// "log"

	// "github.com/BrandokVargas/api-contact/repository"
	// "github.com/BrandokVargas/api-contact/services/user"
	"log"
	"net/http"

	"github.com/BrandokVargas/api-contact/repository"
	"github.com/BrandokVargas/api-contact/routes"
	"github.com/BrandokVargas/api-contact/services/user"
	"github.com/BrandokVargas/api-contact/storage"
)

func main() {
	//Connection db postgress
	storage.NewPostgresDB()

	storageUser := repository.NewRepositoryUser(storage.Pool())
	serviceUser := user.NewService(storageUser)

	//SERVERMUX
	mux := http.NewServeMux()
	routes.RouteUser(mux, serviceUser)

	//Server escuche en el puerto 8080
	log.Println("Server on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error al iniciar el servidor: %v\n", err)
	}

}
