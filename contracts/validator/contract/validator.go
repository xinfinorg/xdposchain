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
const IValidatorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"propose\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

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

// Propose is a paid mutator transaction binding the contract method 0xc198f8ba.
//
// Solidity: function propose() returns()
func (_IValidator *IValidatorTransactor) Propose(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IValidator.contract.Transact(opts, "propose")
}

// Propose is a paid mutator transaction binding the contract method  0xc198f8ba.
//
// Solidity: function propose() returns()
func (_IValidator *IValidatorSession) Propose() (*types.Transaction, error) {
	return _IValidator.Contract.Propose(&_IValidator.TransactOpts)
}

// Propose is a paid mutator transaction binding the contract method  0xc198f8ba.
//
// Solidity: function propose() returns()
func (_IValidator *IValidatorTransactorSession) Propose() (*types.Transaction, error) {
	return _IValidator.Contract.Propose(&_IValidator.TransactOpts)
}

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
const xdcValidatorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"unvote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxCandidateNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"retire\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"propose\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidatorNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"isCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_candidates\",\"type\":\"address[]\"},{\"name\":\"_caps\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Unvote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Propose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Retire\",\"type\":\"event\"}]"

// xdcValidatorBin is the compiled bytecode used for deploying new contracts
const xdcValidatorBin = `0x60606040526000600355341561001457600080fd5b604051610c00380380610c008339810160405280805182019190602001805190910190506000600283805161004d9291602001906100ec565b50600090505b82518110156100e45760408051908101604052600181526020810183838151811061007a57fe5b90602001906020020151905260008085848151811061009557fe5b90602001906020020151600160a060020a0316815260208101919091526040016000208151815460ff191690151517815560208201516001918201556003805482019055919091019050610053565b50505061017a565b828054828255906000526020600020908101928215610143579160200282015b828111156101435782518254600160a060020a031916600160a060020a03919091161782556020929092019160019091019061010c565b5061014f929150610153565b5090565b61017791905b8082111561014f578054600160a060020a0319168155600101610159565b90565b610a77806101896000396000f3006060604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302aa9be281146100c957806306a49fce146100ed5780632d15cc0414610153578063302b6872146101725780633477ee2e146101a957806358e7525f146101db5780636dd7d8ea146101fa5780638198a8dc1461020e578063a4874d7714610221578063c198f8ba14610234578063d09f1ab41461023c578063d51b9e931461024f578063d55b7dff14610282575b600080fd5b34156100d457600080fd5b6100eb600160a060020a0360043516602435610295565b005b34156100f857600080fd5b6101006103db565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561013f578082015183820152602001610127565b505050509050019250505060405180910390f35b341561015e57600080fd5b610100600160a060020a0360043516610444565b341561017d57600080fd5b610197600160a060020a03600435811690602435166104d1565b60405190815260200160405180910390f35b34156101b457600080fd5b6101bf6004356104fe565b604051600160a060020a03909116815260200160405180910390f35b34156101e657600080fd5b610197600160a060020a0360043516610526565b6100eb600160a060020a0360043516610544565b341561021957600080fd5b61019761069b565b341561022c57600080fd5b6100eb6106a1565b6100eb610845565b341561024757600080fd5b610197610999565b341561025a57600080fd5b61026e600160a060020a036004351661099e565b604051901515815260200160405180910390f35b341561028d57600080fd5b6101976109bc565b600160a060020a03808316600090815260208181526040808320339094168352600290930190522054819010156102cb57600080fd5b600160a060020a0382166000908152602081905260409020600101546102f7908263ffffffff6109ca16565b600160a060020a038084166000908152602081815260408083206001810195909555339093168252600290930190925290205461033a908263ffffffff6109ca16565b600160a060020a0380841660009081526020818152604080832033909416808452600290940190915290819020929092559082156108fc0290839051600060405180830381858888f19350505050151561039357600080fd5b7f23ae40ca85f8a7c921ebb4269dc9a81e8a6de8dba614752927e0ff39341392fc8282604051600160a060020a03909216825260208201526040908101905180910390a15050565b6103e36109f2565b600280548060200260200160405190810160405280929190818152602001828054801561043957602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161041b575b505050505090505b90565b61044c6109f2565b6001600083600160a060020a0316600160a060020a031681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156104c557602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116104a7575b50505050509050919050565b600160a060020a039182166000908152602081815260408083209390941682526002909201909152205490565b600280548290811061050c57fe5b600091825260209091200154600160a060020a0316905081565b600160a060020a031660009081526020819052604090206001015490565b600160a060020a03811660009081526020819052604090205460ff16151561056b57600080fd5b600160a060020a038116600090815260208190526040902060010154610597903463ffffffff6109dc16565b600160a060020a03808316600090815260208181526040808320600181019590955533909316825260029093019092529020546105da903463ffffffff6109dc16565b600160a060020a038083166000818152602081815260408083203390951683526002909401815283822094909455908152600192839052208054909181016106228382610a04565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a03161790557ff668ead05c744b9178e571d2edb452e72baf6529c8d72160e64e59b50d865bd08134604051600160a060020a03909216825260208201526040908101905180910390a150565b6101f481565b600160a060020a033316600090815260208190526040812054819060ff1615156106ca57600080fd5b600160a060020a033316600090815260208181526040808320600281018352908320549290915260010154909250610708908363ffffffff6109ca16565b600160a060020a03331660009081526020818152604080832060018101949094556002840182528220829055819052815460ff19169091556003805460001901905590505b6002548110156107cc5733600160a060020a031660028281548110151561077057fe5b600091825260209091200154600160a060020a031614156107c457600280548290811061079957fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191690556107cc565b60010161074d565b600160a060020a03331682156108fc0283604051600060405180830381858888f1935050505015156107fd57600080fd5b7f82b89ed824b293574a2cca050e6e27837b60436d911352f8dca203a9cd35241c3383604051600160a060020a03909216825260208201526040908101905180910390a15050565b69021e19e0c9bab240000034101561085c57600080fd5b600160a060020a03331660009081526020819052604090205460ff161561088257600080fd5b6003546101f490111561089457600080fd5b60028054600181016108a68382610a04565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a03161790556040805190810160409081526001825234602080840191909152600160a060020a033316600090815290819052208151815460ff1916901515178155602082015160019182015533600160a060020a038116600090815260208181526040808320600201909152908190203490819055600380549094019093557f42681fc159c671d489c99c06f4693d03a705a21df2fbb2b84eedda79e0db4cff935090919051600160a060020a03909216825260208201526040908101905180910390a1565b606381565b600160a060020a031660009081526020819052604090205460ff1690565b69021e19e0c9bab240000081565b6000828211156109d657fe5b50900390565b6000828201838110156109eb57fe5b9392505050565b60206040519081016040526000815290565b815481835581811511610a2857600083815260209020610a28918101908301610a2d565b505050565b61044191905b80821115610a475760008155600101610a33565b50905600a165627a7a7230582080c308e459c6cbc6e91ffb692f7e77a6ff3985db345631c572a5e453588bc56c0029`
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




// MaxCandidateNumber is a free data retrieval call binding the contract method 0x8198a8dc.
//
// Solidity: function maxCandidateNumber() constant returns(uint256)
func (_xdcValidator *xdcValidatorCaller) MaxCandidateNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _xdcValidator.contract.Call(opts, out, "maxCandidateNumber")
	return *ret0, err
}

// MaxCandidateNumber is a free data retrieval call binding the contract method 0x8198a8dc.
//
// Solidity: function maxCandidateNumber() constant returns(uint256)
func (_xdcValidator *xdcValidatorSession) MaxCandidateNumber() (*big.Int, error) {
	return _xdcValidator.Contract.MaxCandidateNumber(&_xdcValidator.CallOpts)
}

// MaxCandidateNumber is a free data retrieval call binding the contract method 0x8198a8dc.
//
// Solidity: function maxCandidateNumber() constant returns(uint256)
func (_xdcValidator *xdcValidatorCallerSession) MaxCandidateNumber() (*big.Int, error) {
	return _xdcValidator.Contract.MaxCandidateNumber(&_xdcValidator.CallOpts)
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

// Propose is a paid mutator transaction binding the contract method 0xc198f8ba.
// Solidity: function propose() returns()
func (_xdcValidator *xdcValidatorTransactor) Propose(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _xdcValidator.contract.Transact(opts, "propose")
}

// Propose is a paid mutator transaction binding the contract method 0xc198f8ba.
//
// Solidity: function propose() returns()
func (_xdcValidator *xdcValidatorTransactorSession) Propose() (*types.Transaction, error) {
	return _xdcValidator.Contract.Propose(&_xdcValidator.TransactOpts)
}

// Propose is a paid mutator transaction binding the contract method 0xc198f8ba.
//
// Solidity: function propose() returns()
func (_xdcValidator *xdcValidatorTransactorSession) Propose() (*types.Transaction, error) {
	return _xdcValidator.Contract.Propose(&_xdcValidator.TransactOpts)
}
// Retire is a paid mutator transaction binding the contract method 0xa4874d77.
//
// Solidity: function retire() returns()
func (_xdcValidator *xdcValidatorTransactor) Retire(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _xdcValidator.contract.Transact(opts, "retire")
}

// Retire is a paid mutator transaction binding the contract method 0xa4874d77.
//
// Solidity: function retire() returns()
func (_xdcValidator *xdcValidatorSession) Retire() (*types.Transaction, error) {
	return _xdcValidator.Contract.Retire(&_xdcValidator.TransactOpts)
}

// Retire is a paid mutator transaction binding the contract method 0xa4874d77.
//
// Solidity: function retire() returns()
func (_xdcValidator *xdcValidatorTransactorSession) Retire() (*types.Transaction, error) {
	return _xdcValidator.Contract.Retire(&_xdcValidator.TransactOpts)
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
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPropose is a free log retrieval operation binding the contract event 0x42681fc159c671d489c99c06f4693d03a705a21df2fbb2b84eedda79e0db4cff.
//
// Solidity: event Propose(_candidate address, _cap uint256)
func (_xdcValidator *xdcValidatorFilterer) FilterPropose(opts *bind.FilterOpts) (*xdcValidatorProposeIterator, error) {

	logs, sub, err := _xdcValidator.contract.FilterLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return &xdcValidatorProposeIterator{contract: _xdcValidator.contract, event: "Propose", logs: logs, sub: sub}, nil
}

// WatchPropose is a free log subscription operation binding the contract event 0x42681fc159c671d489c99c06f4693d03a705a21df2fbb2b84eedda79e0db4cff.
//
// Solidity: event Propose(_candidate address, _cap uint256)
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

// xdcValidatorRetireIterator is returned from FilterRetire and is used to iterate over the raw logs and unpacked data for Retire events raised by the xdcValidator contract.
type xdcValidatorRetireIterator struct {
	Event *xdcValidatorRetire // Event containing the contract specifics and raw log

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
func (it *xdcValidatorRetireIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(xdcValidatorRetire)
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
		it.Event = new(xdcValidatorRetire)
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
func (it *xdcValidatorRetireIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *xdcValidatorRetireIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// xdcValidatorRetire represents a Retire event raised by the xdcValidator contract.
type xdcValidatorRetire struct {
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRetire is a free log retrieval operation binding the contract event 0x82b89ed824b293574a2cca050e6e27837b60436d911352f8dca203a9cd35241c.
//
// Solidity: event Retire(_candidate address, _cap uint256)
func (_xdcValidator *xdcValidatorFilterer) FilterRetire(opts *bind.FilterOpts) (*xdcValidatorRetireIterator, error) {

	logs, sub, err := _xdcValidator.contract.FilterLogs(opts, "Retire")
	if err != nil {
		return nil, err
	}
	return &xdcValidatorRetireIterator{contract: _xdcValidator.contract, event: "Retire", logs: logs, sub: sub}, nil
}

// WatchRetire is a free log subscription operation binding the contract event 0x82b89ed824b293574a2cca050e6e27837b60436d911352f8dca203a9cd35241c.
//
// Solidity: event Retire(_candidate address, _cap uint256)
func (_xdcValidator *xdcValidatorFilterer) WatchRetire(opts *bind.WatchOpts, sink chan<- *xdcValidatorRetire) (event.Subscription, error) {

	logs, sub, err := _xdcValidator.contract.WatchLogs(opts, "Retire")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(xdcValidatorRetire)
				if err := _xdcValidator.contract.UnpackLog(event, "Retire", log); err != nil {
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
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnvote is a free log retrieval operation binding the contract event 0x23ae40ca85f8a7c921ebb4269dc9a81e8a6de8dba614752927e0ff39341392fc.
//
// Solidity: event Unvote(_candidate address, _cap uint256)
func (_xdcValidator *xdcValidatorFilterer) FilterUnvote(opts *bind.FilterOpts) (*xdcValidatorUnvoteIterator, error) {

	logs, sub, err := _xdcValidator.contract.FilterLogs(opts, "Unvote")
	if err != nil {
		return nil, err
	}
	return &xdcValidatorUnvoteIterator{contract: _xdcValidator.contract, event: "Unvote", logs: logs, sub: sub}, nil
}

// WatchUnvote is a free log subscription operation binding the contract event 0x23ae40ca85f8a7c921ebb4269dc9a81e8a6de8dba614752927e0ff39341392fc.
//
// Solidity: event Unvote(_candidate address, _cap uint256)
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
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0xf668ead05c744b9178e571d2edb452e72baf6529c8d72160e64e59b50d865bd0.
//
// Solidity: event Vote(_candidate address, _cap uint256)
func (_xdcValidator *xdcValidatorFilterer) FilterVote(opts *bind.FilterOpts) (*xdcValidatorVoteIterator, error) {

	logs, sub, err := _xdcValidator.contract.FilterLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return &xdcValidatorVoteIterator{contract: _xdcValidator.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0xf668ead05c744b9178e571d2edb452e72baf6529c8d72160e64e59b50d865bd0.
//
// Solidity: event Vote(_candidate address, _cap uint256)
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
