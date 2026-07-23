package booking

import (
       "net/http"
       "time"
       "encoding/json"
       jsonx "github.com/YounessBrunno/Cinema-Booking-System/internals/json"
	   )

 
type Handler struct {
  svc *Service
}

func NewHandler(svc *Service) *Handler {
  return &Handler{svc: svc}
}


func (h *Handler) ListSeats(w http.ResponseWriter, r *http.Request)  {
   MovieId := r.PathValue("MovieID")

   bookings := h.svc.store.ListBookings(MovieId)

   seats := make([]seatInfo, 0, len(bookings))
   for _, b := range bookings {
      seats = append(seats, seatInfo{
         SeatID: b.SeatID,
         UserID: b.UserID,
         Booked: true,
      })
   }

   jsonx.WriteJSON(w, http.StatusOK, seats)
}

func (h *Handler) HoldSeat(w http.ResponseWriter, r *http.Request)  {
   movieId := r.PathValue("MovieID")
   seatId := r.PathValue("SeatID")

   type holdRequest struct {
      UserID string `json:"user_id"`
   }

   var req holdRequest

   if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
      jsonx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
      return
   }

   session, err := h.svc.Book(Booking{
      MovieID: movieId,
      SeatID:  seatId,
      UserID:  req.UserID,
   })

   if err != nil {
      jsonx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to book seat"})
      return
   }

   type holdResponse struct {
      SessionID string `json:"session_id"`
      MovieID string `json:"movie_id"`
      SeatID string `json:"seat_id"`
      ExpiresAt string `json:"expires_at"`
   }

   jsonx.WriteJSON(w, http.StatusOK, holdResponse{
      SeatID: seatId,
      MovieID: session.MovieID,
      SessionID: session.ID,
      ExpiresAt: session.ExpiresAt.Format(time.RFC3339),
   })

}

func (h *Handler) ConfirmSeat(w http.ResponseWriter, r *http.Request)  {
   sessionId := r.PathValue("SessionID")  

   
   h.svc.Confirm(sessionId)
}

func (h *Handler) ReleaseSeat(w http.ResponseWriter, r *http.Request)  {
   sessionId := r.PathValue("SessionID")

   h.svc.Release(sessionId)
}


type seatInfo struct {
   SeatID string `json:"seat_id"`
   UserID string `json:"user_id"`
   Booked bool   `json:"booked"`
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

	jsonx.WriteJSON(w, http.StatusOK, movies)
}

