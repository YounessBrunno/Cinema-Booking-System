package booking

import (
	    "net/http"
		"encoding/json"
	   )

 
type Handler struct {
  svc *Service
}

func NewHandler(svc *Service) *Handler {
  return &Handler{svc: svc}
}


func (h *Handler) ListSeats(w http.ResponseWriter, r *http.Request)  {
   
   h.svc.store.ListBookings("")
}

type movieResponse struct {
   ID          int    `json:"id"`
   Title       string `json:"title"`
   Year        int    `json:"year"`
   Rows        int    `json:"rows"`
   SeatsPerRow int    `json:"seats_per_row"`
}

var movies = []movieResponse{
	{ID: 1, Title: "Inception", Year: 2010, Rows: 10, SeatsPerRow: 15},
	{ID: 2, Title: "The Matrix", Year: 1999, Rows: 8, SeatsPerRow: 12},
	{ID: 3, Title: "Interstellar", Year: 2014, Rows: 12, SeatsPerRow: 16},
}

func (h *Handler) ListMovies(w http.ResponseWriter, r *http.Request) {

	WriteJSON(w, http.StatusOK, movies)
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {

   w.Header().Set("Content-Type", "application/json")
   w.WriteHeader(status)
   
   return json.NewEncoder(w).Encode(data)
}
