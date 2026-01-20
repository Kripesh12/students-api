package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func teacherHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Teachers Routes"))
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Students Routes"))
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Execs Routes"))
}

func main() {
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Root Routes"))
	})

	mux.HandleFunc("/teachers/", teacherHandler)

	mux.HandleFunc("/students/", studentHandler)

	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      port,
		Handler:   mux,
		TLSConfig: tlsConfig,
	}
	// if err := http.ListenAndServe(port, nil); err != nil {
	// 	log.Fatal("Error while starting the server: ", err)
	// }

	fmt.Println("Server is running on port: ", port)
	if err := server.ListenAndServeTLS(cert, key); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
