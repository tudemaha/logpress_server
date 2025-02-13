package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/tudemaha/logpress_server/global/dto"
)

func DecompressHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response dto.Response

		r.Body = http.MaxBytesReader(w, r.Body, 6<<30)

		err := r.ParseMultipartForm(512 << 20)
		if err != nil {
			response.DefaultInternalError()
			response.Error = append(response.Error, err.Error())
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			response.DefaultInternalError()
			response.Error = append(response.Error, err.Error())
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}
		defer file.Close()

		dst, err := os.Create("./dump/" + header.Filename)
		if err != nil {
			response.DefaultInternalError()
			response.Error = append(response.Error, err.Error())
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			response.DefaultInternalError()
			response.Error = append(response.Error, err.Error())
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}
		dst.Close()
	}
}
