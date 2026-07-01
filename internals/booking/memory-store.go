package booking

import (
	"errors"
)

type MemoryStore struct {
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {
  return &MemoryStore{
	 bookings: map[string]Booking{},
  }
}

func (s *MemoryStore) Book(b Booking) error {
  if _, exists := s.bookings[b.SeatID]; exists {

    return errors.New("booking already exists")
  }

  s.bookings[b.ID] = b
  
  return nil
}

func (s *MemoryStore) ListBooking(movieID string)  []Booking {
  var bookings []Booking

  for _, booking := range s.bookings {

    if booking.MovieID == movieID {
      bookings = append(bookings, booking)
    }
  }

  return bookings
}