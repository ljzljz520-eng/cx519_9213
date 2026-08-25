package api

import (
	"encoding/json"
	"net/http"
	"soundspace/internal/flow061"
	"soundspace/internal/model"
	"soundspace/internal/query"
	"soundspace/internal/store"
)

type Server struct {
	flow  *flow061.Service
	query *query.Engine
}

func New(s *store.Store) *Server { return &Server{flow: flow061.New(s), query: query.New(s)} }
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/records", s.records)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNoContent) })
	return mux
}
func (s *Server) records(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p, e := s.query.Find(model.SearchFilter{Text: r.URL.Query().Get("q"), Status: r.URL.Query().Get("status")})
		if e != nil {
			http.Error(w, e.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(p)
		return
	}
	if r.Method == http.MethodPost {
		var rec model.Record
		if e := json.NewDecoder(r.Body).Decode(&rec); e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		if e := s.flow.Register(rec, "api"); e != nil {
			http.Error(w, e.Error(), 422)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(rec)
		return
	}
	http.Error(w, "method not allowed", 405)
}
