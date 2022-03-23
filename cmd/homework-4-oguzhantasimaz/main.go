package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/cmd/homework-4-oguzhantasimaz/controllers"
	infrastructure "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/infrastructure/repositories"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/infrastructure/repositories/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz/pkg/infrastructure/repositories/book"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	db, err := infrastructure.NewMySQLDB("root:Ot123456@tcp(127.0.0.1:3306)/homework3?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}

	CORSOptions()
	// tüm response'ları sıkıştırmak için
	r.Use(loggingMiddleware)
	r.Use(authenticationMiddleware)

	bookRepo := book.NewBookRepository(db)
	authorRepo := author.NewAuthorRepository(db)
	authorRepo.Migration()
	bookRepo.Migration()
	authorRepo.InsertSampleData()
	bookRepo.InsertSampleData()

	bookCtrl := controllers.NewBookController(bookRepo)
	authorCtrl := controllers.NewAuthorController(authorRepo)

	b := r.PathPrefix("/books").Subrouter()
	// "/products/{name}/"
	b.HandleFunc("/", bookCtrl.GetAllBooks).Methods("GET")
	// "/products/{id:[0-9]+}"
	b.HandleFunc("/{id:[0-9]+}", bookCtrl.GetBookByID).Methods("GET")

	a := r.PathPrefix("/authors").Subrouter()
	a.HandleFunc("/", authorCtrl.CreateAuthor).Methods("POST")
	a.HandleFunc("/{id:[0-9]+}", authorCtrl.GetAuthorByID).Methods("GET")

	// r.HandleFunc("/products/{name}", ProductNameHandler).
	// 	Methods("GET").
	// 	Schemes("http")
	// r.HandleFunc("/products/{id:[0-9]+}", ProductIdHandler)

	srv := &http.Server{
		Addr:         "0.0.0.0:8090",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	ShutdownServer(srv, time.Second*10)
}

func CORSOptions() {
	handlers.AllowedOrigins([]string{"*"})
	handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	handlers.AllowedMethods([]string{"POST", "GET", "PUT", "PATCH", "OPTIONS"})
}

func ShutdownServer(srv *http.Server, timeout time.Duration) {
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.URL.Query())
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		token := r.Header.Get("Authorization")
		if strings.HasPrefix(r.URL.Path, "/books") {
			if token != "" {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Token not found", http.StatusUnauthorized)
			}
		} else {
			next.ServeHTTP(w, r)
		}

	})
}
