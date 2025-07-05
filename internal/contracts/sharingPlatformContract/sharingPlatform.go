// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sharingPlatformContract

import (
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
const SharingPlatformContractABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetGeneSharing\",\"type\":\"address\"}],\"name\":\"CreateAllFromThirdParty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOfficial\",\"type\":\"bool\"}],\"name\":\"GeneSharingCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ProxyCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MetadataContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"geneSharingAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"dataHash\",\"type\":\"bytes32[]\"}],\"name\":\"addMetadataBatchesFromCreativeWorkSpace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"dataHashs\",\"type\":\"bytes32[]\"}],\"name\":\"createAllFromThirdParty\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createEmptyGeneSharingFromCreativeWorkSpace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"geneSharingAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"}],\"name\":\"obtainViewAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"metadataContract\",\"type\":\"address\"}],\"name\":\"setMetadataContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"setUserAuthStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"updateMetadataSharingStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"viewer\",\"type\":\"address\"}],\"name\":\"verifyViewAccess\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// MetadataContract is a free data retrieval call binding the contract method 0xfa204c96.
//
// Solidity: function MetadataContract() constant returns(address)
func (_SharingPlatformContract *SharingPlatformContractCaller) MetadataContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SharingPlatformContract.contract.Call(opts, out, "MetadataContract")
	return *ret0, err
}

// MetadataContract is a free data retrieval call binding the contract method 0xfa204c96.
//
// Solidity: function MetadataContract() constant returns(address)
func (_SharingPlatformContract *SharingPlatformContractSession) MetadataContract() (common.Address, error) {
	return _SharingPlatformContract.Contract.MetadataContract(&_SharingPlatformContract.CallOpts)
}

// MetadataContract is a free data retrieval call binding the contract method 0xfa204c96.
//
// Solidity: function MetadataContract() constant returns(address)
func (_SharingPlatformContract *SharingPlatformContractCallerSession) MetadataContract() (common.Address, error) {
	return _SharingPlatformContract.Contract.MetadataContract(&_SharingPlatformContract.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_SharingPlatformContract *SharingPlatformContractCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SharingPlatformContract.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_SharingPlatformContract *SharingPlatformContractSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SharingPlatformContract.Contract.Allowance(&_SharingPlatformContract.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_SharingPlatformContract *SharingPlatformContractCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SharingPlatformContract.Contract.Allowance(&_SharingPlatformContract.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SharingPlatformContract *SharingPlatformContractCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SharingPlatformContract.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SharingPlatformContract *SharingPlatformContractSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SharingPlatformContract.Contract.BalanceOf(&_SharingPlatformContract.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SharingPlatformContract *SharingPlatformContractCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SharingPlatformContract.Contract.BalanceOf(&_SharingPlatformContract.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SharingPlatformContract *SharingPlatformContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SharingPlatformContract.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SharingPlatformContract *SharingPlatformContractSession) Decimals() (uint8, error) {
	return _SharingPlatformContract.Contract.Decimals(&_SharingPlatformContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SharingPlatformContract *SharingPlatformContractCallerSession) Decimals() (uint8, error) {
	return _SharingPlatformContract.Contract.Decimals(&_SharingPlatformContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SharingPlatformContract *SharingPlatformContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SharingPlatformContract.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SharingPlatformContract *SharingPlatformContractSession) Name() (string, error) {
	return _SharingPlatformContract.Contract.Name(&_SharingPlatformContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SharingPlatformContract *SharingPlatformContractCallerSession) Name() (string, error) {
	return _SharingPlatformContract.Contract.Name(&_SharingPlatformContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SharingPlatformContract *SharingPlatformContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SharingPlatformContract.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SharingPlatformContract *SharingPlatformContractSession) Symbol() (string, error) {
	return _SharingPlatformContract.Contract.Symbol(&_SharingPlatformContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SharingPlatformContract *SharingPlatformContractCallerSession) Symbol() (string, error) {
	return _SharingPlatformContract.Contract.Symbol(&_SharingPlatformContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SharingPlatformContract *SharingPlatformContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SharingPlatformContract.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SharingPlatformContract *SharingPlatformContractSession) TotalSupply() (*big.Int, error) {
	return _SharingPlatformContract.Contract.TotalSupply(&_SharingPlatformContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SharingPlatformContract *SharingPlatformContractCallerSession) TotalSupply() (*big.Int, error) {
	return _SharingPlatformContract.Contract.TotalSupply(&_SharingPlatformContract.CallOpts)
}

// VerifyViewAccess is a free data retrieval call binding the contract method 0x0d576f17.
//
// Solidity: function verifyViewAccess(bytes32 dataHash, address viewer) constant returns(int256)
func (_SharingPlatformContract *SharingPlatformContractCaller) VerifyViewAccess(opts *bind.CallOpts, dataHash [32]byte, viewer common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SharingPlatformContract.contract.Call(opts, out, "verifyViewAccess", dataHash, viewer)
	return *ret0, err
}

// VerifyViewAccess is a free data retrieval call binding the contract method 0x0d576f17.
//
// Solidity: function verifyViewAccess(bytes32 dataHash, address viewer) constant returns(int256)
func (_SharingPlatformContract *SharingPlatformContractSession) VerifyViewAccess(dataHash [32]byte, viewer common.Address) (*big.Int, error) {
	return _SharingPlatformContract.Contract.VerifyViewAccess(&_SharingPlatformContract.CallOpts, dataHash, viewer)
}

// VerifyViewAccess is a free data retrieval call binding the contract method 0x0d576f17.
//
// Solidity: function verifyViewAccess(bytes32 dataHash, address viewer) constant returns(int256)
func (_SharingPlatformContract *SharingPlatformContractCallerSession) VerifyViewAccess(dataHash [32]byte, viewer common.Address) (*big.Int, error) {
	return _SharingPlatformContract.Contract.VerifyViewAccess(&_SharingPlatformContract.CallOpts, dataHash, viewer)
}

// AddMetadataBatchesFromCreativeWorkSpace is a paid mutator transaction binding the contract method 0xf3d84f0d.
//
// Solidity: function addMetadataBatchesFromCreativeWorkSpace(address geneSharingAddress, bytes32[] dataHash) returns()
func (_SharingPlatformContract *SharingPlatformContractTransactor) AddMetadataBatchesFromCreativeWorkSpace(opts *bind.TransactOpts, geneSharingAddress common.Address, dataHash [][32]byte) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _SharingPlatformContract.contract.TransactWithResult(opts, out, "addMetadataBatchesFromCreativeWorkSpace", geneSharingAddress, dataHash)
	return transaction, receipt, err
}

func (_SharingPlatformContract *SharingPlatformContractTransactor) AsyncAddMetadataBatchesFromCreativeWorkSpace(handler func(*types.Receipt, error), opts *bind.TransactOpts, geneSharingAddress common.Address, dataHash [][32]byte) (*types.Transaction, error) {
	return _SharingPlatformContract.contract.AsyncTransact(opts, handler, "addMetadataBatchesFromCreativeWorkSpace", geneSharingAddress, dataHash)
}

// AddMetadataBatchesFromCreativeWorkSpace is a paid mutator transaction binding the contract method 0xf3d84f0d.
//
// Solidity: function addMetadataBatchesFromCreativeWorkSpace(address geneSharingAddress, bytes32[] dataHash) returns()
func (_SharingPlatformContract *SharingPlatformContractSession) AddMetadataBatchesFromCreativeWorkSpace(geneSharingAddress common.Address, dataHash [][32]byte) (*types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.AddMetadataBatchesFromCreativeWorkSpace(&_SharingPlatformContract.TransactOpts, geneSharingAddress, dataHash)
}

func (_SharingPlatformContract *SharingPlatformContractSession) AsyncAddMetadataBatchesFromCreativeWorkSpace(handler func(*types.Receipt, error), geneSharingAddress common.Address, dataHash [][32]byte) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncAddMetadataBatchesFromCreativeWorkSpace(handler, &_SharingPlatformContract.TransactOpts, geneSharingAddress, dataHash)
}

// AddMetadataBatchesFromCreativeWorkSpace is a paid mutator transaction binding the contract method 0xf3d84f0d.
//
// Solidity: function addMetadataBatchesFromCreativeWorkSpace(address geneSharingAddress, bytes32[] dataHash) returns()
func (_SharingPlatformContract *SharingPlatformContractTransactorSession) AddMetadataBatchesFromCreativeWorkSpace(geneSharingAddress common.Address, dataHash [][32]byte) (*types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.AddMetadataBatchesFromCreativeWorkSpace(&_SharingPlatformContract.TransactOpts, geneSharingAddress, dataHash)
}

func (_SharingPlatformContract *SharingPlatformContractTransactorSession) AsyncAddMetadataBatchesFromCreativeWorkSpace(handler func(*types.Receipt, error), geneSharingAddress common.Address, dataHash [][32]byte) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncAddMetadataBatchesFromCreativeWorkSpace(handler, &_SharingPlatformContract.TransactOpts, geneSharingAddress, dataHash)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SharingPlatformContract *SharingPlatformContractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	transaction, receipt, err := _SharingPlatformContract.contract.TransactWithResult(opts, out, "approve", spender, value)
	return *ret0, transaction, receipt, err
}

func (_SharingPlatformContract *SharingPlatformContractTransactor) AsyncApprove(handler func(*types.Receipt, error), opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SharingPlatformContract.contract.AsyncTransact(opts, handler, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SharingPlatformContract *SharingPlatformContractSession) Approve(spender common.Address, value *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.Approve(&_SharingPlatformContract.TransactOpts, spender, value)
}

func (_SharingPlatformContract *SharingPlatformContractSession) AsyncApprove(handler func(*types.Receipt, error), spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncApprove(handler, &_SharingPlatformContract.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SharingPlatformContract *SharingPlatformContractTransactorSession) Approve(spender common.Address, value *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.Approve(&_SharingPlatformContract.TransactOpts, spender, value)
}

func (_SharingPlatformContract *SharingPlatformContractTransactorSession) AsyncApprove(handler func(*types.Receipt, error), spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncApprove(handler, &_SharingPlatformContract.TransactOpts, spender, value)
}

// CreateAllFromThirdParty is a paid mutator transaction binding the contract method 0x8e8c8ecc.
//
// Solidity: function createAllFromThirdParty(address user, bytes32[] dataHashs) returns(address)
func (_SharingPlatformContract *SharingPlatformContractTransactor) CreateAllFromThirdParty(opts *bind.TransactOpts, user common.Address, dataHashs [][32]byte) (common.Address, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	transaction, receipt, err := _SharingPlatformContract.contract.TransactWithResult(opts, out, "createAllFromThirdParty", user, dataHashs)
	return *ret0, transaction, receipt, err
}

func (_SharingPlatformContract *SharingPlatformContractTransactor) AsyncCreateAllFromThirdParty(handler func(*types.Receipt, error), opts *bind.TransactOpts, user common.Address, dataHashs [][32]byte) (*types.Transaction, error) {
	return _SharingPlatformContract.contract.AsyncTransact(opts, handler, "createAllFromThirdParty", user, dataHashs)
}

// CreateAllFromThirdParty is a paid mutator transaction binding the contract method 0x8e8c8ecc.
//
// Solidity: function createAllFromThirdParty(address user, bytes32[] dataHashs) returns(address)
func (_SharingPlatformContract *SharingPlatformContractSession) CreateAllFromThirdParty(user common.Address, dataHashs [][32]byte) (common.Address, *types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.CreateAllFromThirdParty(&_SharingPlatformContract.TransactOpts, user, dataHashs)
}

func (_SharingPlatformContract *SharingPlatformContractSession) AsyncCreateAllFromThirdParty(handler func(*types.Receipt, error), user common.Address, dataHashs [][32]byte) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncCreateAllFromThirdParty(handler, &_SharingPlatformContract.TransactOpts, user, dataHashs)
}

// CreateAllFromThirdParty is a paid mutator transaction binding the contract method 0x8e8c8ecc.
//
// Solidity: function createAllFromThirdParty(address user, bytes32[] dataHashs) returns(address)
func (_SharingPlatformContract *SharingPlatformContractTransactorSession) CreateAllFromThirdParty(user common.Address, dataHashs [][32]byte) (common.Address, *types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.CreateAllFromThirdParty(&_SharingPlatformContract.TransactOpts, user, dataHashs)
}

func (_SharingPlatformContract *SharingPlatformContractTransactorSession) AsyncCreateAllFromThirdParty(handler func(*types.Receipt, error), user common.Address, dataHashs [][32]byte) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncCreateAllFromThirdParty(handler, &_SharingPlatformContract.TransactOpts, user, dataHashs)
}

// CreateEmptyGeneSharingFromCreativeWorkSpace is a paid mutator transaction binding the contract method 0x233a24d7.
//
// Solidity: function createEmptyGeneSharingFromCreativeWorkSpace() returns(address)
func (_SharingPlatformContract *SharingPlatformContractTransactor) CreateEmptyGeneSharingFromCreativeWorkSpace(opts *bind.TransactOpts) (common.Address, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	transaction, receipt, err := _SharingPlatformContract.contract.TransactWithResult(opts, out, "createEmptyGeneSharingFromCreativeWorkSpace")
	return *ret0, transaction, receipt, err
}

func (_SharingPlatformContract *SharingPlatformContractTransactor) AsyncCreateEmptyGeneSharingFromCreativeWorkSpace(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SharingPlatformContract.contract.AsyncTransact(opts, handler, "createEmptyGeneSharingFromCreativeWorkSpace")
}

// CreateEmptyGeneSharingFromCreativeWorkSpace is a paid mutator transaction binding the contract method 0x233a24d7.
//
// Solidity: function createEmptyGeneSharingFromCreativeWorkSpace() returns(address)
func (_SharingPlatformContract *SharingPlatformContractSession) CreateEmptyGeneSharingFromCreativeWorkSpace() (common.Address, *types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.CreateEmptyGeneSharingFromCreativeWorkSpace(&_SharingPlatformContract.TransactOpts)
}

func (_SharingPlatformContract *SharingPlatformContractSession) AsyncCreateEmptyGeneSharingFromCreativeWorkSpace(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncCreateEmptyGeneSharingFromCreativeWorkSpace(handler, &_SharingPlatformContract.TransactOpts)
}

// CreateEmptyGeneSharingFromCreativeWorkSpace is a paid mutator transaction binding the contract method 0x233a24d7.
//
// Solidity: function createEmptyGeneSharingFromCreativeWorkSpace() returns(address)
func (_SharingPlatformContract *SharingPlatformContractTransactorSession) CreateEmptyGeneSharingFromCreativeWorkSpace() (common.Address, *types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.CreateEmptyGeneSharingFromCreativeWorkSpace(&_SharingPlatformContract.TransactOpts)
}

func (_SharingPlatformContract *SharingPlatformContractTransactorSession) AsyncCreateEmptyGeneSharingFromCreativeWorkSpace(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncCreateEmptyGeneSharingFromCreativeWorkSpace(handler, &_SharingPlatformContract.TransactOpts)
}

// ObtainViewAccess is a paid mutator transaction binding the contract method 0xf95d44ee.
//
// Solidity: function obtainViewAccess(address geneSharingAddress, bytes32 dataHash, string remark) returns()
func (_SharingPlatformContract *SharingPlatformContractTransactor) ObtainViewAccess(opts *bind.TransactOpts, geneSharingAddress common.Address, dataHash [32]byte, remark string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _SharingPlatformContract.contract.TransactWithResult(opts, out, "obtainViewAccess", geneSharingAddress, dataHash, remark)
	return transaction, receipt, err
}

func (_SharingPlatformContract *SharingPlatformContractTransactor) AsyncObtainViewAccess(handler func(*types.Receipt, error), opts *bind.TransactOpts, geneSharingAddress common.Address, dataHash [32]byte, remark string) (*types.Transaction, error) {
	return _SharingPlatformContract.contract.AsyncTransact(opts, handler, "obtainViewAccess", geneSharingAddress, dataHash, remark)
}

// ObtainViewAccess is a paid mutator transaction binding the contract method 0xf95d44ee.
//
// Solidity: function obtainViewAccess(address geneSharingAddress, bytes32 dataHash, string remark) returns()
func (_SharingPlatformContract *SharingPlatformContractSession) ObtainViewAccess(geneSharingAddress common.Address, dataHash [32]byte, remark string) (*types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.ObtainViewAccess(&_SharingPlatformContract.TransactOpts, geneSharingAddress, dataHash, remark)
}

func (_SharingPlatformContract *SharingPlatformContractSession) AsyncObtainViewAccess(handler func(*types.Receipt, error), geneSharingAddress common.Address, dataHash [32]byte, remark string) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncObtainViewAccess(handler, &_SharingPlatformContract.TransactOpts, geneSharingAddress, dataHash, remark)
}

// ObtainViewAccess is a paid mutator transaction binding the contract method 0xf95d44ee.
//
// Solidity: function obtainViewAccess(address geneSharingAddress, bytes32 dataHash, string remark) returns()
func (_SharingPlatformContract *SharingPlatformContractTransactorSession) ObtainViewAccess(geneSharingAddress common.Address, dataHash [32]byte, remark string) (*types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.ObtainViewAccess(&_SharingPlatformContract.TransactOpts, geneSharingAddress, dataHash, remark)
}

func (_SharingPlatformContract *SharingPlatformContractTransactorSession) AsyncObtainViewAccess(handler func(*types.Receipt, error), geneSharingAddress common.Address, dataHash [32]byte, remark string) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncObtainViewAccess(handler, &_SharingPlatformContract.TransactOpts, geneSharingAddress, dataHash, remark)
}

// SetMetadataContract is a paid mutator transaction binding the contract method 0xe5187f43.
//
// Solidity: function setMetadataContract(address metadataContract) returns()
func (_SharingPlatformContract *SharingPlatformContractTransactor) SetMetadataContract(opts *bind.TransactOpts, metadataContract common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _SharingPlatformContract.contract.TransactWithResult(opts, out, "setMetadataContract", metadataContract)
	return transaction, receipt, err
}

func (_SharingPlatformContract *SharingPlatformContractTransactor) AsyncSetMetadataContract(handler func(*types.Receipt, error), opts *bind.TransactOpts, metadataContract common.Address) (*types.Transaction, error) {
	return _SharingPlatformContract.contract.AsyncTransact(opts, handler, "setMetadataContract", metadataContract)
}

// SetMetadataContract is a paid mutator transaction binding the contract method 0xe5187f43.
//
// Solidity: function setMetadataContract(address metadataContract) returns()
func (_SharingPlatformContract *SharingPlatformContractSession) SetMetadataContract(metadataContract common.Address) (*types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.SetMetadataContract(&_SharingPlatformContract.TransactOpts, metadataContract)
}

func (_SharingPlatformContract *SharingPlatformContractSession) AsyncSetMetadataContract(handler func(*types.Receipt, error), metadataContract common.Address) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncSetMetadataContract(handler, &_SharingPlatformContract.TransactOpts, metadataContract)
}

// SetMetadataContract is a paid mutator transaction binding the contract method 0xe5187f43.
//
// Solidity: function setMetadataContract(address metadataContract) returns()
func (_SharingPlatformContract *SharingPlatformContractTransactorSession) SetMetadataContract(metadataContract common.Address) (*types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.SetMetadataContract(&_SharingPlatformContract.TransactOpts, metadataContract)
}

func (_SharingPlatformContract *SharingPlatformContractTransactorSession) AsyncSetMetadataContract(handler func(*types.Receipt, error), metadataContract common.Address) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncSetMetadataContract(handler, &_SharingPlatformContract.TransactOpts, metadataContract)
}

// SetUserAuthStatus is a paid mutator transaction binding the contract method 0xe7129d92.
//
// Solidity: function setUserAuthStatus(address user) returns()
func (_SharingPlatformContract *SharingPlatformContractTransactor) SetUserAuthStatus(opts *bind.TransactOpts, user common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _SharingPlatformContract.contract.TransactWithResult(opts, out, "setUserAuthStatus", user)
	return transaction, receipt, err
}

func (_SharingPlatformContract *SharingPlatformContractTransactor) AsyncSetUserAuthStatus(handler func(*types.Receipt, error), opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _SharingPlatformContract.contract.AsyncTransact(opts, handler, "setUserAuthStatus", user)
}

// SetUserAuthStatus is a paid mutator transaction binding the contract method 0xe7129d92.
//
// Solidity: function setUserAuthStatus(address user) returns()
func (_SharingPlatformContract *SharingPlatformContractSession) SetUserAuthStatus(user common.Address) (*types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.SetUserAuthStatus(&_SharingPlatformContract.TransactOpts, user)
}

func (_SharingPlatformContract *SharingPlatformContractSession) AsyncSetUserAuthStatus(handler func(*types.Receipt, error), user common.Address) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncSetUserAuthStatus(handler, &_SharingPlatformContract.TransactOpts, user)
}

// SetUserAuthStatus is a paid mutator transaction binding the contract method 0xe7129d92.
//
// Solidity: function setUserAuthStatus(address user) returns()
func (_SharingPlatformContract *SharingPlatformContractTransactorSession) SetUserAuthStatus(user common.Address) (*types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.SetUserAuthStatus(&_SharingPlatformContract.TransactOpts, user)
}

func (_SharingPlatformContract *SharingPlatformContractTransactorSession) AsyncSetUserAuthStatus(handler func(*types.Receipt, error), user common.Address) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncSetUserAuthStatus(handler, &_SharingPlatformContract.TransactOpts, user)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SharingPlatformContract *SharingPlatformContractTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	transaction, receipt, err := _SharingPlatformContract.contract.TransactWithResult(opts, out, "transfer", to, value)
	return *ret0, transaction, receipt, err
}

func (_SharingPlatformContract *SharingPlatformContractTransactor) AsyncTransfer(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SharingPlatformContract.contract.AsyncTransact(opts, handler, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SharingPlatformContract *SharingPlatformContractSession) Transfer(to common.Address, value *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.Transfer(&_SharingPlatformContract.TransactOpts, to, value)
}

func (_SharingPlatformContract *SharingPlatformContractSession) AsyncTransfer(handler func(*types.Receipt, error), to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncTransfer(handler, &_SharingPlatformContract.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SharingPlatformContract *SharingPlatformContractTransactorSession) Transfer(to common.Address, value *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.Transfer(&_SharingPlatformContract.TransactOpts, to, value)
}

func (_SharingPlatformContract *SharingPlatformContractTransactorSession) AsyncTransfer(handler func(*types.Receipt, error), to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncTransfer(handler, &_SharingPlatformContract.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SharingPlatformContract *SharingPlatformContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	transaction, receipt, err := _SharingPlatformContract.contract.TransactWithResult(opts, out, "transferFrom", from, to, value)
	return *ret0, transaction, receipt, err
}

func (_SharingPlatformContract *SharingPlatformContractTransactor) AsyncTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SharingPlatformContract.contract.AsyncTransact(opts, handler, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SharingPlatformContract *SharingPlatformContractSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.TransferFrom(&_SharingPlatformContract.TransactOpts, from, to, value)
}

func (_SharingPlatformContract *SharingPlatformContractSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncTransferFrom(handler, &_SharingPlatformContract.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SharingPlatformContract *SharingPlatformContractTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.TransferFrom(&_SharingPlatformContract.TransactOpts, from, to, value)
}

func (_SharingPlatformContract *SharingPlatformContractTransactorSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncTransferFrom(handler, &_SharingPlatformContract.TransactOpts, from, to, value)
}

// UpdateMetadataSharingStatus is a paid mutator transaction binding the contract method 0x50628a99.
//
// Solidity: function updateMetadataSharingStatus(bytes32 dataHash, bool status) returns()
func (_SharingPlatformContract *SharingPlatformContractTransactor) UpdateMetadataSharingStatus(opts *bind.TransactOpts, dataHash [32]byte, status bool) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _SharingPlatformContract.contract.TransactWithResult(opts, out, "updateMetadataSharingStatus", dataHash, status)
	return transaction, receipt, err
}

func (_SharingPlatformContract *SharingPlatformContractTransactor) AsyncUpdateMetadataSharingStatus(handler func(*types.Receipt, error), opts *bind.TransactOpts, dataHash [32]byte, status bool) (*types.Transaction, error) {
	return _SharingPlatformContract.contract.AsyncTransact(opts, handler, "updateMetadataSharingStatus", dataHash, status)
}

// UpdateMetadataSharingStatus is a paid mutator transaction binding the contract method 0x50628a99.
//
// Solidity: function updateMetadataSharingStatus(bytes32 dataHash, bool status) returns()
func (_SharingPlatformContract *SharingPlatformContractSession) UpdateMetadataSharingStatus(dataHash [32]byte, status bool) (*types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.UpdateMetadataSharingStatus(&_SharingPlatformContract.TransactOpts, dataHash, status)
}

func (_SharingPlatformContract *SharingPlatformContractSession) AsyncUpdateMetadataSharingStatus(handler func(*types.Receipt, error), dataHash [32]byte, status bool) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncUpdateMetadataSharingStatus(handler, &_SharingPlatformContract.TransactOpts, dataHash, status)
}

// UpdateMetadataSharingStatus is a paid mutator transaction binding the contract method 0x50628a99.
//
// Solidity: function updateMetadataSharingStatus(bytes32 dataHash, bool status) returns()
func (_SharingPlatformContract *SharingPlatformContractTransactorSession) UpdateMetadataSharingStatus(dataHash [32]byte, status bool) (*types.Transaction, *types.Receipt, error) {
	return _SharingPlatformContract.Contract.UpdateMetadataSharingStatus(&_SharingPlatformContract.TransactOpts, dataHash, status)
}

func (_SharingPlatformContract *SharingPlatformContractTransactorSession) AsyncUpdateMetadataSharingStatus(handler func(*types.Receipt, error), dataHash [32]byte, status bool) (*types.Transaction, error) {
	return _SharingPlatformContract.Contract.AsyncUpdateMetadataSharingStatus(handler, &_SharingPlatformContract.TransactOpts, dataHash, status)
}

// SharingPlatformContractApproval represents a Approval event raised by the SharingPlatformContract contract.
type SharingPlatformContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SharingPlatformContract *SharingPlatformContractFilterer) WatchApproval(fromBlock *int64, handler func(int, []types.Log), owner common.Address, spender common.Address) (string, error) {
	return _SharingPlatformContract.contract.WatchLogs(fromBlock, handler, "Approval", owner, spender)
}

func (_SharingPlatformContract *SharingPlatformContractFilterer) WatchAllApproval(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _SharingPlatformContract.contract.WatchLogs(fromBlock, handler, "Approval")
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SharingPlatformContract *SharingPlatformContractFilterer) ParseApproval(log types.Log) (*SharingPlatformContractApproval, error) {
	event := new(SharingPlatformContractApproval)
	if err := _SharingPlatformContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SharingPlatformContract *SharingPlatformContractSession) WatchApproval(fromBlock *int64, handler func(int, []types.Log), owner common.Address, spender common.Address) (string, error) {
	return _SharingPlatformContract.Contract.WatchApproval(fromBlock, handler, owner, spender)
}

func (_SharingPlatformContract *SharingPlatformContractSession) WatchAllApproval(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _SharingPlatformContract.Contract.WatchAllApproval(fromBlock, handler)
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SharingPlatformContract *SharingPlatformContractSession) ParseApproval(log types.Log) (*SharingPlatformContractApproval, error) {
	return _SharingPlatformContract.Contract.ParseApproval(log)
}

// SharingPlatformContractCreateAllFromThirdParty represents a CreateAllFromThirdParty event raised by the SharingPlatformContract contract.
type SharingPlatformContractCreateAllFromThirdParty struct {
	User              common.Address
	TargetGeneSharing common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// WatchCreateAllFromThirdParty is a free log subscription operation binding the contract event 0x73d757c950d24f916fec89e64a7ef08a55c452e210fdd3e39d9806081af52e2c.
//
// Solidity: event CreateAllFromThirdParty(address indexed user, address indexed targetGeneSharing)
func (_SharingPlatformContract *SharingPlatformContractFilterer) WatchCreateAllFromThirdParty(fromBlock *int64, handler func(int, []types.Log), user common.Address, targetGeneSharing common.Address) (string, error) {
	return _SharingPlatformContract.contract.WatchLogs(fromBlock, handler, "CreateAllFromThirdParty", user, targetGeneSharing)
}

func (_SharingPlatformContract *SharingPlatformContractFilterer) WatchAllCreateAllFromThirdParty(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _SharingPlatformContract.contract.WatchLogs(fromBlock, handler, "CreateAllFromThirdParty")
}

// ParseCreateAllFromThirdParty is a log parse operation binding the contract event 0x73d757c950d24f916fec89e64a7ef08a55c452e210fdd3e39d9806081af52e2c.
//
// Solidity: event CreateAllFromThirdParty(address indexed user, address indexed targetGeneSharing)
func (_SharingPlatformContract *SharingPlatformContractFilterer) ParseCreateAllFromThirdParty(log types.Log) (*SharingPlatformContractCreateAllFromThirdParty, error) {
	event := new(SharingPlatformContractCreateAllFromThirdParty)
	if err := _SharingPlatformContract.contract.UnpackLog(event, "CreateAllFromThirdParty", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchCreateAllFromThirdParty is a free log subscription operation binding the contract event 0x73d757c950d24f916fec89e64a7ef08a55c452e210fdd3e39d9806081af52e2c.
//
// Solidity: event CreateAllFromThirdParty(address indexed user, address indexed targetGeneSharing)
func (_SharingPlatformContract *SharingPlatformContractSession) WatchCreateAllFromThirdParty(fromBlock *int64, handler func(int, []types.Log), user common.Address, targetGeneSharing common.Address) (string, error) {
	return _SharingPlatformContract.Contract.WatchCreateAllFromThirdParty(fromBlock, handler, user, targetGeneSharing)
}

func (_SharingPlatformContract *SharingPlatformContractSession) WatchAllCreateAllFromThirdParty(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _SharingPlatformContract.Contract.WatchAllCreateAllFromThirdParty(fromBlock, handler)
}

// ParseCreateAllFromThirdParty is a log parse operation binding the contract event 0x73d757c950d24f916fec89e64a7ef08a55c452e210fdd3e39d9806081af52e2c.
//
// Solidity: event CreateAllFromThirdParty(address indexed user, address indexed targetGeneSharing)
func (_SharingPlatformContract *SharingPlatformContractSession) ParseCreateAllFromThirdParty(log types.Log) (*SharingPlatformContractCreateAllFromThirdParty, error) {
	return _SharingPlatformContract.Contract.ParseCreateAllFromThirdParty(log)
}

// SharingPlatformContractGeneSharingCreated represents a GeneSharingCreated event raised by the SharingPlatformContract contract.
type SharingPlatformContractGeneSharingCreated struct {
	ContractAddress common.Address
	Creator         common.Address
	IsOfficial      bool
	Raw             types.Log // Blockchain specific contextual infos
}

// WatchGeneSharingCreated is a free log subscription operation binding the contract event 0x77c0221016edc207ab88a77ac3c77f827bb6922b821c579bbc17f80a47743efb.
//
// Solidity: event GeneSharingCreated(address indexed contractAddress, address creator, bool isOfficial)
func (_SharingPlatformContract *SharingPlatformContractFilterer) WatchGeneSharingCreated(fromBlock *int64, handler func(int, []types.Log), contractAddress common.Address) (string, error) {
	return _SharingPlatformContract.contract.WatchLogs(fromBlock, handler, "GeneSharingCreated", contractAddress)
}

func (_SharingPlatformContract *SharingPlatformContractFilterer) WatchAllGeneSharingCreated(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _SharingPlatformContract.contract.WatchLogs(fromBlock, handler, "GeneSharingCreated")
}

// ParseGeneSharingCreated is a log parse operation binding the contract event 0x77c0221016edc207ab88a77ac3c77f827bb6922b821c579bbc17f80a47743efb.
//
// Solidity: event GeneSharingCreated(address indexed contractAddress, address creator, bool isOfficial)
func (_SharingPlatformContract *SharingPlatformContractFilterer) ParseGeneSharingCreated(log types.Log) (*SharingPlatformContractGeneSharingCreated, error) {
	event := new(SharingPlatformContractGeneSharingCreated)
	if err := _SharingPlatformContract.contract.UnpackLog(event, "GeneSharingCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchGeneSharingCreated is a free log subscription operation binding the contract event 0x77c0221016edc207ab88a77ac3c77f827bb6922b821c579bbc17f80a47743efb.
//
// Solidity: event GeneSharingCreated(address indexed contractAddress, address creator, bool isOfficial)
func (_SharingPlatformContract *SharingPlatformContractSession) WatchGeneSharingCreated(fromBlock *int64, handler func(int, []types.Log), contractAddress common.Address) (string, error) {
	return _SharingPlatformContract.Contract.WatchGeneSharingCreated(fromBlock, handler, contractAddress)
}

func (_SharingPlatformContract *SharingPlatformContractSession) WatchAllGeneSharingCreated(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _SharingPlatformContract.Contract.WatchAllGeneSharingCreated(fromBlock, handler)
}

// ParseGeneSharingCreated is a log parse operation binding the contract event 0x77c0221016edc207ab88a77ac3c77f827bb6922b821c579bbc17f80a47743efb.
//
// Solidity: event GeneSharingCreated(address indexed contractAddress, address creator, bool isOfficial)
func (_SharingPlatformContract *SharingPlatformContractSession) ParseGeneSharingCreated(log types.Log) (*SharingPlatformContractGeneSharingCreated, error) {
	return _SharingPlatformContract.Contract.ParseGeneSharingCreated(log)
}

// SharingPlatformContractProxyCallExecuted represents a ProxyCallExecuted event raised by the SharingPlatformContract contract.
type SharingPlatformContractProxyCallExecuted struct {
	User   common.Address
	Target common.Address
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// WatchProxyCallExecuted is a free log subscription operation binding the contract event 0x7b76b72724c269a768f0744ce07ce7c84423347cab83dda9f76d24778993db93.
//
// Solidity: event ProxyCallExecuted(address indexed user, address indexed target, bytes data)
func (_SharingPlatformContract *SharingPlatformContractFilterer) WatchProxyCallExecuted(fromBlock *int64, handler func(int, []types.Log), user common.Address, target common.Address) (string, error) {
	return _SharingPlatformContract.contract.WatchLogs(fromBlock, handler, "ProxyCallExecuted", user, target)
}

func (_SharingPlatformContract *SharingPlatformContractFilterer) WatchAllProxyCallExecuted(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _SharingPlatformContract.contract.WatchLogs(fromBlock, handler, "ProxyCallExecuted")
}

// ParseProxyCallExecuted is a log parse operation binding the contract event 0x7b76b72724c269a768f0744ce07ce7c84423347cab83dda9f76d24778993db93.
//
// Solidity: event ProxyCallExecuted(address indexed user, address indexed target, bytes data)
func (_SharingPlatformContract *SharingPlatformContractFilterer) ParseProxyCallExecuted(log types.Log) (*SharingPlatformContractProxyCallExecuted, error) {
	event := new(SharingPlatformContractProxyCallExecuted)
	if err := _SharingPlatformContract.contract.UnpackLog(event, "ProxyCallExecuted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchProxyCallExecuted is a free log subscription operation binding the contract event 0x7b76b72724c269a768f0744ce07ce7c84423347cab83dda9f76d24778993db93.
//
// Solidity: event ProxyCallExecuted(address indexed user, address indexed target, bytes data)
func (_SharingPlatformContract *SharingPlatformContractSession) WatchProxyCallExecuted(fromBlock *int64, handler func(int, []types.Log), user common.Address, target common.Address) (string, error) {
	return _SharingPlatformContract.Contract.WatchProxyCallExecuted(fromBlock, handler, user, target)
}

func (_SharingPlatformContract *SharingPlatformContractSession) WatchAllProxyCallExecuted(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _SharingPlatformContract.Contract.WatchAllProxyCallExecuted(fromBlock, handler)
}

// ParseProxyCallExecuted is a log parse operation binding the contract event 0x7b76b72724c269a768f0744ce07ce7c84423347cab83dda9f76d24778993db93.
//
// Solidity: event ProxyCallExecuted(address indexed user, address indexed target, bytes data)
func (_SharingPlatformContract *SharingPlatformContractSession) ParseProxyCallExecuted(log types.Log) (*SharingPlatformContractProxyCallExecuted, error) {
	return _SharingPlatformContract.Contract.ParseProxyCallExecuted(log)
}

// SharingPlatformContractTransfer represents a Transfer event raised by the SharingPlatformContract contract.
type SharingPlatformContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SharingPlatformContract *SharingPlatformContractFilterer) WatchTransfer(fromBlock *int64, handler func(int, []types.Log), from common.Address, to common.Address) (string, error) {
	return _SharingPlatformContract.contract.WatchLogs(fromBlock, handler, "Transfer", from, to)
}

func (_SharingPlatformContract *SharingPlatformContractFilterer) WatchAllTransfer(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _SharingPlatformContract.contract.WatchLogs(fromBlock, handler, "Transfer")
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SharingPlatformContract *SharingPlatformContractFilterer) ParseTransfer(log types.Log) (*SharingPlatformContractTransfer, error) {
	event := new(SharingPlatformContractTransfer)
	if err := _SharingPlatformContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SharingPlatformContract *SharingPlatformContractSession) WatchTransfer(fromBlock *int64, handler func(int, []types.Log), from common.Address, to common.Address) (string, error) {
	return _SharingPlatformContract.Contract.WatchTransfer(fromBlock, handler, from, to)
}

func (_SharingPlatformContract *SharingPlatformContractSession) WatchAllTransfer(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _SharingPlatformContract.Contract.WatchAllTransfer(fromBlock, handler)
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SharingPlatformContract *SharingPlatformContractSession) ParseTransfer(log types.Log) (*SharingPlatformContractTransfer, error) {
	return _SharingPlatformContract.Contract.ParseTransfer(log)
}
