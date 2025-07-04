// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sharingPlatformContract

import (
	"fmt"
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

// SharingPlatformContractABI is the input ABI used to generate the binding from.
const SharingPlatformContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"}]"

// SharingPlatformContract is an auto generated Go binding around a Solidity contract.
type SharingPlatformContract struct {
	SharingPlatformContractCaller     // Read-only binding to the contract
	SharingPlatformContractTransactor // Write-only binding to the contract
	SharingPlatformContractFilterer   // Log filterer for contract events
}

// SharingPlatformContractCaller is an auto generated read-only Go binding around a Solidity contract.
type SharingPlatformContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SharingPlatformContractTransactor is an auto generated write-only Go binding around a Solidity contract.
type SharingPlatformContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SharingPlatformContractFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type SharingPlatformContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SharingPlatformContractSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type SharingPlatformContractSession struct {
	Contract     *SharingPlatformContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SharingPlatformContractCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type SharingPlatformContractCallerSession struct {
	Contract *SharingPlatformContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// SharingPlatformContractTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type SharingPlatformContractTransactorSession struct {
	Contract     *SharingPlatformContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// SharingPlatformContractRaw is an auto generated low-level Go binding around a Solidity contract.
type SharingPlatformContractRaw struct {
	Contract *SharingPlatformContract // Generic contract binding to access the raw methods on
}

// SharingPlatformContractCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type SharingPlatformContractCallerRaw struct {
	Contract *SharingPlatformContractCaller // Generic read-only contract binding to access the raw methods on
}

// SharingPlatformContractTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type SharingPlatformContractTransactorRaw struct {
	Contract *SharingPlatformContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSharingPlatformContract creates a new instance of SharingPlatformContract, bound to a specific deployed contract.
func NewSharingPlatformContract(address common.Address, backend bind.ContractBackend) (*SharingPlatformContract, error) {
	contract, err := bindSharingPlatformContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SharingPlatformContract{SharingPlatformContractCaller: SharingPlatformContractCaller{contract: contract}, SharingPlatformContractTransactor: SharingPlatformContractTransactor{contract: contract}, SharingPlatformContractFilterer: SharingPlatformContractFilterer{contract: contract}}, nil
}

// NewSharingPlatformContractCaller creates a new read-only instance of SharingPlatformContract, bound to a specific deployed contract.
func NewSharingPlatformContractCaller(address common.Address, caller bind.ContractCaller) (*SharingPlatformContractCaller, error) {
	contract, err := bindSharingPlatformContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SharingPlatformContractCaller{contract: contract}, nil
}

// NewSharingPlatformContractTransactor creates a new write-only instance of SharingPlatformContract, bound to a specific deployed contract.
func NewSharingPlatformContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SharingPlatformContractTransactor, error) {
	contract, err := bindSharingPlatformContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SharingPlatformContractTransactor{contract: contract}, nil
}

// NewSharingPlatformContractFilterer creates a new log filterer instance of SharingPlatformContract, bound to a specific deployed contract.
func NewSharingPlatformContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SharingPlatformContractFilterer, error) {
	contract, err := bindSharingPlatformContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SharingPlatformContractFilterer{contract: contract}, nil
}

// bindSharingPlatformContract binds a generic wrapper to an already deployed contract.
func bindSharingPlatformContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SharingPlatformContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SharingPlatformContract *SharingPlatformContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SharingPlatformContract.Contract.SharingPlatformContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SharingPlatformContract *SharingPlatformContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.SharingPlatformContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SharingPlatformContract *SharingPlatformContractRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.SharingPlatformContractTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SharingPlatformContract *SharingPlatformContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SharingPlatformContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SharingPlatformContract *SharingPlatformContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SharingPlatformContract *SharingPlatformContractTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.contract.TransactWithResult(opts, result, method, params...)
}
