package bookflight

import (
	"context"
	"errors"
	"math/rand"

	"go.temporal.io/sdk/activity"
)

func Activity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)

	if rand.Intn(2) == 1 {
		return "", errors.New("Random error to see compenstation")
	}

	return "Hello " + name + "!", nil
}
