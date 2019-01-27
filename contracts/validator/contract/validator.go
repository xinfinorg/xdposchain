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

// IValidatorABI is the input ABI used to generate the binding from.
const IValidatorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// IValidatorBin is the compiled bytecode used for deploying new contracts.
const IValidatorBin = `0x`

// DeployIValidator deploys a new Ethereum contract, binding an instance of IValidator to it.
func DeployIValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IValidator, error) {
	parsed, err := abi.JSON(strings.NewReader(IValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IValidatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IValidator{IValidatorCaller: IValidatorCaller{contract: contract}, IValidatorTransactor: IValidatorTransactor{contract: contract}, IValidatorFilterer: IValidatorFilterer{contract: contract}}, nil
}

// IValidator is an auto generated Go binding around an Ethereum contract.
type IValidator struct {
	IValidatorCaller     // Read-only binding to the contract
	IValidatorTransactor // Write-only binding to the contract
	IValidatorFilterer   // Log filterer for contract events
}

// IValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IValidatorSession struct {
	Contract     *IValidator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IValidatorCallerSession struct {
	Contract *IValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IValidatorTransactorSession struct {
	Contract     *IValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IValidatorRaw struct {
	Contract *IValidator // Generic contract binding to access the raw methods on
}

// IValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IValidatorCallerRaw struct {
	Contract *IValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// IValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IValidatorTransactorRaw struct {
	Contract *IValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIValidator creates a new instance of IValidator, bound to a specific deployed contract.
func NewIValidator(address common.Address, backend bind.ContractBackend) (*IValidator, error) {
	contract, err := bindIValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IValidator{IValidatorCaller: IValidatorCaller{contract: contract}, IValidatorTransactor: IValidatorTransactor{contract: contract}, IValidatorFilterer: IValidatorFilterer{contract: contract}}, nil
}

// NewIValidatorCaller creates a new read-only instance of IValidator, bound to a specific deployed contract.
func NewIValidatorCaller(address common.Address, caller bind.ContractCaller) (*IValidatorCaller, error) {
	contract, err := bindIValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IValidatorCaller{contract: contract}, nil
}

// NewIValidatorTransactor creates a new write-only instance of IValidator, bound to a specific deployed contract.
func NewIValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IValidatorTransactor, error) {
	contract, err := bindIValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IValidatorTransactor{contract: contract}, nil
}

// NewIValidatorFilterer creates a new log filterer instance of IValidator, bound to a specific deployed contract.
func NewIValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IValidatorFilterer, error) {
	contract, err := bindIValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IValidatorFilterer{contract: contract}, nil
}

// bindIValidator binds a generic wrapper to an already deployed contract.
func bindIValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IValidator *IValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IValidator.Contract.IValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IValidator *IValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IValidator.Contract.IValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IValidator *IValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IValidator.Contract.IValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IValidator *IValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IValidator *IValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IValidator *IValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IValidator.Contract.contract.Transact(opts, method, params...)
}

// Propose is a paid mutator transaction binding the contract method 0xd6f0948c.
//
// Solidity: function propose(address,  string) returns()
func (_IValidator *IValidatorTransactor) Propose(opts *bind.TransactOpts, arg0 common.Address, arg1 string) (*types.Transaction, error) {
	return _IValidator.contract.Transact(opts, "propose", arg0, arg1)

// Propose is a paid mutator transaction binding the contract method  0xd6f0948c.
//
// Solidity: function propose( address,  string) returns()
func (_IValidator *IValidatorSession) Propose(arg0 common.Address, arg1 string) (*types.Transaction, error) {
	return _IValidator.Contract.Propose(&_IValidator.TransactOpts, arg0, arg1)


// Propose is a paid mutator transaction binding the contract method  0xd6f0948c.
//
// Solidity: function propose( address,  string) returns()
func (_IValidator *IValidatorTransactorSession) Propose(arg0 common.Address, arg1 string) (*types.Transaction, error) {
	return _IValidator.Contract.Propose(&_IValidator.TransactOpts, arg0, arg1)


// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote( address) returns()
func (_IValidator *IValidatorTransactor) Vote(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IValidator.contract.Transact(opts, "vote", arg0)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote( address) returns()
func (_IValidator *IValidatorSession) Vote(arg0 common.Address) (*types.Transaction, error) {
	return _IValidator.Contract.Vote(&_IValidator.TransactOpts, arg0)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote( address) returns()
func (_IValidator *IValidatorTransactorSession) Vote(arg0 common.Address) (*types.Transaction, error) {
	return _IValidator.Contract.Vote(&_IValidator.TransactOpts, arg0)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146060604052600080fd00a165627a7a72305820b9407d48ebc7efee5c9f08b3b3a957df2939281f5913225e8c1291f069b900490029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// xdcValidatorABI is the input ABI used to generate the binding from.
const xdcValidatorABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateBacker\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"unvote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateWithdrawBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_nodeUrl\",\"type\":\"string\"}],\"name\":\"setNodeUrl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"resign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidatorNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"isCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_nodeUrl\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateNodeUrl\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_candidates\",\"type\":\"address[]\"},{\"name\":\"_caps\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Unvote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_backer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Propose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_backer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"Resign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_backer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_nodeUrl\",\"type\":\"string\"}],\"name\":\"SetNodeUrl\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_backer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"

// xdcValidatorBin is the compiled bytecode used for deploying new contracts
const xdcValidatorBin = `0x6060604052600060035534156200001557600080fd5b6040516200142838038062001428833981016040528080518201919060200180519091019050600060028380516200005292916020019062000171565b50600090505b8251811015620001685760a06040519081016040528033600160a060020a0316815260200160206040519081016040908152600082529082526001602083015201838381518110620000a657fe5b9060200190602002015181526020016000815250600080858481518110620000ca57fe5b90602001906020020151600160a060020a03168152602081019190915260400160002081518154600160a060020a031916600160a060020a039190911617815560208201518160010190805162000126929160200190620001dd565b50604082015160028201805460ff1916911515919091179055606082015181600301556080820151600490910155506003805460019081019091550162000058565b505050620002a5565b828054828255906000526020600020908101928215620001cb579160200282015b82811115620001cb5782518254600160a060020a031916600160a060020a03919091161782556020929092019160019091019062000192565b50620001d99291506200025e565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200022057805160ff191683800117855562000250565b8280016001018555821562000250579182015b828111156200025057825182559160200191906001019062000233565b50620001d992915062000288565b6200028591905b80821115620001d9578054600160a060020a031916815560010162000265565b90565b6200028591905b80821115620001d957600081556001016200028f565b61117380620002b56000396000f3006060604052600436106100ef5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041662a9550181146100f457806302aa9be21461012f57806306a49fce1461015357806328265294146101b95780632d15cc04146101ea578063302b6872146102095780633477ee2e1461022e57806351cff8d91461024457806358e7525f146102635780636dd7d8ea14610282578063a3ec796514610296578063ae6e43f5146102f5578063d09f1ab414610314578063d51b9e9314610327578063d55b7dff1461035a578063d6f0948c1461036d578063da67b5991461038d575b600080fd5b34156100ff57600080fd5b610113600160a060020a0360043516610423565b604051600160a060020a03909116815260200160405180910390f35b341561013a57600080fd5b610151600160a060020a0360043516602435610441565b005b341561015e57600080fd5b610166610599565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156101a557808201518382015260200161018d565b505050509050019250505060405180910390f35b34156101c457600080fd5b6101d8600160a060020a0360043516610602565b60405190815260200160405180910390f35b34156101f557600080fd5b610166600160a060020a0360043516610620565b341561021457600080fd5b6101d8600160a060020a03600435811690602435166106ad565b341561023957600080fd5b6101136004356106da565b341561024f57600080fd5b610151600160a060020a0360043516610702565b341561026e57600080fd5b6101d8600160a060020a03600435166108b3565b610151600160a060020a03600435166108d1565b34156102a157600080fd5b61015160048035600160a060020a03169060446024803590810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610a7b95505050505050565b341561030057600080fd5b610151600160a060020a0360043516610b8b565b341561031f57600080fd5b6101d8610d4c565b341561033257600080fd5b610346600160a060020a0360043516610d51565b604051901515815260200160405180910390f35b341561036557600080fd5b6101d8610d72565b61015160048035600160a060020a03169060248035908101910135610d80565b341561039857600080fd5b6103ac600160a060020a0360043516610f87565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156103e85780820151838201526020016103d0565b50505050905090810190601f1680156104155780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600160a060020a039081166000908152602081905260409020541690565b600160a060020a03808316600090815260208181526040808320339094168352600590930190522054829082908190101561047b57600080fd5b600160a060020a0384166000908152602081905260409020600301546104a7908463ffffffff61104c16565b600160a060020a03808616600090815260208181526040808320600381019590955533909316825260059093019092529020546104ea908463ffffffff61104c16565b600160a060020a0380861660009081526020818152604080832033909416808452600590940190915290819020929092559084156108fc0290859051600060405180830381858888f19350505050151561054357600080fd5b7faa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2338585604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a150505050565b6105a1611074565b60028054806020026020016040519081016040528092919081815260200182805480156105f757602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116105d9575b505050505090505b90565b600160a060020a031660009081526020819052604090206004015490565b610628611074565b6001600083600160a060020a0316600160a060020a031681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156106a157602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610683575b50505050509050919050565b600160a060020a039182166000908152602081815260408083209390941682526005909201909152205490565b60028054829081106106e857fe5b600091825260209091200154600160a060020a0316905081565b600160a060020a038181166000908152602081905260408120549091839133821691161461072f57600080fd5b600160a060020a038316600090815260208190526040902060020154839060ff161561075a57600080fd5b600160a060020a0384166000908152602081905260408120600401548591901161078357600080fd5b600160a060020a0381166000908152602081905260409020600401544310156107ab57600080fd5b600160a060020a03808616600081815260208181526040808320339095168352600585018252822054928252526003909101549094506107f1908563ffffffff61104c16565b600160a060020a0380871660008181526020818152604080832060038101969096553390941680835260058601825284832083905592825281905260049093019290925585156108fc0290869051600060405180830381858888f19350505050151561085c57600080fd5b7f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb338686604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15050505050565b600160a060020a031660009081526020819052604090206003015490565b600160a060020a038116600090815260208190526040902060020154819060ff1615156108fd57600080fd5b600160a060020a038216600090815260208190526040902060030154610929903463ffffffff61105e16565b600160a060020a038084166000908152602081815260408083206003810195909555339093168252600590930190925290205415156109c057600160a060020a038216600090815260016020819052604090912080549091810161098d8382611086565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a03161790555b600160a060020a038083166000908152602081815260408083203390941683526005909301905220546109f9903463ffffffff61105e16565b600160a060020a0380841660009081526020818152604080832033948516845260050190915290819020929092557f66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc918490349051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15050565b600160a060020a038281166000908152602081905260409020548391338116911614610aa657600080fd5b600160a060020a0383166000908152602081905260409020600101828051610ad29291602001906110af565b507f63f303264cd4b7a198f0163f96e0b6b1f972f9b73359a70c44241b862879d8a4338484604051600160a060020a0380851682528316602082015260606040820181815290820183818151815260200191508051906020019080838360005b83811015610b4a578082015183820152602001610b32565b50505050905090810190601f168015610b775780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a1505050565b600160a060020a0381811660009081526020819052604081205490918391338216911614610bb857600080fd5b600160a060020a038316600090815260208190526040902060020154839060ff161515610be457600080fd5b600160a060020a0384166000908152602081905260408120600201805460ff191690556003805460001901905592505b600254831015610c965783600160a060020a0316600284815481101515610c3757fe5b600091825260209091200154600160a060020a03161415610c8b576002805484908110610c6057fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff19169055610c96565b600190920191610c14565b600160a060020a038416600090815260208190526040902060040154610cd590606490610cc9904363ffffffff61105e16565b9063ffffffff61105e16565b60008086600160a060020a0316600160a060020a03168152602001908152602001600020600401819055507f4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d33385604051600160a060020a039283168152911660208201526040908101905180910390a150505050565b606381565b600160a060020a031660009081526020819052604090206002015460ff1690565b690a968163f0a57b40000081565b690a968163f0a57b400000341015610d9757600080fd5b600160a060020a038316600090815260208190526040902060020154839060ff1615610dc257600080fd5b6002805460018101610dd48382611086565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03861617905560a06040519081016040528033600160a060020a0316815260200184848080601f0160208091040260200160405190810160405281815292919060208401838380828437505050928452505060016020808401919091523460408085019190915260006060909401849052600160a060020a03891684529083905290912090508151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0391909116178155602082015181600101908051610eca9291602001906110af565b50604082015160028201805460ff191691151591909117905560608201518160030155608082015160049091015550600160a060020a038085166000908152602081815260408083203394851684526005019091529081902034908190556003805460010190557f7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1929187919051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a150505050565b610f8f611074565b60008083600160a060020a0316600160a060020a031681526020019081526020016000206001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106a15780601f1061101f576101008083540402835291602001916106a1565b820191906000526020600020905b81548152906001019060200180831161102d5750939695505050505050565b60008282111561105857fe5b50900390565b60008282018381101561106d57fe5b9392505050565b60206040519081016040526000815290565b8154818355818115116110aa576000838152602090206110aa91810190830161112d565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106110f057805160ff191683800117855561111d565b8280016001018555821561111d579182015b8281111561111d578251825591602001919060010190611102565b5061112992915061112d565b5090565b6105ff91905b8082111561112957600081556001016111335600a165627a7a72305820d2d4c6e641bc033beae9eb0c9fca10f5c9c68a746bc1b0099a8da4c99a3532d30029`

// DeployxdcValidator deploys a new Ethereum contract, binding an instance of xdcValidator to it.
func DeployxdcValidator(auth *bind.TransactOpts, backend bind.ContractBackend, _candidates []common.Address, _caps []*big.Int) (common.Address, *types.Transaction, *xdcValidator, error) {
	parsed, err := abi.JSON(strings.NewReader(xdcValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(xdcValidatorBin), backend, _candidates, _caps)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &xdcValidator{xdcValidatorCaller: xdcValidatorCaller{contract: contract}, xdcValidatorTransactor: xdcValidatorTransactor{contract: contract}, xdcValidatorFilterer: xdcValidatorFilterer{contract: contract}}, nil
}

// xdcValidator is an auto generated Go binding around an Ethereum contract.
type xdcValidator struct {
	xdcValidatorCaller     // Read-only binding to the contract
	xdcValidatorTransactor // Write-only binding to the contract
	xdcValidatorFilterer   // Log filterer for contract events
}

// xdcValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type xdcValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// xdcValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type xdcValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// xdcValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type xdcValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// xdcValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type xdcValidatorSession struct {
	Contract     *xdcValidator    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// xdcValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type xdcValidatorCallerSession struct {
	Contract *xdcValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// xdcValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type xdcValidatorTransactorSession struct {
	Contract     *xdcValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// xdcValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type xdcValidatorRaw struct {
	Contract *xdcValidator // Generic contract binding to access the raw methods on
}

// xdcValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type xdcValidatorCallerRaw struct {
	Contract *xdcValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// xdcValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type xdcValidatorTransactorRaw struct {
	Contract *xdcValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewxdcValidator creates a new instance of xdcValidator, bound to a specific deployed contract.
func NewxdcValidator(address common.Address, backend bind.ContractBackend) (*xdcValidator, error) {
	contract, err := bindxdcValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &xdcValidator{xdcValidatorCaller: xdcValidatorCaller{contract: contract}, xdcValidatorTransactor: xdcValidatorTransactor{contract: contract}, xdcValidatorFilterer: xdcValidatorFilterer{contract: contract}}, nil
}

// NewxdcValidatorCaller creates a new read-only instance of xdcValidator, bound to a specific deployed contract.
func NewxdcValidatorCaller(address common.Address, caller bind.ContractCaller) (*xdcValidatorCaller, error) {
	contract, err := bindxdcValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &xdcValidatorCaller{contract: contract}, nil
}

// NewxdcValidatorTransactor creates a new write-only instance of xdcValidator, bound to a specific deployed contract.
func NewxdcValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*xdcValidatorTransactor, error) {
	contract, err := bindxdcValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &xdcValidatorTransactor{contract: contract}, nil
}

// NewxdcValidatorFilterer creates a new log filterer instance of xdcValidator, bound to a specific deployed contract.
func NewxdcValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*xdcValidatorFilterer, error) {
	contract, err := bindxdcValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &xdcValidatorFilterer{contract: contract}, nil
}

// bindxdcValidator binds a generic wrapper to an already deployed contract.
func bindxdcValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(xdcValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_xdcValidator *xdcValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _xdcValidator.Contract.xdcValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_xdcValidator *xdcValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _xdcValidator.Contract.xdcValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_xdcValidator *xdcValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _xdcValidator.Contract.xdcValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_xdcValidator *xdcValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _xdcValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_xdcValidator *xdcValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _xdcValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_xdcValidator *xdcValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _xdcValidator.Contract.contract.Transact(opts, method, params...)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_xdcValidator *xdcValidatorCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _xdcValidator.contract.Call(opts, out, "candidates", arg0)
	return *ret0, err
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_xdcValidator *xdcValidatorSession) Candidates(arg0 *big.Int) (common.Address, error) {
	return _xdcValidator.Contract.Candidates(&_xdcValidator.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_xdcValidator *xdcValidatorCallerSession) Candidates(arg0 *big.Int) (common.Address, error) {
	return _xdcValidator.Contract.Candidates(&_xdcValidator.CallOpts, arg0)
}
// GetCandidateBacker is a free data retrieval call binding the contract method 0x00a95501.
//
// Solidity: function getCandidateBacker(_candidate address) constant returns(address)
func (_xdcValidator *xdcValidatorCaller) GetCandidateBacker(opts *bind.CallOpts, _candidate common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _xdcValidator.contract.Call(opts, out, "getCandidateBacker", _candidate)
	return *ret0, err
}

// GetCandidateBacker is a free data retrieval call binding the contract method 0x00a95501.
//
// Solidity: function getCandidateBacker(_candidate address) constant returns(address)
func (_xdcValidator *xdcValidatorSession) GetCandidateBacker(_candidate common.Address) (common.Address, error) {
	return _xdcValidator.Contract.GetCandidateBacker(&_xdcValidator.CallOpts, _candidate)
}

// GetCandidateBacker is a free data retrieval call binding the contract method 0x00a95501.
//
// Solidity: function getCandidateBacker(_candidate address) constant returns(address)
func (_xdcValidator *xdcValidatorCallerSession) GetCandidateBacker(_candidate common.Address) (common.Address, error) {
	return _xdcValidator.Contract.GetCandidateBacker(&_xdcValidator.CallOpts, _candidate)
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_xdcValidator *xdcValidatorCaller) GetCandidateCap(opts *bind.CallOpts, _candidate common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _xdcValidator.contract.Call(opts, out, "getCandidateCap", _candidate)
	return *ret0, err
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_xdcValidator *xdcValidatorSession) GetCandidateCap(_candidate common.Address) (*big.Int, error) {
	return _xdcValidator.Contract.GetCandidateCap(&_xdcValidator.CallOpts, _candidate)
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_xdcValidator *xdcValidatorCallerSession) GetCandidateCap(_candidate common.Address) (*big.Int, error) {
	return _xdcValidator.Contract.GetCandidateCap(&_xdcValidator.CallOpts, _candidate)
}

// GetCandidateNodeUrl is a free data retrieval call binding the contract method 0xda67b599.
//
// Solidity: function getCandidateNodeUrl(_candidate address) constant returns(string)
func (_xdcValidator *xdcValidatorCaller) GetCandidateNodeUrl(opts *bind.CallOpts, _candidate common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _xdcValidator.contract.Call(opts, out, "getCandidateNodeUrl", _candidate)
	return *ret0, err
}

// GetCandidateNodeUrl is a free data retrieval call binding the contract method 0xda67b599.
//
// Solidity: function getCandidateNodeUrl(_candidate address) constant returns(string)
func (_xdcValidator *xdcValidatorSession) GetCandidateNodeUrl(_candidate common.Address) (string, error) {
	return _xdcValidator.Contract.GetCandidateNodeUrl(&_xdcValidator.CallOpts, _candidate)
}

// GetCandidateNodeUrl is a free data retrieval call binding the contract method 0xda67b599.
//
// Solidity: function getCandidateNodeUrl(_candidate address) constant returns(string)
func (_xdcValidator *xdcValidatorCallerSession) GetCandidateNodeUrl(_candidate common.Address) (string, error) {
	return _xdcValidator.Contract.GetCandidateNodeUrl(&_xdcValidator.CallOpts, _candidate)
}

// GetCandidateWithdrawBlockNumber is a free data retrieval call binding the contract method 0x28265294.
//
// Solidity: function getCandidateWithdrawBlockNumber(_candidate address) constant returns(uint256)
func (_xdcValidator *xdcValidatorCaller) GetCandidateWithdrawBlockNumber(opts *bind.CallOpts, _candidate common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _xdcValidator.contract.Call(opts, out, "getCandidateWithdrawBlockNumber", _candidate)
	return *ret0, err
}

// GetCandidateWithdrawBlockNumber is a free data retrieval call binding the contract method 0x28265294.
//
// Solidity: function getCandidateWithdrawBlockNumber(_candidate address) constant returns(uint256)
func (_xdcValidator *xdcValidatorSession) GetCandidateWithdrawBlockNumber(_candidate common.Address) (*big.Int, error) {
	return _xdcValidator.Contract.GetCandidateWithdrawBlockNumber(&_xdcValidator.CallOpts, _candidate)
}

// GetCandidateWithdrawBlockNumber is a free data retrieval call binding the contract method 0x28265294.
//
// Solidity: function getCandidateWithdrawBlockNumber(_candidate address) constant returns(uint256)
func (_xdcValidator *xdcValidatorCallerSession) GetCandidateWithdrawBlockNumber(_candidate common.Address) (*big.Int, error) {
	return _xdcValidator.Contract.GetCandidateWithdrawBlockNumber(&_xdcValidator.CallOpts, _candidate)
}
	
// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_xdcValidator *xdcValidatorCaller) GetCandidates(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _xdcValidator.contract.Call(opts, out, "getCandidates")
	return *ret0, err
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_xdcValidator *xdcValidatorSession) GetCandidates() ([]common.Address, error) {
	return _xdcValidator.Contract.GetCandidates(&_xdcValidator.CallOpts)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_xdcValidator *xdcValidatorCallerSession) GetCandidates() ([]common.Address, error) {
	return _xdcValidator.Contract.GetCandidates(&_xdcValidator.CallOpts)
}


// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_xdcValidator *xdcValidatorCaller) GetVoterCap(opts *bind.CallOpts, _candidate common.Address, _voter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _xdcValidator.contract.Call(opts, out, "getVoterCap", _candidate, _voter)
	return *ret0, err
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_xdcValidator *xdcValidatorSession) GetVoterCap(_candidate common.Address, _voter common.Address) (*big.Int, error) {
	return _xdcValidator.Contract.GetVoterCap(&_xdcValidator.CallOpts, _candidate, _voter)
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_xdcValidator *xdcValidatorCallerSession) GetVoterCap(_candidate common.Address, _voter common.Address) (*big.Int, error) {
	return _xdcValidator.Contract.GetVoterCap(&_xdcValidator.CallOpts, _candidate, _voter)
}
// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_xdcValidator *xdcValidatorCaller) GetVoters(opts *bind.CallOpts, _candidate common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _xdcValidator.contract.Call(opts, out, "getVoters", _candidate)
	return *ret0, err
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_xdcValidator *xdcValidatorSession) GetVoters(_candidate common.Address) ([]common.Address, error) {
	return _xdcValidator.Contract.GetVoters(&_xdcValidator.CallOpts, _candidate)
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_xdcValidator *xdcValidatorCallerSession) GetVoters(_candidate common.Address) ([]common.Address, error) {
	return _xdcValidator.Contract.GetVoters(&_xdcValidator.CallOpts, _candidate)
}
// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_xdcValidator *xdcValidatorCaller) IsCandidate(opts *bind.CallOpts, _candidate common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _xdcValidator.contract.Call(opts, out, "isCandidate", _candidate)
	return *ret0, err
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_xdcValidator *xdcValidatorSession) IsCandidate(_candidate common.Address) (bool, error) {
	return _xdcValidator.Contract.IsCandidate(&_xdcValidator.CallOpts, _candidate)
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_xdcValidator *xdcValidatorCallerSession) IsCandidate(_candidate common.Address) (bool, error) {
	return _xdcValidator.Contract.IsCandidate(&_xdcValidator.CallOpts, _candidate)
}


// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_xdcValidator *xdcValidatorCaller) MaxValidatorNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _xdcValidator.contract.Call(opts, out, "maxValidatorNumber")
	return *ret0, err
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_xdcValidator *xdcValidatorSession) MaxValidatorNumber() (*big.Int, error) {
	return _xdcValidator.Contract.MaxValidatorNumber(&_xdcValidator.CallOpts)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_xdcValidator *xdcValidatorCallerSession) MaxValidatorNumber() (*big.Int, error) {
	return _xdcValidator.Contract.MaxValidatorNumber(&_xdcValidator.CallOpts)
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_xdcValidator *xdcValidatorCaller) MinCandidateCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _xdcValidator.contract.Call(opts, out, "minCandidateCap")
	return *ret0, err
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_xdcValidator *xdcValidatorSession) MinCandidateCap() (*big.Int, error) {
	return _xdcValidator.Contract.MinCandidateCap(&_xdcValidator.CallOpts)
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_xdcValidator *xdcValidatorCallerSession) MinCandidateCap() (*big.Int, error) {
	return _xdcValidator.Contract.MinCandidateCap(&_xdcValidator.CallOpts)
}

// Propose is a paid mutator transaction binding the contract method 0xd6f0948c.
//
// Solidity: function propose(_candidate address, _nodeUrl string) returns()
func (_xdcValidator *xdcValidatorTransactor) Propose(opts *bind.TransactOpts, _candidate common.Address, _nodeUrl string) (*types.Transaction, error) {
	return _xdcValidator.contract.Transact(opts, "propose", _candidate, _nodeUrl)
}

// Propose is a paid mutator transaction binding the contract method 0xd6f0948c.
// Solidity: function propose(_candidate address, _nodeUrl string) returns()
func (_xdcValidator *xdcValidatorSession) Propose(_candidate common.Address, _nodeUrl string) (*types.Transaction, error) {
	return _xdcValidator.Contract.Propose(&_xdcValidator.TransactOpts, _candidate, _nodeUrl)
}

// Propose is a paid mutator transaction binding the contract method 0xd6f0948c.
//
// Solidity: function propose(_candidate address, _nodeUrl string) returns()
func (_xdcValidator *xdcValidatorTransactorSession) Propose(_candidate common.Address, _nodeUrl string) (*types.Transaction, error) {
	return _xdcValidator.Contract.Propose(&_xdcValidator.TransactOpts, _candidate, _nodeUrl)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_xdcValidator *xdcValidatorTransactor) Resign(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _xdcValidator.contract.Transact(opts, "resign", _candidate)
}
// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_xdcValidator *xdcValidatorSession) Resign(_candidate common.Address) (*types.Transaction, error) {
	return _xdcValidator.Contract.Resign(&_xdcValidator.TransactOpts, _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5
//
// Solidity: function resign(_candidate address) returns()
func (_xdcValidator *xdcValidatorTransactorSession) Resign(_candidate common.Address) (*types.Transaction, error) {
	return _xdcValidator.Contract.Resign(&_xdcValidator.TransactOpts, _candidate)
}

// SetNodeUrl is a paid mutator transaction binding the contract method 0xa3ec7965.
//
// Solidity: function setNodeUrl(_candidate address, _nodeUrl string) returns()
func (_xdcValidator *xdcValidatorTransactor) SetNodeUrl(opts *bind.TransactOpts, _candidate common.Address, _nodeUrl string) (*types.Transaction, error) {
	return _xdcValidator.contract.Transact(opts, "setNodeUrl", _candidate, _nodeUrl)
}

// SetNodeUrl is a paid mutator transaction binding the contract method 0xa3ec7965.
//
// Solidity: function setNodeUrl(_candidate address, _nodeUrl string) returns()
func (_xdcValidator *xdcValidatorSession) SetNodeUrl(_candidate common.Address, _nodeUrl string) (*types.Transaction, error) {
	return _xdcValidator.Contract.SetNodeUrl(&_xdcValidator.TransactOpts, _candidate, _nodeUrl)
}

// SetNodeUrl is a paid mutator transaction binding the contract method 0xa3ec7965.
//
// Solidity: function setNodeUrl(_candidate address, _nodeUrl string) returns()
func (_xdcValidator *xdcValidatorTransactorSession) SetNodeUrl(_candidate common.Address, _nodeUrl string) (*types.Transaction, error) {
	return _xdcValidator.Contract.SetNodeUrl(&_xdcValidator.TransactOpts, _candidate, _nodeUrl)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_xdcValidator *xdcValidatorTransactor) Unvote(opts *bind.TransactOpts, _candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _xdcValidator.contract.Transact(opts, "unvote", _candidate, _cap)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_xdcValidator *xdcValidatorSession) Unvote(_candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _xdcValidator.Contract.Unvote(&_xdcValidator.TransactOpts, _candidate, _cap)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_xdcValidator *xdcValidatorTransactorSession) Unvote(_candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _xdcValidator.Contract.Unvote(&_xdcValidator.TransactOpts, _candidate, _cap)
}


// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_xdcValidator *xdcValidatorTransactor) Vote(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _xdcValidator.contract.Transact(opts, "vote", _candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_xdcValidator *xdcValidatorSession) Vote(_candidate common.Address) (*types.Transaction, error) {
	return _xdcValidator.Contract.Vote(&_xdcValidator.TransactOpts, _candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_xdcValidator *xdcValidatorTransactorSession) Vote(_candidate common.Address) (*types.Transaction, error) {
	return _xdcValidator.Contract.Vote(&_xdcValidator.TransactOpts, _candidate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(_candidate address) returns()
func (_xdcValidator *xdcValidatorTransactor) Withdraw(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _xdcValidator.contract.Transact(opts, "withdraw", _candidate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(_candidate address) returns()
func (_xdcValidator *xdcValidatorSession) Withdraw(_candidate common.Address) (*types.Transaction, error) {
	return _xdcValidator.Contract.Withdraw(&_xdcValidator.TransactOpts, _candidate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(_candidate address) returns()
func (_xdcValidator *xdcValidatorTransactorSession) Withdraw(_candidate common.Address) (*types.Transaction, error) {
	return _xdcValidator.Contract.Withdraw(&_xdcValidator.TransactOpts, _candidate)
}

// xdcValidatorProposeIterator is returned from FilterPropose and is used to iterate over the raw logs and unpacked data for Propose events raised by the xdcValidator contract.
type xdcValidatorProposeIterator struct {
	Event *xdcValidatorPropose // Event containing the contract specifics and raw log

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
func (it *xdcValidatorProposeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(xdcValidatorPropose)
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
		it.Event = new(xdcValidatorPropose)
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
func (it *xdcValidatorProposeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *xdcValidatorProposeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// xdcValidatorPropose represents a Propose event raised by the xdcValidator contract.
type xdcValidatorPropose struct {
	Backer    common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPropose is a free log retrieval operation binding the contract event 0x7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1
// Solidity: event Propose(_backer address, _candidate address, _cap uint256)
func (_xdcValidator *xdcValidatorFilterer) FilterPropose(opts *bind.FilterOpts) (*xdcValidatorProposeIterator, error) {

	logs, sub, err := _xdcValidator.contract.FilterLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return &xdcValidatorProposeIterator{contract: _xdcValidator.contract, event: "Propose", logs: logs, sub: sub}, nil
}

// WatchPropose is a free log subscription operation binding the contract event 0x7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1.
// Solidity: event Propose(_backer address, _candidate address, _cap uint256)
func (_xdcValidator *xdcValidatorFilterer) WatchPropose(opts *bind.WatchOpts, sink chan<- *xdcValidatorPropose) (event.Subscription, error) {

	logs, sub, err := _xdcValidator.contract.WatchLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(xdcValidatorPropose)
				if err := _xdcValidator.contract.UnpackLog(event, "Propose", log); err != nil {
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

// xdcValidatorResignIterator is returned from FilterResign and is used to iterate over the raw logs and unpacked data for Resign events raised by the xdcValidator contract.
type xdcValidatorResignIterator struct {
	Event *xdcValidatorResign // Event containing the contract specifics and raw log

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
func (it *xdcValidatorResignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(xdcValidatorResign)
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
		it.Event = new(xdcValidatorResign)
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
func (it *xdcValidatorResignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *xdcValidatorResignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
// xdcValidatorResign represents a Resign event raised by the xdcValidator contract.
type xdcValidatorResign struct {
	Backer    common.Address
	Candidate common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResign is a free log retrieval operation binding the contract event 0x4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3.
//
// Solidity: event Resign(_backer address, _candidate address)
func (_xdcValidator *xdcValidatorFilterer) FilterResign(opts *bind.FilterOpts) (*xdcValidatorResignIterator, error) {

	logs, sub, err := _xdcValidator.contract.FilterLogs(opts, "Resign")
	if err != nil {
		return nil, err
	}
	return &xdcValidatorResignIterator{contract: _xdcValidator.contract, event: "Resign", logs: logs, sub: sub}, nil
}

// WatchResign is a free log subscription operation binding the contract event 0x4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3.
//
// Solidity: event Resign(_backer address, _candidate address)
func (_xdcValidator *xdcValidatorFilterer) WatchResign(opts *bind.WatchOpts, sink chan<- *xdcValidatorResign) (event.Subscription, error) {
	logs, sub, err := _xdcValidator.contract.WatchLogs(opts, "Resign")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(xdcValidatorResign)
				if err := _xdcValidator.contract.UnpackLog(event, "Resign", log); err != nil {
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

// xdcValidatorSetNodeUrlIterator is returned from FilterSetNodeUrl and is used to iterate over the raw logs and unpacked data for SetNodeUrl events raised by the xdcValidator contract.
type xdcValidatorSetNodeUrlIterator struct {
	Event *xdcValidatorSetNodeUrl // Event containing the contract specifics and raw log

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
func (it *xdcValidatorSetNodeUrlIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(xdcValidatorSetNodeUrl)
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
		it.Event = new(xdcValidatorSetNodeUrl)
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
func (it *xdcValidatorSetNodeUrlIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *xdcValidatorSetNodeUrlIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// xdcValidatorSetNodeUrl represents a SetNodeUrl event raised by the xdcValidator contract.
type xdcValidatorSetNodeUrl struct {
	Backer    common.Address
	Candidate common.Address
	NodeUrl   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetNodeUrl is a free log retrieval operation binding the contract event 0x63f303264cd4b7a198f0163f96e0b6b1f972f9b73359a70c44241b862879d8a4.
//
// Solidity: event SetNodeUrl(_backer address, _candidate address, _nodeUrl string)
func (_xdcValidator *xdcValidatorFilterer) FilterSetNodeUrl(opts *bind.FilterOpts) (*xdcValidatorSetNodeUrlIterator, error) {

	logs, sub, err := _xdcValidator.contract.FilterLogs(opts, "SetNodeUrl")
	if err != nil {
		return nil, err
	}
	return &xdcValidatorSetNodeUrlIterator{contract: _xdcValidator.contract, event: "SetNodeUrl", logs: logs, sub: sub}, nil
}

// WatchSetNodeUrl is a free log subscription operation binding the contract event 0x63f303264cd4b7a198f0163f96e0b6b1f972f9b73359a70c44241b862879d8a4.
//
// Solidity: event SetNodeUrl(_backer address, _candidate address, _nodeUrl string)
func (_xdcValidator *xdcValidatorFilterer) WatchSetNodeUrl(opts *bind.WatchOpts, sink chan<- *xdcValidatorSetNodeUrl) (event.Subscription, error) {

	logs, sub, err := _xdcValidator.contract.WatchLogs(opts, "SetNodeUrl")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(xdcValidatorSetNodeUrl)
				if err := _xdcValidator.contract.UnpackLog(event, "SetNodeUrl", log); err != nil {
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


// xdcValidatorUnvoteIterator is returned from FilterUnvote and is used to iterate over the raw logs and unpacked data for Unvote events raised by the xdcValidator contract.
type xdcValidatorUnvoteIterator struct {
	Event *xdcValidatorUnvote // Event containing the contract specifics and raw log

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
func (it *xdcValidatorUnvoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(xdcValidatorUnvote)
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
		it.Event = new(xdcValidatorUnvote)
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
func (it *xdcValidatorUnvoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *xdcValidatorUnvoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// xdcValidatorUnvote represents a Unvote event raised by the xdcValidator contract.
type xdcValidatorUnvote struct {
	Voter     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnvote is a free log retrieval operation binding the contract event 0xaa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2.
//
// Solidity: event Unvote(_voter address, _candidate address, _cap uint256)
func (_xdcValidator *xdcValidatorFilterer) FilterUnvote(opts *bind.FilterOpts) (*xdcValidatorUnvoteIterator, error) {

	logs, sub, err := _xdcValidator.contract.FilterLogs(opts, "Unvote")
	if err != nil {
		return nil, err
	}
	return &xdcValidatorUnvoteIterator{contract: _xdcValidator.contract, event: "Unvote", logs: logs, sub: sub}, nil
}

// WatchUnvote is a free log subscription operation binding the contract event 0xaa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2.
//
// Solidity: event Unvote(_voter address, _candidate address, _cap uint256)
func (_xdcValidator *xdcValidatorFilterer) WatchUnvote(opts *bind.WatchOpts, sink chan<- *xdcValidatorUnvote) (event.Subscription, error) {

	logs, sub, err := _xdcValidator.contract.WatchLogs(opts, "Unvote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(xdcValidatorUnvote)
				if err := _xdcValidator.contract.UnpackLog(event, "Unvote", log); err != nil {
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

// xdcValidatorVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the xdcValidator contract.
type xdcValidatorVoteIterator struct {
	Event *xdcValidatorVote // Event containing the contract specifics and raw log

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
func (it *xdcValidatorVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(xdcValidatorVote)
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
		it.Event = new(xdcValidatorVote)
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
func (it *xdcValidatorVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *xdcValidatorVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// xdcValidatorVote represents a Vote event raised by the xdcValidator contract.
type xdcValidatorVote struct {
	Voter     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(_voter address, _candidate address, _cap uint256)
func (_xdcValidator *xdcValidatorFilterer) FilterVote(opts *bind.FilterOpts) (*xdcValidatorVoteIterator, error) {

	logs, sub, err := _xdcValidator.contract.FilterLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return &xdcValidatorVoteIterator{contract: _xdcValidator.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(_voter address, _candidate address, _cap uint256)
func (_xdcValidator *xdcValidatorFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *xdcValidatorVote) (event.Subscription, error) {

	logs, sub, err := _xdcValidator.contract.WatchLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(xdcValidatorVote)
				if err := _xdcValidator.contract.UnpackLog(event, "Vote", log); err != nil {
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

// xdcValidatorWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the xdcValidator contract.
type xdcValidatorWithdrawIterator struct {
	Event *xdcValidatorWithdraw // Event containing the contract specifics and raw log

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
func (it *xdcValidatorWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(xdcValidatorWithdraw)
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
		it.Event = new(xdcValidatorWithdraw)
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
func (it *xdcValidatorWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *xdcValidatorWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// xdcValidatorWithdraw represents a Withdraw event raised by the xdcValidator contract.
type xdcValidatorWithdraw struct {
	Backer    common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(_backer address, _candidate address, _cap uint256)
func (_xdcValidator *xdcValidatorFilterer) FilterWithdraw(opts *bind.FilterOpts) (*xdcValidatorWithdrawIterator, error) {

	logs, sub, err := _xdcValidator.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &xdcValidatorWithdrawIterator{contract: _xdcValidator.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(_backer address, _candidate address, _cap uint256)
func (_xdcValidator *xdcValidatorFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *xdcValidatorWithdraw) (event.Subscription, error) {

	logs, sub, err := _xdcValidator.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(xdcValidatorWithdraw)
				if err := _xdcValidator.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
