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
const TelcoABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"userDirectWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_phoneNumber\",\"type\":\"string\"},{\"name\":\"_password\",\"type\":\"bytes\"}],\"name\":\"userInDirectWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_phoneNumber\",\"type\":\"string\"},{\"name\":\"_password\",\"type\":\"bytes\"}],\"name\":\"checkPasswordValid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_phoneNumber\",\"type\":\"string\"},{\"name\":\"_password\",\"type\":\"bytes\"}],\"name\":\"getUserTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_toPhoneNumber\",\"type\":\"string\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_phoneNumber\",\"type\":\"string\"},{\"name\":\"_password\",\"type\":\"bytes\"}],\"name\":\"authTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_phoneNumber\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// TelcoBin is the compiled bytecode used for deploying new contracts.
const TelcoBin = `0x608060405234801561001057600080fd5b50600080546001600160a01b03191633178082556040516001600160a01b039190911691907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a36110b6806100696000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063b523ba9f11610066578063b523ba9f146101ec578063c8a8948214610315578063ef91a4ef14610453578063f2c298be14610575578063f2fde38b146105e35761009e565b80631872b42c146100a3578063715018a6146100d15780637799c341146100d95780638da5cb5b146101ac5780638f32d59b146101d0575b600080fd5b6100cf600480360360408110156100b957600080fd5b506001600160a01b038135169060200135610609565b005b6100cf61070f565b6100cf600480360360808110156100ef57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561011e57600080fd5b82018360208201111561013057600080fd5b803590602001918460018302840111600160201b8311171561015157600080fd5b919390929091602081019035600160201b81111561016e57600080fd5b82018360208201111561018057600080fd5b803590602001918460018302840111600160201b831117156101a157600080fd5b50909250905061076a565b6101b46108de565b604080516001600160a01b039092168252519081900360200190f35b6101d86108ee565b604080519115158252519081900360200190f35b6101d86004803603604081101561020257600080fd5b810190602081018135600160201b81111561021c57600080fd5b82018360208201111561022e57600080fd5b803590602001918460018302840111600160201b8311171561024f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156102a157600080fd5b8201836020820111156102b357600080fd5b803590602001918460018302840111600160201b831117156102d457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108ff945050505050565b6100cf6004803603608081101561032b57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561035a57600080fd5b82018360208201111561036c57600080fd5b803590602001918460018302840111600160201b8311171561038d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156103df57600080fd5b8201836020820111156103f157600080fd5b803590602001918460018302840111600160201b8311171561041257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610aa6945050505050565b6100cf600480360360a081101561046957600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561049357600080fd5b8201836020820111156104a557600080fd5b803590602001918460018302840111600160201b831117156104c657600080fd5b91939092823592604081019060200135600160201b8111156104e757600080fd5b8201836020820111156104f957600080fd5b803590602001918460018302840111600160201b8311171561051a57600080fd5b919390929091602081019035600160201b81111561053757600080fd5b82018360208201111561054957600080fd5b803590602001918460018302840111600160201b8311171561056a57600080fd5b509092509050610c90565b6101d86004803603602081101561058b57600080fd5b810190602081018135600160201b8111156105a557600080fd5b8201836020820111156105b757600080fd5b803590602001918460018302840111600160201b831117156105d857600080fd5b509092509050610eb1565b6100cf600480360360208110156105f957600080fd5b50356001600160a01b0316610f67565b3360009081526001602081815260408084206001600160a01b0387168552928301909152909120548211156106885760408051600160e51b62461bcd02815260206004820152601c60248201527f6e6f7420656e6f7567682066756e647320746f20776974686472617700000000604482015290519081900360640190fd5b60408051600160e01b63a9059cbb02815233600482015260248101849052905184916001600160a01b0383169163a9059cbb916044808201926020929091908290030181600087803b1580156106dd57600080fd5b505af11580156106f1573d6000803e3d6000fd5b505050506040513d602081101561070757600080fd5b505050505050565b6107176108ee565b61072057600080fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6107dd84848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f880181900481028201810190925286815292508691508590819084018382808284376000920191909152506108ff92505050565b6108295760408051600160e51b62461bcd0281526020600482015260126024820152600160721b711c185cdcdddbdc99081b9bdd081d985b1a5902604482015290519081900360640190fd5b6000600285856040518083838082843791909101948552505060408051938490036020908101852054600160e01b63a9059cbb0286526001600160a01b0390811660048701819052602487018d905292519296508c95908616945063a9059cbb9350604480820193918290030181600087803b1580156108a857600080fd5b505af11580156108bc573d6000803e3d6000fd5b505050506040513d60208110156108d257600080fd5b50505050505050505050565b6000546001600160a01b03165b90565b6000546001600160a01b0316331490565b6000806002846040518082805190602001908083835b602083106109345780518252601f199092019160209182019101610915565b51815160001960209485036101000a019081169019919091161790529201948552506040519384900301909220546001600160a01b031692505050806109c45760408051600160e51b62461bcd02815260206004820152601360248201527f7573657220646f6573206e6f7420657869737400000000000000000000000000604482015290519081900360640190fd5b604051600160e01b63923905a702815260206004820181815285516024840152855184936001600160a01b0385169363923905a7938993909283926044909101919085019080838360005b83811015610a27578181015183820152602001610a0f565b50505050905090810190601f168015610a545780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b158015610a7157600080fd5b505afa158015610a85573d6000803e3d6000fd5b505050506040513d6020811015610a9b57600080fd5b505195945050505050565b610aae6108ee565b610ab757600080fd5b610ac182826108ff565b610b0d5760408051600160e51b62461bcd0281526020600482015260126024820152600160721b711c185cdcdddbdc99081b9bdd081d985b1a5902604482015290519081900360640190fd5b60006002836040518082805190602001908083835b60208310610b415780518252601f199092019160209182019101610b22565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382018520546001600160a01b03908116600081815260018552928320600160e01b6311173ecd0288523060048901818152938e1660248a0152604489018d9052608060648a019081528b5160848b01528b51939a5091988a98508897506311173ecd9691958f958f958e959194919360a40192918601918190849084905b83811015610bff578181015183820152602001610be7565b50505050905090810190601f168015610c2c5780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b158015610c4e57600080fd5b505af1158015610c62573d6000803e3d6000fd5b5050506001600160a01b03909716600090815260019092016020525060409020805490940190935550505050565b610d0384848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f880181900481028201810190925286815292508691508590819084018382808284376000920191909152506108ff92505050565b610d4f5760408051600160e51b62461bcd0281526020600482015260126024820152600160721b711c185cdcdddbdc99081b9bdd081d985b1a5902604482015290519081900360640190fd5b600060028585604051808383808284378083019250505092505050908152602001604051809103902060009054906101000a90046001600160a01b03169050600060016000836001600160a01b03166001600160a01b031681526020019081526020016000209050868160010160008c6001600160a01b03166001600160a01b03168152602001908152602001600020541015610deb57600080fd5b6001600160a01b038a16600090815260018201602052604080822080548a90039055516002908b908b90808383808284378083019250505092505050908152602001604051809103902060009054906101000a90046001600160a01b031690508760016000836001600160a01b03166001600160a01b0316815260200190815260200160002060010160008d6001600160a01b03166001600160a01b03168152602001908152602001600020600082825401925050819055505050505050505050505050565b604080516020601f8401819004810282018301835281018381526000928291908690869081908501838280828437600092018290525093909452505033815260016020908152604090912083518051919350610f11928492910190610ff2565b509050503360028484604051808383808284379190910194855250506040519283900360200190922080546001600160a01b03949094166001600160a01b03199094169390931790925550600191505092915050565b610f6f6108ee565b610f7857600080fd5b610f8181610f84565b50565b6001600160a01b038116610f9757600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061103357805160ff1916838001178555611060565b82800160010185558215611060579182015b82811115611060578251825591602001919060010190611045565b5061106c929150611070565b5090565b6108eb91905b8082111561106c576000815560010161107656fea165627a7a7230582084732f0780ffd42182c046fcb557ed81d531f58baf462be4ae6fd699b484ca340029`

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

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string _phoneNumber) returns(bool)
func (_Telco *TelcoTransactor) Register(opts *bind.TransactOpts, _phoneNumber string) (*types.Transaction, error) {
	return _Telco.contract.Transact(opts, "register", _phoneNumber)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string _phoneNumber) returns(bool)
func (_Telco *TelcoSession) Register(_phoneNumber string) (*types.Transaction, error) {
	return _Telco.Contract.Register(&_Telco.TransactOpts, _phoneNumber)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string _phoneNumber) returns(bool)
func (_Telco *TelcoTransactorSession) Register(_phoneNumber string) (*types.Transaction, error) {
	return _Telco.Contract.Register(&_Telco.TransactOpts, _phoneNumber)
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
const UserABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_erc20Address\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_password\",\"type\":\"bytes\"}],\"name\":\"trustedEntityTransferERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hashedPasswords\",\"type\":\"bytes32[]\"}],\"name\":\"addPasswords\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_password\",\"type\":\"bytes\"}],\"name\":\"trustedEntityTransferEth\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_trustedEntity\",\"type\":\"address\"}],\"name\":\"addTrustedEntity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_password\",\"type\":\"bytes\"}],\"name\":\"isPasswordValid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_erc20Address\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_trustedEntity\",\"type\":\"address\"}],\"name\":\"removeTrustedEntity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferEth\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_phoneNumber\",\"type\":\"string\"},{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// UserBin is the compiled bytecode used for deploying new contracts.
const UserBin = `0x608060405234801561001057600080fd5b5060405162000d0238038062000d028339810180604052604081101561003557600080fd5b81019080805164010000000081111561004d57600080fd5b8201602081018481111561006057600080fd5b815164010000000081118282018710171561007a57600080fd5b5050929190602001805164010000000081111561009657600080fd5b820160208101848111156100a957600080fd5b81516401000000008111828201871017156100c357600080fd5b5050600080546001600160a01b03191633178082556040519295506001600160a01b0316935091507f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3815161012390600190602085019061013f565b50805161013790600290602084019061013f565b5050506101da565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061018057805160ff19168380011785556101ad565b828001600101855582156101ad579182015b828111156101ad578251825591602001919060010190610192565b506101b99291506101bd565b5090565b6101d791905b808211156101b957600081556001016101c3565b90565b610b1880620001ea6000396000f3fe6080604052600436106100a75760003560e01c80638f32d59b116100645780638f32d59b1461032c578063923905a7146103555780639db5dbe414610406578063dea612d414610449578063e9bb84c21461047c578063f2fde38b146104a8576100a7565b806311173ecd146100ac5780633feb3c4f1461017f57806340076c3d146101fa5780635168bf3e146102b3578063715018a6146102e65780638da5cb5b146102fb575b600080fd5b3480156100b857600080fd5b5061017d600480360360808110156100cf57600080fd5b6001600160a01b03823581169260208101359091169160408201359190810190608081016060820135600160201b81111561010957600080fd5b82018360208201111561011b57600080fd5b803590602001918460018302840111600160201b8311171561013c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104db945050505050565b005b34801561018b57600080fd5b5061017d600480360360208110156101a257600080fd5b810190602081018135600160201b8111156101bc57600080fd5b8201836020820111156101ce57600080fd5b803590602001918460208302840111600160201b831117156101ef57600080fd5b5090925090506106a7565b61017d6004803603606081101561021057600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561023f57600080fd5b82018360208201111561025157600080fd5b803590602001918460018302840111600160201b8311171561027257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061070a945050505050565b3480156102bf57600080fd5b5061017d600480360360208110156102d657600080fd5b50356001600160a01b0316610834565b3480156102f257600080fd5b5061017d610869565b34801561030757600080fd5b506103106108c4565b604080516001600160a01b039092168252519081900360200190f35b34801561033857600080fd5b506103416108d3565b604080519115158252519081900360200190f35b34801561036157600080fd5b506103416004803603602081101561037857600080fd5b810190602081018135600160201b81111561039257600080fd5b8201836020820111156103a457600080fd5b803590602001918460018302840111600160201b831117156103c557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108e4945050505050565b34801561041257600080fd5b5061017d6004803603606081101561042957600080fd5b506001600160a01b03813581169160208101359091169060400135610901565b34801561045557600080fd5b5061017d6004803603602081101561046c57600080fd5b50356001600160a01b03166109e8565b61017d6004803603604081101561049257600080fd5b506001600160a01b038135169060200135610a1a565b3480156104b457600080fd5b5061017d600480360360208110156104cb57600080fd5b50356001600160a01b0316610a61565b3360009081526005602052604090205460ff1661053a5760408051600160e51b62461bcd0281526020600482015260126024820152600160701b716e6f74207472757374656420656e7469747902604482015290519081900360640190fd5b805160208083019190912060009081526003909152604090205460ff166105ab5760408051600160e51b62461bcd02815260206004820152601560248201527f70617373776f7264206973206e6f742076616c69640000000000000000000000604482015290519081900360640190fd5b80516020808301919091206000908152600382526040808220805460ff191690558051600160e01b63a9059cbb0281526001600160a01b03888116600483015260248201879052915187949285169363a9059cbb93604480850194919392918390030190829087803b15801561062057600080fd5b505af1158015610634573d6000803e3d6000fd5b505050506040513d602081101561064a57600080fd5b50516106a05760408051600160e51b62461bcd02815260206004820152601560248201527f70726f626c656d2073656e64696e672045524332300000000000000000000000604482015290519081900360640190fd5b5050505050565b6106af6108d3565b6106b857600080fd5b60005b81811015610705576001600360008585858181106106d557fe5b60209081029290920135835250810191909152604001600020805460ff19169115159190911790556001016106bb565b505050565b3360009081526005602052604090205460ff166107695760408051600160e51b62461bcd0281526020600482015260126024820152600160701b716e6f74207472757374656420656e7469747902604482015290519081900360640190fd5b805160208083019190912060009081526003909152604090205460ff166107da5760408051600160e51b62461bcd02815260206004820152601560248201527f70617373776f7264206973206e6f742076616c69640000000000000000000000604482015290519081900360640190fd5b8051602080830191909120600090815260039091526040808220805460ff19169055516001600160a01b0385169184156108fc02918591818181858888f1935050505015801561082e573d6000803e3d6000fd5b50505050565b61083c6108d3565b61084557600080fd5b6001600160a01b03166000908152600560205260409020805460ff19166001179055565b6108716108d3565b61087a57600080fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b6000546001600160a01b0316331490565b805160209182012060009081526003909152604090205460ff1690565b6109096108d3565b61091257600080fd5b60408051600160e01b63a9059cbb0281526001600160a01b038581166004830152602482018490529151849283169163a9059cbb9160448083019260209291908290030181600087803b15801561096857600080fd5b505af115801561097c573d6000803e3d6000fd5b505050506040513d602081101561099257600080fd5b505161082e5760408051600160e51b62461bcd02815260206004820152601560248201527f70726f626c656d2073656e64696e672045524332300000000000000000000000604482015290519081900360640190fd5b6109f06108d3565b6109f957600080fd5b6001600160a01b03166000908152600560205260409020805460ff19169055565b610a226108d3565b610a2b57600080fd5b6040516001600160a01b0383169082156108fc029083906000818181858888f19350505050158015610705573d6000803e3d6000fd5b610a696108d3565b610a7257600080fd5b610a7b81610a7e565b50565b6001600160a01b038116610a9157600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea165627a7a723058208e4de78f2e8eb5bf681622c54eace9a5eb741b332ec514a4f676d885eef6eaed0029`

// DeployUser deploys a new Ethereum contract, binding an instance of User to it.
func DeployUser(auth *bind.TransactOpts, backend bind.ContractBackend, _phoneNumber string, _name string) (common.Address, *types.Transaction, *User, error) {
	parsed, err := abi.JSON(strings.NewReader(UserABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserBin), backend, _phoneNumber, _name)
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

