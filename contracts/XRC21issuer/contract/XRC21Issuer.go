// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// AbstractTokenXRC21ABI is the input ABI used to generate the binding from.
const AbstractTokenXRC21ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AbstractTokenXRC21Bin is the compiled bytecode used for deploying new contracts.
const AbstractTokenXRC21Bin = `0x`

// DeployAbstractTokenXRC21 deploys a new Ethereum contract, binding an instance of AbstractTokenXRC21 to it.
func DeployAbstractTokenXRC21(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AbstractTokenXRC21, error) {
	parsed, err := abi.JSON(strings.NewReader(AbstractTokenXRC21ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbstractTokenXRC21Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AbstractTokenXRC21{AbstractTokenXRC21Caller: AbstractTokenXRC21Caller{contract: contract}, AbstractTokenXRC21Transactor: AbstractTokenXRC21Transactor{contract: contract}, AbstractTokenXRC21Filterer: AbstractTokenXRC21Filterer{contract: contract}}, nil
}

// AbstractTokenXRC21 is an auto generated Go binding around an Ethereum contract.
type AbstractTokenXRC21 struct {
	AbstractTokenXRC21Caller     // Read-only binding to the contract
	AbstractTokenXRC21Transactor // Write-only binding to the contract
	AbstractTokenXRC21Filterer   // Log filterer for contract events
}

// AbstractTokenXRC21Caller is an auto generated read-only Go binding around an Ethereum contract.
type AbstractTokenXRC21Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractTokenXRC21Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AbstractTokenXRC21Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractTokenXRC21Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbstractTokenXRC21Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractTokenXRC21Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbstractTokenXRC21Session struct {
	Contract     *AbstractTokenXRC21 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AbstractTokenXRC21CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbstractTokenXRC21CallerSession struct {
	Contract *AbstractTokenXRC21Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AbstractTokenXRC21TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbstractTokenXRC21TransactorSession struct {
	Contract     *AbstractTokenXRC21Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AbstractTokenXRC21Raw is an auto generated low-level Go binding around an Ethereum contract.
type AbstractTokenXRC21Raw struct {
	Contract *AbstractTokenXRC21 // Generic contract binding to access the raw methods on
}

// AbstractTokenXRC21CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbstractTokenXRC21CallerRaw struct {
	Contract *AbstractTokenXRC21Caller // Generic read-only contract binding to access the raw methods on
}

// AbstractTokenXRC21TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbstractTokenXRC21TransactorRaw struct {
	Contract *AbstractTokenXRC21Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAbstractTokenXRC21 creates a new instance of AbstractTokenXRC21, bound to a specific deployed contract.
func NewAbstractTokenXRC21(address common.Address, backend bind.ContractBackend) (*AbstractTokenXRC21, error) {
	contract, err := bindAbstractTokenXRC21(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbstractTokenXRC21{AbstractTokenXRC21Caller: AbstractTokenXRC21Caller{contract: contract}, AbstractTokenXRC21Transactor: AbstractTokenXRC21Transactor{contract: contract}, AbstractTokenXRC21Filterer: AbstractTokenXRC21Filterer{contract: contract}}, nil
}

// NewAbstractTokenXRC21Caller creates a new read-only instance of AbstractTokenXRC21, bound to a specific deployed contract.
func NewAbstractTokenXRC21Caller(address common.Address, caller bind.ContractCaller) (*AbstractTokenXRC21Caller, error) {
	contract, err := bindAbstractTokenXRC21(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractTokenXRC21Caller{contract: contract}, nil
}

// NewAbstractTokenXRC21Transactor creates a new write-only instance of AbstractTokenXRC21, bound to a specific deployed contract.
func NewAbstractTokenXRC21Transactor(address common.Address, transactor bind.ContractTransactor) (*AbstractTokenXRC21Transactor, error) {
	contract, err := bindAbstractTokenXRC21(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractTokenXRC21Transactor{contract: contract}, nil
}

// NewAbstractTokenXRC21Filterer creates a new log filterer instance of AbstractTokenXRC21, bound to a specific deployed contract.
func NewAbstractTokenXRC21Filterer(address common.Address, filterer bind.ContractFilterer) (*AbstractTokenXRC21Filterer, error) {
	contract, err := bindAbstractTokenXRC21(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbstractTokenXRC21Filterer{contract: contract}, nil
}

// bindAbstractTokenXRC21 binds a generic wrapper to an already deployed contract.
func bindAbstractTokenXRC21(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbstractTokenXRC21ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractTokenXRC21 *AbstractTokenXRC21Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AbstractTokenXRC21.Contract.AbstractTokenXRC21Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractTokenXRC21 *AbstractTokenXRC21Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractTokenXRC21.Contract.AbstractTokenXRC21Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractTokenXRC21 *AbstractTokenXRC21Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractTokenXRC21.Contract.AbstractTokenXRC21Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractTokenXRC21 *AbstractTokenXRC21CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AbstractTokenXRC21.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractTokenXRC21 *AbstractTokenXRC21TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractTokenXRC21.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractTokenXRC21 *AbstractTokenXRC21TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractTokenXRC21.Contract.contract.Transact(opts, method, params...)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_AbstractTokenXRC21 *AbstractTokenXRC21Caller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AbstractTokenXRC21.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_AbstractTokenXRC21 *AbstractTokenXRC21Session) Issuer() (common.Address, error) {
	return _AbstractTokenXRC21.Contract.Issuer(&_AbstractTokenXRC21.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_AbstractTokenXRC21 *AbstractTokenXRC21CallerSession) Issuer() (common.Address, error) {
	return _AbstractTokenXRC21.Contract.Issuer(&_AbstractTokenXRC21.CallOpts)
}

// XRC21IssuerABI is the input ABI used to generate the binding from.
const XRC21IssuerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"minCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenCapacity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"apply\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"charge\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Apply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"supporter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Charge\",\"type\":\"event\"}]"

// XRC21IssuerBin is the compiled bytecode used for deploying new contracts.
const XRC21IssuerBin = `0x608060405234801561001057600080fd5b5060405160208061047d8339810160405251600055610449806100346000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633fa615b081146100715780638f3a981c146100985780639d63848a146100b9578063c6b32f341461011e578063fc6bd76a14610134575b600080fd5b34801561007d57600080fd5b50610086610148565b60408051918252519081900360200190f35b3480156100a457600080fd5b50610086600160a060020a036004351661014e565b3480156100c557600080fd5b506100ce610169565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561010a5781810151838201526020016100f2565b505050509050019250505060405180910390f35b610132600160a060020a03600435166101cb565b005b610132600160a060020a036004351661035d565b60005490565b600160a060020a031660009081526002602052604090205490565b606060018054806020026020016040519081016040528092919081815260200182805480156101c157602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116101a3575b5050505050905090565b600081600160a060020a03811615156101e357600080fd5b6000543410156101f257600080fd5b82915033600160a060020a031682600160a060020a0316631d1438486040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561025657600080fd5b505af115801561026a573d6000803e3d6000fd5b505050506040513d602081101561028057600080fd5b5051600160a060020a03161461029557600080fd5b600180548082019091557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385169081179091556000908152600260205260409020546103039034610404565b600160a060020a0384166000818152600260209081526040918290209390935580513481529051919233927f2d485624158277d5113a56388c3abf5c20e3511dd112123ba376d16b21e4d7169281900390910190a3505050565b80600160a060020a038116151561037357600080fd5b60005434101561038257600080fd5b600160a060020a0382166000908152600260205260409020546103ab903463ffffffff61040416565b600160a060020a0383166000818152600260209081526040918290209390935580513481529051919233927f5cffac866325fd9b2a8ea8df2f110a0058313b279402d15ae28dd324a2282e069281900390910190a35050565b60008282018381101561041657600080fd5b93925050505600a165627a7a7230582017d25c61c6dd745630f85ec71140358aef38c084b91e940943bb4a15314b5e660029`

// DeployXRC21Issuer deploys a new Ethereum contract, binding an instance of XRC21Issuer to it.
func DeployXRC21Issuer(auth *bind.TransactOpts, backend bind.ContractBackend, value *big.Int) (common.Address, *types.Transaction, *XRC21Issuer, error) {
	parsed, err := abi.JSON(strings.NewReader(XRC21IssuerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(XRC21IssuerBin), backend, value)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XRC21Issuer{XRC21IssuerCaller: XRC21IssuerCaller{contract: contract}, XRC21IssuerTransactor: XRC21IssuerTransactor{contract: contract}, XRC21IssuerFilterer: XRC21IssuerFilterer{contract: contract}}, nil
}

// XRC21Issuer is an auto generated Go binding around an Ethereum contract.
type XRC21Issuer struct {
	XRC21IssuerCaller     // Read-only binding to the contract
	XRC21IssuerTransactor // Write-only binding to the contract
	XRC21IssuerFilterer   // Log filterer for contract events
}

// XRC21IssuerCaller is an auto generated read-only Go binding around an Ethereum contract.
type XRC21IssuerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XRC21IssuerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XRC21IssuerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XRC21IssuerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XRC21IssuerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XRC21IssuerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XRC21IssuerSession struct {
	Contract     *XRC21Issuer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XRC21IssuerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XRC21IssuerCallerSession struct {
	Contract *XRC21IssuerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// XRC21IssuerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XRC21IssuerTransactorSession struct {
	Contract     *XRC21IssuerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// XRC21IssuerRaw is an auto generated low-level Go binding around an Ethereum contract.
type XRC21IssuerRaw struct {
	Contract *XRC21Issuer // Generic contract binding to access the raw methods on
}

// XRC21IssuerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XRC21IssuerCallerRaw struct {
	Contract *XRC21IssuerCaller // Generic read-only contract binding to access the raw methods on
}

// XRC21IssuerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XRC21IssuerTransactorRaw struct {
	Contract *XRC21IssuerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXRC21Issuer creates a new instance of XRC21Issuer, bound to a specific deployed contract.
func NewXRC21Issuer(address common.Address, backend bind.ContractBackend) (*XRC21Issuer, error) {
	contract, err := bindXRC21Issuer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XRC21Issuer{XRC21IssuerCaller: XRC21IssuerCaller{contract: contract}, XRC21IssuerTransactor: XRC21IssuerTransactor{contract: contract}, XRC21IssuerFilterer: XRC21IssuerFilterer{contract: contract}}, nil
}

// NewXRC21IssuerCaller creates a new read-only instance of XRC21Issuer, bound to a specific deployed contract.
func NewXRC21IssuerCaller(address common.Address, caller bind.ContractCaller) (*XRC21IssuerCaller, error) {
	contract, err := bindXRC21Issuer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XRC21IssuerCaller{contract: contract}, nil
}

// NewXRC21IssuerTransactor creates a new write-only instance of XRC21Issuer, bound to a specific deployed contract.
func NewXRC21IssuerTransactor(address common.Address, transactor bind.ContractTransactor) (*XRC21IssuerTransactor, error) {
	contract, err := bindXRC21Issuer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XRC21IssuerTransactor{contract: contract}, nil
}

// NewXRC21IssuerFilterer creates a new log filterer instance of XRC21Issuer, bound to a specific deployed contract.
func NewXRC21IssuerFilterer(address common.Address, filterer bind.ContractFilterer) (*XRC21IssuerFilterer, error) {
	contract, err := bindXRC21Issuer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XRC21IssuerFilterer{contract: contract}, nil
}

// bindXRC21Issuer binds a generic wrapper to an already deployed contract.
func bindXRC21Issuer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XRC21IssuerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XRC21Issuer *XRC21IssuerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XRC21Issuer.Contract.XRC21IssuerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XRC21Issuer *XRC21IssuerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XRC21Issuer.Contract.XRC21IssuerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XRC21Issuer *XRC21IssuerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XRC21Issuer.Contract.XRC21IssuerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XRC21Issuer *XRC21IssuerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XRC21Issuer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XRC21Issuer *XRC21IssuerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XRC21Issuer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XRC21Issuer *XRC21IssuerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XRC21Issuer.Contract.contract.Transact(opts, method, params...)
}

// GetTokenCapacity is a free data retrieval call binding the contract method 0x8f3a981c.
//
// Solidity: function getTokenCapacity(token address) constant returns(uint256)
func (_XRC21Issuer *XRC21IssuerCaller) GetTokenCapacity(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XRC21Issuer.contract.Call(opts, out, "getTokenCapacity", token)
	return *ret0, err
}

// GetTokenCapacity is a free data retrieval call binding the contract method 0x8f3a981c.
//
// Solidity: function getTokenCapacity(token address) constant returns(uint256)
func (_XRC21Issuer *XRC21IssuerSession) GetTokenCapacity(token common.Address) (*big.Int, error) {
	return _XRC21Issuer.Contract.GetTokenCapacity(&_XRC21Issuer.CallOpts, token)
}

// GetTokenCapacity is a free data retrieval call binding the contract method 0x8f3a981c.
//
// Solidity: function getTokenCapacity(token address) constant returns(uint256)
func (_XRC21Issuer *XRC21IssuerCallerSession) GetTokenCapacity(token common.Address) (*big.Int, error) {
	return _XRC21Issuer.Contract.GetTokenCapacity(&_XRC21Issuer.CallOpts, token)
}

// MinCap is a free data retrieval call binding the contract method 0x3fa615b0.
//
// Solidity: function minCap() constant returns(uint256)
func (_XRC21Issuer *XRC21IssuerCaller) MinCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XRC21Issuer.contract.Call(opts, out, "minCap")
	return *ret0, err
}

// MinCap is a free data retrieval call binding the contract method 0x3fa615b0.
//
// Solidity: function minCap() constant returns(uint256)
func (_XRC21Issuer *XRC21IssuerSession) MinCap() (*big.Int, error) {
	return _XRC21Issuer.Contract.MinCap(&_XRC21Issuer.CallOpts)
}

// MinCap is a free data retrieval call binding the contract method 0x3fa615b0.
//
// Solidity: function minCap() constant returns(uint256)
func (_XRC21Issuer *XRC21IssuerCallerSession) MinCap() (*big.Int, error) {
	return _XRC21Issuer.Contract.MinCap(&_XRC21Issuer.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_XRC21Issuer *XRC21IssuerCaller) Tokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _XRC21Issuer.contract.Call(opts, out, "tokens")
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_XRC21Issuer *XRC21IssuerSession) Tokens() ([]common.Address, error) {
	return _XRC21Issuer.Contract.Tokens(&_XRC21Issuer.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_XRC21Issuer *XRC21IssuerCallerSession) Tokens() ([]common.Address, error) {
	return _XRC21Issuer.Contract.Tokens(&_XRC21Issuer.CallOpts)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_XRC21Issuer *XRC21IssuerTransactor) Apply(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _XRC21Issuer.contract.Transact(opts, "apply", token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_XRC21Issuer *XRC21IssuerSession) Apply(token common.Address) (*types.Transaction, error) {
	return _XRC21Issuer.Contract.Apply(&_XRC21Issuer.TransactOpts, token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_XRC21Issuer *XRC21IssuerTransactorSession) Apply(token common.Address) (*types.Transaction, error) {
	return _XRC21Issuer.Contract.Apply(&_XRC21Issuer.TransactOpts, token)
}

// Charge is a paid mutator transaction binding the contract method 0xfc6bd76a.
//
// Solidity: function charge(token address) returns()
func (_XRC21Issuer *XRC21IssuerTransactor) Charge(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _XRC21Issuer.contract.Transact(opts, "charge", token)
}

// Charge is a paid mutator transaction binding the contract method 0xfc6bd76a.
//
// Solidity: function charge(token address) returns()
func (_XRC21Issuer *XRC21IssuerSession) Charge(token common.Address) (*types.Transaction, error) {
	return _XRC21Issuer.Contract.Charge(&_XRC21Issuer.TransactOpts, token)
}

// Charge is a paid mutator transaction binding the contract method 0xfc6bd76a.
//
// Solidity: function charge(token address) returns()
func (_XRC21Issuer *XRC21IssuerTransactorSession) Charge(token common.Address) (*types.Transaction, error) {
	return _XRC21Issuer.Contract.Charge(&_XRC21Issuer.TransactOpts, token)
}

// XRC21IssuerApplyIterator is returned from FilterApply and is used to iterate over the raw logs and unpacked data for Apply events raised by the XRC21Issuer contract.
type XRC21IssuerApplyIterator struct {
	Event *XRC21IssuerApply // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XRC21IssuerApplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XRC21IssuerApply)
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
		it.Event = new(XRC21IssuerApply)
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
func (it *XRC21IssuerApplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XRC21IssuerApplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XRC21IssuerApply represents a Apply event raised by the XRC21Issuer contract.
type XRC21IssuerApply struct {
	Issuer common.Address
	Token  common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApply is a free log retrieval operation binding the contract event 0x2d485624158277d5113a56388c3abf5c20e3511dd112123ba376d16b21e4d716.
//
// Solidity: event Apply(issuer indexed address, token indexed address, value uint256)
func (_XRC21Issuer *XRC21IssuerFilterer) FilterApply(opts *bind.FilterOpts, issuer []common.Address, token []common.Address) (*XRC21IssuerApplyIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _XRC21Issuer.contract.FilterLogs(opts, "Apply", issuerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &XRC21IssuerApplyIterator{contract: _XRC21Issuer.contract, event: "Apply", logs: logs, sub: sub}, nil
}

// WatchApply is a free log subscription operation binding the contract event 0x2d485624158277d5113a56388c3abf5c20e3511dd112123ba376d16b21e4d716.
//
// Solidity: event Apply(issuer indexed address, token indexed address, value uint256)
func (_XRC21Issuer *XRC21IssuerFilterer) WatchApply(opts *bind.WatchOpts, sink chan<- *XRC21IssuerApply, issuer []common.Address, token []common.Address) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _XRC21Issuer.contract.WatchLogs(opts, "Apply", issuerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XRC21IssuerApply)
				if err := _XRC21Issuer.contract.UnpackLog(event, "Apply", log); err != nil {
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

// XRC21IssuerChargeIterator is returned from FilterCharge and is used to iterate over the raw logs and unpacked data for Charge events raised by the XRC21Issuer contract.
type XRC21IssuerChargeIterator struct {
	Event *XRC21IssuerCharge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *XRC21IssuerChargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XRC21IssuerCharge)
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
		it.Event = new(XRC21IssuerCharge)
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
func (it *XRC21IssuerChargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XRC21IssuerChargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XRC21IssuerCharge represents a Charge event raised by the XRC21Issuer contract.
type XRC21IssuerCharge struct {
	Supporter common.Address
	Token     common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCharge is a free log retrieval operation binding the contract event 0x5cffac866325fd9b2a8ea8df2f110a0058313b279402d15ae28dd324a2282e06.
//
// Solidity: event Charge(supporter indexed address, token indexed address, value uint256)
func (_XRC21Issuer *XRC21IssuerFilterer) FilterCharge(opts *bind.FilterOpts, supporter []common.Address, token []common.Address) (*XRC21IssuerChargeIterator, error) {

	var supporterRule []interface{}
	for _, supporterItem := range supporter {
		supporterRule = append(supporterRule, supporterItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _XRC21Issuer.contract.FilterLogs(opts, "Charge", supporterRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &XRC21IssuerChargeIterator{contract: _XRC21Issuer.contract, event: "Charge", logs: logs, sub: sub}, nil
}

// WatchCharge is a free log subscription operation binding the contract event 0x5cffac866325fd9b2a8ea8df2f110a0058313b279402d15ae28dd324a2282e06.
//
// Solidity: event Charge(supporter indexed address, token indexed address, value uint256)
func (_XRC21Issuer *XRC21IssuerFilterer) WatchCharge(opts *bind.WatchOpts, sink chan<- *XRC21IssuerCharge, supporter []common.Address, token []common.Address) (event.Subscription, error) {

	var supporterRule []interface{}
	for _, supporterItem := range supporter {
		supporterRule = append(supporterRule, supporterItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _XRC21Issuer.contract.WatchLogs(opts, "Charge", supporterRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XRC21IssuerCharge)
				if err := _XRC21Issuer.contract.UnpackLog(event, "Charge", log); err != nil {
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
