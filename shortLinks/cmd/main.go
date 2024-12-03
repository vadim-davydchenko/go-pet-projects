package main

import (
	"fmt"
	"net/http"
	"shortLinks/configs"
	"shortLinks/internal/auth"
	"shortLinks/link"
	"shortLinks/pkg/db"
	"shortLinks/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	linkRepository := link.NewLinkRepository(db)

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logging(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
