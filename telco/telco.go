// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package telco

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

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// IERC20Bin is the compiled bytecode used for deploying new contracts.
const IERC20Bin = `0x`

// DeployIERC20 deploys a new Ethereum contract, binding an instance of IERC20 to it.
func DeployIERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// TelcoABI is the input ABI used to generate the binding from.
const TelcoABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"userDirectWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"getUserBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_phoneNumber\",\"type\":\"string\"},{\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"registerNewUser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_phoneNumber\",\"type\":\"string\"},{\"name\":\"_password\",\"type\":\"bytes\"}],\"name\":\"userInDirectWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"users\",\"outputs\":[{\"name\":\"phoneNumber\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_phoneNumber\",\"type\":\"string\"},{\"name\":\"_password\",\"type\":\"bytes\"}],\"name\":\"checkPasswordValid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_phoneNumber\",\"type\":\"string\"},{\"name\":\"_password\",\"type\":\"bytes\"}],\"name\":\"getUserTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_phoneNumber\",\"type\":\"string\"}],\"name\":\"getFromPhoneNumberToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getUserPhoneNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_toPhoneNumber\",\"type\":\"string\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_phoneNumber\",\"type\":\"string\"},{\"name\":\"_password\",\"type\":\"bytes\"}],\"name\":\"authTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// TelcoBin is the compiled bytecode used for deploying new contracts.
const TelcoBin = `0x608060405234801561001057600080fd5b50600080546001600160a01b03191633178082556040516001600160a01b039190911691907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3611562806100696000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063a87430ba1161008c578063ca62cba311610066578063ca62cba3146105f1578063cc5662391461065f578063ef91a4ef14610685578063f2fde38b146107a7576100ea565b8063a87430ba146102ef578063b523ba9f1461038a578063c8a89482146104b3576100ea565b80637479601b116100c85780637479601b146101655780637799c341146101f05780638da5cb5b146102c35780638f32d59b146102e7576100ea565b80631872b42c146100ef5780636805d6ad1461011d578063715018a61461015d575b600080fd5b61011b6004803603604081101561010557600080fd5b506001600160a01b0381351690602001356107cd565b005b61014b6004803603604081101561013357600080fd5b506001600160a01b03813581169160200135166108e9565b60408051918252519081900360200190f35b61011b610915565b6101dc6004803603604081101561017b57600080fd5b810190602081018135600160201b81111561019557600080fd5b8201836020820111156101a757600080fd5b803590602001918460018302840111600160201b831117156101c857600080fd5b9193509150356001600160a01b0316610970565b604080519115158252519081900360200190f35b61011b6004803603608081101561020657600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561023557600080fd5b82018360208201111561024757600080fd5b803590602001918460018302840111600160201b8311171561026857600080fd5b919390929091602081019035600160201b81111561028557600080fd5b82018360208201111561029757600080fd5b803590602001918460018302840111600160201b831117156102b857600080fd5b509092509050610a30565b6102cb610c22565b604080516001600160a01b039092168252519081900360200190f35b6101dc610c32565b6103156004803603602081101561030557600080fd5b50356001600160a01b0316610c43565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561034f578181015183820152602001610337565b50505050905090810190601f16801561037c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101dc600480360360408110156103a057600080fd5b810190602081018135600160201b8111156103ba57600080fd5b8201836020820111156103cc57600080fd5b803590602001918460018302840111600160201b831117156103ed57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561043f57600080fd5b82018360208201111561045157600080fd5b803590602001918460018302840111600160201b8311171561047257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610cdf945050505050565b61011b600480360360808110156104c957600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b8111156104f857600080fd5b82018360208201111561050a57600080fd5b803590602001918460018302840111600160201b8311171561052b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561057d57600080fd5b82018360208201111561058f57600080fd5b803590602001918460018302840111600160201b831117156105b057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610e86945050505050565b6102cb6004803603602081101561060757600080fd5b810190602081018135600160201b81111561062157600080fd5b82018360208201111561063357600080fd5b803590602001918460018302840111600160201b8311171561065457600080fd5b509092509050611070565b6103156004803603602081101561067557600080fd5b50356001600160a01b03166110aa565b61011b600480360360a081101561069b57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156106c557600080fd5b8201836020820111156106d757600080fd5b803590602001918460018302840111600160201b831117156106f857600080fd5b91939092823592604081019060200135600160201b81111561071957600080fd5b82018360208201111561072b57600080fd5b803590602001918460018302840111600160201b8311171561074c57600080fd5b919390929091602081019035600160201b81111561076957600080fd5b82018360208201111561077b57600080fd5b803590602001918460018302840111600160201b8311171561079c57600080fd5b509092509050611154565b61011b600480360360208110156107bd57600080fd5b50356001600160a01b0316611413565b3360009081526001602081815260408084206001600160a01b03871685529283019091529091205482111561084c5760408051600160e51b62461bcd02815260206004820152601c60248201527f6e6f7420656e6f7567682066756e647320746f20776974686472617700000000604482015290519081900360640190fd5b6001600160a01b03831660008181526001830160209081526040808320805487900390558051600160e01b63a9059cbb02815233600482015260248101879052905187949363a9059cbb93604480850194919392918390030190829087803b1580156108b757600080fd5b505af11580156108cb573d6000803e3d6000fd5b505050506040513d60208110156108e157600080fd5b505050505050565b6001600160a01b0391821660009081526001602081815260408084209490951683529201909152205490565b61091d610c32565b61092657600080fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b604080516020601f850181900481028201830183528101848152600092829190879087908190850183828082843760009201829052509390945250506001600160a01b0385168152600160209081526040909120835180519193506109d992849291019061149e565b509050508160028585604051808383808284379190910194855250506040519283900360200190922080546001600160a01b03949094166001600160a01b0319909416939093179092555060019150509392505050565b610aa384848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f88018190048102820181019092528681529250869150859081908401838280828437600092019190915250610cdf92505050565b610aef5760408051600160e51b62461bcd0281526020600482015260126024820152600160721b711c185cdcdddbdc99081b9bdd081d985b1a5902604482015290519081900360640190fd5b600060028585604051808383808284378083019250505092505050908152602001604051809103902060009054906101000a90046001600160a01b0316905060008790508660016000846001600160a01b03166001600160a01b0316815260200190815260200160002060010160008a6001600160a01b03166001600160a01b0316815260200190815260200160002060008282540392505081905550806001600160a01b031663a9059cbb83896040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015610bec57600080fd5b505af1158015610c00573d6000803e3d6000fd5b505050506040513d6020811015610c1657600080fd5b50505050505050505050565b6000546001600160a01b03165b90565b6000546001600160a01b0316331490565b60016020818152600092835260409283902080548451600294821615610100026000190190911693909304601f8101839004830284018301909452838352928391830182828015610cd55780601f10610caa57610100808354040283529160200191610cd5565b820191906000526020600020905b815481529060010190602001808311610cb857829003601f168201915b5050505050905081565b6000806002846040518082805190602001908083835b60208310610d145780518252601f199092019160209182019101610cf5565b51815160001960209485036101000a019081169019919091161790529201948552506040519384900301909220546001600160a01b03169250505080610da45760408051600160e51b62461bcd02815260206004820152601360248201527f7573657220646f6573206e6f7420657869737400000000000000000000000000604482015290519081900360640190fd5b604051600160e01b63923905a702815260206004820181815285516024840152855184936001600160a01b0385169363923905a7938993909283926044909101919085019080838360005b83811015610e07578181015183820152602001610def565b50505050905090810190601f168015610e345780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b158015610e5157600080fd5b505afa158015610e65573d6000803e3d6000fd5b505050506040513d6020811015610e7b57600080fd5b505195945050505050565b610e8e610c32565b610e9757600080fd5b610ea18282610cdf565b610eed5760408051600160e51b62461bcd0281526020600482015260126024820152600160721b711c185cdcdddbdc99081b9bdd081d985b1a5902604482015290519081900360640190fd5b60006002836040518082805190602001908083835b60208310610f215780518252601f199092019160209182019101610f02565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382018520546001600160a01b03908116600081815260018552928320600160e01b6311173ecd0288523060048901818152938e1660248a0152604489018d9052608060648a019081528b5160848b01528b51939a5091988a98508897506311173ecd9691958f958f958e959194919360a40192918601918190849084905b83811015610fdf578181015183820152602001610fc7565b50505050905090810190601f16801561100c5780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b15801561102e57600080fd5b505af1158015611042573d6000803e3d6000fd5b5050506001600160a01b03909716600090815260019092016020525060409020805490940190935550505050565b6000600283836040518083838082843791909101948552505060405192839003602001909220546001600160a01b03169250505092915050565b6001600160a01b03811660009081526001602081815260409283902080548451600294821615610100026000190190911693909304601f810183900483028401830190945283835260609390918301828280156111485780601f1061111d57610100808354040283529160200191611148565b820191906000526020600020905b81548152906001019060200180831161112b57829003601f168201915b50505050509050919050565b6111c784848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f88018190048102820181019092528681529250869150859081908401838280828437600092019190915250610cdf92505050565b6112135760408051600160e51b62461bcd0281526020600482015260126024820152600160721b711c185cdcdddbdc99081b9bdd081d985b1a5902604482015290519081900360640190fd5b600060028585604051808383808284379190910194855250506040805160209481900385018120546001600160a01b031660008181526001875292909220600160e01b639da6874502825260048201958652602482018890529195509093859350839250639da6874591889188918190604401848480828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b1580156112c257600080fd5b505af11580156112d6573d6000803e3d6000fd5b505050506001600160a01b038b16600090815260018301602052604090205488111561134c5760408051600160e51b62461bcd02815260206004820152601260248201527f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b038b16600090815260018301602052604080822080548b90039055516002908c908c90808383808284378083019250505092505050908152602001604051809103902060009054906101000a90046001600160a01b031690508860016000836001600160a01b03166001600160a01b0316815260200190815260200160002060010160008e6001600160a01b03166001600160a01b0316815260200190815260200160002060008282540192505081905550505050505050505050505050565b61141b610c32565b61142457600080fd5b61142d81611430565b50565b6001600160a01b03811661144357600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106114df57805160ff191683800117855561150c565b8280016001018555821561150c579182015b8281111561150c5782518255916020019190600101906114f1565b5061151892915061151c565b5090565b610c2f91905b80821115611518576000815560010161152256fea165627a7a72305820a09d032fe95bef472bc156a5c19b4160ba9e00bc419f7c730a21a73fe567a4940029`

// DeployTelco deploys a new Ethereum contract, binding an instance of Telco to it.
func DeployTelco(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Telco, error) {
	parsed, err := abi.JSON(strings.NewReader(TelcoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TelcoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Telco{TelcoCaller: TelcoCaller{contract: contract}, TelcoTransactor: TelcoTransactor{contract: contract}, TelcoFilterer: TelcoFilterer{contract: contract}}, nil
}

// Telco is an auto generated Go binding around an Ethereum contract.
type Telco struct {
	TelcoCaller     // Read-only binding to the contract
	TelcoTransactor // Write-only binding to the contract
	TelcoFilterer   // Log filterer for contract events
}

// TelcoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TelcoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TelcoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TelcoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TelcoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TelcoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TelcoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TelcoSession struct {
	Contract     *Telco            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TelcoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TelcoCallerSession struct {
	Contract *TelcoCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TelcoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TelcoTransactorSession struct {
	Contract     *TelcoTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TelcoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TelcoRaw struct {
	Contract *Telco // Generic contract binding to access the raw methods on
}

// TelcoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TelcoCallerRaw struct {
	Contract *TelcoCaller // Generic read-only contract binding to access the raw methods on
}

// TelcoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TelcoTransactorRaw struct {
	Contract *TelcoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTelco creates a new instance of Telco, bound to a specific deployed contract.
func NewTelco(address common.Address, backend bind.ContractBackend) (*Telco, error) {
	contract, err := bindTelco(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Telco{TelcoCaller: TelcoCaller{contract: contract}, TelcoTransactor: TelcoTransactor{contract: contract}, TelcoFilterer: TelcoFilterer{contract: contract}}, nil
}

// NewTelcoCaller creates a new read-only instance of Telco, bound to a specific deployed contract.
func NewTelcoCaller(address common.Address, caller bind.ContractCaller) (*TelcoCaller, error) {
	contract, err := bindTelco(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TelcoCaller{contract: contract}, nil
}

// NewTelcoTransactor creates a new write-only instance of Telco, bound to a specific deployed contract.
func NewTelcoTransactor(address common.Address, transactor bind.ContractTransactor) (*TelcoTransactor, error) {
	contract, err := bindTelco(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TelcoTransactor{contract: contract}, nil
}

// NewTelcoFilterer creates a new log filterer instance of Telco, bound to a specific deployed contract.
func NewTelcoFilterer(address common.Address, filterer bind.ContractFilterer) (*TelcoFilterer, error) {
	contract, err := bindTelco(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TelcoFilterer{contract: contract}, nil
}

// bindTelco binds a generic wrapper to an already deployed contract.
func bindTelco(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TelcoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Telco *TelcoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Telco.Contract.TelcoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Telco *TelcoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Telco.Contract.TelcoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Telco *TelcoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Telco.Contract.TelcoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Telco *TelcoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Telco.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Telco *TelcoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Telco.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Telco *TelcoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Telco.Contract.contract.Transact(opts, method, params...)
}

// CheckPasswordValid is a free data retrieval call binding the contract method 0xb523ba9f.
//
// Solidity: function checkPasswordValid(string _phoneNumber, bytes _password) constant returns(bool)
func (_Telco *TelcoCaller) CheckPasswordValid(opts *bind.CallOpts, _phoneNumber string, _password []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Telco.contract.Call(opts, out, "checkPasswordValid", _phoneNumber, _password)
	return *ret0, err
}

// CheckPasswordValid is a free data retrieval call binding the contract method 0xb523ba9f.
//
// Solidity: function checkPasswordValid(string _phoneNumber, bytes _password) constant returns(bool)
func (_Telco *TelcoSession) CheckPasswordValid(_phoneNumber string, _password []byte) (bool, error) {
	return _Telco.Contract.CheckPasswordValid(&_Telco.CallOpts, _phoneNumber, _password)
}

// CheckPasswordValid is a free data retrieval call binding the contract method 0xb523ba9f.
//
// Solidity: function checkPasswordValid(string _phoneNumber, bytes _password) constant returns(bool)
func (_Telco *TelcoCallerSession) CheckPasswordValid(_phoneNumber string, _password []byte) (bool, error) {
	return _Telco.Contract.CheckPasswordValid(&_Telco.CallOpts, _phoneNumber, _password)
}

// GetFromPhoneNumberToAddress is a free data retrieval call binding the contract method 0xca62cba3.
//
// Solidity: function getFromPhoneNumberToAddress(string _phoneNumber) constant returns(address)
func (_Telco *TelcoCaller) GetFromPhoneNumberToAddress(opts *bind.CallOpts, _phoneNumber string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Telco.contract.Call(opts, out, "getFromPhoneNumberToAddress", _phoneNumber)
	return *ret0, err
}

// GetFromPhoneNumberToAddress is a free data retrieval call binding the contract method 0xca62cba3.
//
// Solidity: function getFromPhoneNumberToAddress(string _phoneNumber) constant returns(address)
func (_Telco *TelcoSession) GetFromPhoneNumberToAddress(_phoneNumber string) (common.Address, error) {
	return _Telco.Contract.GetFromPhoneNumberToAddress(&_Telco.CallOpts, _phoneNumber)
}

// GetFromPhoneNumberToAddress is a free data retrieval call binding the contract method 0xca62cba3.
//
// Solidity: function getFromPhoneNumberToAddress(string _phoneNumber) constant returns(address)
func (_Telco *TelcoCallerSession) GetFromPhoneNumberToAddress(_phoneNumber string) (common.Address, error) {
	return _Telco.Contract.GetFromPhoneNumberToAddress(&_Telco.CallOpts, _phoneNumber)
}

// GetUserBalance is a free data retrieval call binding the contract method 0x6805d6ad.
//
// Solidity: function getUserBalance(address _addr, address _tokenAddress) constant returns(uint256)
func (_Telco *TelcoCaller) GetUserBalance(opts *bind.CallOpts, _addr common.Address, _tokenAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Telco.contract.Call(opts, out, "getUserBalance", _addr, _tokenAddress)
	return *ret0, err
}

// GetUserBalance is a free data retrieval call binding the contract method 0x6805d6ad.
//
// Solidity: function getUserBalance(address _addr, address _tokenAddress) constant returns(uint256)
func (_Telco *TelcoSession) GetUserBalance(_addr common.Address, _tokenAddress common.Address) (*big.Int, error) {
	return _Telco.Contract.GetUserBalance(&_Telco.CallOpts, _addr, _tokenAddress)
}

// GetUserBalance is a free data retrieval call binding the contract method 0x6805d6ad.
//
// Solidity: function getUserBalance(address _addr, address _tokenAddress) constant returns(uint256)
func (_Telco *TelcoCallerSession) GetUserBalance(_addr common.Address, _tokenAddress common.Address) (*big.Int, error) {
	return _Telco.Contract.GetUserBalance(&_Telco.CallOpts, _addr, _tokenAddress)
}

// GetUserPhoneNumber is a free data retrieval call binding the contract method 0xcc566239.
//
// Solidity: function getUserPhoneNumber(address _addr) constant returns(string)
func (_Telco *TelcoCaller) GetUserPhoneNumber(opts *bind.CallOpts, _addr common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Telco.contract.Call(opts, out, "getUserPhoneNumber", _addr)
	return *ret0, err
}

// GetUserPhoneNumber is a free data retrieval call binding the contract method 0xcc566239.
//
// Solidity: function getUserPhoneNumber(address _addr) constant returns(string)
func (_Telco *TelcoSession) GetUserPhoneNumber(_addr common.Address) (string, error) {
	return _Telco.Contract.GetUserPhoneNumber(&_Telco.CallOpts, _addr)
}

// GetUserPhoneNumber is a free data retrieval call binding the contract method 0xcc566239.
//
// Solidity: function getUserPhoneNumber(address _addr) constant returns(string)
func (_Telco *TelcoCallerSession) GetUserPhoneNumber(_addr common.Address) (string, error) {
	return _Telco.Contract.GetUserPhoneNumber(&_Telco.CallOpts, _addr)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Telco *TelcoCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Telco.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Telco *TelcoSession) IsOwner() (bool, error) {
	return _Telco.Contract.IsOwner(&_Telco.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Telco *TelcoCallerSession) IsOwner() (bool, error) {
	return _Telco.Contract.IsOwner(&_Telco.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Telco *TelcoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Telco.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Telco *TelcoSession) Owner() (common.Address, error) {
	return _Telco.Contract.Owner(&_Telco.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Telco *TelcoCallerSession) Owner() (common.Address, error) {
	return _Telco.Contract.Owner(&_Telco.CallOpts)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) constant returns(string phoneNumber)
func (_Telco *TelcoCaller) Users(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Telco.contract.Call(opts, out, "users", arg0)
	return *ret0, err
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) constant returns(string phoneNumber)
func (_Telco *TelcoSession) Users(arg0 common.Address) (string, error) {
	return _Telco.Contract.Users(&_Telco.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) constant returns(string phoneNumber)
func (_Telco *TelcoCallerSession) Users(arg0 common.Address) (string, error) {
	return _Telco.Contract.Users(&_Telco.CallOpts, arg0)
}

// AuthTransfer is a paid mutator transaction binding the contract method 0xef91a4ef.
//
// Solidity: function authTransfer(address _tokenAddress, string _toPhoneNumber, uint256 _amount, string _phoneNumber, bytes _password) returns()
func (_Telco *TelcoTransactor) AuthTransfer(opts *bind.TransactOpts, _tokenAddress common.Address, _toPhoneNumber string, _amount *big.Int, _phoneNumber string, _password []byte) (*types.Transaction, error) {
	return _Telco.contract.Transact(opts, "authTransfer", _tokenAddress, _toPhoneNumber, _amount, _phoneNumber, _password)
}

// AuthTransfer is a paid mutator transaction binding the contract method 0xef91a4ef.
//
// Solidity: function authTransfer(address _tokenAddress, string _toPhoneNumber, uint256 _amount, string _phoneNumber, bytes _password) returns()
func (_Telco *TelcoSession) AuthTransfer(_tokenAddress common.Address, _toPhoneNumber string, _amount *big.Int, _phoneNumber string, _password []byte) (*types.Transaction, error) {
	return _Telco.Contract.AuthTransfer(&_Telco.TransactOpts, _tokenAddress, _toPhoneNumber, _amount, _phoneNumber, _password)
}

// AuthTransfer is a paid mutator transaction binding the contract method 0xef91a4ef.
//
// Solidity: function authTransfer(address _tokenAddress, string _toPhoneNumber, uint256 _amount, string _phoneNumber, bytes _password) returns()
func (_Telco *TelcoTransactorSession) AuthTransfer(_tokenAddress common.Address, _toPhoneNumber string, _amount *big.Int, _phoneNumber string, _password []byte) (*types.Transaction, error) {
	return _Telco.Contract.AuthTransfer(&_Telco.TransactOpts, _tokenAddress, _toPhoneNumber, _amount, _phoneNumber, _password)
}

// GetUserTokens is a paid mutator transaction binding the contract method 0xc8a89482.
//
// Solidity: function getUserTokens(address _tokenAddress, uint256 _amount, string _phoneNumber, bytes _password) returns()
func (_Telco *TelcoTransactor) GetUserTokens(opts *bind.TransactOpts, _tokenAddress common.Address, _amount *big.Int, _phoneNumber string, _password []byte) (*types.Transaction, error) {
	return _Telco.contract.Transact(opts, "getUserTokens", _tokenAddress, _amount, _phoneNumber, _password)
}

// GetUserTokens is a paid mutator transaction binding the contract method 0xc8a89482.
//
// Solidity: function getUserTokens(address _tokenAddress, uint256 _amount, string _phoneNumber, bytes _password) returns()
func (_Telco *TelcoSession) GetUserTokens(_tokenAddress common.Address, _amount *big.Int, _phoneNumber string, _password []byte) (*types.Transaction, error) {
	return _Telco.Contract.GetUserTokens(&_Telco.TransactOpts, _tokenAddress, _amount, _phoneNumber, _password)
}

// GetUserTokens is a paid mutator transaction binding the contract method 0xc8a89482.
//
// Solidity: function getUserTokens(address _tokenAddress, uint256 _amount, string _phoneNumber, bytes _password) returns()
func (_Telco *TelcoTransactorSession) GetUserTokens(_tokenAddress common.Address, _amount *big.Int, _phoneNumber string, _password []byte) (*types.Transaction, error) {
	return _Telco.Contract.GetUserTokens(&_Telco.TransactOpts, _tokenAddress, _amount, _phoneNumber, _password)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0x7479601b.
//
// Solidity: function registerNewUser(string _phoneNumber, address userAddress) returns(bool)
func (_Telco *TelcoTransactor) RegisterNewUser(opts *bind.TransactOpts, _phoneNumber string, userAddress common.Address) (*types.Transaction, error) {
	return _Telco.contract.Transact(opts, "registerNewUser", _phoneNumber, userAddress)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0x7479601b.
//
// Solidity: function registerNewUser(string _phoneNumber, address userAddress) returns(bool)
func (_Telco *TelcoSession) RegisterNewUser(_phoneNumber string, userAddress common.Address) (*types.Transaction, error) {
	return _Telco.Contract.RegisterNewUser(&_Telco.TransactOpts, _phoneNumber, userAddress)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0x7479601b.
//
// Solidity: function registerNewUser(string _phoneNumber, address userAddress) returns(bool)
func (_Telco *TelcoTransactorSession) RegisterNewUser(_phoneNumber string, userAddress common.Address) (*types.Transaction, error) {
	return _Telco.Contract.RegisterNewUser(&_Telco.TransactOpts, _phoneNumber, userAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Telco *TelcoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Telco.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Telco *TelcoSession) RenounceOwnership() (*types.Transaction, error) {
	return _Telco.Contract.RenounceOwnership(&_Telco.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Telco *TelcoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Telco.Contract.RenounceOwnership(&_Telco.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Telco *TelcoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Telco.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Telco *TelcoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Telco.Contract.TransferOwnership(&_Telco.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Telco *TelcoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Telco.Contract.TransferOwnership(&_Telco.TransactOpts, newOwner)
}

// UserDirectWithdraw is a paid mutator transaction binding the contract method 0x1872b42c.
//
// Solidity: function userDirectWithdraw(address _tokenAddress, uint256 _amount) returns()
func (_Telco *TelcoTransactor) UserDirectWithdraw(opts *bind.TransactOpts, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Telco.contract.Transact(opts, "userDirectWithdraw", _tokenAddress, _amount)
}

// UserDirectWithdraw is a paid mutator transaction binding the contract method 0x1872b42c.
//
// Solidity: function userDirectWithdraw(address _tokenAddress, uint256 _amount) returns()
func (_Telco *TelcoSession) UserDirectWithdraw(_tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Telco.Contract.UserDirectWithdraw(&_Telco.TransactOpts, _tokenAddress, _amount)
}

// UserDirectWithdraw is a paid mutator transaction binding the contract method 0x1872b42c.
//
// Solidity: function userDirectWithdraw(address _tokenAddress, uint256 _amount) returns()
func (_Telco *TelcoTransactorSession) UserDirectWithdraw(_tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Telco.Contract.UserDirectWithdraw(&_Telco.TransactOpts, _tokenAddress, _amount)
}

// UserInDirectWithdraw is a paid mutator transaction binding the contract method 0x7799c341.
//
// Solidity: function userInDirectWithdraw(address _tokenAddress, uint256 _amount, string _phoneNumber, bytes _password) returns()
func (_Telco *TelcoTransactor) UserInDirectWithdraw(opts *bind.TransactOpts, _tokenAddress common.Address, _amount *big.Int, _phoneNumber string, _password []byte) (*types.Transaction, error) {
	return _Telco.contract.Transact(opts, "userInDirectWithdraw", _tokenAddress, _amount, _phoneNumber, _password)
}

// UserInDirectWithdraw is a paid mutator transaction binding the contract method 0x7799c341.
//
// Solidity: function userInDirectWithdraw(address _tokenAddress, uint256 _amount, string _phoneNumber, bytes _password) returns()
func (_Telco *TelcoSession) UserInDirectWithdraw(_tokenAddress common.Address, _amount *big.Int, _phoneNumber string, _password []byte) (*types.Transaction, error) {
	return _Telco.Contract.UserInDirectWithdraw(&_Telco.TransactOpts, _tokenAddress, _amount, _phoneNumber, _password)
}

// UserInDirectWithdraw is a paid mutator transaction binding the contract method 0x7799c341.
//
// Solidity: function userInDirectWithdraw(address _tokenAddress, uint256 _amount, string _phoneNumber, bytes _password) returns()
func (_Telco *TelcoTransactorSession) UserInDirectWithdraw(_tokenAddress common.Address, _amount *big.Int, _phoneNumber string, _password []byte) (*types.Transaction, error) {
	return _Telco.Contract.UserInDirectWithdraw(&_Telco.TransactOpts, _tokenAddress, _amount, _phoneNumber, _password)
}

// TelcoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Telco contract.
type TelcoOwnershipTransferredIterator struct {
	Event *TelcoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TelcoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TelcoOwnershipTransferred)
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
		it.Event = new(TelcoOwnershipTransferred)
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
func (it *TelcoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TelcoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TelcoOwnershipTransferred represents a OwnershipTransferred event raised by the Telco contract.
type TelcoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Telco *TelcoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TelcoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Telco.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TelcoOwnershipTransferredIterator{contract: _Telco.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Telco *TelcoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TelcoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Telco.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TelcoOwnershipTransferred)
				if err := _Telco.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// UserABI is the input ABI used to generate the binding from.
const UserABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_erc20Address\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_password\",\"type\":\"bytes\"}],\"name\":\"trustedEntityTransferERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"passwords\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hashedPasswords\",\"type\":\"bytes32[]\"}],\"name\":\"addPasswords\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_password\",\"type\":\"bytes\"}],\"name\":\"trustedEntityTransferEth\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_trustedEntity\",\"type\":\"address\"}],\"name\":\"addTrustedEntity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"phoneNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_password\",\"type\":\"bytes\"}],\"name\":\"isPasswordValid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_password\",\"type\":\"bytes\"}],\"name\":\"retirePassword\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_erc20Address\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustedEntities\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_trustedEntity\",\"type\":\"address\"}],\"name\":\"removeTrustedEntity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferEth\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalPasswordsLeft\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_phoneNumber\",\"type\":\"string\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_trustedEntity\",\"type\":\"address\"},{\"name\":\"_hashedPasswords\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// UserBin is the compiled bytecode used for deploying new contracts.
const UserBin = `0x60806040523480156200001157600080fd5b506040516200122938038062001229833981018060405260808110156200003757600080fd5b8101908080516401000000008111156200005057600080fd5b820160208101848111156200006457600080fd5b81516401000000008111828201871017156200007f57600080fd5b505092919060200180516401000000008111156200009c57600080fd5b82016020810184811115620000b057600080fd5b8151640100000000811182820187101715620000cb57600080fd5b50506020820151604090920180519194929391640100000000811115620000f157600080fd5b820160208101848111156200010557600080fd5b81518560208202830111640100000000821117156200012357600080fd5b5050600080546001600160a01b03191633178082556040519295506001600160a01b0316935091507f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a383516200018590600190602087019062000291565b5082516200019b90600290602086019062000291565b50620001ad81620001c860201b60201c565b620001be826200024160201b60201c565b5050505062000333565b620001d86200027f60201b60201c565b620001e257600080fd5b60005b815181101562000234576001600460008484815181106200020257fe5b6020908102919091018101518252810191909152604001600020805460ff1916911515919091179055600101620001e5565b5051600380549091019055565b620002516200027f60201b60201c565b6200025b57600080fd5b6001600160a01b03166000908152600560205260409020805460ff19166001179055565b6000546001600160a01b031633145b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002d457805160ff191683800117855562000304565b8280016001018555821562000304579182015b8281111562000304578251825591602001919060010190620002e7565b506200031292915062000316565b5090565b6200028e91905b808211156200031257600081556001016200031d565b610ee680620003436000396000f3fe6080604052600436106101095760003560e01c80638f32d59b11610095578063db73012b11610064578063db73012b14610658578063dea612d41461068b578063e9bb84c2146106be578063ef9428d1146106ea578063f2fde38b1461071157610109565b80638f32d59b1461049e578063923905a7146104b35780639da68745146105645780639db5dbe41461061557610109565b806340076c3d116100dc57806340076c3d146103575780635168bf3e14610410578063715018a614610443578063747f3380146104585780638da5cb5b1461046d57610109565b806306fdde031461010e57806311173ecd1461019857806334c9cab81461026b5780633feb3c4f146102a9575b600080fd5b34801561011a57600080fd5b50610123610744565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561015d578181015183820152602001610145565b50505050905090810190601f16801561018a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101a457600080fd5b50610269600480360360808110156101bb57600080fd5b6001600160a01b03823581169260208101359091169160408201359190810190608081016060820135600160201b8111156101f557600080fd5b82018360208201111561020757600080fd5b803590602001918460018302840111600160201b8311171561022857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506107cf945050505050565b005b34801561027757600080fd5b506102956004803603602081101561028e57600080fd5b50356109a8565b604080519115158252519081900360200190f35b3480156102b557600080fd5b50610269600480360360208110156102cc57600080fd5b810190602081018135600160201b8111156102e657600080fd5b8201836020820111156102f857600080fd5b803590602001918460208302840111600160201b8311171561031957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506109bd945050505050565b6102696004803603606081101561036d57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561039c57600080fd5b8201836020820111156103ae57600080fd5b803590602001918460018302840111600160201b831117156103cf57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a2a945050505050565b34801561041c57600080fd5b506102696004803603602081101561043357600080fd5b50356001600160a01b0316610b5e565b34801561044f57600080fd5b50610269610b93565b34801561046457600080fd5b50610123610bee565b34801561047957600080fd5b50610482610c48565b604080516001600160a01b039092168252519081900360200190f35b3480156104aa57600080fd5b50610295610c57565b3480156104bf57600080fd5b50610295600480360360208110156104d657600080fd5b810190602081018135600160201b8111156104f057600080fd5b82018360208201111561050257600080fd5b803590602001918460018302840111600160201b8311171561052357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c68945050505050565b34801561057057600080fd5b506102696004803603602081101561058757600080fd5b810190602081018135600160201b8111156105a157600080fd5b8201836020820111156105b357600080fd5b803590602001918460018302840111600160201b831117156105d457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c85945050505050565b34801561062157600080fd5b506102696004803603606081101561063857600080fd5b506001600160a01b03813581169160208101359091169060400135610caf565b34801561066457600080fd5b506102956004803603602081101561067b57600080fd5b50356001600160a01b0316610d96565b34801561069757600080fd5b50610269600480360360208110156106ae57600080fd5b50356001600160a01b0316610dab565b610269600480360360408110156106d457600080fd5b506001600160a01b038135169060200135610ddd565b3480156106f657600080fd5b506106ff610e29565b60408051918252519081900360200190f35b34801561071d57600080fd5b506102696004803603602081101561073457600080fd5b50356001600160a01b0316610e2f565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156107c75780601f1061079c576101008083540402835291602001916107c7565b820191906000526020600020905b8154815290600101906020018083116107aa57829003601f168201915b505050505081565b3360009081526005602052604090205460ff1661082e5760408051600160e51b62461bcd0281526020600482015260126024820152600160701b716e6f74207472757374656420656e7469747902604482015290519081900360640190fd5b805160208083019190912060009081526004909152604090205460ff1661089f5760408051600160e51b62461bcd02815260206004820152601560248201527f70617373776f7264206973206e6f742076616c69640000000000000000000000604482015290519081900360640190fd5b8051602080830191909120600090815260048083526040808320805460ff19169055600380546000190190558051600160e01b63a9059cbb0281526001600160a01b038981169382019390935260248101879052905187949285169363a9059cbb93604480850194919392918390030190829087803b15801561092157600080fd5b505af1158015610935573d6000803e3d6000fd5b505050506040513d602081101561094b57600080fd5b50516109a15760408051600160e51b62461bcd02815260206004820152601560248201527f70726f626c656d2073656e64696e672045524332300000000000000000000000604482015290519081900360640190fd5b5050505050565b60046020526000908152604090205460ff1681565b6109c5610c57565b6109ce57600080fd5b60005b8151811015610a1d576001600460008484815181106109ec57fe5b6020908102919091018101518252810191909152604001600020805460ff19169115159190911790556001016109d1565b5051600380549091019055565b3360009081526005602052604090205460ff16610a895760408051600160e51b62461bcd0281526020600482015260126024820152600160701b716e6f74207472757374656420656e7469747902604482015290519081900360640190fd5b805160208083019190912060009081526004909152604090205460ff16610afa5760408051600160e51b62461bcd02815260206004820152601560248201527f70617373776f7264206973206e6f742076616c69640000000000000000000000604482015290519081900360640190fd5b8051602080830191909120600090815260049091526040808220805460ff1916905560038054600019019055516001600160a01b0385169184156108fc02918591818181858888f19350505050158015610b58573d6000803e3d6000fd5b50505050565b610b66610c57565b610b6f57600080fd5b6001600160a01b03166000908152600560205260409020805460ff19166001179055565b610b9b610c57565b610ba457600080fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156107c75780601f1061079c576101008083540402835291602001916107c7565b6000546001600160a01b031690565b6000546001600160a01b0316331490565b805160209182012060009081526004909152604090205460ff1690565b80516020918201206000908152600490915260409020805460ff1916905560038054600019019055565b610cb7610c57565b610cc057600080fd5b60408051600160e01b63a9059cbb0281526001600160a01b038581166004830152602482018490529151849283169163a9059cbb9160448083019260209291908290030181600087803b158015610d1657600080fd5b505af1158015610d2a573d6000803e3d6000fd5b505050506040513d6020811015610d4057600080fd5b5051610b585760408051600160e51b62461bcd02815260206004820152601560248201527f70726f626c656d2073656e64696e672045524332300000000000000000000000604482015290519081900360640190fd5b60056020526000908152604090205460ff1681565b610db3610c57565b610dbc57600080fd5b6001600160a01b03166000908152600560205260409020805460ff19169055565b610de5610c57565b610dee57600080fd5b6040516001600160a01b0383169082156108fc029083906000818181858888f19350505050158015610e24573d6000803e3d6000fd5b505050565b60035481565b610e37610c57565b610e4057600080fd5b610e4981610e4c565b50565b6001600160a01b038116610e5f57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea165627a7a723058202c6eb174dbcfaf01f016e5381d42f3dd1d0ac77c09d2d551af6ca6f1be00c8e30029`

// DeployUser deploys a new Ethereum contract, binding an instance of User to it.
func DeployUser(auth *bind.TransactOpts, backend bind.ContractBackend, _phoneNumber string, _name string, _trustedEntity common.Address, _hashedPasswords [][32]byte) (common.Address, *types.Transaction, *User, error) {
	parsed, err := abi.JSON(strings.NewReader(UserABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserBin), backend, _phoneNumber, _name, _trustedEntity, _hashedPasswords)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &User{UserCaller: UserCaller{contract: contract}, UserTransactor: UserTransactor{contract: contract}, UserFilterer: UserFilterer{contract: contract}}, nil
}

// User is an auto generated Go binding around an Ethereum contract.
type User struct {
	UserCaller     // Read-only binding to the contract
	UserTransactor // Write-only binding to the contract
	UserFilterer   // Log filterer for contract events
}

// UserCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserSession struct {
	Contract     *User             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserCallerSession struct {
	Contract *UserCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserTransactorSession struct {
	Contract     *UserTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserRaw struct {
	Contract *User // Generic contract binding to access the raw methods on
}

// UserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserCallerRaw struct {
	Contract *UserCaller // Generic read-only contract binding to access the raw methods on
}

// UserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserTransactorRaw struct {
	Contract *UserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUser creates a new instance of User, bound to a specific deployed contract.
func NewUser(address common.Address, backend bind.ContractBackend) (*User, error) {
	contract, err := bindUser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &User{UserCaller: UserCaller{contract: contract}, UserTransactor: UserTransactor{contract: contract}, UserFilterer: UserFilterer{contract: contract}}, nil
}

// NewUserCaller creates a new read-only instance of User, bound to a specific deployed contract.
func NewUserCaller(address common.Address, caller bind.ContractCaller) (*UserCaller, error) {
	contract, err := bindUser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserCaller{contract: contract}, nil
}

// NewUserTransactor creates a new write-only instance of User, bound to a specific deployed contract.
func NewUserTransactor(address common.Address, transactor bind.ContractTransactor) (*UserTransactor, error) {
	contract, err := bindUser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserTransactor{contract: contract}, nil
}

// NewUserFilterer creates a new log filterer instance of User, bound to a specific deployed contract.
func NewUserFilterer(address common.Address, filterer bind.ContractFilterer) (*UserFilterer, error) {
	contract, err := bindUser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserFilterer{contract: contract}, nil
}

// bindUser binds a generic wrapper to an already deployed contract.
func bindUser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _User.Contract.UserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _User.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_User *UserCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _User.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_User *UserSession) IsOwner() (bool, error) {
	return _User.Contract.IsOwner(&_User.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_User *UserCallerSession) IsOwner() (bool, error) {
	return _User.Contract.IsOwner(&_User.CallOpts)
}

// IsPasswordValid is a free data retrieval call binding the contract method 0x923905a7.
//
// Solidity: function isPasswordValid(bytes _password) constant returns(bool)
func (_User *UserCaller) IsPasswordValid(opts *bind.CallOpts, _password []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _User.contract.Call(opts, out, "isPasswordValid", _password)
	return *ret0, err
}

// IsPasswordValid is a free data retrieval call binding the contract method 0x923905a7.
//
// Solidity: function isPasswordValid(bytes _password) constant returns(bool)
func (_User *UserSession) IsPasswordValid(_password []byte) (bool, error) {
	return _User.Contract.IsPasswordValid(&_User.CallOpts, _password)
}

// IsPasswordValid is a free data retrieval call binding the contract method 0x923905a7.
//
// Solidity: function isPasswordValid(bytes _password) constant returns(bool)
func (_User *UserCallerSession) IsPasswordValid(_password []byte) (bool, error) {
	return _User.Contract.IsPasswordValid(&_User.CallOpts, _password)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_User *UserCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _User.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_User *UserSession) Name() (string, error) {
	return _User.Contract.Name(&_User.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_User *UserCallerSession) Name() (string, error) {
	return _User.Contract.Name(&_User.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_User *UserCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _User.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_User *UserSession) Owner() (common.Address, error) {
	return _User.Contract.Owner(&_User.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_User *UserCallerSession) Owner() (common.Address, error) {
	return _User.Contract.Owner(&_User.CallOpts)
}

// Passwords is a free data retrieval call binding the contract method 0x34c9cab8.
//
// Solidity: function passwords(bytes32 ) constant returns(bool)
func (_User *UserCaller) Passwords(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _User.contract.Call(opts, out, "passwords", arg0)
	return *ret0, err
}

// Passwords is a free data retrieval call binding the contract method 0x34c9cab8.
//
// Solidity: function passwords(bytes32 ) constant returns(bool)
func (_User *UserSession) Passwords(arg0 [32]byte) (bool, error) {
	return _User.Contract.Passwords(&_User.CallOpts, arg0)
}

// Passwords is a free data retrieval call binding the contract method 0x34c9cab8.
//
// Solidity: function passwords(bytes32 ) constant returns(bool)
func (_User *UserCallerSession) Passwords(arg0 [32]byte) (bool, error) {
	return _User.Contract.Passwords(&_User.CallOpts, arg0)
}

// PhoneNumber is a free data retrieval call binding the contract method 0x747f3380.
//
// Solidity: function phoneNumber() constant returns(string)
func (_User *UserCaller) PhoneNumber(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _User.contract.Call(opts, out, "phoneNumber")
	return *ret0, err
}

// PhoneNumber is a free data retrieval call binding the contract method 0x747f3380.
//
// Solidity: function phoneNumber() constant returns(string)
func (_User *UserSession) PhoneNumber() (string, error) {
	return _User.Contract.PhoneNumber(&_User.CallOpts)
}

// PhoneNumber is a free data retrieval call binding the contract method 0x747f3380.
//
// Solidity: function phoneNumber() constant returns(string)
func (_User *UserCallerSession) PhoneNumber() (string, error) {
	return _User.Contract.PhoneNumber(&_User.CallOpts)
}

// TotalPasswordsLeft is a free data retrieval call binding the contract method 0xef9428d1.
//
// Solidity: function totalPasswordsLeft() constant returns(uint256)
func (_User *UserCaller) TotalPasswordsLeft(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _User.contract.Call(opts, out, "totalPasswordsLeft")
	return *ret0, err
}

// TotalPasswordsLeft is a free data retrieval call binding the contract method 0xef9428d1.
//
// Solidity: function totalPasswordsLeft() constant returns(uint256)
func (_User *UserSession) TotalPasswordsLeft() (*big.Int, error) {
	return _User.Contract.TotalPasswordsLeft(&_User.CallOpts)
}

// TotalPasswordsLeft is a free data retrieval call binding the contract method 0xef9428d1.
//
// Solidity: function totalPasswordsLeft() constant returns(uint256)
func (_User *UserCallerSession) TotalPasswordsLeft() (*big.Int, error) {
	return _User.Contract.TotalPasswordsLeft(&_User.CallOpts)
}

// TrustedEntities is a free data retrieval call binding the contract method 0xdb73012b.
//
// Solidity: function trustedEntities(address ) constant returns(bool)
func (_User *UserCaller) TrustedEntities(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _User.contract.Call(opts, out, "trustedEntities", arg0)
	return *ret0, err
}

// TrustedEntities is a free data retrieval call binding the contract method 0xdb73012b.
//
// Solidity: function trustedEntities(address ) constant returns(bool)
func (_User *UserSession) TrustedEntities(arg0 common.Address) (bool, error) {
	return _User.Contract.TrustedEntities(&_User.CallOpts, arg0)
}

// TrustedEntities is a free data retrieval call binding the contract method 0xdb73012b.
//
// Solidity: function trustedEntities(address ) constant returns(bool)
func (_User *UserCallerSession) TrustedEntities(arg0 common.Address) (bool, error) {
	return _User.Contract.TrustedEntities(&_User.CallOpts, arg0)
}

// AddPasswords is a paid mutator transaction binding the contract method 0x3feb3c4f.
//
// Solidity: function addPasswords(bytes32[] _hashedPasswords) returns()
func (_User *UserTransactor) AddPasswords(opts *bind.TransactOpts, _hashedPasswords [][32]byte) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "addPasswords", _hashedPasswords)
}

// AddPasswords is a paid mutator transaction binding the contract method 0x3feb3c4f.
//
// Solidity: function addPasswords(bytes32[] _hashedPasswords) returns()
func (_User *UserSession) AddPasswords(_hashedPasswords [][32]byte) (*types.Transaction, error) {
	return _User.Contract.AddPasswords(&_User.TransactOpts, _hashedPasswords)
}

// AddPasswords is a paid mutator transaction binding the contract method 0x3feb3c4f.
//
// Solidity: function addPasswords(bytes32[] _hashedPasswords) returns()
func (_User *UserTransactorSession) AddPasswords(_hashedPasswords [][32]byte) (*types.Transaction, error) {
	return _User.Contract.AddPasswords(&_User.TransactOpts, _hashedPasswords)
}

// AddTrustedEntity is a paid mutator transaction binding the contract method 0x5168bf3e.
//
// Solidity: function addTrustedEntity(address _trustedEntity) returns()
func (_User *UserTransactor) AddTrustedEntity(opts *bind.TransactOpts, _trustedEntity common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "addTrustedEntity", _trustedEntity)
}

// AddTrustedEntity is a paid mutator transaction binding the contract method 0x5168bf3e.
//
// Solidity: function addTrustedEntity(address _trustedEntity) returns()
func (_User *UserSession) AddTrustedEntity(_trustedEntity common.Address) (*types.Transaction, error) {
	return _User.Contract.AddTrustedEntity(&_User.TransactOpts, _trustedEntity)
}

// AddTrustedEntity is a paid mutator transaction binding the contract method 0x5168bf3e.
//
// Solidity: function addTrustedEntity(address _trustedEntity) returns()
func (_User *UserTransactorSession) AddTrustedEntity(_trustedEntity common.Address) (*types.Transaction, error) {
	return _User.Contract.AddTrustedEntity(&_User.TransactOpts, _trustedEntity)
}

// RemoveTrustedEntity is a paid mutator transaction binding the contract method 0xdea612d4.
//
// Solidity: function removeTrustedEntity(address _trustedEntity) returns()
func (_User *UserTransactor) RemoveTrustedEntity(opts *bind.TransactOpts, _trustedEntity common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "removeTrustedEntity", _trustedEntity)
}

// RemoveTrustedEntity is a paid mutator transaction binding the contract method 0xdea612d4.
//
// Solidity: function removeTrustedEntity(address _trustedEntity) returns()
func (_User *UserSession) RemoveTrustedEntity(_trustedEntity common.Address) (*types.Transaction, error) {
	return _User.Contract.RemoveTrustedEntity(&_User.TransactOpts, _trustedEntity)
}

// RemoveTrustedEntity is a paid mutator transaction binding the contract method 0xdea612d4.
//
// Solidity: function removeTrustedEntity(address _trustedEntity) returns()
func (_User *UserTransactorSession) RemoveTrustedEntity(_trustedEntity common.Address) (*types.Transaction, error) {
	return _User.Contract.RemoveTrustedEntity(&_User.TransactOpts, _trustedEntity)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_User *UserTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_User *UserSession) RenounceOwnership() (*types.Transaction, error) {
	return _User.Contract.RenounceOwnership(&_User.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_User *UserTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _User.Contract.RenounceOwnership(&_User.TransactOpts)
}

// RetirePassword is a paid mutator transaction binding the contract method 0x9da68745.
//
// Solidity: function retirePassword(bytes _password) returns()
func (_User *UserTransactor) RetirePassword(opts *bind.TransactOpts, _password []byte) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "retirePassword", _password)
}

// RetirePassword is a paid mutator transaction binding the contract method 0x9da68745.
//
// Solidity: function retirePassword(bytes _password) returns()
func (_User *UserSession) RetirePassword(_password []byte) (*types.Transaction, error) {
	return _User.Contract.RetirePassword(&_User.TransactOpts, _password)
}

// RetirePassword is a paid mutator transaction binding the contract method 0x9da68745.
//
// Solidity: function retirePassword(bytes _password) returns()
func (_User *UserTransactorSession) RetirePassword(_password []byte) (*types.Transaction, error) {
	return _User.Contract.RetirePassword(&_User.TransactOpts, _password)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address _to, address _erc20Address, uint256 _amount) returns()
func (_User *UserTransactor) TransferERC20(opts *bind.TransactOpts, _to common.Address, _erc20Address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "transferERC20", _to, _erc20Address, _amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address _to, address _erc20Address, uint256 _amount) returns()
func (_User *UserSession) TransferERC20(_to common.Address, _erc20Address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _User.Contract.TransferERC20(&_User.TransactOpts, _to, _erc20Address, _amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address _to, address _erc20Address, uint256 _amount) returns()
func (_User *UserTransactorSession) TransferERC20(_to common.Address, _erc20Address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _User.Contract.TransferERC20(&_User.TransactOpts, _to, _erc20Address, _amount)
}

// TransferEth is a paid mutator transaction binding the contract method 0xe9bb84c2.
//
// Solidity: function transferEth(address _to, uint256 _amount) returns()
func (_User *UserTransactor) TransferEth(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "transferEth", _to, _amount)
}

// TransferEth is a paid mutator transaction binding the contract method 0xe9bb84c2.
//
// Solidity: function transferEth(address _to, uint256 _amount) returns()
func (_User *UserSession) TransferEth(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _User.Contract.TransferEth(&_User.TransactOpts, _to, _amount)
}

// TransferEth is a paid mutator transaction binding the contract method 0xe9bb84c2.
//
// Solidity: function transferEth(address _to, uint256 _amount) returns()
func (_User *UserTransactorSession) TransferEth(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _User.Contract.TransferEth(&_User.TransactOpts, _to, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_User *UserTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_User *UserSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _User.Contract.TransferOwnership(&_User.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_User *UserTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _User.Contract.TransferOwnership(&_User.TransactOpts, newOwner)
}

// TrustedEntityTransferERC20 is a paid mutator transaction binding the contract method 0x11173ecd.
//
// Solidity: function trustedEntityTransferERC20(address _to, address _erc20Address, uint256 _amount, bytes _password) returns()
func (_User *UserTransactor) TrustedEntityTransferERC20(opts *bind.TransactOpts, _to common.Address, _erc20Address common.Address, _amount *big.Int, _password []byte) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "trustedEntityTransferERC20", _to, _erc20Address, _amount, _password)
}

// TrustedEntityTransferERC20 is a paid mutator transaction binding the contract method 0x11173ecd.
//
// Solidity: function trustedEntityTransferERC20(address _to, address _erc20Address, uint256 _amount, bytes _password) returns()
func (_User *UserSession) TrustedEntityTransferERC20(_to common.Address, _erc20Address common.Address, _amount *big.Int, _password []byte) (*types.Transaction, error) {
	return _User.Contract.TrustedEntityTransferERC20(&_User.TransactOpts, _to, _erc20Address, _amount, _password)
}

// TrustedEntityTransferERC20 is a paid mutator transaction binding the contract method 0x11173ecd.
//
// Solidity: function trustedEntityTransferERC20(address _to, address _erc20Address, uint256 _amount, bytes _password) returns()
func (_User *UserTransactorSession) TrustedEntityTransferERC20(_to common.Address, _erc20Address common.Address, _amount *big.Int, _password []byte) (*types.Transaction, error) {
	return _User.Contract.TrustedEntityTransferERC20(&_User.TransactOpts, _to, _erc20Address, _amount, _password)
}

// TrustedEntityTransferEth is a paid mutator transaction binding the contract method 0x40076c3d.
//
// Solidity: function trustedEntityTransferEth(address _to, uint256 _amount, bytes _password) returns()
func (_User *UserTransactor) TrustedEntityTransferEth(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _password []byte) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "trustedEntityTransferEth", _to, _amount, _password)
}

// TrustedEntityTransferEth is a paid mutator transaction binding the contract method 0x40076c3d.
//
// Solidity: function trustedEntityTransferEth(address _to, uint256 _amount, bytes _password) returns()
func (_User *UserSession) TrustedEntityTransferEth(_to common.Address, _amount *big.Int, _password []byte) (*types.Transaction, error) {
	return _User.Contract.TrustedEntityTransferEth(&_User.TransactOpts, _to, _amount, _password)
}

// TrustedEntityTransferEth is a paid mutator transaction binding the contract method 0x40076c3d.
//
// Solidity: function trustedEntityTransferEth(address _to, uint256 _amount, bytes _password) returns()
func (_User *UserTransactorSession) TrustedEntityTransferEth(_to common.Address, _amount *big.Int, _password []byte) (*types.Transaction, error) {
	return _User.Contract.TrustedEntityTransferEth(&_User.TransactOpts, _to, _amount, _password)
}

// UserOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the User contract.
type UserOwnershipTransferredIterator struct {
	Event *UserOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UserOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserOwnershipTransferred)
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
		it.Event = new(UserOwnershipTransferred)
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
func (it *UserOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserOwnershipTransferred represents a OwnershipTransferred event raised by the User contract.
type UserOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_User *UserFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UserOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _User.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UserOwnershipTransferredIterator{contract: _User.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_User *UserFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UserOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _User.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserOwnershipTransferred)
				if err := _User.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

