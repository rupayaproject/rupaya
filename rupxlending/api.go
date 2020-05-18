package rupxlending

import (
	"context"
	"errors"
	"sync"
	"time"
)


// List of errors
var (
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicRupXLendingAPI provides the rupX RPC service that can be
// use publicly without security implications.
type PublicRupXLendingAPI struct {
	t        *Lending
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicRupXLendingAPI create a new RPC rupX service.
func NewPublicRupXLendingAPI(t *Lending) *PublicRupXLendingAPI {
	api := &PublicRupXLendingAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the Lending sub-protocol version.
func (api *PublicRupXLendingAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}
