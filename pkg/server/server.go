package server

import (
	"log"
	"net/http"
	"os"
)

func StartServer() {
	port := os.Getenv("PORT")
	log.Println("INFO StartServer: server started at :", port)
	http.ListenAndServe(":"+port, nil)
}
