package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tnlitv/mindMapServer/auth/handlers"
)

func main() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":8000"),
		Handler: handlers.New(),
	}

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
