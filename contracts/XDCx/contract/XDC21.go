// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/XinFinOrg/XDPoSChain"
	"github.com/XinFinOrg/XDPoSChain/accounts/abi"
	"github.com/XinFinOrg/XDPoSChain/accounts/abi/bind"
	"github.com/XinFinOrg/XDPoSChain/common"
	"github.com/XinFinOrg/XDPoSChain/core/types"
	"github.com/XinFinOrg/XDPoSChain/event"
)

// IXDC21ABI is the input ABI used to generate the binding from.
const IXDC21ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Fee\",\"type\":\"event\"}]"

// IXDC21Bin is the compiled bytecode used for deploying new contracts.
const IXDC21Bin = `0x`

// DeployIXDC21 deploys a new Ethereum contract, binding an instance of IXDC21 to it.
func DeployIXDC21(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IXDC21, error) {
	parsed, err := abi.JSON(strings.NewReader(IXDC21ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IXDC21Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IXDC21{IXDC21Caller: IXDC21Caller{contract: contract}, IXDC21Transactor: IXDC21Transactor{contract: contract}, IXDC21Filterer: IXDC21Filterer{contract: contract}}, nil
}

// IXDC21 is an auto generated Go binding around an Ethereum contract.
type IXDC21 struct {
	IXDC21Caller     // Read-only binding to the contract
	IXDC21Transactor // Write-only binding to the contract
	IXDC21Filterer   // Log filterer for contract events
}

// IXDC21Caller is an auto generated read-only Go binding around an Ethereum contract.
type IXDC21Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IXDC21Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IXDC21Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IXDC21Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IXDC21Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IXDC21Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IXDC21Session struct {
	Contract     *IXDC21           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IXDC21CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IXDC21CallerSession struct {
	Contract *IXDC21Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IXDC21TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IXDC21TransactorSession struct {
	Contract     *IXDC21Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IXDC21Raw is an auto generated low-level Go binding around an Ethereum contract.
type IXDC21Raw struct {
	Contract *IXDC21 // Generic contract binding to access the raw methods on
}

// IXDC21CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IXDC21CallerRaw struct {
	Contract *IXDC21Caller // Generic read-only contract binding to access the raw methods on
}

// IXDC21TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IXDC21TransactorRaw struct {
	Contract *IXDC21Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIXDC21 creates a new instance of IXDC21, bound to a specific deployed contract.
func NewIXDC21(address common.Address, backend bind.ContractBackend) (*IXDC21, error) {
	contract, err := bindIXDC21(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IXDC21{IXDC21Caller: IXDC21Caller{contract: contract}, IXDC21Transactor: IXDC21Transactor{contract: contract}, IXDC21Filterer: IXDC21Filterer{contract: contract}}, nil
}

// NewIXDC21Caller creates a new read-only instance of IXDC21, bound to a specific deployed contract.
func NewIXDC21Caller(address common.Address, caller bind.ContractCaller) (*IXDC21Caller, error) {
	contract, err := bindIXDC21(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IXDC21Caller{contract: contract}, nil
}

// NewIXDC21Transactor creates a new write-only instance of IXDC21, bound to a specific deployed contract.
func NewIXDC21Transactor(address common.Address, transactor bind.ContractTransactor) (*IXDC21Transactor, error) {
	contract, err := bindIXDC21(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IXDC21Transactor{contract: contract}, nil
}

// NewIXDC21Filterer creates a new log filterer instance of IXDC21, bound to a specific deployed contract.
func NewIXDC21Filterer(address common.Address, filterer bind.ContractFilterer) (*IXDC21Filterer, error) {
	contract, err := bindIXDC21(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IXDC21Filterer{contract: contract}, nil
}

// bindIXDC21 binds a generic wrapper to an already deployed contract.
func bindIXDC21(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IXDC21ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IXDC21 *IXDC21Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IXDC21.Contract.IXDC21Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IXDC21 *IXDC21Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IXDC21.Contract.IXDC21Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IXDC21 *IXDC21Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IXDC21.Contract.IXDC21Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IXDC21 *IXDC21CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IXDC21.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IXDC21 *IXDC21TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IXDC21.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IXDC21 *IXDC21TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IXDC21.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_IXDC21 *IXDC21Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IXDC21.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_IXDC21 *IXDC21Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IXDC21.Contract.Allowance(&_IXDC21.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_IXDC21 *IXDC21CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IXDC21.Contract.Allowance(&_IXDC21.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_IXDC21 *IXDC21Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IXDC21.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_IXDC21 *IXDC21Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _IXDC21.Contract.BalanceOf(&_IXDC21.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_IXDC21 *IXDC21CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _IXDC21.Contract.BalanceOf(&_IXDC21.CallOpts, who)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_IXDC21 *IXDC21Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IXDC21.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_IXDC21 *IXDC21Session) Decimals() (uint8, error) {
	return _IXDC21.Contract.Decimals(&_IXDC21.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_IXDC21 *IXDC21CallerSession) Decimals() (uint8, error) {
	return _IXDC21.Contract.Decimals(&_IXDC21.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_IXDC21 *IXDC21Caller) EstimateFee(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IXDC21.contract.Call(opts, out, "estimateFee", value)
	return *ret0, err
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_IXDC21 *IXDC21Session) EstimateFee(value *big.Int) (*big.Int, error) {
	return _IXDC21.Contract.EstimateFee(&_IXDC21.CallOpts, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_IXDC21 *IXDC21CallerSession) EstimateFee(value *big.Int) (*big.Int, error) {
	return _IXDC21.Contract.EstimateFee(&_IXDC21.CallOpts, value)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_IXDC21 *IXDC21Caller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IXDC21.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_IXDC21 *IXDC21Session) Issuer() (common.Address, error) {
	return _IXDC21.Contract.Issuer(&_IXDC21.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_IXDC21 *IXDC21CallerSession) Issuer() (common.Address, error) {
	return _IXDC21.Contract.Issuer(&_IXDC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IXDC21 *IXDC21Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IXDC21.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IXDC21 *IXDC21Session) TotalSupply() (*big.Int, error) {
	return _IXDC21.Contract.TotalSupply(&_IXDC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IXDC21 *IXDC21CallerSession) TotalSupply() (*big.Int, error) {
	return _IXDC21.Contract.TotalSupply(&_IXDC21.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IXDC21 *IXDC21Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IXDC21.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IXDC21 *IXDC21Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IXDC21.Contract.Approve(&_IXDC21.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IXDC21 *IXDC21TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IXDC21.Contract.Approve(&_IXDC21.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IXDC21 *IXDC21Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IXDC21.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IXDC21 *IXDC21Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IXDC21.Contract.Transfer(&_IXDC21.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IXDC21 *IXDC21TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IXDC21.Contract.Transfer(&_IXDC21.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IXDC21 *IXDC21Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IXDC21.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IXDC21 *IXDC21Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IXDC21.Contract.TransferFrom(&_IXDC21.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IXDC21 *IXDC21TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IXDC21.Contract.TransferFrom(&_IXDC21.TransactOpts, from, to, value)
}

// IXDC21ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IXDC21 contract.
type IXDC21ApprovalIterator struct {
	Event *IXDC21Approval // Event containing the contract specifics and raw log

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
func (it *IXDC21ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IXDC21Approval)
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
		it.Event = new(IXDC21Approval)
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
func (it *IXDC21ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IXDC21ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IXDC21Approval represents a Approval event raised by the IXDC21 contract.
type IXDC21Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_IXDC21 *IXDC21Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IXDC21ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IXDC21.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IXDC21ApprovalIterator{contract: _IXDC21.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_IXDC21 *IXDC21Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IXDC21Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IXDC21.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IXDC21Approval)
				if err := _IXDC21.contract.UnpackLog(event, "Approval", log); err != nil {
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

// IXDC21FeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the IXDC21 contract.
type IXDC21FeeIterator struct {
	Event *IXDC21Fee // Event containing the contract specifics and raw log

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
func (it *IXDC21FeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IXDC21Fee)
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
		it.Event = new(IXDC21Fee)
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
func (it *IXDC21FeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IXDC21FeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IXDC21Fee represents a Fee event raised by the IXDC21 contract.
type IXDC21Fee struct {
	From   common.Address
	To     common.Address
	Issuer common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_IXDC21 *IXDC21Filterer) FilterFee(opts *bind.FilterOpts, from []common.Address, to []common.Address, issuer []common.Address) (*IXDC21FeeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _IXDC21.contract.FilterLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &IXDC21FeeIterator{contract: _IXDC21.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_IXDC21 *IXDC21Filterer) WatchFee(opts *bind.WatchOpts, sink chan<- *IXDC21Fee, from []common.Address, to []common.Address, issuer []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _IXDC21.contract.WatchLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IXDC21Fee)
				if err := _IXDC21.contract.UnpackLog(event, "Fee", log); err != nil {
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

// IXDC21TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IXDC21 contract.
type IXDC21TransferIterator struct {
	Event *IXDC21Transfer // Event containing the contract specifics and raw log

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
func (it *IXDC21TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IXDC21Transfer)
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
		it.Event = new(IXDC21Transfer)
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
func (it *IXDC21TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IXDC21TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IXDC21Transfer represents a Transfer event raised by the IXDC21 contract.
type IXDC21Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_IXDC21 *IXDC21Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IXDC21TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IXDC21.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IXDC21TransferIterator{contract: _IXDC21.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_IXDC21 *IXDC21Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IXDC21Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IXDC21.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IXDC21Transfer)
				if err := _IXDC21.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// MyXDC21ABI is the input ABI used to generate the binding from.
const MyXDC21ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"burnId\",\"type\":\"uint256\"}],\"name\":\"getBurn\",\"outputs\":[{\"name\":\"_burner\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setMinFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"depositFee\",\"type\":\"uint256\"}],\"name\":\"setDepositFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"burnList\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"burner\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmationCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WITHDRAW_FEE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\"},{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionIds\",\"outputs\":[{\"name\":\"_transactionIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmations\",\"outputs\":[{\"name\":\"_confirmations\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"withdrawFee\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_required\",\"type\":\"uint256\"}],\"name\":\"changeRequirement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferContractIssuer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OWNER_COUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"required\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEPOSIT_FEE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"replaceOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBurnCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"_required\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"},{\"name\":\"cap\",\"type\":\"uint256\"},{\"name\":\"minFee\",\"type\":\"uint256\"},{\"name\":\"depositFee\",\"type\":\"uint256\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Execution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"RequirementChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burnID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TokenBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Fee\",\"type\":\"event\"}]"

// MyXDC21Bin is the compiled bytecode used for deploying new contracts.
const MyXDC21Bin = `0x6080604052600060085560006009553480156200001b57600080fd5b5060405162002a7b38038062002a7b8339810160409081528151602080840151928401516060850151608086015160a087015160c088015160e08901516101008a0151958a018051988b019a909895019693959294919390929160009189918991899162000090916005919086019062000356565b508151620000a690600690602085019062000356565b506007805460ff191660ff92909216919091179055505089518960328211801590620000d25750818111155b8015620000de57508015155b8015620000ea57508115155b1515620000f657600080fd5b6200010b338864010000000062000240810204565b6200011f33640100000000620002ff810204565b620001338664010000000062000337810204565b600092505b8b518310156200020b57600c60008d858151811015156200015557fe5b6020908102909101810151600160a060020a031682528101919091526040016000205460ff16158015620001ab57508b838151811015156200019357fe5b90602001906020020151600160a060020a0316600014155b1515620001b757600080fd5b6001600c60008e86815181101515620001cc57fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff19169115159190911790556001929092019162000138565b8b516200022090600d9060208f0190620003db565b505050600e98909855600991909155600855506200048895505050505050565b600160a060020a03821615156200025657600080fd5b6004546200027390826401000000006200221b6200033c82021704565b600455600160a060020a038216600090815260208190526040902054620002a990826401000000006200221b6200033c82021704565b600160a060020a0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b600160a060020a03811615156200031557600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b600155565b6000828201838110156200034f57600080fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200039957805160ff1916838001178555620003c9565b82800160010185558215620003c9579182015b82811115620003c9578251825591602001919060010190620003ac565b50620003d792915062000441565b5090565b82805482825590600052602060002090810192821562000433579160200282015b82811115620004335782518254600160a060020a031916600160a060020a03909116178255602090920191600190910190620003fc565b50620003d792915062000461565b6200045e91905b80821115620003d7576000815560010162000448565b90565b6200045e91905b80821115620003d7578054600160a060020a031916815560010162000468565b6125e380620004986000396000f30060806040526004361061020e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025e7c27811461021357806306fdde0314610247578063095ea7b3146102d1578063127e8e4d14610309578063173825d91461033357806318160ddd146103565780631d1438481461036b57806320ea8d861461038057806323b872dd1461039857806324ec7590146103c25780632eb3f49b146103d75780632f54bf6e14610487578063313ce567146104a857806331ac9920146104d35780633411c81c146104eb578063490ae2101461050f57806354741525146105275780637065cb481461054657806370a082311461056757806377661c6414610588578063784547a7146105fa5780638b51d13f1461061257806395d89b411461062a5780639ace38c21461063f5780639bff5ddb146106fa578063a0e67e2b1461070f578063a8abe69a14610774578063a9059cbb14610799578063b5dc40c3146107bd578063b6ac642a146107d5578063b77bf600146107ed578063ba51a6df14610802578063c01a8c841461081a578063c28ce6ff14610832578063c642747414610853578063d74f8edd146108bc578063dc8452cd146108d1578063dd62ed3e146108e6578063de363e651461090d578063e20056e614610922578063e7cf548c14610949578063ee22610b1461095e578063fe9d930314610976575b600080fd5b34801561021f57600080fd5b5061022b6004356109d4565b60408051600160a060020a039092168252519081900360200190f35b34801561025357600080fd5b5061025c6109fc565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561029657818101518382015260200161027e565b50505050905090810190601f1680156102c35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102dd57600080fd5b506102f5600160a060020a0360043516602435610a93565b604080519115158252519081900360200190f35b34801561031557600080fd5b50610321600435610b4d565b60408051918252519081900360200190f35b34801561033f57600080fd5b50610354600160a060020a0360043516610b79565b005b34801561036257600080fd5b50610321610cf0565b34801561037757600080fd5b5061022b610cf6565b34801561038c57600080fd5b50610354600435610d05565b3480156103a457600080fd5b506102f5600160a060020a0360043581169060243516604435610dbf565b3480156103ce57600080fd5b50610321610f03565b3480156103e357600080fd5b506103ef600435610f09565b6040518084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561044a578181015183820152602001610432565b50505050905090810190601f1680156104775780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561049357600080fd5b506102f5600160a060020a0360043516611016565b3480156104b457600080fd5b506104bd61102b565b6040805160ff9092168252519081900360200190f35b3480156104df57600080fd5b50610354600435611034565b3480156104f757600080fd5b506102f5600435600160a060020a0360243516611057565b34801561051b57600080fd5b50610354600435611077565b34801561053357600080fd5b5061032160043515156024351515611098565b34801561055257600080fd5b50610354600160a060020a0360043516611104565b34801561057357600080fd5b50610321600160a060020a0360043516611229565b34801561059457600080fd5b506105a0600435611244565b6040518084815260200183600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019080838360008381101561044a578181015183820152602001610432565b34801561060657600080fd5b506102f560043561130d565b34801561061e57600080fd5b50610321600435611391565b34801561063657600080fd5b5061025c611400565b34801561064b57600080fd5b50610657600435611461565b6040518085600160a060020a0316600160a060020a031681526020018481526020018060200183151515158152602001828103825284818151815260200191508051906020019080838360005b838110156106bc5781810151838201526020016106a4565b50505050905090810190601f1680156106e95780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34801561070657600080fd5b50610321611520565b34801561071b57600080fd5b50610724611526565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610760578181015183820152602001610748565b505050509050019250505060405180910390f35b34801561078057600080fd5b5061072460043560243560443515156064351515611587565b3480156107a557600080fd5b506102f5600160a060020a03600435166024356116d3565b3480156107c957600080fd5b506107246004356117a5565b3480156107e157600080fd5b5061035460043561191e565b3480156107f957600080fd5b5061032161193f565b34801561080e57600080fd5b50610354600435611945565b34801561082657600080fd5b506103546004356119c4565b34801561083e57600080fd5b50610354600160a060020a0360043516611a8d565b34801561085f57600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610321948235600160a060020a0316946024803595369594606494920191908190840183828082843750949750611ac19650505050505050565b3480156108c857600080fd5b50610321611ae0565b3480156108dd57600080fd5b50610321611ae5565b3480156108f257600080fd5b50610321600160a060020a0360043581169060243516611aeb565b34801561091957600080fd5b50610321611b16565b34801561092e57600080fd5b50610354600160a060020a0360043581169060243516611b1c565b34801561095557600080fd5b50610321611ca6565b34801561096a57600080fd5b50610354600435611cac565b34801561098257600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610354958335953695604494919390910191908190840183828082843750949750611f079650505050505050565b600d8054829081106109e257fe5b600091825260209091200154600160a060020a0316905081565b60058054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610a885780601f10610a5d57610100808354040283529160200191610a88565b820191906000526020600020905b815481529060010190602001808311610a6b57829003601f168201915b505050505090505b90565b6000600160a060020a0383161515610aaa57600080fd5b600154336000908152602081905260409020541015610ac857600080fd5b336000818152600360209081526040808320600160a060020a0388811685529252909120849055600254600154610b04939291909116906120fb565b604080518381529051600160a060020a0385169133917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360200190a350600192915050565b600154600090610b7390610b67848463ffffffff6121ed16565b9063ffffffff61221b16565b92915050565b6000333014610b8757600080fd5b600160a060020a0382166000908152600c6020526040902054829060ff161515610bb057600080fd5b600160a060020a0383166000908152600c60205260408120805460ff1916905591505b600d5460001901821015610c8b5782600160a060020a0316600d83815481101515610bfa57fe5b600091825260209091200154600160a060020a03161415610c8057600d80546000198101908110610c2757fe5b600091825260209091200154600d8054600160a060020a039092169184908110610c4d57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550610c8b565b600190910190610bd3565b600d80546000190190610c9e90826124f6565b50600d54600e541115610cb757600d54610cb790611945565b604051600160a060020a038416907f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9090600090a2505050565b60045490565b600254600160a060020a031690565b336000818152600c602052604090205460ff161515610d2357600080fd5b6000828152600b602090815260408083203380855292529091205483919060ff161515610d4f57600080fd5b6000848152600a6020526040902060030154849060ff1615610d7057600080fd5b6000858152600b60209081526040808320338085529252808320805460ff191690555187927ff6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e991a35050505050565b600080610dd76001548461221b90919063ffffffff16565b600160a060020a038616600090815260208190526040902054909150811115610dff57600080fd5b600160a060020a0385166000908152600360209081526040808320338452909152902054831115610e2f57600080fd5b600160a060020a0385166000908152600360209081526040808320338452909152902054610e63908263ffffffff61222d16565b600160a060020a0386166000908152600360209081526040808320338452909152902055610e928585856120fb565b600254600154610eaf918791600160a060020a03909116906120fb565b6002546001546040805191825251600160a060020a039283169287169133917ffcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd999181900360200190a4506001949350505050565b60015490565b6000806060601084815481101515610f1d57fe5b600091825260209091206001600390920201015460108054600160a060020a0390921694509085908110610f4d57fe5b9060005260206000209060030201600001549150601084815481101515610f7057fe5b600091825260209182902060026003909202018101805460408051601f6000196101006001861615020190931694909404918201859004850284018501905280835291929091908301828280156110085780601f10610fdd57610100808354040283529160200191611008565b820191906000526020600020905b815481529060010190602001808311610feb57829003601f168201915b505050505090509193909250565b600c6020526000908152604090205460ff1681565b60075460ff1690565b600254600160a060020a0316331461104b57600080fd5b61105481612244565b50565b600b60209081526000928352604080842090915290825290205460ff1681565b61107f610cf6565b600160a060020a0316331461109357600080fd5b600955565b6000805b600f548110156110fd578380156110c557506000818152600a602052604090206003015460ff16155b806110e957508280156110e957506000818152600a602052604090206003015460ff165b156110f5576001820191505b60010161109c565b5092915050565b33301461111057600080fd5b600160a060020a0381166000908152600c6020526040902054819060ff161561113857600080fd5b81600160a060020a038116151561114e57600080fd5b600d80549050600101600e546032821115801561116b5750818111155b801561117657508015155b801561118157508115155b151561118c57600080fd5b600160a060020a0385166000818152600c6020526040808220805460ff19166001908117909155600d8054918201815583527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb501805473ffffffffffffffffffffffffffffffffffffffff191684179055517ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d9190a25050505050565b600160a060020a031660009081526020819052604090205490565b601080548290811061125257fe5b6000918252602091829020600391909102018054600180830154600280850180546040805161010096831615969096026000190190911692909204601f8101889004880285018801909252818452939650600160a060020a03909116949192918301828280156113035780601f106112d857610100808354040283529160200191611303565b820191906000526020600020905b8154815290600101906020018083116112e657829003601f168201915b5050505050905083565b600080805b600d5481101561138a576000848152600b60205260408120600d80549192918490811061133b57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff161561136f576001820191505b600e54821415611382576001925061138a565b600101611312565b5050919050565b6000805b600d548110156113fa576000838152600b60205260408120600d8054919291849081106113be57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff16156113f2576001820191505b600101611395565b50919050565b60068054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610a885780601f10610a5d57610100808354040283529160200191610a88565b600a6020908152600091825260409182902080546001808301546002808501805488516101009582161595909502600019011691909104601f8101879004870284018701909752868352600160a060020a0390931695909491929183018282801561150d5780601f106114e25761010080835404028352916020019161150d565b820191906000526020600020905b8154815290600101906020018083116114f057829003601f168201915b5050506003909301549192505060ff1684565b60085481565b6060600d805480602002602001604051908101604052809291908181526020018280548015610a8857602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311611560575050505050905090565b606060006060600080600f54881161159f57876115a3565b600f545b93508884036040519080825280602002602001820160405280156115d1578160200160208202803883390190505b509250600091508890505b878110156116555786801561160357506000818152600a602052604090206003015460ff16155b80611627575085801561162757506000818152600a602052604090206003015460ff165b1561164d5780838381518110151561163b57fe5b60209081029091010152600191909101905b6001016115dc565b8160405190808252806020026020018201604052801561167f578160200160208202803883390190505b509450600090505b818110156116c757828181518110151561169d57fe5b9060200190602002015185828151811015156116b557fe5b60209081029091010152600101611687565b50505050949350505050565b600080600160a060020a03841615156116eb57600080fd5b6001546116ff90849063ffffffff61221b16565b3360009081526020819052604090205490915081111561171e57600080fd5b6117293385856120fb565b6000600154111561179b57600254600154611751913391600160a060020a03909116906120fb565b6002546001546040805191825251600160a060020a039283169287169133917ffcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd999181900360200190a45b5060019392505050565b606080600080600d805490506040519080825280602002602001820160405280156117da578160200160208202803883390190505b50925060009150600090505b600d54811015611897576000858152600b60205260408120600d80549192918490811061180f57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff161561188f57600d80548290811061184a57fe5b6000918252602090912001548351600160a060020a039091169084908490811061187057fe5b600160a060020a03909216602092830290910190910152600191909101905b6001016117e6565b816040519080825280602002602001820160405280156118c1578160200160208202803883390190505b509350600090505b818110156119165782818151811015156118df57fe5b9060200190602002015184828151811015156118f757fe5b600160a060020a039092166020928302909101909101526001016118c9565b505050919050565b611926610cf6565b600160a060020a0316331461193a57600080fd5b600855565b600f5481565b33301461195157600080fd5b600d5481603282118015906119665750818111155b801561197157508015155b801561197c57508115155b151561198757600080fd5b600e8390556040805184815290517fa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a9181900360200190a1505050565b336000818152600c602052604090205460ff1615156119e257600080fd5b6000828152600a60205260409020548290600160a060020a03161515611a0757600080fd5b6000838152600b602090815260408083203380855292529091205484919060ff1615611a3257600080fd5b6000858152600b60209081526040808320338085529252808320805460ff191660011790555187927f4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef91a3611a8685611cac565b5050505050565b611a95610cf6565b600160a060020a03163314611aa957600080fd5b600160a060020a038116156110545761105433612249565b6000611ace84848461228d565b9050611ad9816119c4565b9392505050565b603281565b600e5481565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b60095481565b6000333014611b2a57600080fd5b600160a060020a0383166000908152600c6020526040902054839060ff161515611b5357600080fd5b600160a060020a0383166000908152600c6020526040902054839060ff1615611b7b57600080fd5b600092505b600d54831015611c0c5784600160a060020a0316600d84815481101515611ba357fe5b600091825260209091200154600160a060020a03161415611c015783600d84815481101515611bce57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550611c0c565b600190920191611b80565b600160a060020a038086166000818152600c6020526040808220805460ff1990811690915593881682528082208054909416600117909355915190917f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9091a2604051600160a060020a038516907ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d90600090a25050505050565b60105490565b336000818152600c602052604081205490919060ff161515611ccd57600080fd5b6000838152600b602090815260408083203380855292529091205484919060ff161515611cf957600080fd5b6000858152600a6020526040902060030154859060ff1615611d1a57600080fd5b611d238661130d565b15611eff576000868152600a6020526040902060038101805460ff19166001908117909155600280830154929750600019918316156101000291909101909116041515611ded576009546001860154611d819163ffffffff61222d16565b600186018190558554611d9f91600160a060020a039091169061237e565b60006009541115611dbd57611dbd611db5610cf6565b60095461237e565b60405186907f33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed7590600090a2611eff565b8460000160009054906101000a9004600160a060020a0316600160a060020a03168560010154866002016040518082805460018160011615610100020316600290048015611e7c5780601f10611e5157610100808354040283529160200191611e7c565b820191906000526020600020905b815481529060010190602001808311611e5f57829003601f168201915b505091505060006040518083038185875af19250505015611ec75760405186907f33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed7590600090a2611eff565b60405186907f526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b7923690600090a260038501805460ff191690555b505050505050565b6008546000908311611f1857600080fd5b611f223384612428565b60006008541115611f4057611f40611f38610cf6565b60085461237e565b600854611f5490849063ffffffff61222d16565b604080516060810182528281523360208083019182529282018681526010805460018101808355600092909252845160039091027f1b6847dc741a1b0cd08d278845f9d819d87b734759afb55fe2de5cb82a9ae672810191825593517f1b6847dc741a1b0cd08d278845f9d819d87b734759afb55fe2de5cb82a9ae67385018054600160a060020a039290921673ffffffffffffffffffffffffffffffffffffffff199092169190911790559151805196975090959394919361203e937f1b6847dc741a1b0cd08d278845f9d819d87b734759afb55fe2de5cb82a9ae6740192919091019061251f565b5050505033600160a060020a03166001601080549050037f6905852b196f81e7e03058512a599446c358027fc943c1e193b6649a39379bb583856040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156120bb5781810151838201526020016120a3565b50505050905090810190601f1680156120e85780820380516001836020036101000a031916815260200191505b50935050505060405180910390a3505050565b600160a060020a03831660009081526020819052604090205481111561212057600080fd5b600160a060020a038216151561213557600080fd5b600160a060020a03831660009081526020819052604090205461215e908263ffffffff61222d16565b600160a060020a038085166000908152602081905260408082209390935590841681522054612193908263ffffffff61221b16565b600160a060020a038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008083151561220057600091506110fd565b5082820282848281151561221057fe5b0414611ad957600080fd5b600082820183811015611ad957600080fd5b6000808383111561223d57600080fd5b5050900390565b600155565b600160a060020a038116151561225e57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600083600160a060020a03811615156122a557600080fd5b600f5460408051608081018252600160a060020a0388811682526020808301898152838501898152600060608601819052878152600a8452959095208451815473ffffffffffffffffffffffffffffffffffffffff19169416939093178355516001830155925180519496509193909261232692600285019291019061251f565b50606091909101516003909101805460ff1916911515919091179055600f8054600101905560405182907fc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e5190600090a2509392505050565b600160a060020a038216151561239357600080fd5b6004546123a6908263ffffffff61221b16565b600455600160a060020a0382166000908152602081905260409020546123d2908263ffffffff61221b16565b600160a060020a0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b600160a060020a038216151561243d57600080fd5b600160a060020a03821660009081526020819052604090205481111561246257600080fd5b600454612475908263ffffffff61222d16565b600455600160a060020a0382166000908152602081905260409020546124a1908263ffffffff61222d16565b600160a060020a038316600081815260208181526040808320949094558351858152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35050565b81548183558181111561251a5760008381526020902061251a91810190830161259d565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061256057805160ff191683800117855561258d565b8280016001018555821561258d579182015b8281111561258d578251825591602001919060010190612572565b5061259992915061259d565b5090565b610a9091905b8082111561259957600081556001016125a35600a165627a7a72305820bffa6044329c4005991204d5fa011824a3408b8d46a641ee00f98556e9536c350029`

// DeployMyXDC21 deploys a new Ethereum contract, binding an instance of MyXDC21 to it.
func DeployMyXDC21(auth *bind.TransactOpts, backend bind.ContractBackend, _owners []common.Address, _required *big.Int, _name string, _symbol string, _decimals uint8, cap *big.Int, minFee *big.Int, depositFee *big.Int, withdrawFee *big.Int) (common.Address, *types.Transaction, *MyXDC21, error) {
	parsed, err := abi.JSON(strings.NewReader(MyXDC21ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MyXDC21Bin), backend, _owners, _required, _name, _symbol, _decimals, cap, minFee, depositFee, withdrawFee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyXDC21{MyXDC21Caller: MyXDC21Caller{contract: contract}, MyXDC21Transactor: MyXDC21Transactor{contract: contract}, MyXDC21Filterer: MyXDC21Filterer{contract: contract}}, nil
}

// MyXDC21 is an auto generated Go binding around an Ethereum contract.
type MyXDC21 struct {
	MyXDC21Caller     // Read-only binding to the contract
	MyXDC21Transactor // Write-only binding to the contract
	MyXDC21Filterer   // Log filterer for contract events
}

// MyXDC21Caller is an auto generated read-only Go binding around an Ethereum contract.
type MyXDC21Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyXDC21Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MyXDC21Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyXDC21Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyXDC21Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyXDC21Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyXDC21Session struct {
	Contract     *MyXDC21          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyXDC21CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyXDC21CallerSession struct {
	Contract *MyXDC21Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MyXDC21TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyXDC21TransactorSession struct {
	Contract     *MyXDC21Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MyXDC21Raw is an auto generated low-level Go binding around an Ethereum contract.
type MyXDC21Raw struct {
	Contract *MyXDC21 // Generic contract binding to access the raw methods on
}

// MyXDC21CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyXDC21CallerRaw struct {
	Contract *MyXDC21Caller // Generic read-only contract binding to access the raw methods on
}

// MyXDC21TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyXDC21TransactorRaw struct {
	Contract *MyXDC21Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMyXDC21 creates a new instance of MyXDC21, bound to a specific deployed contract.
func NewMyXDC21(address common.Address, backend bind.ContractBackend) (*MyXDC21, error) {
	contract, err := bindMyXDC21(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyXDC21{MyXDC21Caller: MyXDC21Caller{contract: contract}, MyXDC21Transactor: MyXDC21Transactor{contract: contract}, MyXDC21Filterer: MyXDC21Filterer{contract: contract}}, nil
}

// NewMyXDC21Caller creates a new read-only instance of MyXDC21, bound to a specific deployed contract.
func NewMyXDC21Caller(address common.Address, caller bind.ContractCaller) (*MyXDC21Caller, error) {
	contract, err := bindMyXDC21(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyXDC21Caller{contract: contract}, nil
}

// NewMyXDC21Transactor creates a new write-only instance of MyXDC21, bound to a specific deployed contract.
func NewMyXDC21Transactor(address common.Address, transactor bind.ContractTransactor) (*MyXDC21Transactor, error) {
	contract, err := bindMyXDC21(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyXDC21Transactor{contract: contract}, nil
}

// NewMyXDC21Filterer creates a new log filterer instance of MyXDC21, bound to a specific deployed contract.
func NewMyXDC21Filterer(address common.Address, filterer bind.ContractFilterer) (*MyXDC21Filterer, error) {
	contract, err := bindMyXDC21(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyXDC21Filterer{contract: contract}, nil
}

// bindMyXDC21 binds a generic wrapper to an already deployed contract.
func bindMyXDC21(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyXDC21ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyXDC21 *MyXDC21Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyXDC21.Contract.MyXDC21Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyXDC21 *MyXDC21Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyXDC21.Contract.MyXDC21Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyXDC21 *MyXDC21Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyXDC21.Contract.MyXDC21Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyXDC21 *MyXDC21CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyXDC21.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyXDC21 *MyXDC21TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyXDC21.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyXDC21 *MyXDC21TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyXDC21.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITFEE is a free data retrieval call binding the contract method 0xde363e65.
//
// Solidity: function DEPOSIT_FEE() constant returns(uint256)
func (_MyXDC21 *MyXDC21Caller) DEPOSITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "DEPOSIT_FEE")
	return *ret0, err
}

// DEPOSITFEE is a free data retrieval call binding the contract method 0xde363e65.
//
// Solidity: function DEPOSIT_FEE() constant returns(uint256)
func (_MyXDC21 *MyXDC21Session) DEPOSITFEE() (*big.Int, error) {
	return _MyXDC21.Contract.DEPOSITFEE(&_MyXDC21.CallOpts)
}

// DEPOSITFEE is a free data retrieval call binding the contract method 0xde363e65.
//
// Solidity: function DEPOSIT_FEE() constant returns(uint256)
func (_MyXDC21 *MyXDC21CallerSession) DEPOSITFEE() (*big.Int, error) {
	return _MyXDC21.Contract.DEPOSITFEE(&_MyXDC21.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MyXDC21 *MyXDC21Caller) MAXOWNERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "MAX_OWNER_COUNT")
	return *ret0, err
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MyXDC21 *MyXDC21Session) MAXOWNERCOUNT() (*big.Int, error) {
	return _MyXDC21.Contract.MAXOWNERCOUNT(&_MyXDC21.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MyXDC21 *MyXDC21CallerSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MyXDC21.Contract.MAXOWNERCOUNT(&_MyXDC21.CallOpts)
}

// WITHDRAWFEE is a free data retrieval call binding the contract method 0x9bff5ddb.
//
// Solidity: function WITHDRAW_FEE() constant returns(uint256)
func (_MyXDC21 *MyXDC21Caller) WITHDRAWFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "WITHDRAW_FEE")
	return *ret0, err
}

// WITHDRAWFEE is a free data retrieval call binding the contract method 0x9bff5ddb.
//
// Solidity: function WITHDRAW_FEE() constant returns(uint256)
func (_MyXDC21 *MyXDC21Session) WITHDRAWFEE() (*big.Int, error) {
	return _MyXDC21.Contract.WITHDRAWFEE(&_MyXDC21.CallOpts)
}

// WITHDRAWFEE is a free data retrieval call binding the contract method 0x9bff5ddb.
//
// Solidity: function WITHDRAW_FEE() constant returns(uint256)
func (_MyXDC21 *MyXDC21CallerSession) WITHDRAWFEE() (*big.Int, error) {
	return _MyXDC21.Contract.WITHDRAWFEE(&_MyXDC21.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MyXDC21 *MyXDC21Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MyXDC21 *MyXDC21Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MyXDC21.Contract.Allowance(&_MyXDC21.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MyXDC21 *MyXDC21CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MyXDC21.Contract.Allowance(&_MyXDC21.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MyXDC21 *MyXDC21Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MyXDC21 *MyXDC21Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MyXDC21.Contract.BalanceOf(&_MyXDC21.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MyXDC21 *MyXDC21CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MyXDC21.Contract.BalanceOf(&_MyXDC21.CallOpts, owner)
}

// BurnList is a free data retrieval call binding the contract method 0x77661c64.
//
// Solidity: function burnList( uint256) constant returns(value uint256, burner address, data bytes)
func (_MyXDC21 *MyXDC21Caller) BurnList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Value  *big.Int
	Burner common.Address
	Data   []byte
}, error) {
	ret := new(struct {
		Value  *big.Int
		Burner common.Address
		Data   []byte
	})
	out := ret
	err := _MyXDC21.contract.Call(opts, out, "burnList", arg0)
	return *ret, err
}

// BurnList is a free data retrieval call binding the contract method 0x77661c64.
//
// Solidity: function burnList( uint256) constant returns(value uint256, burner address, data bytes)
func (_MyXDC21 *MyXDC21Session) BurnList(arg0 *big.Int) (struct {
	Value  *big.Int
	Burner common.Address
	Data   []byte
}, error) {
	return _MyXDC21.Contract.BurnList(&_MyXDC21.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x77661c64.
//
// Solidity: function burnList( uint256) constant returns(value uint256, burner address, data bytes)
func (_MyXDC21 *MyXDC21CallerSession) BurnList(arg0 *big.Int) (struct {
	Value  *big.Int
	Burner common.Address
	Data   []byte
}, error) {
	return _MyXDC21.Contract.BurnList(&_MyXDC21.CallOpts, arg0)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations( uint256,  address) constant returns(bool)
func (_MyXDC21 *MyXDC21Caller) Confirmations(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "confirmations", arg0, arg1)
	return *ret0, err
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations( uint256,  address) constant returns(bool)
func (_MyXDC21 *MyXDC21Session) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MyXDC21.Contract.Confirmations(&_MyXDC21.CallOpts, arg0, arg1)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations( uint256,  address) constant returns(bool)
func (_MyXDC21 *MyXDC21CallerSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MyXDC21.Contract.Confirmations(&_MyXDC21.CallOpts, arg0, arg1)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyXDC21 *MyXDC21Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyXDC21 *MyXDC21Session) Decimals() (uint8, error) {
	return _MyXDC21.Contract.Decimals(&_MyXDC21.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyXDC21 *MyXDC21CallerSession) Decimals() (uint8, error) {
	return _MyXDC21.Contract.Decimals(&_MyXDC21.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_MyXDC21 *MyXDC21Caller) EstimateFee(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "estimateFee", value)
	return *ret0, err
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_MyXDC21 *MyXDC21Session) EstimateFee(value *big.Int) (*big.Int, error) {
	return _MyXDC21.Contract.EstimateFee(&_MyXDC21.CallOpts, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_MyXDC21 *MyXDC21CallerSession) EstimateFee(value *big.Int) (*big.Int, error) {
	return _MyXDC21.Contract.EstimateFee(&_MyXDC21.CallOpts, value)
}

// GetBurn is a free data retrieval call binding the contract method 0x2eb3f49b.
//
// Solidity: function getBurn(burnId uint256) constant returns(_burner address, _value uint256, _data bytes)
func (_MyXDC21 *MyXDC21Caller) GetBurn(opts *bind.CallOpts, burnId *big.Int) (struct {
	Burner common.Address
	Value  *big.Int
	Data   []byte
}, error) {
	ret := new(struct {
		Burner common.Address
		Value  *big.Int
		Data   []byte
	})
	out := ret
	err := _MyXDC21.contract.Call(opts, out, "getBurn", burnId)
	return *ret, err
}

// GetBurn is a free data retrieval call binding the contract method 0x2eb3f49b.
//
// Solidity: function getBurn(burnId uint256) constant returns(_burner address, _value uint256, _data bytes)
func (_MyXDC21 *MyXDC21Session) GetBurn(burnId *big.Int) (struct {
	Burner common.Address
	Value  *big.Int
	Data   []byte
}, error) {
	return _MyXDC21.Contract.GetBurn(&_MyXDC21.CallOpts, burnId)
}

// GetBurn is a free data retrieval call binding the contract method 0x2eb3f49b.
//
// Solidity: function getBurn(burnId uint256) constant returns(_burner address, _value uint256, _data bytes)
func (_MyXDC21 *MyXDC21CallerSession) GetBurn(burnId *big.Int) (struct {
	Burner common.Address
	Value  *big.Int
	Data   []byte
}, error) {
	return _MyXDC21.Contract.GetBurn(&_MyXDC21.CallOpts, burnId)
}

// GetBurnCount is a free data retrieval call binding the contract method 0xe7cf548c.
//
// Solidity: function getBurnCount() constant returns(uint256)
func (_MyXDC21 *MyXDC21Caller) GetBurnCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "getBurnCount")
	return *ret0, err
}

// GetBurnCount is a free data retrieval call binding the contract method 0xe7cf548c.
//
// Solidity: function getBurnCount() constant returns(uint256)
func (_MyXDC21 *MyXDC21Session) GetBurnCount() (*big.Int, error) {
	return _MyXDC21.Contract.GetBurnCount(&_MyXDC21.CallOpts)
}

// GetBurnCount is a free data retrieval call binding the contract method 0xe7cf548c.
//
// Solidity: function getBurnCount() constant returns(uint256)
func (_MyXDC21 *MyXDC21CallerSession) GetBurnCount() (*big.Int, error) {
	return _MyXDC21.Contract.GetBurnCount(&_MyXDC21.CallOpts)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(transactionId uint256) constant returns(count uint256)
func (_MyXDC21 *MyXDC21Caller) GetConfirmationCount(opts *bind.CallOpts, transactionId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "getConfirmationCount", transactionId)
	return *ret0, err
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(transactionId uint256) constant returns(count uint256)
func (_MyXDC21 *MyXDC21Session) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MyXDC21.Contract.GetConfirmationCount(&_MyXDC21.CallOpts, transactionId)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(transactionId uint256) constant returns(count uint256)
func (_MyXDC21 *MyXDC21CallerSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MyXDC21.Contract.GetConfirmationCount(&_MyXDC21.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(transactionId uint256) constant returns(_confirmations address[])
func (_MyXDC21 *MyXDC21Caller) GetConfirmations(opts *bind.CallOpts, transactionId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "getConfirmations", transactionId)
	return *ret0, err
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(transactionId uint256) constant returns(_confirmations address[])
func (_MyXDC21 *MyXDC21Session) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MyXDC21.Contract.GetConfirmations(&_MyXDC21.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(transactionId uint256) constant returns(_confirmations address[])
func (_MyXDC21 *MyXDC21CallerSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MyXDC21.Contract.GetConfirmations(&_MyXDC21.CallOpts, transactionId)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MyXDC21 *MyXDC21Caller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "getOwners")
	return *ret0, err
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MyXDC21 *MyXDC21Session) GetOwners() ([]common.Address, error) {
	return _MyXDC21.Contract.GetOwners(&_MyXDC21.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MyXDC21 *MyXDC21CallerSession) GetOwners() ([]common.Address, error) {
	return _MyXDC21.Contract.GetOwners(&_MyXDC21.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(pending bool, executed bool) constant returns(count uint256)
func (_MyXDC21 *MyXDC21Caller) GetTransactionCount(opts *bind.CallOpts, pending bool, executed bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "getTransactionCount", pending, executed)
	return *ret0, err
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(pending bool, executed bool) constant returns(count uint256)
func (_MyXDC21 *MyXDC21Session) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MyXDC21.Contract.GetTransactionCount(&_MyXDC21.CallOpts, pending, executed)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(pending bool, executed bool) constant returns(count uint256)
func (_MyXDC21 *MyXDC21CallerSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MyXDC21.Contract.GetTransactionCount(&_MyXDC21.CallOpts, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(from uint256, to uint256, pending bool, executed bool) constant returns(_transactionIds uint256[])
func (_MyXDC21 *MyXDC21Caller) GetTransactionIds(opts *bind.CallOpts, from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "getTransactionIds", from, to, pending, executed)
	return *ret0, err
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(from uint256, to uint256, pending bool, executed bool) constant returns(_transactionIds uint256[])
func (_MyXDC21 *MyXDC21Session) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MyXDC21.Contract.GetTransactionIds(&_MyXDC21.CallOpts, from, to, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(from uint256, to uint256, pending bool, executed bool) constant returns(_transactionIds uint256[])
func (_MyXDC21 *MyXDC21CallerSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MyXDC21.Contract.GetTransactionIds(&_MyXDC21.CallOpts, from, to, pending, executed)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(transactionId uint256) constant returns(bool)
func (_MyXDC21 *MyXDC21Caller) IsConfirmed(opts *bind.CallOpts, transactionId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "isConfirmed", transactionId)
	return *ret0, err
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(transactionId uint256) constant returns(bool)
func (_MyXDC21 *MyXDC21Session) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MyXDC21.Contract.IsConfirmed(&_MyXDC21.CallOpts, transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(transactionId uint256) constant returns(bool)
func (_MyXDC21 *MyXDC21CallerSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MyXDC21.Contract.IsConfirmed(&_MyXDC21.CallOpts, transactionId)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner( address) constant returns(bool)
func (_MyXDC21 *MyXDC21Caller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "isOwner", arg0)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner( address) constant returns(bool)
func (_MyXDC21 *MyXDC21Session) IsOwner(arg0 common.Address) (bool, error) {
	return _MyXDC21.Contract.IsOwner(&_MyXDC21.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner( address) constant returns(bool)
func (_MyXDC21 *MyXDC21CallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MyXDC21.Contract.IsOwner(&_MyXDC21.CallOpts, arg0)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_MyXDC21 *MyXDC21Caller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_MyXDC21 *MyXDC21Session) Issuer() (common.Address, error) {
	return _MyXDC21.Contract.Issuer(&_MyXDC21.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_MyXDC21 *MyXDC21CallerSession) Issuer() (common.Address, error) {
	return _MyXDC21.Contract.Issuer(&_MyXDC21.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_MyXDC21 *MyXDC21Caller) MinFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "minFee")
	return *ret0, err
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_MyXDC21 *MyXDC21Session) MinFee() (*big.Int, error) {
	return _MyXDC21.Contract.MinFee(&_MyXDC21.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_MyXDC21 *MyXDC21CallerSession) MinFee() (*big.Int, error) {
	return _MyXDC21.Contract.MinFee(&_MyXDC21.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyXDC21 *MyXDC21Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyXDC21 *MyXDC21Session) Name() (string, error) {
	return _MyXDC21.Contract.Name(&_MyXDC21.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyXDC21 *MyXDC21CallerSession) Name() (string, error) {
	return _MyXDC21.Contract.Name(&_MyXDC21.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_MyXDC21 *MyXDC21Caller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_MyXDC21 *MyXDC21Session) Owners(arg0 *big.Int) (common.Address, error) {
	return _MyXDC21.Contract.Owners(&_MyXDC21.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_MyXDC21 *MyXDC21CallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MyXDC21.Contract.Owners(&_MyXDC21.CallOpts, arg0)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MyXDC21 *MyXDC21Caller) Required(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "required")
	return *ret0, err
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MyXDC21 *MyXDC21Session) Required() (*big.Int, error) {
	return _MyXDC21.Contract.Required(&_MyXDC21.CallOpts)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MyXDC21 *MyXDC21CallerSession) Required() (*big.Int, error) {
	return _MyXDC21.Contract.Required(&_MyXDC21.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyXDC21 *MyXDC21Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyXDC21 *MyXDC21Session) Symbol() (string, error) {
	return _MyXDC21.Contract.Symbol(&_MyXDC21.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyXDC21 *MyXDC21CallerSession) Symbol() (string, error) {
	return _MyXDC21.Contract.Symbol(&_MyXDC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MyXDC21 *MyXDC21Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MyXDC21 *MyXDC21Session) TotalSupply() (*big.Int, error) {
	return _MyXDC21.Contract.TotalSupply(&_MyXDC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MyXDC21 *MyXDC21CallerSession) TotalSupply() (*big.Int, error) {
	return _MyXDC21.Contract.TotalSupply(&_MyXDC21.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MyXDC21 *MyXDC21Caller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyXDC21.contract.Call(opts, out, "transactionCount")
	return *ret0, err
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MyXDC21 *MyXDC21Session) TransactionCount() (*big.Int, error) {
	return _MyXDC21.Contract.TransactionCount(&_MyXDC21.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MyXDC21 *MyXDC21CallerSession) TransactionCount() (*big.Int, error) {
	return _MyXDC21.Contract.TransactionCount(&_MyXDC21.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions( uint256) constant returns(destination address, value uint256, data bytes, executed bool)
func (_MyXDC21 *MyXDC21Caller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	ret := new(struct {
		Destination common.Address
		Value       *big.Int
		Data        []byte
		Executed    bool
	})
	out := ret
	err := _MyXDC21.contract.Call(opts, out, "transactions", arg0)
	return *ret, err
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions( uint256) constant returns(destination address, value uint256, data bytes, executed bool)
func (_MyXDC21 *MyXDC21Session) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MyXDC21.Contract.Transactions(&_MyXDC21.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions( uint256) constant returns(destination address, value uint256, data bytes, executed bool)
func (_MyXDC21 *MyXDC21CallerSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MyXDC21.Contract.Transactions(&_MyXDC21.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(owner address) returns()
func (_MyXDC21 *MyXDC21Transactor) AddOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MyXDC21.contract.Transact(opts, "addOwner", owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(owner address) returns()
func (_MyXDC21 *MyXDC21Session) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MyXDC21.Contract.AddOwner(&_MyXDC21.TransactOpts, owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(owner address) returns()
func (_MyXDC21 *MyXDC21TransactorSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MyXDC21.Contract.AddOwner(&_MyXDC21.TransactOpts, owner)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MyXDC21 *MyXDC21Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyXDC21.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MyXDC21 *MyXDC21Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.Approve(&_MyXDC21.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MyXDC21 *MyXDC21TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.Approve(&_MyXDC21.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(value uint256, data bytes) returns()
func (_MyXDC21 *MyXDC21Transactor) Burn(opts *bind.TransactOpts, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyXDC21.contract.Transact(opts, "burn", value, data)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(value uint256, data bytes) returns()
func (_MyXDC21 *MyXDC21Session) Burn(value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyXDC21.Contract.Burn(&_MyXDC21.TransactOpts, value, data)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(value uint256, data bytes) returns()
func (_MyXDC21 *MyXDC21TransactorSession) Burn(value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyXDC21.Contract.Burn(&_MyXDC21.TransactOpts, value, data)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(_required uint256) returns()
func (_MyXDC21 *MyXDC21Transactor) ChangeRequirement(opts *bind.TransactOpts, _required *big.Int) (*types.Transaction, error) {
	return _MyXDC21.contract.Transact(opts, "changeRequirement", _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(_required uint256) returns()
func (_MyXDC21 *MyXDC21Session) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.ChangeRequirement(&_MyXDC21.TransactOpts, _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(_required uint256) returns()
func (_MyXDC21 *MyXDC21TransactorSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.ChangeRequirement(&_MyXDC21.TransactOpts, _required)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(transactionId uint256) returns()
func (_MyXDC21 *MyXDC21Transactor) ConfirmTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MyXDC21.contract.Transact(opts, "confirmTransaction", transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(transactionId uint256) returns()
func (_MyXDC21 *MyXDC21Session) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.ConfirmTransaction(&_MyXDC21.TransactOpts, transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(transactionId uint256) returns()
func (_MyXDC21 *MyXDC21TransactorSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.ConfirmTransaction(&_MyXDC21.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(transactionId uint256) returns()
func (_MyXDC21 *MyXDC21Transactor) ExecuteTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MyXDC21.contract.Transact(opts, "executeTransaction", transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(transactionId uint256) returns()
func (_MyXDC21 *MyXDC21Session) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.ExecuteTransaction(&_MyXDC21.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(transactionId uint256) returns()
func (_MyXDC21 *MyXDC21TransactorSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.ExecuteTransaction(&_MyXDC21.TransactOpts, transactionId)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(owner address) returns()
func (_MyXDC21 *MyXDC21Transactor) RemoveOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MyXDC21.contract.Transact(opts, "removeOwner", owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(owner address) returns()
func (_MyXDC21 *MyXDC21Session) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MyXDC21.Contract.RemoveOwner(&_MyXDC21.TransactOpts, owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(owner address) returns()
func (_MyXDC21 *MyXDC21TransactorSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MyXDC21.Contract.RemoveOwner(&_MyXDC21.TransactOpts, owner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(owner address, newOwner address) returns()
func (_MyXDC21 *MyXDC21Transactor) ReplaceOwner(opts *bind.TransactOpts, owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MyXDC21.contract.Transact(opts, "replaceOwner", owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(owner address, newOwner address) returns()
func (_MyXDC21 *MyXDC21Session) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MyXDC21.Contract.ReplaceOwner(&_MyXDC21.TransactOpts, owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(owner address, newOwner address) returns()
func (_MyXDC21 *MyXDC21TransactorSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MyXDC21.Contract.ReplaceOwner(&_MyXDC21.TransactOpts, owner, newOwner)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(transactionId uint256) returns()
func (_MyXDC21 *MyXDC21Transactor) RevokeConfirmation(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MyXDC21.contract.Transact(opts, "revokeConfirmation", transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(transactionId uint256) returns()
func (_MyXDC21 *MyXDC21Session) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.RevokeConfirmation(&_MyXDC21.TransactOpts, transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(transactionId uint256) returns()
func (_MyXDC21 *MyXDC21TransactorSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.RevokeConfirmation(&_MyXDC21.TransactOpts, transactionId)
}

// SetDepositFee is a paid mutator transaction binding the contract method 0x490ae210.
//
// Solidity: function setDepositFee(depositFee uint256) returns()
func (_MyXDC21 *MyXDC21Transactor) SetDepositFee(opts *bind.TransactOpts, depositFee *big.Int) (*types.Transaction, error) {
	return _MyXDC21.contract.Transact(opts, "setDepositFee", depositFee)
}

// SetDepositFee is a paid mutator transaction binding the contract method 0x490ae210.
//
// Solidity: function setDepositFee(depositFee uint256) returns()
func (_MyXDC21 *MyXDC21Session) SetDepositFee(depositFee *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.SetDepositFee(&_MyXDC21.TransactOpts, depositFee)
}

// SetDepositFee is a paid mutator transaction binding the contract method 0x490ae210.
//
// Solidity: function setDepositFee(depositFee uint256) returns()
func (_MyXDC21 *MyXDC21TransactorSession) SetDepositFee(depositFee *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.SetDepositFee(&_MyXDC21.TransactOpts, depositFee)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x31ac9920.
//
// Solidity: function setMinFee(value uint256) returns()
func (_MyXDC21 *MyXDC21Transactor) SetMinFee(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _MyXDC21.contract.Transact(opts, "setMinFee", value)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x31ac9920.
//
// Solidity: function setMinFee(value uint256) returns()
func (_MyXDC21 *MyXDC21Session) SetMinFee(value *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.SetMinFee(&_MyXDC21.TransactOpts, value)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x31ac9920.
//
// Solidity: function setMinFee(value uint256) returns()
func (_MyXDC21 *MyXDC21TransactorSession) SetMinFee(value *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.SetMinFee(&_MyXDC21.TransactOpts, value)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(withdrawFee uint256) returns()
func (_MyXDC21 *MyXDC21Transactor) SetWithdrawFee(opts *bind.TransactOpts, withdrawFee *big.Int) (*types.Transaction, error) {
	return _MyXDC21.contract.Transact(opts, "setWithdrawFee", withdrawFee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(withdrawFee uint256) returns()
func (_MyXDC21 *MyXDC21Session) SetWithdrawFee(withdrawFee *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.SetWithdrawFee(&_MyXDC21.TransactOpts, withdrawFee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(withdrawFee uint256) returns()
func (_MyXDC21 *MyXDC21TransactorSession) SetWithdrawFee(withdrawFee *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.SetWithdrawFee(&_MyXDC21.TransactOpts, withdrawFee)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(destination address, value uint256, data bytes) returns(transactionId uint256)
func (_MyXDC21 *MyXDC21Transactor) SubmitTransaction(opts *bind.TransactOpts, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyXDC21.contract.Transact(opts, "submitTransaction", destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(destination address, value uint256, data bytes) returns(transactionId uint256)
func (_MyXDC21 *MyXDC21Session) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyXDC21.Contract.SubmitTransaction(&_MyXDC21.TransactOpts, destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(destination address, value uint256, data bytes) returns(transactionId uint256)
func (_MyXDC21 *MyXDC21TransactorSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyXDC21.Contract.SubmitTransaction(&_MyXDC21.TransactOpts, destination, value, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MyXDC21 *MyXDC21Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyXDC21.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MyXDC21 *MyXDC21Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.Transfer(&_MyXDC21.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MyXDC21 *MyXDC21TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.Transfer(&_MyXDC21.TransactOpts, to, value)
}

// TransferContractIssuer is a paid mutator transaction binding the contract method 0xc28ce6ff.
//
// Solidity: function transferContractIssuer(newOwner address) returns()
func (_MyXDC21 *MyXDC21Transactor) TransferContractIssuer(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MyXDC21.contract.Transact(opts, "transferContractIssuer", newOwner)
}

// TransferContractIssuer is a paid mutator transaction binding the contract method 0xc28ce6ff.
//
// Solidity: function transferContractIssuer(newOwner address) returns()
func (_MyXDC21 *MyXDC21Session) TransferContractIssuer(newOwner common.Address) (*types.Transaction, error) {
	return _MyXDC21.Contract.TransferContractIssuer(&_MyXDC21.TransactOpts, newOwner)
}

// TransferContractIssuer is a paid mutator transaction binding the contract method 0xc28ce6ff.
//
// Solidity: function transferContractIssuer(newOwner address) returns()
func (_MyXDC21 *MyXDC21TransactorSession) TransferContractIssuer(newOwner common.Address) (*types.Transaction, error) {
	return _MyXDC21.Contract.TransferContractIssuer(&_MyXDC21.TransactOpts, newOwner)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MyXDC21 *MyXDC21Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyXDC21.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MyXDC21 *MyXDC21Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.TransferFrom(&_MyXDC21.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MyXDC21 *MyXDC21TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyXDC21.Contract.TransferFrom(&_MyXDC21.TransactOpts, from, to, value)
}

// MyXDC21ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MyXDC21 contract.
type MyXDC21ApprovalIterator struct {
	Event *MyXDC21Approval // Event containing the contract specifics and raw log

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
func (it *MyXDC21ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXDC21Approval)
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
		it.Event = new(MyXDC21Approval)
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
func (it *MyXDC21ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXDC21ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXDC21Approval represents a Approval event raised by the MyXDC21 contract.
type MyXDC21Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_MyXDC21 *MyXDC21Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MyXDC21ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MyXDC21.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MyXDC21ApprovalIterator{contract: _MyXDC21.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_MyXDC21 *MyXDC21Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MyXDC21Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MyXDC21.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXDC21Approval)
				if err := _MyXDC21.contract.UnpackLog(event, "Approval", log); err != nil {
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

// MyXDC21ConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the MyXDC21 contract.
type MyXDC21ConfirmationIterator struct {
	Event *MyXDC21Confirmation // Event containing the contract specifics and raw log

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
func (it *MyXDC21ConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXDC21Confirmation)
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
		it.Event = new(MyXDC21Confirmation)
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
func (it *MyXDC21ConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXDC21ConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXDC21Confirmation represents a Confirmation event raised by the MyXDC21 contract.
type MyXDC21Confirmation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(sender indexed address, transactionId indexed uint256)
func (_MyXDC21 *MyXDC21Filterer) FilterConfirmation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MyXDC21ConfirmationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyXDC21.contract.FilterLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MyXDC21ConfirmationIterator{contract: _MyXDC21.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(sender indexed address, transactionId indexed uint256)
func (_MyXDC21 *MyXDC21Filterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *MyXDC21Confirmation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyXDC21.contract.WatchLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXDC21Confirmation)
				if err := _MyXDC21.contract.UnpackLog(event, "Confirmation", log); err != nil {
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

// MyXDC21ExecutionIterator is returned from FilterExecution and is used to iterate over the raw logs and unpacked data for Execution events raised by the MyXDC21 contract.
type MyXDC21ExecutionIterator struct {
	Event *MyXDC21Execution // Event containing the contract specifics and raw log

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
func (it *MyXDC21ExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXDC21Execution)
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
		it.Event = new(MyXDC21Execution)
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
func (it *MyXDC21ExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXDC21ExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXDC21Execution represents a Execution event raised by the MyXDC21 contract.
type MyXDC21Execution struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecution is a free log retrieval operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(transactionId indexed uint256)
func (_MyXDC21 *MyXDC21Filterer) FilterExecution(opts *bind.FilterOpts, transactionId []*big.Int) (*MyXDC21ExecutionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyXDC21.contract.FilterLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MyXDC21ExecutionIterator{contract: _MyXDC21.contract, event: "Execution", logs: logs, sub: sub}, nil
}

// WatchExecution is a free log subscription operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(transactionId indexed uint256)
func (_MyXDC21 *MyXDC21Filterer) WatchExecution(opts *bind.WatchOpts, sink chan<- *MyXDC21Execution, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyXDC21.contract.WatchLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXDC21Execution)
				if err := _MyXDC21.contract.UnpackLog(event, "Execution", log); err != nil {
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

// MyXDC21ExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the MyXDC21 contract.
type MyXDC21ExecutionFailureIterator struct {
	Event *MyXDC21ExecutionFailure // Event containing the contract specifics and raw log

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
func (it *MyXDC21ExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXDC21ExecutionFailure)
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
		it.Event = new(MyXDC21ExecutionFailure)
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
func (it *MyXDC21ExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXDC21ExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXDC21ExecutionFailure represents a ExecutionFailure event raised by the MyXDC21 contract.
type MyXDC21ExecutionFailure struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(transactionId indexed uint256)
func (_MyXDC21 *MyXDC21Filterer) FilterExecutionFailure(opts *bind.FilterOpts, transactionId []*big.Int) (*MyXDC21ExecutionFailureIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyXDC21.contract.FilterLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MyXDC21ExecutionFailureIterator{contract: _MyXDC21.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(transactionId indexed uint256)
func (_MyXDC21 *MyXDC21Filterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *MyXDC21ExecutionFailure, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyXDC21.contract.WatchLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXDC21ExecutionFailure)
				if err := _MyXDC21.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
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

// MyXDC21FeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the MyXDC21 contract.
type MyXDC21FeeIterator struct {
	Event *MyXDC21Fee // Event containing the contract specifics and raw log

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
func (it *MyXDC21FeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXDC21Fee)
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
		it.Event = new(MyXDC21Fee)
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
func (it *MyXDC21FeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXDC21FeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXDC21Fee represents a Fee event raised by the MyXDC21 contract.
type MyXDC21Fee struct {
	From   common.Address
	To     common.Address
	Issuer common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_MyXDC21 *MyXDC21Filterer) FilterFee(opts *bind.FilterOpts, from []common.Address, to []common.Address, issuer []common.Address) (*MyXDC21FeeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _MyXDC21.contract.FilterLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &MyXDC21FeeIterator{contract: _MyXDC21.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_MyXDC21 *MyXDC21Filterer) WatchFee(opts *bind.WatchOpts, sink chan<- *MyXDC21Fee, from []common.Address, to []common.Address, issuer []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _MyXDC21.contract.WatchLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXDC21Fee)
				if err := _MyXDC21.contract.UnpackLog(event, "Fee", log); err != nil {
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

// MyXDC21OwnerAdditionIterator is returned from FilterOwnerAddition and is used to iterate over the raw logs and unpacked data for OwnerAddition events raised by the MyXDC21 contract.
type MyXDC21OwnerAdditionIterator struct {
	Event *MyXDC21OwnerAddition // Event containing the contract specifics and raw log

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
func (it *MyXDC21OwnerAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXDC21OwnerAddition)
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
		it.Event = new(MyXDC21OwnerAddition)
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
func (it *MyXDC21OwnerAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXDC21OwnerAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXDC21OwnerAddition represents a OwnerAddition event raised by the MyXDC21 contract.
type MyXDC21OwnerAddition struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerAddition is a free log retrieval operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(owner indexed address)
func (_MyXDC21 *MyXDC21Filterer) FilterOwnerAddition(opts *bind.FilterOpts, owner []common.Address) (*MyXDC21OwnerAdditionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MyXDC21.contract.FilterLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MyXDC21OwnerAdditionIterator{contract: _MyXDC21.contract, event: "OwnerAddition", logs: logs, sub: sub}, nil
}

// WatchOwnerAddition is a free log subscription operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(owner indexed address)
func (_MyXDC21 *MyXDC21Filterer) WatchOwnerAddition(opts *bind.WatchOpts, sink chan<- *MyXDC21OwnerAddition, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MyXDC21.contract.WatchLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXDC21OwnerAddition)
				if err := _MyXDC21.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
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

// MyXDC21OwnerRemovalIterator is returned from FilterOwnerRemoval and is used to iterate over the raw logs and unpacked data for OwnerRemoval events raised by the MyXDC21 contract.
type MyXDC21OwnerRemovalIterator struct {
	Event *MyXDC21OwnerRemoval // Event containing the contract specifics and raw log

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
func (it *MyXDC21OwnerRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXDC21OwnerRemoval)
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
		it.Event = new(MyXDC21OwnerRemoval)
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
func (it *MyXDC21OwnerRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXDC21OwnerRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXDC21OwnerRemoval represents a OwnerRemoval event raised by the MyXDC21 contract.
type MyXDC21OwnerRemoval struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoval is a free log retrieval operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(owner indexed address)
func (_MyXDC21 *MyXDC21Filterer) FilterOwnerRemoval(opts *bind.FilterOpts, owner []common.Address) (*MyXDC21OwnerRemovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MyXDC21.contract.FilterLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MyXDC21OwnerRemovalIterator{contract: _MyXDC21.contract, event: "OwnerRemoval", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoval is a free log subscription operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(owner indexed address)
func (_MyXDC21 *MyXDC21Filterer) WatchOwnerRemoval(opts *bind.WatchOpts, sink chan<- *MyXDC21OwnerRemoval, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MyXDC21.contract.WatchLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXDC21OwnerRemoval)
				if err := _MyXDC21.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
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

// MyXDC21RequirementChangeIterator is returned from FilterRequirementChange and is used to iterate over the raw logs and unpacked data for RequirementChange events raised by the MyXDC21 contract.
type MyXDC21RequirementChangeIterator struct {
	Event *MyXDC21RequirementChange // Event containing the contract specifics and raw log

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
func (it *MyXDC21RequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXDC21RequirementChange)
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
		it.Event = new(MyXDC21RequirementChange)
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
func (it *MyXDC21RequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXDC21RequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXDC21RequirementChange represents a RequirementChange event raised by the MyXDC21 contract.
type MyXDC21RequirementChange struct {
	Required *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequirementChange is a free log retrieval operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(required uint256)
func (_MyXDC21 *MyXDC21Filterer) FilterRequirementChange(opts *bind.FilterOpts) (*MyXDC21RequirementChangeIterator, error) {

	logs, sub, err := _MyXDC21.contract.FilterLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return &MyXDC21RequirementChangeIterator{contract: _MyXDC21.contract, event: "RequirementChange", logs: logs, sub: sub}, nil
}

// WatchRequirementChange is a free log subscription operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(required uint256)
func (_MyXDC21 *MyXDC21Filterer) WatchRequirementChange(opts *bind.WatchOpts, sink chan<- *MyXDC21RequirementChange) (event.Subscription, error) {

	logs, sub, err := _MyXDC21.contract.WatchLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXDC21RequirementChange)
				if err := _MyXDC21.contract.UnpackLog(event, "RequirementChange", log); err != nil {
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

// MyXDC21RevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the MyXDC21 contract.
type MyXDC21RevocationIterator struct {
	Event *MyXDC21Revocation // Event containing the contract specifics and raw log

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
func (it *MyXDC21RevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXDC21Revocation)
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
		it.Event = new(MyXDC21Revocation)
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
func (it *MyXDC21RevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXDC21RevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXDC21Revocation represents a Revocation event raised by the MyXDC21 contract.
type MyXDC21Revocation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(sender indexed address, transactionId indexed uint256)
func (_MyXDC21 *MyXDC21Filterer) FilterRevocation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MyXDC21RevocationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyXDC21.contract.FilterLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MyXDC21RevocationIterator{contract: _MyXDC21.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(sender indexed address, transactionId indexed uint256)
func (_MyXDC21 *MyXDC21Filterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *MyXDC21Revocation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyXDC21.contract.WatchLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXDC21Revocation)
				if err := _MyXDC21.contract.UnpackLog(event, "Revocation", log); err != nil {
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

// MyXDC21SubmissionIterator is returned from FilterSubmission and is used to iterate over the raw logs and unpacked data for Submission events raised by the MyXDC21 contract.
type MyXDC21SubmissionIterator struct {
	Event *MyXDC21Submission // Event containing the contract specifics and raw log

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
func (it *MyXDC21SubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXDC21Submission)
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
		it.Event = new(MyXDC21Submission)
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
func (it *MyXDC21SubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXDC21SubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXDC21Submission represents a Submission event raised by the MyXDC21 contract.
type MyXDC21Submission struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmission is a free log retrieval operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(transactionId indexed uint256)
func (_MyXDC21 *MyXDC21Filterer) FilterSubmission(opts *bind.FilterOpts, transactionId []*big.Int) (*MyXDC21SubmissionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyXDC21.contract.FilterLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MyXDC21SubmissionIterator{contract: _MyXDC21.contract, event: "Submission", logs: logs, sub: sub}, nil
}

// WatchSubmission is a free log subscription operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(transactionId indexed uint256)
func (_MyXDC21 *MyXDC21Filterer) WatchSubmission(opts *bind.WatchOpts, sink chan<- *MyXDC21Submission, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyXDC21.contract.WatchLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXDC21Submission)
				if err := _MyXDC21.contract.UnpackLog(event, "Submission", log); err != nil {
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

// MyXDC21TokenBurnIterator is returned from FilterTokenBurn and is used to iterate over the raw logs and unpacked data for TokenBurn events raised by the MyXDC21 contract.
type MyXDC21TokenBurnIterator struct {
	Event *MyXDC21TokenBurn // Event containing the contract specifics and raw log

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
func (it *MyXDC21TokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXDC21TokenBurn)
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
		it.Event = new(MyXDC21TokenBurn)
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
func (it *MyXDC21TokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXDC21TokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXDC21TokenBurn represents a TokenBurn event raised by the MyXDC21 contract.
type MyXDC21TokenBurn struct {
	BurnID *big.Int
	Burner common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenBurn is a free log retrieval operation binding the contract event 0x6905852b196f81e7e03058512a599446c358027fc943c1e193b6649a39379bb5.
//
// Solidity: event TokenBurn(burnID indexed uint256, burner indexed address, value uint256, data bytes)
func (_MyXDC21 *MyXDC21Filterer) FilterTokenBurn(opts *bind.FilterOpts, burnID []*big.Int, burner []common.Address) (*MyXDC21TokenBurnIterator, error) {

	var burnIDRule []interface{}
	for _, burnIDItem := range burnID {
		burnIDRule = append(burnIDRule, burnIDItem)
	}
	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _MyXDC21.contract.FilterLogs(opts, "TokenBurn", burnIDRule, burnerRule)
	if err != nil {
		return nil, err
	}
	return &MyXDC21TokenBurnIterator{contract: _MyXDC21.contract, event: "TokenBurn", logs: logs, sub: sub}, nil
}

// WatchTokenBurn is a free log subscription operation binding the contract event 0x6905852b196f81e7e03058512a599446c358027fc943c1e193b6649a39379bb5.
//
// Solidity: event TokenBurn(burnID indexed uint256, burner indexed address, value uint256, data bytes)
func (_MyXDC21 *MyXDC21Filterer) WatchTokenBurn(opts *bind.WatchOpts, sink chan<- *MyXDC21TokenBurn, burnID []*big.Int, burner []common.Address) (event.Subscription, error) {

	var burnIDRule []interface{}
	for _, burnIDItem := range burnID {
		burnIDRule = append(burnIDRule, burnIDItem)
	}
	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _MyXDC21.contract.WatchLogs(opts, "TokenBurn", burnIDRule, burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXDC21TokenBurn)
				if err := _MyXDC21.contract.UnpackLog(event, "TokenBurn", log); err != nil {
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

// MyXDC21TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MyXDC21 contract.
type MyXDC21TransferIterator struct {
	Event *MyXDC21Transfer // Event containing the contract specifics and raw log

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
func (it *MyXDC21TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyXDC21Transfer)
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
		it.Event = new(MyXDC21Transfer)
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
func (it *MyXDC21TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyXDC21TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyXDC21Transfer represents a Transfer event raised by the MyXDC21 contract.
type MyXDC21Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_MyXDC21 *MyXDC21Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MyXDC21TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyXDC21.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MyXDC21TransferIterator{contract: _MyXDC21.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_MyXDC21 *MyXDC21Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyXDC21Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyXDC21.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyXDC21Transfer)
				if err := _MyXDC21.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// XDC21ABI is the input ABI used to generate the binding from.
const XDC21ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setMinFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Fee\",\"type\":\"event\"}]"

// XDC21Bin is the compiled bytecode used for deploying new contracts.
const XDC21Bin = `0x608060405234801561001057600080fd5b50604051610a2c380380610a2c83398101604090815281516020808401519284015191840180519094939093019261004e916005919086019061007f565b50815161006290600690602085019061007f565b506007805460ff191660ff929092169190911790555061011a9050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100c057805160ff19168380011785556100ed565b828001600101855582156100ed579182015b828111156100ed5782518255916020019190600101906100d2565b506100f99291506100fd565b5090565b61011791905b808211156100f95760008155600101610103565b90565b610903806101296000396000f3006080604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100c9578063095ea7b314610153578063127e8e4d1461018b57806318160ddd146101b55780631d143848146101ca57806323b872dd146101fb57806324ec759014610225578063313ce5671461023a57806331ac99201461026557806370a082311461027f57806395d89b41146102a0578063a9059cbb146102b5578063dd62ed3e146102d9575b600080fd5b3480156100d557600080fd5b506100de610300565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610118578181015183820152602001610100565b50505050905090810190601f1680156101455780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015f57600080fd5b50610177600160a060020a0360043516602435610396565b604080519115158252519081900360200190f35b34801561019757600080fd5b506101a3600435610450565b60408051918252519081900360200190f35b3480156101c157600080fd5b506101a361047c565b3480156101d657600080fd5b506101df610482565b60408051600160a060020a039092168252519081900360200190f35b34801561020757600080fd5b50610177600160a060020a0360043581169060243516604435610491565b34801561023157600080fd5b506101a36105d5565b34801561024657600080fd5b5061024f6105db565b6040805160ff9092168252519081900360200190f35b34801561027157600080fd5b5061027d6004356105e4565b005b34801561028b57600080fd5b506101a3600160a060020a0360043516610607565b3480156102ac57600080fd5b506100de610622565b3480156102c157600080fd5b50610177600160a060020a0360043516602435610683565b3480156102e557600080fd5b506101a3600160a060020a0360043581169060243516610757565b60058054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561038c5780601f106103615761010080835404028352916020019161038c565b820191906000526020600020905b81548152906001019060200180831161036f57829003601f168201915b5050505050905090565b6000600160a060020a03831615156103ad57600080fd5b6001543360009081526020819052604090205410156103cb57600080fd5b336000818152600360209081526040808320600160a060020a038881168552925290912084905560025460015461040793929190911690610782565b604080518381529051600160a060020a0385169133917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360200190a350600192915050565b6001546000906104769061046a848463ffffffff61087416565b9063ffffffff6108a916565b92915050565b60045490565b600254600160a060020a031690565b6000806104a9600154846108a990919063ffffffff16565b600160a060020a0386166000908152602081905260409020549091508111156104d157600080fd5b600160a060020a038516600090815260036020908152604080832033845290915290205483111561050157600080fd5b600160a060020a0385166000908152600360209081526040808320338452909152902054610535908263ffffffff6108bb16565b600160a060020a0386166000908152600360209081526040808320338452909152902055610564858585610782565b600254600154610581918791600160a060020a0390911690610782565b6002546001546040805191825251600160a060020a039283169287169133917ffcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd999181900360200190a4506001949350505050565b60015490565b60075460ff1690565b600254600160a060020a031633146105fb57600080fd5b610604816108d2565b50565b600160a060020a031660009081526020819052604090205490565b60068054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561038c5780601f106103615761010080835404028352916020019161038c565b600080600160a060020a038416151561069b57600080fd5b6001546106af90849063ffffffff6108a916565b336000908152602081905260409020549091508111156106ce57600080fd5b6106d9338585610782565b6000600154111561074b57600254600154610701913391600160a060020a0390911690610782565b6002546001546040805191825251600160a060020a039283169287169133917ffcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd999181900360200190a45b600191505b5092915050565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b600160a060020a0383166000908152602081905260409020548111156107a757600080fd5b600160a060020a03821615156107bc57600080fd5b600160a060020a0383166000908152602081905260409020546107e5908263ffffffff6108bb16565b600160a060020a03808516600090815260208190526040808220939093559084168152205461081a908263ffffffff6108a916565b600160a060020a038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000808315156108875760009150610750565b5082820282848281151561089757fe5b04146108a257600080fd5b9392505050565b6000828201838110156108a257600080fd5b600080838311156108cb57600080fd5b5050900390565b6001555600a165627a7a723058206c24d49045155ad83d904de717409715bdb93edbc379076a4a1d5f03d3ae00900029`

// DeployXDC21 deploys a new Ethereum contract, binding an instance of XDC21 to it.
func DeployXDC21(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals uint8) (common.Address, *types.Transaction, *XDC21, error) {
	parsed, err := abi.JSON(strings.NewReader(XDC21ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(XDC21Bin), backend, name, symbol, decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XDC21{XDC21Caller: XDC21Caller{contract: contract}, XDC21Transactor: XDC21Transactor{contract: contract}, XDC21Filterer: XDC21Filterer{contract: contract}}, nil
}

// XDC21 is an auto generated Go binding around an Ethereum contract.
type XDC21 struct {
	XDC21Caller     // Read-only binding to the contract
	XDC21Transactor // Write-only binding to the contract
	XDC21Filterer   // Log filterer for contract events
}

// XDC21Caller is an auto generated read-only Go binding around an Ethereum contract.
type XDC21Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XDC21Transactor is an auto generated write-only Go binding around an Ethereum contract.
type XDC21Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XDC21Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XDC21Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XDC21Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XDC21Session struct {
	Contract     *XDC21            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XDC21CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XDC21CallerSession struct {
	Contract *XDC21Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// XDC21TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XDC21TransactorSession struct {
	Contract     *XDC21Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XDC21Raw is an auto generated low-level Go binding around an Ethereum contract.
type XDC21Raw struct {
	Contract *XDC21 // Generic contract binding to access the raw methods on
}

// XDC21CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XDC21CallerRaw struct {
	Contract *XDC21Caller // Generic read-only contract binding to access the raw methods on
}

// XDC21TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XDC21TransactorRaw struct {
	Contract *XDC21Transactor // Generic write-only contract binding to access the raw methods on
}

// NewXDC21 creates a new instance of XDC21, bound to a specific deployed contract.
func NewXDC21(address common.Address, backend bind.ContractBackend) (*XDC21, error) {
	contract, err := bindXDC21(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XDC21{XDC21Caller: XDC21Caller{contract: contract}, XDC21Transactor: XDC21Transactor{contract: contract}, XDC21Filterer: XDC21Filterer{contract: contract}}, nil
}

// NewXDC21Caller creates a new read-only instance of XDC21, bound to a specific deployed contract.
func NewXDC21Caller(address common.Address, caller bind.ContractCaller) (*XDC21Caller, error) {
	contract, err := bindXDC21(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XDC21Caller{contract: contract}, nil
}

// NewXDC21Transactor creates a new write-only instance of XDC21, bound to a specific deployed contract.
func NewXDC21Transactor(address common.Address, transactor bind.ContractTransactor) (*XDC21Transactor, error) {
	contract, err := bindXDC21(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XDC21Transactor{contract: contract}, nil
}

// NewXDC21Filterer creates a new log filterer instance of XDC21, bound to a specific deployed contract.
func NewXDC21Filterer(address common.Address, filterer bind.ContractFilterer) (*XDC21Filterer, error) {
	contract, err := bindXDC21(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XDC21Filterer{contract: contract}, nil
}

// bindXDC21 binds a generic wrapper to an already deployed contract.
func bindXDC21(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XDC21ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XDC21 *XDC21Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XDC21.Contract.XDC21Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XDC21 *XDC21Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XDC21.Contract.XDC21Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XDC21 *XDC21Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XDC21.Contract.XDC21Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XDC21 *XDC21CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XDC21.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XDC21 *XDC21TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XDC21.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XDC21 *XDC21TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XDC21.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_XDC21 *XDC21Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XDC21.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_XDC21 *XDC21Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _XDC21.Contract.Allowance(&_XDC21.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_XDC21 *XDC21CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _XDC21.Contract.Allowance(&_XDC21.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_XDC21 *XDC21Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XDC21.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_XDC21 *XDC21Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _XDC21.Contract.BalanceOf(&_XDC21.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_XDC21 *XDC21CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _XDC21.Contract.BalanceOf(&_XDC21.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_XDC21 *XDC21Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _XDC21.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_XDC21 *XDC21Session) Decimals() (uint8, error) {
	return _XDC21.Contract.Decimals(&_XDC21.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_XDC21 *XDC21CallerSession) Decimals() (uint8, error) {
	return _XDC21.Contract.Decimals(&_XDC21.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_XDC21 *XDC21Caller) EstimateFee(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XDC21.contract.Call(opts, out, "estimateFee", value)
	return *ret0, err
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_XDC21 *XDC21Session) EstimateFee(value *big.Int) (*big.Int, error) {
	return _XDC21.Contract.EstimateFee(&_XDC21.CallOpts, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_XDC21 *XDC21CallerSession) EstimateFee(value *big.Int) (*big.Int, error) {
	return _XDC21.Contract.EstimateFee(&_XDC21.CallOpts, value)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_XDC21 *XDC21Caller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _XDC21.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_XDC21 *XDC21Session) Issuer() (common.Address, error) {
	return _XDC21.Contract.Issuer(&_XDC21.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_XDC21 *XDC21CallerSession) Issuer() (common.Address, error) {
	return _XDC21.Contract.Issuer(&_XDC21.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_XDC21 *XDC21Caller) MinFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XDC21.contract.Call(opts, out, "minFee")
	return *ret0, err
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_XDC21 *XDC21Session) MinFee() (*big.Int, error) {
	return _XDC21.Contract.MinFee(&_XDC21.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_XDC21 *XDC21CallerSession) MinFee() (*big.Int, error) {
	return _XDC21.Contract.MinFee(&_XDC21.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_XDC21 *XDC21Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _XDC21.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_XDC21 *XDC21Session) Name() (string, error) {
	return _XDC21.Contract.Name(&_XDC21.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_XDC21 *XDC21CallerSession) Name() (string, error) {
	return _XDC21.Contract.Name(&_XDC21.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_XDC21 *XDC21Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _XDC21.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_XDC21 *XDC21Session) Symbol() (string, error) {
	return _XDC21.Contract.Symbol(&_XDC21.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_XDC21 *XDC21CallerSession) Symbol() (string, error) {
	return _XDC21.Contract.Symbol(&_XDC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_XDC21 *XDC21Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XDC21.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_XDC21 *XDC21Session) TotalSupply() (*big.Int, error) {
	return _XDC21.Contract.TotalSupply(&_XDC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_XDC21 *XDC21CallerSession) TotalSupply() (*big.Int, error) {
	return _XDC21.Contract.TotalSupply(&_XDC21.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_XDC21 *XDC21Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _XDC21.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_XDC21 *XDC21Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _XDC21.Contract.Approve(&_XDC21.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_XDC21 *XDC21TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _XDC21.Contract.Approve(&_XDC21.TransactOpts, spender, value)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x31ac9920.
//
// Solidity: function setMinFee(value uint256) returns()
func (_XDC21 *XDC21Transactor) SetMinFee(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _XDC21.contract.Transact(opts, "setMinFee", value)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x31ac9920.
//
// Solidity: function setMinFee(value uint256) returns()
func (_XDC21 *XDC21Session) SetMinFee(value *big.Int) (*types.Transaction, error) {
	return _XDC21.Contract.SetMinFee(&_XDC21.TransactOpts, value)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x31ac9920.
//
// Solidity: function setMinFee(value uint256) returns()
func (_XDC21 *XDC21TransactorSession) SetMinFee(value *big.Int) (*types.Transaction, error) {
	return _XDC21.Contract.SetMinFee(&_XDC21.TransactOpts, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_XDC21 *XDC21Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _XDC21.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_XDC21 *XDC21Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _XDC21.Contract.Transfer(&_XDC21.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_XDC21 *XDC21TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _XDC21.Contract.Transfer(&_XDC21.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_XDC21 *XDC21Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _XDC21.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_XDC21 *XDC21Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _XDC21.Contract.TransferFrom(&_XDC21.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_XDC21 *XDC21TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _XDC21.Contract.TransferFrom(&_XDC21.TransactOpts, from, to, value)
}

// XDC21ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the XDC21 contract.
type XDC21ApprovalIterator struct {
	Event *XDC21Approval // Event containing the contract specifics and raw log

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
func (it *XDC21ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XDC21Approval)
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
		it.Event = new(XDC21Approval)
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
func (it *XDC21ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XDC21ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XDC21Approval represents a Approval event raised by the XDC21 contract.
type XDC21Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_XDC21 *XDC21Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*XDC21ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _XDC21.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &XDC21ApprovalIterator{contract: _XDC21.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_XDC21 *XDC21Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *XDC21Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _XDC21.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XDC21Approval)
				if err := _XDC21.contract.UnpackLog(event, "Approval", log); err != nil {
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

// XDC21FeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the XDC21 contract.
type XDC21FeeIterator struct {
	Event *XDC21Fee // Event containing the contract specifics and raw log

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
func (it *XDC21FeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XDC21Fee)
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
		it.Event = new(XDC21Fee)
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
func (it *XDC21FeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XDC21FeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XDC21Fee represents a Fee event raised by the XDC21 contract.
type XDC21Fee struct {
	From   common.Address
	To     common.Address
	Issuer common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_XDC21 *XDC21Filterer) FilterFee(opts *bind.FilterOpts, from []common.Address, to []common.Address, issuer []common.Address) (*XDC21FeeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _XDC21.contract.FilterLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &XDC21FeeIterator{contract: _XDC21.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_XDC21 *XDC21Filterer) WatchFee(opts *bind.WatchOpts, sink chan<- *XDC21Fee, from []common.Address, to []common.Address, issuer []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _XDC21.contract.WatchLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XDC21Fee)
				if err := _XDC21.contract.UnpackLog(event, "Fee", log); err != nil {
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

// XDC21TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the XDC21 contract.
type XDC21TransferIterator struct {
	Event *XDC21Transfer // Event containing the contract specifics and raw log

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
func (it *XDC21TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XDC21Transfer)
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
		it.Event = new(XDC21Transfer)
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
func (it *XDC21TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XDC21TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XDC21Transfer represents a Transfer event raised by the XDC21 contract.
type XDC21Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_XDC21 *XDC21Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*XDC21TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _XDC21.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &XDC21TransferIterator{contract: _XDC21.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_XDC21 *XDC21Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *XDC21Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _XDC21.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XDC21Transfer)
				if err := _XDC21.contract.UnpackLog(event, "Transfer", log); err != nil {
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
