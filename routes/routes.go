package routes

import (
	"log"
	"net/http"

	"github.com/tudemaha/logpress_server/internal/common"
)

func LoadRoutes() {
	log.Println("INFO LoadRoutes: loading routes...")

	http.HandleFunc("/ping", common.PingHandler())

	log.Println("INFO LoadRoutes: routes loaded.")
}
