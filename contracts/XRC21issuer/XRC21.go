package XRC21issuer

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/XRC21issuer/contract"
	"math/big"
)

type MyXRC21 struct {
	*contract.MyXRC21Session
	contractBackend bind.ContractBackend
}

func NewXRC21(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*MyXRC21, error) {
	smartContract, err := contract.NewMyXRC21(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &MyXRC21{
		&contract.MyXRC21Session{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployXRC21(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, name string, symbol string, decimals uint8, cap, fee *big.Int) (common.Address, *MyXRC21, error) {
	contractAddr, _, _, err := contract.DeployMyXRC21(transactOpts, contractBackend, name, symbol, decimals, cap, fee)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewXRC21(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}
