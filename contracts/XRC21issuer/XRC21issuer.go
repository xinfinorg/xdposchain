package XRC21issuer

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/XRC21issuer/contract"
	"math/big"
)

type XRC21Issuer struct {
	*contract.XRC21IssuerSession
	contractBackend bind.ContractBackend
}

func NewXRC21Issuer(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*XRC21Issuer, error) {
	contractObject, err := contract.NewXRC21Issuer(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &XRC21Issuer{
		&contract.XRC21IssuerSession{
			Contract:     contractObject,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployXRC21Issuer(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, minApply *big.Int) (common.Address, *XRC21Issuer, error) {
	contractAddr, _, _, err := contract.DeployXRC21Issuer(transactOpts, contractBackend, minApply)
	if err != nil {
		return contractAddr, nil, err
	}
	contractObject, err := NewXRC21Issuer(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, contractObject, nil
}
