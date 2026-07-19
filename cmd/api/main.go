package main

import (
	"log"
	"net/http"
)


func main() {
   mux := http.NewServeMux()

   mux.HandleFunc("/bookings", func(w http.ResponseWriter, r *http.Request) {})

   if err := http.ListenAndServe(":8080", mux); err != nil {
      log.Fatal(err)
   }

}