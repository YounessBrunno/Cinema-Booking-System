package booking


import "errors"

var (
	ErrBookingAlreadyExists = errors.New("booking already exists")
)


type Booking struct {
	ID string
	MovieID string
	SeatID string
	UserID string
	Status string
}

type BookingStore interface {
	Book(b Booking) error
	ListBooking(movieID string)  []Booking
}