package application

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	srv    *http.Server
}

// StartApplication sets up the router and middleware
func StartApplication() {
	s := server{
		router: mux.NewRouter(),
	}
	s.routes()

	port := os.Getenv("PORT")

	s.srv = &http.Server{
		Addr:           ":" + port,
		Handler:        s.router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Print("Connected to port " + port)
	err := s.srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func (s *server) routes() {
	// checks server status
	s.router.HandleFunc("/health_check", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("OK")
	}).Methods("GET")

}
