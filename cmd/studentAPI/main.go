package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/piyushk8/StudentAPI/internal/config"
)

func main() {
	// loaidng configs
	cfg := config.MUSTLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write(([]byte("welcome to student api")))
	})
	fmt.Print("ddd",cfg.HTTPServer,cfg.Addr)
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	err := server.ListenAndServe()
	
	if err != nil {
		log.Fatal("failed to start server")
	}

	fmt.Println(("server started"))
}	
