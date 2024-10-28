package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	//
	"github.com/ahmad1702/v2/seed"
	//
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func sendErrorRes(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json := `{"error": "` + msg + `"}`
	w.Write([]byte(json))
}

func main() {
	zips := seed.GetZipInfos()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	// example: http://localhost:4321/zip-info/12345
	r.Get("/zip-info/{zip}", func(w http.ResponseWriter, r *http.Request) {
		zip_str := chi.URLParam(r, "zip")
		zip_code, err := strconv.Atoi(zip_str)
		if err != nil {
			fmt.Println("Error parsing zip:", err)
			sendErrorRes(w, 400, "Error parsing zip")
			return
		}

		// if zip_code is a key of zips, then return the value
		if val, ok := zips[zip_code]; ok {
			zip_info_json, err := json.Marshal(val)
			if err != nil {
				fmt.Println("Error marshalling zip info:", err)
				sendErrorRes(w, 500, "Error marshalling zip info")
				return
			}
			w.Write(zip_info_json)
			return
		} else {
			sendErrorRes(w, 404, "Zip not found")
			return
		}
	})
	fmt.Println("Server listening at http://localhost:4321")

	err := http.ListenAndServe(":4321", r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
