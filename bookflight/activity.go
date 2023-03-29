package bookflight

import (
	"context"
	"errors"
	"math/rand"

	"github.com/cv65kr/saga-temporal/sdk"
)

func Activity(ctx context.Context, booking sdk.Booking) (string, error) {

	if rand.Intn(2) == 1 {
		return "", errors.New("Random error to see compenstation")
	}

	return "Hello booking" + booking.Id + "!", nil
}
