package booking



type Service struct {
	store BookingStore
}

func NewService(store BookingStore) *Service {
  return &Service{
	 store: store,
  }
}

func (s *Service) Book(booking Booking) error {

   if err := s.store.Book(booking); err != nil {
	 return err
   }

   return nil
}


func (s *Service) ListBookings(movieID string) []Booking {
   return s.store.ListBookings(movieID)
}