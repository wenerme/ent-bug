package ulids

import (
	"crypto/rand"
	"time"

	ulid "github.com/oklog/ulid/v2"
)

// The default entropy source.
var defaultEntropySource *ulid.MonotonicEntropy

func init() {
	// Seed the default entropy source.
	// TODO: To improve testability, this package should allow control of entropy sources and the time.Now implementation.
	defaultEntropySource = ulid.Monotonic(rand.Reader, 0)
}

func New() (ulid.ULID, error) {
	return ulid.New(ulid.Timestamp(time.Now()), defaultEntropySource)
}
