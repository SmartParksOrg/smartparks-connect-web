package web

import (
	"log"
	"net/http"
)

func StartHttpServer() {

	http.Handle("/api/v1/login", cors(http.HandlerFunc(handleLogin)))
	http.Handle("/api/v1/device/queue", cors(http.HandlerFunc(handleAPI)))
	http.Handle("/api/v1/list", cors(http.HandlerFunc(handleList)))

	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("public/dist/assets"))))
	http.HandleFunc("/", handleRoot)
	log.Println("http server runing in port : 8881")
	log.Fatal(http.ListenAndServe(":8881", nil))
}
func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers to allow all CORS requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

		// If it's an OPTIONS request, return immediately with a 200 status code
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next middleware in the chain
		next.ServeHTTP(w, r)
	})
}