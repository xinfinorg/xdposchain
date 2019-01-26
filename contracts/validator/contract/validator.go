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
const xdcValidatorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"propose\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"unvote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"isCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_validators\",\"type\":\"address[]\"},{\"name\":\"_caps\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// xdcValidatorBin is the compiled bytecode used for deploying new contracts.
const xdcValidatorBin = `0x6060604052683635c9adc5dea00000600355341561001c57600080fd5b604051610a06380380610a068339810160405280805182019190602001805190910190506000600183805161005592916020019061011f565b50600283805161006992916020019061011f565b50600090505b8251811015610117576060604051908101604090815260018083526020830152810183838151811061009d57fe5b9060200190602002015190526000808584815181106100b857fe5b90602001906020020151600160a060020a0316815260208101919091526040016000208151815460ff1916901515178155602082015181549015156101000261ff0019909116178155604082015160019182015591909101905061006f565b5050506101ad565b828054828255906000526020600020908101928215610176579160200282015b828111156101765782518254600160a060020a031916600160a060020a03919091161782556020929092019160019091019061013f565b50610182929150610186565b5090565b6101aa91905b80821115610182578054600160a060020a031916815560010161018c565b90565b61084a806101bc6000396000f3006060604052600436106100ae5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630126795181146100b357806302aa9be2146100c957806306a49fce146100eb578063302b6872146101515780633477ee2e1461018857806335aa2e44146101ba57806358e7525f146101d05780636dd7d8ea146101ef578063b7ab4db514610203578063d51b9e9314610216578063facd743b14610249575b600080fd5b6100c7600160a060020a0360043516610268565b005b34156100d457600080fd5b6100c7600160a060020a0360043516602435610362565b34156100f657600080fd5b6100fe6104b6565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561013d578082015183820152602001610125565b505050509050019250505060405180910390f35b341561015c57600080fd5b610176600160a060020a036004358116906024351661051f565b60405190815260200160405180910390f35b341561019357600080fd5b61019e60043561054c565b604051600160a060020a03909116815260200160405180910390f35b34156101c557600080fd5b61019e600435610574565b34156101db57600080fd5b610176600160a060020a0360043516610582565b6100c7600160a060020a03600435166105a0565b341561020e57600080fd5b6100fe6106f6565b341561022157600080fd5b610235600160a060020a036004351661075c565b604051901515815260200160405180910390f35b341561025457600080fd5b610235600160a060020a036004351661077f565b600160a060020a03331660009081526020819052604090205460ff16151561028f57600080fd5b600160a060020a038116600090815260208190526040902054610100900460ff1615156102fb5760028054600181016102c883826107c5565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b60606040519081016040908152600080835260016020808501919091523483850152600160a060020a0385168252819052208151815460ff1916901515178155602082015181549015156101000261ff001990911617815560408201516001909101555050565b600160a060020a03821660009081526020819052604090205460ff161561038857600080fd5b600160a060020a038216600090815260208190526040902054610100900460ff1615156103b457600080fd5b600160a060020a03808316600090815260208181526040808320339094168352600290930190522054819010156103ea57600080fd5b600160a060020a038216600090815260208190526040902060010154610416908263ffffffff61079d16565b600160a060020a0380841660009081526020818152604080832060018101959095553390931682526002909301909252902054610459908263ffffffff61079d16565b600160a060020a0380841660009081526020818152604080832033909416808452600290940190915290819020929092559082156108fc0290839051600060405180830381858888f1935050505015156104b257600080fd5b5050565b6104be6107ee565b600280548060200260200160405190810160405280929190818152602001828054801561051457602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116104f6575b505050505090505b90565b600160a060020a039182166000908152602081815260408083209390941682526002909201909152205490565b600280548290811061055a57fe5b600091825260209091200154600160a060020a0316905081565b600180548290811061055a57fe5b600160a060020a031660009081526020819052604090206001015490565b600160a060020a038116600090815260208190526040902054610100900460ff1615156105cc57600080fd5b600160a060020a0381166000908152602081905260409020600101546105f8903463ffffffff6107af16565b600160a060020a038216600090815260208190526040902060010181905560035490106106f357600160a060020a03811660009081526020819052604090205460ff161515610685576001805480820161065283826107c5565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b600160a060020a03808216600090815260208181526040808320805460ff191660011781553390941683526002909301905220546106c9903463ffffffff6107af16565b600160a060020a038083166000908152602081815260408083203390941683526002909301905220555b50565b6106fe6107ee565b600180548060200260200160405190810160405280929190818152602001828054801561051457602002820191906000526020600020908154600160a060020a031681526001909101906020018083116104f6575050505050905090565b600160a060020a0316600090815260208190526040902054610100900460ff1690565b600160a060020a031660009081526020819052604090205460ff1690565b6000828211156107a957fe5b50900390565b6000828201838110156107be57fe5b9392505050565b8154818355818115116107e9576000838152602090206107e9918101908301610800565b505050565b60206040519081016040526000815290565b61051c91905b8082111561081a5760008155600101610806565b50905600a165627a7a723058202ad1e9fda0e95a1d0040266c6438e7a86b19a47ec3debbaaf3a2ce2608178ff50029`

// DeployxdcValidator deploys a new Ethereum contract, binding an instance of xdcValidator to it.
func DeployxdcValidator(auth *bind.TransactOpts, backend bind.ContractBackend, _validators []common.Address, _caps []*big.Int) (common.Address, *types.Transaction, *xdcValidator, error) {
	parsed, err := abi.JSON(strings.NewReader(xdcValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(xdcValidatorBin), backend, _validators, _caps)
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

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() constant returns(address[])
func (_xdcValidator *xdcValidatorCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _xdcValidator.contract.Call(opts, out, "getValidators")
	return *ret0, err
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() constant returns(address[])
func (_xdcValidator *xdcValidatorSession) GetValidators() ([]common.Address, error) {
	return _xdcValidator.Contract.GetValidators(&_xdcValidator.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() constant returns(address[])
func (_xdcValidator *xdcValidatorCallerSession) GetValidators() ([]common.Address, error) {
	return _xdcValidator.Contract.GetValidators(&_xdcValidator.CallOpts)
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

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(_candidate address) constant returns(bool)
func (_xdcValidator *xdcValidatorCaller) IsValidator(opts *bind.CallOpts, _candidate common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _xdcValidator.contract.Call(opts, out, "isValidator", _candidate)
	return *ret0, err
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(_candidate address) constant returns(bool)
func (_xdcValidator *xdcValidatorSession) IsValidator(_candidate common.Address) (bool, error) {
	return _xdcValidator.Contract.IsValidator(&_xdcValidator.CallOpts, _candidate)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(_candidate address) constant returns(bool)
func (_xdcValidator *xdcValidatorCallerSession) IsValidator(_candidate common.Address) (bool, error) {
	return _xdcValidator.Contract.IsValidator(&_xdcValidator.CallOpts, _candidate)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators( uint256) constant returns(address)
func (_xdcValidator *xdcValidatorCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _xdcValidator.contract.Call(opts, out, "validators", arg0)
	return *ret0, err
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators( uint256) constant returns(address)
func (_xdcValidator *xdcValidatorSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _xdcValidator.Contract.Validators(&_xdcValidator.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators( uint256) constant returns(address)
func (_xdcValidator *xdcValidatorCallerSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _xdcValidator.Contract.Validators(&_xdcValidator.CallOpts, arg0)
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

