// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rewards

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RewardsABI is the input ABI used to generate the binding from.
const RewardsABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blockHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockSignerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainReward\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"blockHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"standbyNodes\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"calculateRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochsWithRewardsCalculated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochsWithRewardsCalculatedForNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blockSignerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardTransferEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initialEpoch\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"masterNodeBlocksConfirmedHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingRewardTransferAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingRewardsTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardTransferEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"setCurrentEpochByOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"setRewardTransferEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasuryAddress\",\"type\":\"address\"}],\"name\":\"setTreasuryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"slashedNodes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSlashed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"untilEpoch\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"standbyNodeBlocksConfirmedHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"standbyNodeHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Rewards is an auto generated Go binding around an Ethereum contract.
type Rewards struct {
	RewardsCaller     // Read-only binding to the contract
	RewardsTransactor // Write-only binding to the contract
	RewardsFilterer   // Log filterer for contract events
}

// RewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardsSession struct {
	Contract     *Rewards          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardsCallerSession struct {
	Contract *RewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardsTransactorSession struct {
	Contract     *RewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardsRaw struct {
	Contract *Rewards // Generic contract binding to access the raw methods on
}

// RewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardsCallerRaw struct {
	Contract *RewardsCaller // Generic read-only contract binding to access the raw methods on
}

// RewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardsTransactorRaw struct {
	Contract *RewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewards creates a new instance of Rewards, bound to a specific deployed contract.
func NewRewards(address common.Address, backend bind.ContractBackend) (*Rewards, error) {
	contract, err := bindRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rewards{RewardsCaller: RewardsCaller{contract: contract}, RewardsTransactor: RewardsTransactor{contract: contract}, RewardsFilterer: RewardsFilterer{contract: contract}}, nil
}

// NewRewardsCaller creates a new read-only instance of Rewards, bound to a specific deployed contract.
func NewRewardsCaller(address common.Address, caller bind.ContractCaller) (*RewardsCaller, error) {
	contract, err := bindRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsCaller{contract: contract}, nil
}

// NewRewardsTransactor creates a new write-only instance of Rewards, bound to a specific deployed contract.
func NewRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardsTransactor, error) {
	contract, err := bindRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsTransactor{contract: contract}, nil
}

// NewRewardsFilterer creates a new log filterer instance of Rewards, bound to a specific deployed contract.
func NewRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardsFilterer, error) {
	contract, err := bindRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardsFilterer{contract: contract}, nil
}

// bindRewards binds a generic wrapper to an already deployed contract.
func bindRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewards *RewardsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rewards.Contract.RewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewards *RewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.Contract.RewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewards *RewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewards.Contract.RewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewards *RewardsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewards *RewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewards *RewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewards.Contract.contract.Transact(opts, method, params...)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address _address) returns()
func (_Rewards *RewardsTransactor) AddWhitelisted(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "addWhitelisted", _address)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address _address) returns()
func (_Rewards *RewardsSession) AddWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.AddWhitelisted(&_Rewards.TransactOpts, _address)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address _address) returns()
func (_Rewards *RewardsTransactorSession) AddWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.AddWhitelisted(&_Rewards.TransactOpts, _address)
}

// BlockHistory is a paid mutator transaction binding the contract method 0x648729ae.
//
// Solidity: function blockHistory(address , uint256 ) returns(uint256)
func (_Rewards *RewardsTransactor) BlockHistory(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "blockHistory", arg0, arg1)
}

// BlockHistory is a paid mutator transaction binding the contract method 0x648729ae.
//
// Solidity: function blockHistory(address , uint256 ) returns(uint256)
func (_Rewards *RewardsSession) BlockHistory(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.BlockHistory(&_Rewards.TransactOpts, arg0, arg1)
}

// BlockHistory is a paid mutator transaction binding the contract method 0x648729ae.
//
// Solidity: function blockHistory(address , uint256 ) returns(uint256)
func (_Rewards *RewardsTransactorSession) BlockHistory(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.BlockHistory(&_Rewards.TransactOpts, arg0, arg1)
}

// BlockSignerAddress is a paid mutator transaction binding the contract method 0x78a13230.
//
// Solidity: function blockSignerAddress() returns(address)
func (_Rewards *RewardsTransactor) BlockSignerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "blockSignerAddress")
}

// BlockSignerAddress is a paid mutator transaction binding the contract method 0x78a13230.
//
// Solidity: function blockSignerAddress() returns(address)
func (_Rewards *RewardsSession) BlockSignerAddress() (*types.Transaction, error) {
	return _Rewards.Contract.BlockSignerAddress(&_Rewards.TransactOpts)
}

// BlockSignerAddress is a paid mutator transaction binding the contract method 0x78a13230.
//
// Solidity: function blockSignerAddress() returns(address)
func (_Rewards *RewardsTransactorSession) BlockSignerAddress() (*types.Transaction, error) {
	return _Rewards.Contract.BlockSignerAddress(&_Rewards.TransactOpts)
}

// CalculateRewards is a paid mutator transaction binding the contract method 0xc016f923.
//
// Solidity: function calculateRewards(uint256 chainReward, bytes32[] blockHashes, address[] standbyNodes, uint256 epoch) returns()
func (_Rewards *RewardsTransactor) CalculateRewards(opts *bind.TransactOpts, chainReward *big.Int, blockHashes [][32]byte, standbyNodes []common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "calculateRewards", chainReward, blockHashes, standbyNodes, epoch)
}

// CalculateRewards is a paid mutator transaction binding the contract method 0xc016f923.
//
// Solidity: function calculateRewards(uint256 chainReward, bytes32[] blockHashes, address[] standbyNodes, uint256 epoch) returns()
func (_Rewards *RewardsSession) CalculateRewards(chainReward *big.Int, blockHashes [][32]byte, standbyNodes []common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.CalculateRewards(&_Rewards.TransactOpts, chainReward, blockHashes, standbyNodes, epoch)
}

// CalculateRewards is a paid mutator transaction binding the contract method 0xc016f923.
//
// Solidity: function calculateRewards(uint256 chainReward, bytes32[] blockHashes, address[] standbyNodes, uint256 epoch) returns()
func (_Rewards *RewardsTransactorSession) CalculateRewards(chainReward *big.Int, blockHashes [][32]byte, standbyNodes []common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.CalculateRewards(&_Rewards.TransactOpts, chainReward, blockHashes, standbyNodes, epoch)
}

// CurrentEpoch is a paid mutator transaction binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() returns(uint256)
func (_Rewards *RewardsTransactor) CurrentEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "currentEpoch")
}

// CurrentEpoch is a paid mutator transaction binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() returns(uint256)
func (_Rewards *RewardsSession) CurrentEpoch() (*types.Transaction, error) {
	return _Rewards.Contract.CurrentEpoch(&_Rewards.TransactOpts)
}

// CurrentEpoch is a paid mutator transaction binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() returns(uint256)
func (_Rewards *RewardsTransactorSession) CurrentEpoch() (*types.Transaction, error) {
	return _Rewards.Contract.CurrentEpoch(&_Rewards.TransactOpts)
}

// EpochsWithRewardsCalculated is a paid mutator transaction binding the contract method 0xdba6ba43.
//
// Solidity: function epochsWithRewardsCalculated(uint256 ) returns(bool)
func (_Rewards *RewardsTransactor) EpochsWithRewardsCalculated(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "epochsWithRewardsCalculated", arg0)
}

// EpochsWithRewardsCalculated is a paid mutator transaction binding the contract method 0xdba6ba43.
//
// Solidity: function epochsWithRewardsCalculated(uint256 ) returns(bool)
func (_Rewards *RewardsSession) EpochsWithRewardsCalculated(arg0 *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.EpochsWithRewardsCalculated(&_Rewards.TransactOpts, arg0)
}

// EpochsWithRewardsCalculated is a paid mutator transaction binding the contract method 0xdba6ba43.
//
// Solidity: function epochsWithRewardsCalculated(uint256 ) returns(bool)
func (_Rewards *RewardsTransactorSession) EpochsWithRewardsCalculated(arg0 *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.EpochsWithRewardsCalculated(&_Rewards.TransactOpts, arg0)
}

// EpochsWithRewardsCalculatedForNode is a paid mutator transaction binding the contract method 0x38a32416.
//
// Solidity: function epochsWithRewardsCalculatedForNode(uint256 , address ) returns(bool)
func (_Rewards *RewardsTransactor) EpochsWithRewardsCalculatedForNode(opts *bind.TransactOpts, arg0 *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "epochsWithRewardsCalculatedForNode", arg0, arg1)
}

// EpochsWithRewardsCalculatedForNode is a paid mutator transaction binding the contract method 0x38a32416.
//
// Solidity: function epochsWithRewardsCalculatedForNode(uint256 , address ) returns(bool)
func (_Rewards *RewardsSession) EpochsWithRewardsCalculatedForNode(arg0 *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.EpochsWithRewardsCalculatedForNode(&_Rewards.TransactOpts, arg0, arg1)
}

// EpochsWithRewardsCalculatedForNode is a paid mutator transaction binding the contract method 0x38a32416.
//
// Solidity: function epochsWithRewardsCalculatedForNode(uint256 , address ) returns(bool)
func (_Rewards *RewardsTransactorSession) EpochsWithRewardsCalculatedForNode(arg0 *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.EpochsWithRewardsCalculatedForNode(&_Rewards.TransactOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _blockSignerAddress, uint256 _rewardTransferEpoch, uint256 _initialEpoch) returns()
func (_Rewards *RewardsTransactor) Initialize(opts *bind.TransactOpts, _blockSignerAddress common.Address, _rewardTransferEpoch *big.Int, _initialEpoch *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "initialize", _blockSignerAddress, _rewardTransferEpoch, _initialEpoch)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _blockSignerAddress, uint256 _rewardTransferEpoch, uint256 _initialEpoch) returns()
func (_Rewards *RewardsSession) Initialize(_blockSignerAddress common.Address, _rewardTransferEpoch *big.Int, _initialEpoch *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.Initialize(&_Rewards.TransactOpts, _blockSignerAddress, _rewardTransferEpoch, _initialEpoch)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _blockSignerAddress, uint256 _rewardTransferEpoch, uint256 _initialEpoch) returns()
func (_Rewards *RewardsTransactorSession) Initialize(_blockSignerAddress common.Address, _rewardTransferEpoch *big.Int, _initialEpoch *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.Initialize(&_Rewards.TransactOpts, _blockSignerAddress, _rewardTransferEpoch, _initialEpoch)
}

// MasterNodeBlocksConfirmedHistory is a paid mutator transaction binding the contract method 0xd7a2099d.
//
// Solidity: function masterNodeBlocksConfirmedHistory(uint256 ) returns(uint256)
func (_Rewards *RewardsTransactor) MasterNodeBlocksConfirmedHistory(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "masterNodeBlocksConfirmedHistory", arg0)
}

// MasterNodeBlocksConfirmedHistory is a paid mutator transaction binding the contract method 0xd7a2099d.
//
// Solidity: function masterNodeBlocksConfirmedHistory(uint256 ) returns(uint256)
func (_Rewards *RewardsSession) MasterNodeBlocksConfirmedHistory(arg0 *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.MasterNodeBlocksConfirmedHistory(&_Rewards.TransactOpts, arg0)
}

// MasterNodeBlocksConfirmedHistory is a paid mutator transaction binding the contract method 0xd7a2099d.
//
// Solidity: function masterNodeBlocksConfirmedHistory(uint256 ) returns(uint256)
func (_Rewards *RewardsTransactorSession) MasterNodeBlocksConfirmedHistory(arg0 *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.MasterNodeBlocksConfirmedHistory(&_Rewards.TransactOpts, arg0)
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Rewards *RewardsTransactor) Owner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "owner")
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Rewards *RewardsSession) Owner() (*types.Transaction, error) {
	return _Rewards.Contract.Owner(&_Rewards.TransactOpts)
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Rewards *RewardsTransactorSession) Owner() (*types.Transaction, error) {
	return _Rewards.Contract.Owner(&_Rewards.TransactOpts)
}

// PendingRewardTransferAddresses is a paid mutator transaction binding the contract method 0x7c6b6611.
//
// Solidity: function pendingRewardTransferAddresses(uint256 ) returns(address)
func (_Rewards *RewardsTransactor) PendingRewardTransferAddresses(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "pendingRewardTransferAddresses", arg0)
}

// PendingRewardTransferAddresses is a paid mutator transaction binding the contract method 0x7c6b6611.
//
// Solidity: function pendingRewardTransferAddresses(uint256 ) returns(address)
func (_Rewards *RewardsSession) PendingRewardTransferAddresses(arg0 *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.PendingRewardTransferAddresses(&_Rewards.TransactOpts, arg0)
}

// PendingRewardTransferAddresses is a paid mutator transaction binding the contract method 0x7c6b6611.
//
// Solidity: function pendingRewardTransferAddresses(uint256 ) returns(address)
func (_Rewards *RewardsTransactorSession) PendingRewardTransferAddresses(arg0 *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.PendingRewardTransferAddresses(&_Rewards.TransactOpts, arg0)
}

// PendingRewardsTransaction is a paid mutator transaction binding the contract method 0xa150604d.
//
// Solidity: function pendingRewardsTransaction(address ) returns(uint256)
func (_Rewards *RewardsTransactor) PendingRewardsTransaction(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "pendingRewardsTransaction", arg0)
}

// PendingRewardsTransaction is a paid mutator transaction binding the contract method 0xa150604d.
//
// Solidity: function pendingRewardsTransaction(address ) returns(uint256)
func (_Rewards *RewardsSession) PendingRewardsTransaction(arg0 common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.PendingRewardsTransaction(&_Rewards.TransactOpts, arg0)
}

// PendingRewardsTransaction is a paid mutator transaction binding the contract method 0xa150604d.
//
// Solidity: function pendingRewardsTransaction(address ) returns(uint256)
func (_Rewards *RewardsTransactorSession) PendingRewardsTransaction(arg0 common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.PendingRewardsTransaction(&_Rewards.TransactOpts, arg0)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns()
func (_Rewards *RewardsTransactor) RemoveWhitelisted(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "removeWhitelisted", _address)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns()
func (_Rewards *RewardsSession) RemoveWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.RemoveWhitelisted(&_Rewards.TransactOpts, _address)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address _address) returns()
func (_Rewards *RewardsTransactorSession) RemoveWhitelisted(_address common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.RemoveWhitelisted(&_Rewards.TransactOpts, _address)
}

// RewardTransferEpoch is a paid mutator transaction binding the contract method 0x29c6f90c.
//
// Solidity: function rewardTransferEpoch() returns(uint256)
func (_Rewards *RewardsTransactor) RewardTransferEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "rewardTransferEpoch")
}

// RewardTransferEpoch is a paid mutator transaction binding the contract method 0x29c6f90c.
//
// Solidity: function rewardTransferEpoch() returns(uint256)
func (_Rewards *RewardsSession) RewardTransferEpoch() (*types.Transaction, error) {
	return _Rewards.Contract.RewardTransferEpoch(&_Rewards.TransactOpts)
}

// RewardTransferEpoch is a paid mutator transaction binding the contract method 0x29c6f90c.
//
// Solidity: function rewardTransferEpoch() returns(uint256)
func (_Rewards *RewardsTransactorSession) RewardTransferEpoch() (*types.Transaction, error) {
	return _Rewards.Contract.RewardTransferEpoch(&_Rewards.TransactOpts)
}

// SetCurrentEpochByOwner is a paid mutator transaction binding the contract method 0xdf0426f6.
//
// Solidity: function setCurrentEpochByOwner(uint256 epoch) returns()
func (_Rewards *RewardsTransactor) SetCurrentEpochByOwner(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "setCurrentEpochByOwner", epoch)
}

// SetCurrentEpochByOwner is a paid mutator transaction binding the contract method 0xdf0426f6.
//
// Solidity: function setCurrentEpochByOwner(uint256 epoch) returns()
func (_Rewards *RewardsSession) SetCurrentEpochByOwner(epoch *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.SetCurrentEpochByOwner(&_Rewards.TransactOpts, epoch)
}

// SetCurrentEpochByOwner is a paid mutator transaction binding the contract method 0xdf0426f6.
//
// Solidity: function setCurrentEpochByOwner(uint256 epoch) returns()
func (_Rewards *RewardsTransactorSession) SetCurrentEpochByOwner(epoch *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.SetCurrentEpochByOwner(&_Rewards.TransactOpts, epoch)
}

// SetRewardTransferEpoch is a paid mutator transaction binding the contract method 0x8040ea70.
//
// Solidity: function setRewardTransferEpoch(uint256 _epoch) returns()
func (_Rewards *RewardsTransactor) SetRewardTransferEpoch(opts *bind.TransactOpts, _epoch *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "setRewardTransferEpoch", _epoch)
}

// SetRewardTransferEpoch is a paid mutator transaction binding the contract method 0x8040ea70.
//
// Solidity: function setRewardTransferEpoch(uint256 _epoch) returns()
func (_Rewards *RewardsSession) SetRewardTransferEpoch(_epoch *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.SetRewardTransferEpoch(&_Rewards.TransactOpts, _epoch)
}

// SetRewardTransferEpoch is a paid mutator transaction binding the contract method 0x8040ea70.
//
// Solidity: function setRewardTransferEpoch(uint256 _epoch) returns()
func (_Rewards *RewardsTransactorSession) SetRewardTransferEpoch(_epoch *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.SetRewardTransferEpoch(&_Rewards.TransactOpts, _epoch)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_Rewards *RewardsTransactor) SetTreasuryAddress(opts *bind.TransactOpts, _treasuryAddress common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "setTreasuryAddress", _treasuryAddress)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_Rewards *RewardsSession) SetTreasuryAddress(_treasuryAddress common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SetTreasuryAddress(&_Rewards.TransactOpts, _treasuryAddress)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_Rewards *RewardsTransactorSession) SetTreasuryAddress(_treasuryAddress common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SetTreasuryAddress(&_Rewards.TransactOpts, _treasuryAddress)
}

// SlashedNodes is a paid mutator transaction binding the contract method 0xfb2594be.
//
// Solidity: function slashedNodes(address ) returns(bool isSlashed, uint256 untilEpoch)
func (_Rewards *RewardsTransactor) SlashedNodes(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "slashedNodes", arg0)
}

// SlashedNodes is a paid mutator transaction binding the contract method 0xfb2594be.
//
// Solidity: function slashedNodes(address ) returns(bool isSlashed, uint256 untilEpoch)
func (_Rewards *RewardsSession) SlashedNodes(arg0 common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SlashedNodes(&_Rewards.TransactOpts, arg0)
}

// SlashedNodes is a paid mutator transaction binding the contract method 0xfb2594be.
//
// Solidity: function slashedNodes(address ) returns(bool isSlashed, uint256 untilEpoch)
func (_Rewards *RewardsTransactorSession) SlashedNodes(arg0 common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SlashedNodes(&_Rewards.TransactOpts, arg0)
}

// StandbyNodeBlocksConfirmedHistory is a paid mutator transaction binding the contract method 0x5614cf56.
//
// Solidity: function standbyNodeBlocksConfirmedHistory(uint256 ) returns(uint256)
func (_Rewards *RewardsTransactor) StandbyNodeBlocksConfirmedHistory(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "standbyNodeBlocksConfirmedHistory", arg0)
}

// StandbyNodeBlocksConfirmedHistory is a paid mutator transaction binding the contract method 0x5614cf56.
//
// Solidity: function standbyNodeBlocksConfirmedHistory(uint256 ) returns(uint256)
func (_Rewards *RewardsSession) StandbyNodeBlocksConfirmedHistory(arg0 *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.StandbyNodeBlocksConfirmedHistory(&_Rewards.TransactOpts, arg0)
}

// StandbyNodeBlocksConfirmedHistory is a paid mutator transaction binding the contract method 0x5614cf56.
//
// Solidity: function standbyNodeBlocksConfirmedHistory(uint256 ) returns(uint256)
func (_Rewards *RewardsTransactorSession) StandbyNodeBlocksConfirmedHistory(arg0 *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.StandbyNodeBlocksConfirmedHistory(&_Rewards.TransactOpts, arg0)
}

// StandbyNodeHistory is a paid mutator transaction binding the contract method 0x3f090603.
//
// Solidity: function standbyNodeHistory(address ) returns(uint256)
func (_Rewards *RewardsTransactor) StandbyNodeHistory(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "standbyNodeHistory", arg0)
}

// StandbyNodeHistory is a paid mutator transaction binding the contract method 0x3f090603.
//
// Solidity: function standbyNodeHistory(address ) returns(uint256)
func (_Rewards *RewardsSession) StandbyNodeHistory(arg0 common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.StandbyNodeHistory(&_Rewards.TransactOpts, arg0)
}

// StandbyNodeHistory is a paid mutator transaction binding the contract method 0x3f090603.
//
// Solidity: function standbyNodeHistory(address ) returns(uint256)
func (_Rewards *RewardsTransactorSession) StandbyNodeHistory(arg0 common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.StandbyNodeHistory(&_Rewards.TransactOpts, arg0)
}

// TreasuryAddress is a paid mutator transaction binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() returns(address)
func (_Rewards *RewardsTransactor) TreasuryAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "treasuryAddress")
}

// TreasuryAddress is a paid mutator transaction binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() returns(address)
func (_Rewards *RewardsSession) TreasuryAddress() (*types.Transaction, error) {
	return _Rewards.Contract.TreasuryAddress(&_Rewards.TransactOpts)
}

// TreasuryAddress is a paid mutator transaction binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() returns(address)
func (_Rewards *RewardsTransactorSession) TreasuryAddress() (*types.Transaction, error) {
	return _Rewards.Contract.TreasuryAddress(&_Rewards.TransactOpts)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) returns(bool)
func (_Rewards *RewardsTransactor) Whitelist(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "whitelist", arg0)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) returns(bool)
func (_Rewards *RewardsSession) Whitelist(arg0 common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.Whitelist(&_Rewards.TransactOpts, arg0)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) returns(bool)
func (_Rewards *RewardsTransactorSession) Whitelist(arg0 common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.Whitelist(&_Rewards.TransactOpts, arg0)
}
