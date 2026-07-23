package booking

import "context"


type Service struct {
	store BookingStore
}

func NewService(store BookingStore) *Service {
  return &Service{
	 store: store,
  }
}

func (s *Service) Book(booking Booking) (Booking, error) {

   session, err := s.store.Book(booking); 
   
   if err != nil {
      return Booking{}, err
   }
	 
  

   return session, nil
}

func (s *Service) Hold(booking Booking) error {

   

   return nil
}

func (s *Service) ConfirmSeat(ctx context.Context, sessionID string, userID string) (Booking, error) {
	return s.store.Confirm(ctx, sessionID, userID)
}

func (s *Service) ReleaseSeat(ctx context.Context, sessionID string, userID string) error {
	return s.store.Release(ctx, sessionID, userID)
}

func (s *Service) ListBookings(movieID string) []Booking {
   return s.store.ListBookings(movieID)
}