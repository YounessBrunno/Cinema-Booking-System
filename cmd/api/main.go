package main

import (
	"log"
	"net/http"

	"github.com/YounessBrunno/Cinema-Booking-System/internals/booking"
	"github.com/redis/go-redis/v9"
)

func main() {
	mux := http.NewServeMux()
	

	store := booking.NewRedisStore(redis.NewClient(&redis.Options{Addr: "localhost:6379"}))
	svc := booking.NewService(store)
    bookingHandler := booking.NewHandler(svc)

    mux.HandleFunc("GET /movies", bookingHandler.ListMovies)
    mux.Handle("GET /", http.FileServer(http.Dir("static")))
	mux.HandleFunc("GET /movies/{MovieID}/seats", bookingHandler.ListSeats)
	mux.HandleFunc("POST /movies/{MovieID}/seats/{SeatID}/hold", bookingHandler.HoldSeat)
	mux.HandleFunc("POST /sessions/{SessionID}/confirm", bookingHandler.ConfirmSession)
	mux.HandleFunc("DELETE /sessions/{SessionID}", bookingHandler.ReleaseSession)
  


  
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
