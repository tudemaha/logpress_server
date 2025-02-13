package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/tudemaha/logpress_server/global/dto"
	"github.com/tudemaha/logpress_server/internal/decompress/service"
)

func DecompressHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var response dto.Response

		if r.Method != "POST" {
			response.DefaultNotAllowed()
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		startTime := time.Now()

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

		filename := strings.Split(header.Filename, ".")
		if filename[len(filename)-1] != "gz" && filename[len(filename)-1] != "sql" {
			response.DefaultBadRequest()
			response.Error = append(response.Error, "uploaded file extension must .gz or .sql")
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		var fullpath string
		if filename[len(filename)-1] == "gz" {
			fullpath = "./dump/compressed/" + header.Filename
		} else {
			fullpath = "./dump/uncompressed/" + header.Filename
		}

		dst, err := os.Create(fullpath)
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

		transferTime := time.Now()

		var decompressTime time.Time
		if filename[len(filename)-1] == "gz" {
			err := service.DecompressGZIP(filename[0])
			if err != nil {
				response.DefaultInternalError()
				response.Error = append(response.Error, err.Error())
				w.WriteHeader(response.Code)
				json.NewEncoder(w).Encode(response)
				return
			}
			decompressTime = time.Now()
		}

		err = service.MergeDump(filename[0])
		if err != nil {
			response.DefaultInternalError()
			response.Error = append(response.Error, err.Error())
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		mergeTime := time.Now()

		response.DefaultOK()
		response.Data = map[string]string{
			"start_time":      startTime.String(),
			"transfer_time":   transferTime.String(),
			"decompress_time": decompressTime.String(),
			"merge_time":      mergeTime.String(),
		}
		json.NewEncoder(w).Encode(response)
	}
}
