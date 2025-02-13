package routes

import (
	"log"
	"net/http"

	"github.com/tudemaha/logpress_server/internal/common"
	"github.com/tudemaha/logpress_server/internal/decompress/controller"
)

func LoadRoutes() {
	log.Println("INFO LoadRoutes: loading routes...")

	http.HandleFunc("/ping", common.PingHandler())
	http.HandleFunc("/upload", controller.DecompressHandler())

	log.Println("INFO LoadRoutes: routes loaded.")
}
