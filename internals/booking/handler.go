package booking

import (
       "net/http"
       "github.com/YounessBrunno/Cinema-Booking-System/internals/json"
	   )

 
type Handler struct {
  svc *Service
}

func NewHandler(svc *Service) *Handler {
  return &Handler{svc: svc}
}


func (h *Handler) ListSeats(w http.ResponseWriter, r *http.Request)  {
   MovieId := r.PathValue("MovieId")

   bookings := h.svc.store.ListBookings(MovieId)

   json.WriteJSON(w, http.StatusOK, bookings)
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

	json.WriteJSON(w, http.StatusOK, movies)
}

