package pkg

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-kivik/kivik/v4"
)

type IngestionServer struct {
	*http.ServeMux
	Config    AppConfig
	Db        *kivik.DB
	startTime time.Time
}

func NewIngestionServer(config AppConfig, db *kivik.DB) *IngestionServer {
	sv := http.NewServeMux()
	return &IngestionServer{
		ServeMux:  sv,
		Config:    config,
		Db:        db,
		startTime: time.Now(),
	}
}

func (s *IngestionServer) Prepare() {
	s.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		body := map[string]any{
			"status":         "healthy",
			"uptime_minutes": time.Since(s.startTime).Minutes(),
			"app_name":       s.Config.Name,
			"description":    s.Config.Description,
		}

		if err := json.NewEncoder(w).Encode(body); err != nil {
			log.Printf("Write Error: %s", err.Error())
		}
	})

	s.HandleFunc("POST /ingest/json", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Ingesting log")
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		_, _, err := s.Db.CreateDoc(r.Context(), body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	s.HandleFunc("GET /logs", func(w http.ResponseWriter, r *http.Request) {
		resultSet := s.Db.AllDocs(r.Context(), kivik.IncludeDocs())
		var docs []map[string]any
		var doc map[string]any
		for resultSet.Next() {
			if err := resultSet.ScanDoc(&doc); err == nil {
				docs = append(docs, doc)
			}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"docs": docs,
		})
	})
}

func (s *IngestionServer) Listen() {
	log.Printf("Ingestion Server running on %s\n", s.Config.Addr)
	http.ListenAndServe(s.Config.Addr, s)
}
