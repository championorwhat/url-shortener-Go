package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"url-shortener/internal/config"
	"url-shortener/internal/handler"
	"url-shortener/internal/middleware"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"

	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"
)

func main() {
	// --------------------------------------------------
	// Open database (relative to project root)
	// --------------------------------------------------
	db, err := sql.Open("sqlite", "urls.db")
	if err != nil {
		log.Fatal("failed to open database:", err)
	}

	// --------------------------------------------------
	// Load and apply schema (NO SILENT FAILURES)
	// --------------------------------------------------
	schemaPath := "migrations/schema.sql"
	log.Println("Loading schema from:", schemaPath)

	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		log.Fatal("FAILED TO READ SCHEMA:", err)
	}

	log.Println("Schema loaded successfully")

	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatal("FAILED TO EXECUTE SCHEMA:", err)
	}

	log.Println("Schema applied successfully")

	// --------------------------------------------------
	// Dependency wiring
	// --------------------------------------------------
	repo := repository.NewURLRepository(db)
	svc := service.NewURLService(repo)
	h := handler.NewURLHandler(svc)

	// --------------------------------------------------
	// Router setup
	// --------------------------------------------------
	r := mux.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CORS)


	// Health check (required for deployment)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// API routes
	r.HandleFunc("/api/shorten", h.Shorten).Methods("POST", "OPTIONS")
	r.HandleFunc("/{code}", h.Redirect).Methods("GET", "OPTIONS")

	// --------------------------------------------------
	// Start server
	// --------------------------------------------------
	port := config.GetPort()
	log.Println("Server running on port", port)

	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("server failed:", err)
	}
}