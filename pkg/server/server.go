package server

import (
	"log"
	"net/http"
	"os"
)

func StartServer() {
	log.Println("INFO StartServer: server started at :8080")
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}
