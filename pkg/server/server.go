package server

import (
	"log"
	"net/http"
)

func StartServer() {
	log.Println("INFO StartServer: server started at :8080")
	http.ListenAndServe(":8080", nil)
}
