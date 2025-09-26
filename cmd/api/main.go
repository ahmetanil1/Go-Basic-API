package main

import (
	"context"
	"fmt"
	"library-management/internal/db"
	"library-management/internal/handlers"
	"library-management/internal/repos"
	"library-management/internal/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// CROSS-CUTTING CONCERN --> basics of middleware
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	// Connect to mongo
	db.Connect()
	// Repository ve service
	// Repository ve service
	bookRepo := repos.NewBookRepository()
	bookService := services.NewBookService(bookRepo)

	// Handler örneği
	bookHandler := &handlers.BookHandler{
		Service: bookService,
	}

	mux := http.NewServeMux()

	// Tüm kitaplar
	mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			bookHandler.GetAllBooksHandler(w, r)
		} else if r.Method == http.MethodPost {
			bookHandler.CreateBookHandler(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})

	// ID ile kitap
	mux.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/books/"):]
		if id == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			bookHandler.GetByIDBookHandler(w, r, id)
		case http.MethodPut:
			bookHandler.UpdateBookHandler(w, r, id)
		case http.MethodDelete:
			bookHandler.DeleteBookHandler(w, r, id)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})

	// setup server
	server := http.Server{
		Addr:    ":8080",
		Handler: logMiddleware(mux),
	}

	// Start a different goroutine for the server other than the main goroutine
	go func() {
		fmt.Println("Server listening on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")

	// arrange the timeout before shutting down goroutines
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// shutdown after time that arranged up
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	fmt.Println("Server cleanly stopped")
}
