// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mkr

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

// SaiProxyCreateAndExecuteABI is the input ABI used to generate the binding from.
const SaiProxyCreateAndExecuteABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"registry_\",\"type\":\"address\"},{\"name\":\"tub_\",\"type\":\"address\"}],\"name\":\"createAndOpen\",\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registry_\",\"type\":\"address\"},{\"name\":\"tub_\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"createOpenLockAndDraw\",\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registry_\",\"type\":\"address\"},{\"name\":\"tub_\",\"type\":\"address\"}],\"name\":\"createOpenAndLock\",\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// SaiProxyCreateAndExecuteBin is the compiled bytecode used for deploying new contracts.
const SaiProxyCreateAndExecuteBin = `0x`

// DeploySaiProxyCreateAndExecute deploys a new Ethereum contract, binding an instance of SaiProxyCreateAndExecute to it.
func DeploySaiProxyCreateAndExecute(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SaiProxyCreateAndExecute, error) {
	parsed, err := abi.JSON(strings.NewReader(SaiProxyCreateAndExecuteABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SaiProxyCreateAndExecuteBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SaiProxyCreateAndExecute{SaiProxyCreateAndExecuteCaller: SaiProxyCreateAndExecuteCaller{contract: contract}, SaiProxyCreateAndExecuteTransactor: SaiProxyCreateAndExecuteTransactor{contract: contract}, SaiProxyCreateAndExecuteFilterer: SaiProxyCreateAndExecuteFilterer{contract: contract}}, nil
}

// SaiProxyCreateAndExecute is an auto generated Go binding around an Ethereum contract.
type SaiProxyCreateAndExecute struct {
	SaiProxyCreateAndExecuteCaller     // Read-only binding to the contract
	SaiProxyCreateAndExecuteTransactor // Write-only binding to the contract
	SaiProxyCreateAndExecuteFilterer   // Log filterer for contract events
}

// SaiProxyCreateAndExecuteCaller is an auto generated read-only Go binding around an Ethereum contract.
type SaiProxyCreateAndExecuteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaiProxyCreateAndExecuteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SaiProxyCreateAndExecuteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaiProxyCreateAndExecuteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SaiProxyCreateAndExecuteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaiProxyCreateAndExecuteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SaiProxyCreateAndExecuteSession struct {
	Contract     *SaiProxyCreateAndExecute // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SaiProxyCreateAndExecuteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SaiProxyCreateAndExecuteCallerSession struct {
	Contract *SaiProxyCreateAndExecuteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// SaiProxyCreateAndExecuteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SaiProxyCreateAndExecuteTransactorSession struct {
	Contract     *SaiProxyCreateAndExecuteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// SaiProxyCreateAndExecuteRaw is an auto generated low-level Go binding around an Ethereum contract.
type SaiProxyCreateAndExecuteRaw struct {
	Contract *SaiProxyCreateAndExecute // Generic contract binding to access the raw methods on
}

// SaiProxyCreateAndExecuteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SaiProxyCreateAndExecuteCallerRaw struct {
	Contract *SaiProxyCreateAndExecuteCaller // Generic read-only contract binding to access the raw methods on
}

// SaiProxyCreateAndExecuteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SaiProxyCreateAndExecuteTransactorRaw struct {
	Contract *SaiProxyCreateAndExecuteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSaiProxyCreateAndExecute creates a new instance of SaiProxyCreateAndExecute, bound to a specific deployed contract.
func NewSaiProxyCreateAndExecute(address common.Address, backend bind.ContractBackend) (*SaiProxyCreateAndExecute, error) {
	contract, err := bindSaiProxyCreateAndExecute(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SaiProxyCreateAndExecute{SaiProxyCreateAndExecuteCaller: SaiProxyCreateAndExecuteCaller{contract: contract}, SaiProxyCreateAndExecuteTransactor: SaiProxyCreateAndExecuteTransactor{contract: contract}, SaiProxyCreateAndExecuteFilterer: SaiProxyCreateAndExecuteFilterer{contract: contract}}, nil
}

// NewSaiProxyCreateAndExecuteCaller creates a new read-only instance of SaiProxyCreateAndExecute, bound to a specific deployed contract.
func NewSaiProxyCreateAndExecuteCaller(address common.Address, caller bind.ContractCaller) (*SaiProxyCreateAndExecuteCaller, error) {
	contract, err := bindSaiProxyCreateAndExecute(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SaiProxyCreateAndExecuteCaller{contract: contract}, nil
}

// NewSaiProxyCreateAndExecuteTransactor creates a new write-only instance of SaiProxyCreateAndExecute, bound to a specific deployed contract.
func NewSaiProxyCreateAndExecuteTransactor(address common.Address, transactor bind.ContractTransactor) (*SaiProxyCreateAndExecuteTransactor, error) {
	contract, err := bindSaiProxyCreateAndExecute(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SaiProxyCreateAndExecuteTransactor{contract: contract}, nil
}

// NewSaiProxyCreateAndExecuteFilterer creates a new log filterer instance of SaiProxyCreateAndExecute, bound to a specific deployed contract.
func NewSaiProxyCreateAndExecuteFilterer(address common.Address, filterer bind.ContractFilterer) (*SaiProxyCreateAndExecuteFilterer, error) {
	contract, err := bindSaiProxyCreateAndExecute(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SaiProxyCreateAndExecuteFilterer{contract: contract}, nil
}

// bindSaiProxyCreateAndExecute binds a generic wrapper to an already deployed contract.
func bindSaiProxyCreateAndExecute(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SaiProxyCreateAndExecuteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SaiProxyCreateAndExecute *SaiProxyCreateAndExecuteRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SaiProxyCreateAndExecute.Contract.SaiProxyCreateAndExecuteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SaiProxyCreateAndExecute *SaiProxyCreateAndExecuteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaiProxyCreateAndExecute.Contract.SaiProxyCreateAndExecuteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SaiProxyCreateAndExecute *SaiProxyCreateAndExecuteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SaiProxyCreateAndExecute.Contract.SaiProxyCreateAndExecuteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SaiProxyCreateAndExecute *SaiProxyCreateAndExecuteCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SaiProxyCreateAndExecute.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SaiProxyCreateAndExecute *SaiProxyCreateAndExecuteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaiProxyCreateAndExecute.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SaiProxyCreateAndExecute *SaiProxyCreateAndExecuteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SaiProxyCreateAndExecute.Contract.contract.Transact(opts, method, params...)
}

// CreateAndOpen is a paid mutator transaction binding the contract method 0x581f3c50.
//
// Solidity: function createAndOpen(address registry_, address tub_) returns(address proxy, bytes32 cup)
func (_SaiProxyCreateAndExecute *SaiProxyCreateAndExecuteTransactor) CreateAndOpen(opts *bind.TransactOpts, registry_ common.Address, tub_ common.Address) (*types.Transaction, error) {
	return _SaiProxyCreateAndExecute.contract.Transact(opts, "createAndOpen", registry_, tub_)
}

// CreateAndOpen is a paid mutator transaction binding the contract method 0x581f3c50.
//
// Solidity: function createAndOpen(address registry_, address tub_) returns(address proxy, bytes32 cup)
func (_SaiProxyCreateAndExecute *SaiProxyCreateAndExecuteSession) CreateAndOpen(registry_ common.Address, tub_ common.Address) (*types.Transaction, error) {
	return _SaiProxyCreateAndExecute.Contract.CreateAndOpen(&_SaiProxyCreateAndExecute.TransactOpts, registry_, tub_)
}

// CreateAndOpen is a paid mutator transaction binding the contract method 0x581f3c50.
//
// Solidity: function createAndOpen(address registry_, address tub_) returns(address proxy, bytes32 cup)
func (_SaiProxyCreateAndExecute *SaiProxyCreateAndExecuteTransactorSession) CreateAndOpen(registry_ common.Address, tub_ common.Address) (*types.Transaction, error) {
	return _SaiProxyCreateAndExecute.Contract.CreateAndOpen(&_SaiProxyCreateAndExecute.TransactOpts, registry_, tub_)
}

// CreateOpenAndLock is a paid mutator transaction binding the contract method 0xeefe3818.
//
// Solidity: function createOpenAndLock(address registry_, address tub_) returns(address proxy, bytes32 cup)
func (_SaiProxyCreateAndExecute *SaiProxyCreateAndExecuteTransactor) CreateOpenAndLock(opts *bind.TransactOpts, registry_ common.Address, tub_ common.Address) (*types.Transaction, error) {
	return _SaiProxyCreateAndExecute.contract.Transact(opts, "createOpenAndLock", registry_, tub_)
}

// CreateOpenAndLock is a paid mutator transaction binding the contract method 0xeefe3818.
//
// Solidity: function createOpenAndLock(address registry_, address tub_) returns(address proxy, bytes32 cup)
func (_SaiProxyCreateAndExecute *SaiProxyCreateAndExecuteSession) CreateOpenAndLock(registry_ common.Address, tub_ common.Address) (*types.Transaction, error) {
	return _SaiProxyCreateAndExecute.Contract.CreateOpenAndLock(&_SaiProxyCreateAndExecute.TransactOpts, registry_, tub_)
}

// CreateOpenAndLock is a paid mutator transaction binding the contract method 0xeefe3818.
//
// Solidity: function createOpenAndLock(address registry_, address tub_) returns(address proxy, bytes32 cup)
func (_SaiProxyCreateAndExecute *SaiProxyCreateAndExecuteTransactorSession) CreateOpenAndLock(registry_ common.Address, tub_ common.Address) (*types.Transaction, error) {
	return _SaiProxyCreateAndExecute.Contract.CreateOpenAndLock(&_SaiProxyCreateAndExecute.TransactOpts, registry_, tub_)
}

// CreateOpenLockAndDraw is a paid mutator transaction binding the contract method 0xd3140a65.
//
// Solidity: function createOpenLockAndDraw(address registry_, address tub_, uint256 wad) returns(address proxy, bytes32 cup)
func (_SaiProxyCreateAndExecute *SaiProxyCreateAndExecuteTransactor) CreateOpenLockAndDraw(opts *bind.TransactOpts, registry_ common.Address, tub_ common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SaiProxyCreateAndExecute.contract.Transact(opts, "createOpenLockAndDraw", registry_, tub_, wad)
}

// CreateOpenLockAndDraw is a paid mutator transaction binding the contract method 0xd3140a65.
//
// Solidity: function createOpenLockAndDraw(address registry_, address tub_, uint256 wad) returns(address proxy, bytes32 cup)
func (_SaiProxyCreateAndExecute *SaiProxyCreateAndExecuteSession) CreateOpenLockAndDraw(registry_ common.Address, tub_ common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SaiProxyCreateAndExecute.Contract.CreateOpenLockAndDraw(&_SaiProxyCreateAndExecute.TransactOpts, registry_, tub_, wad)
}

// CreateOpenLockAndDraw is a paid mutator transaction binding the contract method 0xd3140a65.
//
// Solidity: function createOpenLockAndDraw(address registry_, address tub_, uint256 wad) returns(address proxy, bytes32 cup)
func (_SaiProxyCreateAndExecute *SaiProxyCreateAndExecuteTransactorSession) CreateOpenLockAndDraw(registry_ common.Address, tub_ common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SaiProxyCreateAndExecute.Contract.CreateOpenLockAndDraw(&_SaiProxyCreateAndExecute.TransactOpts, registry_, tub_, wad)
}
