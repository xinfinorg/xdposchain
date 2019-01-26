// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IValidatorABI is the input ABI used to generate the binding from.
const IValidatorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"propose\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

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

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose( address) returns()
func (_IValidator *IValidatorTransactor) Propose(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IValidator.contract.Transact(opts, "propose", arg0)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose( address) returns()
func (_IValidator *IValidatorSession) Propose(arg0 common.Address) (*types.Transaction, error) {
	return _IValidator.Contract.Propose(&_IValidator.TransactOpts, arg0)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose( address) returns()
func (_IValidator *IValidatorTransactorSession) Propose(arg0 common.Address) (*types.Transaction, error) {
	return _IValidator.Contract.Propose(&_IValidator.TransactOpts, arg0)
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
const xdcValidatorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"propose\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"unvote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"isCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_candidates\",\"type\":\"address[]\"},{\"name\":\"_caps\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
// xdcValidatorBin is the compiled bytecode used for deploying new contracts.
const xdcValidatorBin = `0x60606040526000600255341561001457600080fd5b6040516107d63803806107d68339810160405280805182019190602001805190910190506000600183805161004d9291602001906100ec565b50600090505b82518110156100e45760408051908101604052600181526020810183838151811061007a57fe5b90602001906020020151905260008085848151811061009557fe5b90602001906020020151600160a060020a0316815260208101919091526040016000208151815460ff191690151517815560208201516001918201556002805482019055919091019050610053565b50505061017a565b828054828255906000526020600020908101928215610143579160200282015b828111156101435782518254600160a060020a031916600160a060020a03919091161782556020929092019160019091019061010c565b5061014f929150610153565b5090565b61017791905b8082111561014f578054600160a060020a0319168155600101610159565b90565b61064d806101896000396000f30060606040526004361061008d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301267951811461009257806302aa9be2146100a857806306a49fce146100ca578063302b6872146101305780633477ee2e1461016757806358e7525f146101995780636dd7d8ea146101b8578063d51b9e93146101cc575b600080fd5b6100a6600160a060020a03600435166101ff565b005b34156100b357600080fd5b6100a6600160a060020a03600435166024356102b9565b34156100d557600080fd5b6100dd6103e2565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561011c578082015183820152602001610104565b505050509050019250505060405180910390f35b341561013b57600080fd5b610155600160a060020a036004358116906024351661044b565b60405190815260200160405180910390f35b341561017257600080fd5b61017d600435610478565b604051600160a060020a03909116815260200160405180910390f35b34156101a457600080fd5b610155600160a060020a03600435166104a0565b6100a6600160a060020a03600435166104be565b34156101d757600080fd5b6101eb600160a060020a0360043516610582565b604051901515815260200160405180910390f35b600160a060020a03811660009081526020819052604090205460ff161515610265576001805480820161023283826105c8565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b6040805190810160409081526001825234602080840191909152600160a060020a038416600090815290819052208151815460ff191690151517815560208201516001918201556002805490910190555050565b600160a060020a03821660009081526020819052604090205460ff1615156102e057600080fd5b600160a060020a038083166000908152602081815260408083203390941683526002909301905220548190101561031657600080fd5b600160a060020a038216600090815260208190526040902060010154610342908263ffffffff6105a016565b600160a060020a0380841660009081526020818152604080832060018101959095553390931682526002909301909252902054610385908263ffffffff6105a016565b600160a060020a0380841660009081526020818152604080832033909416808452600290940190915290819020929092559082156108fc0290839051600060405180830381858888f1935050505015156103de57600080fd5b5050565b6103ea6105f1565b600180548060200260200160405190810160405280929190818152602001828054801561044057602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610422575b505050505090505b90565b600160a060020a039182166000908152602081815260408083209390941682526002909201909152205490565b600180548290811061048657fe5b600091825260209091200154600160a060020a0316905081565b600160a060020a031660009081526020819052604090206001015490565b600160a060020a03811660009081526020819052604090205460ff1615156104e557600080fd5b600160a060020a038116600090815260208190526040902060010154610511903463ffffffff6105b216565b600160a060020a0380831660009081526020818152604080832060018101959095553390931682526002909301909252902054610554903463ffffffff6105b216565b600160a060020a03918216600090815260208181526040808320339095168352600290940190529190912055565b600160a060020a031660009081526020819052604090205460ff1690565b6000828211156105ac57fe5b50900390565b6000828201838110156105c157fe5b9392505050565b8154818355818115116105ec576000838152602090206105ec918101908301610603565b505050565b60206040519081016040526000815290565b61044891905b8082111561061d5760008155600101610609565b50905600a165627a7a723058204f35ebf7b5775f52742c989407a76898f3dcbfa948c7b047bd1768f7a5e1f2130029`

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



// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_xdcValidator *xdcValidatorTransactor) Propose(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _xdcValidator.contract.Transact(opts, "propose", _candidate)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_xdcValidator *xdcValidatorSession) Propose(_candidate common.Address) (*types.Transaction, error) {
	return _xdcValidator.Contract.Propose(&_xdcValidator.TransactOpts, _candidate)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_xdcValidator *xdcValidatorTransactorSession) Propose(_candidate common.Address) (*types.Transaction, error) {
	return _xdcValidator.Contract.Propose(&_xdcValidator.TransactOpts, _candidate)
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

