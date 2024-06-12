package main

import (
	"CrudProject/internal/configs"
	"CrudProject/pkg/handlers"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"log"
	"net/http"
)

func main() {
	log.Println("[Server] Loading Environment Variables")
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	log.Println("[Server] Setting Middlewares")
	r.Use(cors.AllowAll().Handler)

	log.Println("[Server] Setting Routes")
	r.Post("/api/", handlers.VerifyApplication)

	log.Println("[Server] Starting Server")
	err = http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
	if err != nil {
		log.Println(err)
	}

}
