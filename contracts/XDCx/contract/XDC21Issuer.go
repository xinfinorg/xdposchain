// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"github.com/XinFinOrg/XDPoSChain"
	"github.com/XinFinOrg/XDPoSChain/accounts/abi"
	"github.com/XinFinOrg/XDPoSChain/accounts/abi/bind"
	"github.com/XinFinOrg/XDPoSChain/common"
	"github.com/XinFinOrg/XDPoSChain/core/types"
	"github.com/XinFinOrg/XDPoSChain/event"
	"math/big"
	"strings"
)

// AbstractTokenXDC21ABI is the input ABI used to generate the binding from.
const AbstractTokenXDC21ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AbstractTokenXDC21Bin is the compiled bytecode used for deploying new contracts.
const AbstractTokenXDC21Bin = `0x`

// DeployAbstractTokenXDC21 deploys a new Ethereum contract, binding an instance of AbstractTokenXDC21 to it.
func DeployAbstractTokenXDC21(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AbstractTokenXDC21, error) {
	parsed, err := abi.JSON(strings.NewReader(AbstractTokenXDC21ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbstractTokenXDC21Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AbstractTokenXDC21{AbstractTokenXDC21Caller: AbstractTokenXDC21Caller{contract: contract}, AbstractTokenXDC21Transactor: AbstractTokenXDC21Transactor{contract: contract}, AbstractTokenXDC21Filterer: AbstractTokenXDC21Filterer{contract: contract}}, nil
}

// AbstractTokenXDC21 is an auto generated Go binding around an Ethereum contract.
type AbstractTokenXDC21 struct {
	AbstractTokenXDC21Caller     // Read-only binding to the contract
	AbstractTokenXDC21Transactor // Write-only binding to the contract
	AbstractTokenXDC21Filterer   // Log filterer for contract events
}

// AbstractTokenXDC21Caller is an auto generated read-only Go binding around an Ethereum contract.
type AbstractTokenXDC21Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractTokenXDC21Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AbstractTokenXDC21Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractTokenXDC21Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbstractTokenXDC21Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractTokenXDC21Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbstractTokenXDC21Session struct {
	Contract     *AbstractTokenXDC21 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AbstractTokenXDC21CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbstractTokenXDC21CallerSession struct {
	Contract *AbstractTokenXDC21Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AbstractTokenXDC21TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbstractTokenXDC21TransactorSession struct {
	Contract     *AbstractTokenXDC21Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AbstractTokenXDC21Raw is an auto generated low-level Go binding around an Ethereum contract.
type AbstractTokenXDC21Raw struct {
	Contract *AbstractTokenXDC21 // Generic contract binding to access the raw methods on
}

// AbstractTokenXDC21CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbstractTokenXDC21CallerRaw struct {
	Contract *AbstractTokenXDC21Caller // Generic read-only contract binding to access the raw methods on
}

// AbstractTokenXDC21TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbstractTokenXDC21TransactorRaw struct {
	Contract *AbstractTokenXDC21Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAbstractTokenXDC21 creates a new instance of AbstractTokenXDC21, bound to a specific deployed contract.
func NewAbstractTokenXDC21(address common.Address, backend bind.ContractBackend) (*AbstractTokenXDC21, error) {
	contract, err := bindAbstractTokenXDC21(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbstractTokenXDC21{AbstractTokenXDC21Caller: AbstractTokenXDC21Caller{contract: contract}, AbstractTokenXDC21Transactor: AbstractTokenXDC21Transactor{contract: contract}, AbstractTokenXDC21Filterer: AbstractTokenXDC21Filterer{contract: contract}}, nil
}

// NewAbstractTokenXDC21Caller creates a new read-only instance of AbstractTokenXDC21, bound to a specific deployed contract.
func NewAbstractTokenXDC21Caller(address common.Address, caller bind.ContractCaller) (*AbstractTokenXDC21Caller, error) {
	contract, err := bindAbstractTokenXDC21(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractTokenXDC21Caller{contract: contract}, nil
}

// NewAbstractTokenXDC21Transactor creates a new write-only instance of AbstractTokenXDC21, bound to a specific deployed contract.
func NewAbstractTokenXDC21Transactor(address common.Address, transactor bind.ContractTransactor) (*AbstractTokenXDC21Transactor, error) {
	contract, err := bindAbstractTokenXDC21(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractTokenXDC21Transactor{contract: contract}, nil
}

// NewAbstractTokenXDC21Filterer creates a new log filterer instance of AbstractTokenXDC21, bound to a specific deployed contract.
func NewAbstractTokenXDC21Filterer(address common.Address, filterer bind.ContractFilterer) (*AbstractTokenXDC21Filterer, error) {
	contract, err := bindAbstractTokenXDC21(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbstractTokenXDC21Filterer{contract: contract}, nil
}

// bindAbstractTokenXDC21 binds a generic wrapper to an already deployed contract.
func bindAbstractTokenXDC21(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbstractTokenXDC21ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractTokenXDC21 *AbstractTokenXDC21Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AbstractTokenXDC21.Contract.AbstractTokenXDC21Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractTokenXDC21 *AbstractTokenXDC21Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractTokenXDC21.Contract.AbstractTokenXDC21Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractTokenXDC21 *AbstractTokenXDC21Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractTokenXDC21.Contract.AbstractTokenXDC21Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractTokenXDC21 *AbstractTokenXDC21CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AbstractTokenXDC21.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractTokenXDC21 *AbstractTokenXDC21TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractTokenXDC21.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractTokenXDC21 *AbstractTokenXDC21TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractTokenXDC21.Contract.contract.Transact(opts, method, params...)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_AbstractTokenXDC21 *AbstractTokenXDC21Caller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AbstractTokenXDC21.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_AbstractTokenXDC21 *AbstractTokenXDC21Session) Issuer() (common.Address, error) {
	return _AbstractTokenXDC21.Contract.Issuer(&_AbstractTokenXDC21.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_AbstractTokenXDC21 *AbstractTokenXDC21CallerSession) Issuer() (common.Address, error) {
	return _AbstractTokenXDC21.Contract.Issuer(&_AbstractTokenXDC21.CallOpts)
}

// XDC21IssuerABI is the input ABI used to generate the binding from.
const XDC21IssuerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"minCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenCapacity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"apply\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"charge\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Apply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"supporter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Charge\",\"type\":\"event\"}]"

// XDC21IssuerBin is the compiled bytecode used for deploying new contracts.
const XDC21IssuerBin = `0x608060405234801561001057600080fd5b506040516020806104578339810160405251600055610423806100346000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633fa615b081146100715780638f3a981c146100985780639d63848a146100b9578063c6b32f341461011e578063fc6bd76a14610134575b600080fd5b34801561007d57600080fd5b50610086610148565b60408051918252519081900360200190f35b3480156100a457600080fd5b50610086600160a060020a036004351661014e565b3480156100c557600080fd5b506100ce610169565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561010a5781810151838201526020016100f2565b505050509050019250505060405180910390f35b610132600160a060020a03600435166101cb565b005b610132600160a060020a036004351661035d565b60005490565b600160a060020a031660009081526002602052604090205490565b606060018054806020026020016040519081016040528092919081815260200182805480156101c157602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116101a3575b5050505050905090565b600081600160a060020a03811615156101e357600080fd5b6000543410156101f257600080fd5b82915033600160a060020a031682600160a060020a0316631d1438486040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561025657600080fd5b505af115801561026a573d6000803e3d6000fd5b505050506040513d602081101561028057600080fd5b5051600160a060020a03161461029557600080fd5b600180548082019091557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03851690811790915560009081526002602052604090205461030390346103de565b600160a060020a0384166000818152600260209081526040918290209390935580513481529051919233927f2d485624158277d5113a56388c3abf5c20e3511dd112123ba376d16b21e4d7169281900390910190a3505050565b600160a060020a038116600090815260026020526040902054610386903463ffffffff6103de16565b600160a060020a0382166000818152600260209081526040918290209390935580513481529051919233927f5cffac866325fd9b2a8ea8df2f110a0058313b279402d15ae28dd324a2282e069281900390910190a350565b6000828201838110156103f057600080fd5b93925050505600a165627a7a7230582005dc9504c7a156980fbaadfe03ffb20a475e65b947f9a8ef3e6d6beee52325a80029`

// DeployXDC21Issuer deploys a new Ethereum contract, binding an instance of XDC21Issuer to it.
func DeployXDC21Issuer(auth *bind.TransactOpts, backend bind.ContractBackend, value *big.Int) (common.Address, *types.Transaction, *XDC21Issuer, error) {
	parsed, err := abi.JSON(strings.NewReader(XDC21IssuerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(XDC21IssuerBin), backend, value)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XDC21Issuer{XDC21IssuerCaller: XDC21IssuerCaller{contract: contract}, XDC21IssuerTransactor: XDC21IssuerTransactor{contract: contract}, XDC21IssuerFilterer: XDC21IssuerFilterer{contract: contract}}, nil
}

// XDC21Issuer is an auto generated Go binding around an Ethereum contract.
type XDC21Issuer struct {
	XDC21IssuerCaller     // Read-only binding to the contract
	XDC21IssuerTransactor // Write-only binding to the contract
	XDC21IssuerFilterer   // Log filterer for contract events
}

// XDC21IssuerCaller is an auto generated read-only Go binding around an Ethereum contract.
type XDC21IssuerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XDC21IssuerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XDC21IssuerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XDC21IssuerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XDC21IssuerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XDC21IssuerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XDC21IssuerSession struct {
	Contract     *XDC21Issuer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XDC21IssuerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XDC21IssuerCallerSession struct {
	Contract *XDC21IssuerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// XDC21IssuerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XDC21IssuerTransactorSession struct {
	Contract     *XDC21IssuerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// XDC21IssuerRaw is an auto generated low-level Go binding around an Ethereum contract.
type XDC21IssuerRaw struct {
	Contract *XDC21Issuer // Generic contract binding to access the raw methods on
}

// XDC21IssuerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XDC21IssuerCallerRaw struct {
	Contract *XDC21IssuerCaller // Generic read-only contract binding to access the raw methods on
}

// XDC21IssuerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XDC21IssuerTransactorRaw struct {
	Contract *XDC21IssuerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXDC21Issuer creates a new instance of XDC21Issuer, bound to a specific deployed contract.
func NewXDC21Issuer(address common.Address, backend bind.ContractBackend) (*XDC21Issuer, error) {
	contract, err := bindXDC21Issuer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XDC21Issuer{XDC21IssuerCaller: XDC21IssuerCaller{contract: contract}, XDC21IssuerTransactor: XDC21IssuerTransactor{contract: contract}, XDC21IssuerFilterer: XDC21IssuerFilterer{contract: contract}}, nil
}

// NewXDC21IssuerCaller creates a new read-only instance of XDC21Issuer, bound to a specific deployed contract.
func NewXDC21IssuerCaller(address common.Address, caller bind.ContractCaller) (*XDC21IssuerCaller, error) {
	contract, err := bindXDC21Issuer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XDC21IssuerCaller{contract: contract}, nil
}

// NewXDC21IssuerTransactor creates a new write-only instance of XDC21Issuer, bound to a specific deployed contract.
func NewXDC21IssuerTransactor(address common.Address, transactor bind.ContractTransactor) (*XDC21IssuerTransactor, error) {
	contract, err := bindXDC21Issuer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XDC21IssuerTransactor{contract: contract}, nil
}

// NewXDC21IssuerFilterer creates a new log filterer instance of XDC21Issuer, bound to a specific deployed contract.
func NewXDC21IssuerFilterer(address common.Address, filterer bind.ContractFilterer) (*XDC21IssuerFilterer, error) {
	contract, err := bindXDC21Issuer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XDC21IssuerFilterer{contract: contract}, nil
}

// bindXDC21Issuer binds a generic wrapper to an already deployed contract.
func bindXDC21Issuer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XDC21IssuerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XDC21Issuer *XDC21IssuerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XDC21Issuer.Contract.XDC21IssuerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XDC21Issuer *XDC21IssuerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XDC21Issuer.Contract.XDC21IssuerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XDC21Issuer *XDC21IssuerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XDC21Issuer.Contract.XDC21IssuerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XDC21Issuer *XDC21IssuerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XDC21Issuer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XDC21Issuer *XDC21IssuerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XDC21Issuer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XDC21Issuer *XDC21IssuerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XDC21Issuer.Contract.contract.Transact(opts, method, params...)
}

// GetTokenCapacity is a free data retrieval call binding the contract method 0x8f3a981c.
//
// Solidity: function getTokenCapacity(token address) constant returns(uint256)
func (_XDC21Issuer *XDC21IssuerCaller) GetTokenCapacity(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XDC21Issuer.contract.Call(opts, out, "getTokenCapacity", token)
	return *ret0, err
}

// GetTokenCapacity is a free data retrieval call binding the contract method 0x8f3a981c.
//
// Solidity: function getTokenCapacity(token address) constant returns(uint256)
func (_XDC21Issuer *XDC21IssuerSession) GetTokenCapacity(token common.Address) (*big.Int, error) {
	return _XDC21Issuer.Contract.GetTokenCapacity(&_XDC21Issuer.CallOpts, token)
}

// GetTokenCapacity is a free data retrieval call binding the contract method 0x8f3a981c.
//
// Solidity: function getTokenCapacity(token address) constant returns(uint256)
func (_XDC21Issuer *XDC21IssuerCallerSession) GetTokenCapacity(token common.Address) (*big.Int, error) {
	return _XDC21Issuer.Contract.GetTokenCapacity(&_XDC21Issuer.CallOpts, token)
}

// MinCap is a free data retrieval call binding the contract method 0x3fa615b0.
//
// Solidity: function minCap() constant returns(uint256)
func (_XDC21Issuer *XDC21IssuerCaller) MinCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XDC21Issuer.contract.Call(opts, out, "minCap")
	return *ret0, err
}

// MinCap is a free data retrieval call binding the contract method 0x3fa615b0.
//
// Solidity: function minCap() constant returns(uint256)
func (_XDC21Issuer *XDC21IssuerSession) MinCap() (*big.Int, error) {
	return _XDC21Issuer.Contract.MinCap(&_XDC21Issuer.CallOpts)
}

// MinCap is a free data retrieval call binding the contract method 0x3fa615b0.
//
// Solidity: function minCap() constant returns(uint256)
func (_XDC21Issuer *XDC21IssuerCallerSession) MinCap() (*big.Int, error) {
	return _XDC21Issuer.Contract.MinCap(&_XDC21Issuer.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_XDC21Issuer *XDC21IssuerCaller) Tokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _XDC21Issuer.contract.Call(opts, out, "tokens")
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_XDC21Issuer *XDC21IssuerSession) Tokens() ([]common.Address, error) {
	return _XDC21Issuer.Contract.Tokens(&_XDC21Issuer.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_XDC21Issuer *XDC21IssuerCallerSession) Tokens() ([]common.Address, error) {
	return _XDC21Issuer.Contract.Tokens(&_XDC21Issuer.CallOpts)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_XDC21Issuer *XDC21IssuerTransactor) Apply(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _XDC21Issuer.contract.Transact(opts, "apply", token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_XDC21Issuer *XDC21IssuerSession) Apply(token common.Address) (*types.Transaction, error) {
	return _XDC21Issuer.Contract.Apply(&_XDC21Issuer.TransactOpts, token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_XDC21Issuer *XDC21IssuerTransactorSession) Apply(token common.Address) (*types.Transaction, error) {
	return _XDC21Issuer.Contract.Apply(&_XDC21Issuer.TransactOpts, token)
}

// Charge is a paid mutator transaction binding the contract method 0xfc6bd76a.
//
// Solidity: function charge(token address) returns()
func (_XDC21Issuer *XDC21IssuerTransactor) Charge(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _XDC21Issuer.contract.Transact(opts, "charge", token)
}

// Charge is a paid mutator transaction binding the contract method 0xfc6bd76a.
//
// Solidity: function charge(token address) returns()
func (_XDC21Issuer *XDC21IssuerSession) Charge(token common.Address) (*types.Transaction, error) {
	return _XDC21Issuer.Contract.Charge(&_XDC21Issuer.TransactOpts, token)
}

// Charge is a paid mutator transaction binding the contract method 0xfc6bd76a.
//
// Solidity: function charge(token address) returns()
func (_XDC21Issuer *XDC21IssuerTransactorSession) Charge(token common.Address) (*types.Transaction, error) {
	return _XDC21Issuer.Contract.Charge(&_XDC21Issuer.TransactOpts, token)
}

// XDC21IssuerApplyIterator is returned from FilterApply and is used to iterate over the raw logs and unpacked data for Apply events raised by the XDC21Issuer contract.
type XDC21IssuerApplyIterator struct {
	Event *XDC21IssuerApply // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  XDPoSChain.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XDC21IssuerApplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XDC21IssuerApply)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(XDC21IssuerApply)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XDC21IssuerApplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XDC21IssuerApplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XDC21IssuerApply represents a Apply event raised by the XDC21Issuer contract.
type XDC21IssuerApply struct {
	Issuer common.Address
	Token  common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApply is a free log retrieval operation binding the contract event 0x2d485624158277d5113a56388c3abf5c20e3511dd112123ba376d16b21e4d716.
//
// Solidity: event Apply(issuer indexed address, token indexed address, value uint256)
func (_XDC21Issuer *XDC21IssuerFilterer) FilterApply(opts *bind.FilterOpts, issuer []common.Address, token []common.Address) (*XDC21IssuerApplyIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _XDC21Issuer.contract.FilterLogs(opts, "Apply", issuerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &XDC21IssuerApplyIterator{contract: _XDC21Issuer.contract, event: "Apply", logs: logs, sub: sub}, nil
}

// WatchApply is a free log subscription operation binding the contract event 0x2d485624158277d5113a56388c3abf5c20e3511dd112123ba376d16b21e4d716.
//
// Solidity: event Apply(issuer indexed address, token indexed address, value uint256)
func (_XDC21Issuer *XDC21IssuerFilterer) WatchApply(opts *bind.WatchOpts, sink chan<- *XDC21IssuerApply, issuer []common.Address, token []common.Address) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _XDC21Issuer.contract.WatchLogs(opts, "Apply", issuerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XDC21IssuerApply)
				if err := _XDC21Issuer.contract.UnpackLog(event, "Apply", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// XDC21IssuerChargeIterator is returned from FilterCharge and is used to iterate over the raw logs and unpacked data for Charge events raised by the XDC21Issuer contract.
type XDC21IssuerChargeIterator struct {
	Event *XDC21IssuerCharge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  XDPoSChain.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XDC21IssuerChargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XDC21IssuerCharge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(XDC21IssuerCharge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *XDC21IssuerChargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XDC21IssuerChargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XDC21IssuerCharge represents a Charge event raised by the XDC21Issuer contract.
type XDC21IssuerCharge struct {
	Supporter common.Address
	Token     common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCharge is a free log retrieval operation binding the contract event 0x5cffac866325fd9b2a8ea8df2f110a0058313b279402d15ae28dd324a2282e06.
//
// Solidity: event Charge(supporter indexed address, token indexed address, value uint256)
func (_XDC21Issuer *XDC21IssuerFilterer) FilterCharge(opts *bind.FilterOpts, supporter []common.Address, token []common.Address) (*XDC21IssuerChargeIterator, error) {

	var supporterRule []interface{}
	for _, supporterItem := range supporter {
		supporterRule = append(supporterRule, supporterItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _XDC21Issuer.contract.FilterLogs(opts, "Charge", supporterRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &XDC21IssuerChargeIterator{contract: _XDC21Issuer.contract, event: "Charge", logs: logs, sub: sub}, nil
}

// WatchCharge is a free log subscription operation binding the contract event 0x5cffac866325fd9b2a8ea8df2f110a0058313b279402d15ae28dd324a2282e06.
//
// Solidity: event Charge(supporter indexed address, token indexed address, value uint256)
func (_XDC21Issuer *XDC21IssuerFilterer) WatchCharge(opts *bind.WatchOpts, sink chan<- *XDC21IssuerCharge, supporter []common.Address, token []common.Address) (event.Subscription, error) {

	var supporterRule []interface{}
	for _, supporterItem := range supporter {
		supporterRule = append(supporterRule, supporterItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _XDC21Issuer.contract.WatchLogs(opts, "Charge", supporterRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XDC21IssuerCharge)
				if err := _XDC21Issuer.contract.UnpackLog(event, "Charge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
