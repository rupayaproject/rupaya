package rupx

import (
	"github.com/rupayaproject/rupaya/common"
	"github.com/rupayaproject/rupaya/ethdb"
	"github.com/rupayaproject/rupaya/rupx/rupx_state"
	"github.com/globalsign/mgo"
)

type OrderDao interface {
	// for both leveldb and mongodb
	IsEmptyKey(key []byte) bool
	Close()

	// mongodb methods
	HasObject(hash common.Hash) (bool, error)
	GetObject(hash common.Hash, val interface{}) (interface{}, error)
	PutObject(hash common.Hash, val interface{}) error
	DeleteObject(hash common.Hash) error // won't return error if key not found
	GetOrderByTxHash(txhash common.Hash) []*rupx_state.OrderItem
	GetListOrderByHashes(hashes []string) []*rupx_state.OrderItem
	DeleteTradeByTxHash(txhash common.Hash)
	InitBulk() *mgo.Session
	CommitBulk() error

	// leveldb methods
	Put(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Has(key []byte) (bool, error)
	Delete(key []byte) error
	NewBatch() ethdb.Batch
}
