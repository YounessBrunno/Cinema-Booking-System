package booking

import (
	"sync"
)


type ConcurrentStore struct {
	bookings map[string]Booking
	sync.RWMutex
}

func NewConcurrentStore() *ConcurrentStore {
  return &ConcurrentStore{
	 bookings: map[string]Booking{},
  }
}

func (s *ConcurrentStore) Book(b Booking) error {

  s.Lock()
  defer s.Unlock()

  if _, exists := s.bookings[b.SeatID]; exists {

    return ErrBookingAlreadyExists
  }

  s.bookings[b.SeatID] = b
  
  return nil
}

func (s *ConcurrentStore) ListBooking(movieID string)  []Booking {
  s.RLock()
  defer s.RUnlock()

  var bookings []Booking

  for _, booking := range s.bookings {

    if booking.MovieID == movieID {
      bookings = append(bookings, booking)
    }
  }

  return bookings
}