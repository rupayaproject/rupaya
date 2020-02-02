package rupx

import (
	"context"
	"errors"
	"github.com/rupayaproject/go-rupaya/rupx/rupx_state"
	"math/big"
	"sync"
	"time"

	"github.com/rupayaproject/go-rupaya/common"
)

const (
	LimitThresholdOrderNonceInQueue = 100
)

// List of errors
var (
	ErrNoTopics          = errors.New("missing topic(s)")
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicRupXAPI provides the rupX RPC service that can be
// use publicly without security implications.
type PublicRupXAPI struct {
	t        *RupX
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicRupXAPI create a new RPC rupX service.
func NewPublicRupXAPI(t *RupX) *PublicRupXAPI {
	api := &PublicRupXAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the RupX sub-protocol version.
func (api *PublicRupXAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}

// GetOrderNonce returns the latest orderNonce of the given address
func (api *PublicRupXAPI) GetOrderNonce(address common.Address) (*big.Int, error) {
	//TODO: getOrderNonce from state
	return big.NewInt(0), nil
}

// GetPendingOrders returns pending orders of the given pair
func (api *PublicRupXAPI) GetPendingOrders(pairName string) ([]*rupx_state.OrderItem, error) {
	result := []*rupx_state.OrderItem{}
	//TODO: get pending orders from orderpool
	return result, nil
}
