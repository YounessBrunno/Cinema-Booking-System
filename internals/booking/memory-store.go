package booking 


type MemoryStore struct {
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {
  return &MemoryStore{
	 bookings: map[string]Booking{},
  }
}

func (s *MemoryStore) Book(b Booking) error {
  
}

func (s *MemoryStore) ListBooking(movieID string)  []Booking {

}