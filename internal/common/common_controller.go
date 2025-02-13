package common

import (
	"encoding/json"
	"net/http"

	"github.com/tudemaha/logpress_server/global/dto"
)

func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var response dto.Response
		if r.Method != "GET" {
			response.DefaultNotAllowed()
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		response.DefaultOK()
		response.Data = map[string]string{
			"ping": "pong",
		}

		json.NewEncoder(w).Encode(response)
	}
}
