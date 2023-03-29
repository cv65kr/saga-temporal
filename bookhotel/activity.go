package bookhotel

import (
	"context"

	"github.com/cv65kr/saga-temporal/sdk"
)

func Activity(ctx context.Context, booking sdk.Booking) (string, error) {
	return "Hello book " + booking.Id + "!", nil
}
