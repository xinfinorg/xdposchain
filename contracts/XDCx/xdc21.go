package XDCx

import (
	"github.com/XinFinOrg/XDPoSChain/accounts/abi/bind"
	"github.com/XinFinOrg/XDPoSChain/common"
	"github.com/XinFinOrg/XDPoSChain/contracts/XDCx/contract"
	"math/big"
)

type MyXDC21 struct {
	*contract.MyXDC21Session
	contractBackend bind.ContractBackend
}

func NewXDC21(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*MyXDC21, error) {
	smartContract, err := contract.NewMyXDC21(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &MyXDC21{
		&contract.MyXDC21Session{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployXDC21(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, owners []common.Address, required *big.Int, name string, symbol string, decimals uint8, cap, fee, depositFee, withdrawFee *big.Int) (common.Address, *MyXDC21, error) {
	contractAddr, _, _, err := contract.DeployMyXDC21(transactOpts, contractBackend, owners, required, name, symbol, decimals, cap, fee, depositFee, withdrawFee)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewXDC21(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}
