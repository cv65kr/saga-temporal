package bookcar

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
)

func PollStatusActivity(ctx context.Context) (string, error) {
	if rand.Intn(2) == 1 {
		return "", errors.New("Random error to see compenstation")
	}

	return fmt.Sprintf("test value %d", rand.Intn(30)), nil
}
