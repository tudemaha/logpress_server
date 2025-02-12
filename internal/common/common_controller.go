package common

import (
	"encoding/json"
	"net/http"

	"github.com/tudemaha/logpress_server/global/dto"
)

func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response dto.Response
		response.DefaultOK()
		response.Data = map[string]string{
			"ping": "pong",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
