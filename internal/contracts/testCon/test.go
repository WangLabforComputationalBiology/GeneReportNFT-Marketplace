// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testCon

import (
	_ "fmt"
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/v3/abi"
	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
)

// TestABI is the input ABI used to generate the binding from.
const TestABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"DataUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Test is an auto generated Go binding around a Solidity contract.
type Test struct {
	TestCaller     // Read-only binding to the contract
	TestTransactor // Write-only binding to the contract
	TestFilterer   // Log filterer for contract events
}

// TestCaller is an auto generated read-only Go binding around a Solidity contract.
type TestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTransactor is an auto generated write-only Go binding around a Solidity contract.
type TestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TestSession struct {
	Contract     *Test             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TestCallerSession struct {
	Contract *TestCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TestTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TestTransactorSession struct {
	Contract     *TestTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestRaw is an auto generated low-level Go binding around a Solidity contract.
type TestRaw struct {
	Contract *Test // Generic contract binding to access the raw methods on
}

// TestCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TestCallerRaw struct {
	Contract *TestCaller // Generic read-only contract binding to access the raw methods on
}

// TestTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TestTransactorRaw struct {
	Contract *TestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest creates a new instance of Test, bound to a specific deployed contract.
func NewTest(address common.Address, backend bind.ContractBackend) (*Test, error) {
	contract, err := bindTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// NewTestCaller creates a new read-only instance of Test, bound to a specific deployed contract.
func NewTestCaller(address common.Address, caller bind.ContractCaller) (*TestCaller, error) {
	contract, err := bindTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestCaller{contract: contract}, nil
}

// NewTestTransactor creates a new write-only instance of Test, bound to a specific deployed contract.
func NewTestTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTransactor, error) {
	contract, err := bindTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTransactor{contract: contract}, nil
}

// NewTestFilterer creates a new log filterer instance of Test, bound to a specific deployed contract.
func NewTestFilterer(address common.Address, filterer bind.ContractFilterer) (*TestFilterer, error) {
	contract, err := bindTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestFilterer{contract: contract}, nil
}

// bindTest binds a generic wrapper to an already deployed contract.
func bindTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Test.Contract.TestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Test.Contract.TestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Test.Contract.TestTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Test.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Test.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Test.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 key) constant returns(uint256)
func (_Test *TestCaller) GetData(opts *bind.CallOpts, key [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Test.contract.Call(opts, out, "getData", key)
	return *ret0, err
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 key) constant returns(uint256)
func (_Test *TestSession) GetData(key [32]byte) (*big.Int, error) {
	return _Test.Contract.GetData(&_Test.CallOpts, key)
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 key) constant returns(uint256)
func (_Test *TestCallerSession) GetData(key [32]byte) (*big.Int, error) {
	return _Test.Contract.GetData(&_Test.CallOpts, key)
}

// StoreData is a paid mutator transaction binding the contract method 0x379d7010.
//
// Solidity: function storeData(bytes32 key, uint256 value) returns()
func (_Test *TestTransactor) StoreData(opts *bind.TransactOpts, key [32]byte, value *big.Int) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Test.contract.TransactWithResult(opts, out, "storeData", key, value)
	return transaction, receipt, err
}

func (_Test *TestTransactor) AsyncStoreData(handler func(*types.Receipt, error), opts *bind.TransactOpts, key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Test.contract.AsyncTransact(opts, handler, "storeData", key, value)
}

// StoreData is a paid mutator transaction binding the contract method 0x379d7010.
//
// Solidity: function storeData(bytes32 key, uint256 value) returns()
func (_Test *TestSession) StoreData(key [32]byte, value *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Test.Contract.StoreData(&_Test.TransactOpts, key, value)
}

func (_Test *TestSession) AsyncStoreData(handler func(*types.Receipt, error), key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Test.Contract.AsyncStoreData(handler, &_Test.TransactOpts, key, value)
}

// StoreData is a paid mutator transaction binding the contract method 0x379d7010.
//
// Solidity: function storeData(bytes32 key, uint256 value) returns()
func (_Test *TestTransactorSession) StoreData(key [32]byte, value *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Test.Contract.StoreData(&_Test.TransactOpts, key, value)
}

func (_Test *TestTransactorSession) AsyncStoreData(handler func(*types.Receipt, error), key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Test.Contract.AsyncStoreData(handler, &_Test.TransactOpts, key, value)
}

// TestDataUpdated represents a DataUpdated event raised by the Test contract.
type TestDataUpdated struct {
	Key   [32]byte
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// WatchDataUpdated is a free log subscription operation binding the contract event 0x15d5385033e48d1ad0053bcf550121622ed5fb6c81e02f351891b13f826b62bd.
//
// Solidity: event DataUpdated(bytes32 indexed key, uint256 value)
func (_Test *TestFilterer) WatchDataUpdated(fromBlock *int64, handler func(int, []types.Log), key [32]byte) (string, error) {
	return _Test.contract.WatchLogs(fromBlock, handler, "DataUpdated", key)
}

func (_Test *TestFilterer) WatchAllDataUpdated(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Test.contract.WatchLogs(fromBlock, handler, "DataUpdated")
}

// ParseDataUpdated is a log parse operation binding the contract event 0x15d5385033e48d1ad0053bcf550121622ed5fb6c81e02f351891b13f826b62bd.
//
// Solidity: event DataUpdated(bytes32 indexed key, uint256 value)
func (_Test *TestFilterer) ParseDataUpdated(log types.Log) (*TestDataUpdated, error) {
	event := new(TestDataUpdated)
	if err := _Test.contract.UnpackLog(event, "DataUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchDataUpdated is a free log subscription operation binding the contract event 0x15d5385033e48d1ad0053bcf550121622ed5fb6c81e02f351891b13f826b62bd.
//
// Solidity: event DataUpdated(bytes32 indexed key, uint256 value)
func (_Test *TestSession) WatchDataUpdated(fromBlock *int64, handler func(int, []types.Log), key [32]byte) (string, error) {
	return _Test.Contract.WatchDataUpdated(fromBlock, handler, key)
}

func (_Test *TestSession) WatchAllDataUpdated(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _Test.Contract.WatchAllDataUpdated(fromBlock, handler)
}

// ParseDataUpdated is a log parse operation binding the contract event 0x15d5385033e48d1ad0053bcf550121622ed5fb6c81e02f351891b13f826b62bd.
//
// Solidity: event DataUpdated(bytes32 indexed key, uint256 value)
func (_Test *TestSession) ParseDataUpdated(log types.Log) (*TestDataUpdated, error) {
	return _Test.Contract.ParseDataUpdated(log)
}
