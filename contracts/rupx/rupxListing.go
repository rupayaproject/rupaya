package rupx

import (
	"github.com/rupayaproject/rupaya/accounts/abi/bind"
	"github.com/rupayaproject/rupaya/common"
	"github.com/rupayaproject/rupaya/contracts/rupx/contract"
)

type RUPXListing struct {
	*contract.RUPXListingSession
	contractBackend bind.ContractBackend
}

func NewMyRUPXListing(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*RUPXListing, error) {
	smartContract, err := contract.NewRUPXListing(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &RUPXListing{
		&contract.RUPXListingSession{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployRUPXListing(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *RUPXListing, error) {
	contractAddr, _, _, err := contract.DeployRUPXListing(transactOpts, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewMyRUPXListing(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}
