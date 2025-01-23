package xdc21issuer

import (
	"github.com/XinFinOrg/XDPoSChain/accounts/abi/bind"
	"github.com/XinFinOrg/XDPoSChain/common"
	"github.com/XinFinOrg/XDPoSChain/contracts/xdc21issuer/contract"
	"math/big"
)

type XDC21Issuer struct {
	*contract.XDC21IssuerSession
	contractBackend bind.ContractBackend
}

func NewXDC21Issuer(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*XDC21Issuer, error) {
	contractObject, err := contract.NewXDC21Issuer(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &XDC21Issuer{
		&contract.XDC21IssuerSession{
			Contract:     contractObject,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployXDC21Issuer(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, minApply *big.Int) (common.Address, *XDC21Issuer, error) {
	contractAddr, _, _, err := contract.DeployXDC21Issuer(transactOpts, contractBackend, minApply)
	if err != nil {
		return contractAddr, nil, err
	}
	contractObject, err := NewXDC21Issuer(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, contractObject, nil
}
