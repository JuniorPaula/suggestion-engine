package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"suggestion-engine/engine"
	"time"
)

type Server struct {
	engine  *engine.SuggestionEngine
	history *engine.History
	learner *engine.Learner
}

func NewServer() *Server {
	e := engine.NewSuggestionEngine()
	if err := engine.LoadFromFile("../../data/searches.txt", e); err != nil {
		log.Fatalf("[ERROR] could not load dataset: %v\n", err)
	}

	h := engine.NewHistory("../../data/search_log.txt")
	h.Load()

	l := engine.NewLearner(e, "../../data/searches.txt")

	// save the word freq each 60 seconds
	go func() {
		for {
			time.Sleep(60 * time.Second)
			if err := l.Save(); err != nil {
				log.Println(err)
			}
		}
	}()

	return &Server{
		engine:  e,
		history: h,
		learner: l,
	}
}

func (s *Server) HandleSuggest(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing query parameter: q", http.StatusBadRequest)
		return
	}

	start := time.Now()
	suggestions := s.engine.Suggest(query, 5)
	elapsed := time.Since(start)

	// Learning and history
	s.history.Add(query)
	s.learner.Learn(query)

	resp := map[string]interface{}{
		"query":             query,
		"suggestions":       suggestions,
		"took_ms":           elapsed.Milliseconds(),
		"suggestions_count": len(suggestions),
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ok")
}
