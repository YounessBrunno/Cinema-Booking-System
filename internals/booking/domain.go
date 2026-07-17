package booking


import ( "errors"
       "time"
      )
var (
	ErrBookingAlreadyExists = errors.New("booking already exists")
)


type Booking struct {
	ID string
	MovieID string
	SeatID string
	UserID string
	Status string
	ExpiresAt time.Time
}


type BookingStore interface {
	Book(b Booking) error
	ListBookings(movieID string) []Booking
}