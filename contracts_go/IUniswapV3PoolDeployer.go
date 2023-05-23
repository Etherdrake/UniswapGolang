// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dexabi

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DexabiMetaData contains all meta data concerning the Dexabi contract.
var DexabiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"parameters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DexabiABI is the input ABI used to generate the binding from.
// Deprecated: Use DexabiMetaData.ABI instead.
var DexabiABI = DexabiMetaData.ABI

// Dexabi is an auto generated Go binding around an Ethereum contract.
type Dexabi struct {
	DexabiCaller     // Read-only binding to the contract
	DexabiTransactor // Write-only binding to the contract
	DexabiFilterer   // Log filterer for contract events
}

// DexabiCaller is an auto generated read-only Go binding around an Ethereum contract.
type DexabiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexabiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DexabiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexabiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DexabiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexabiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DexabiSession struct {
	Contract     *Dexabi           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexabiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DexabiCallerSession struct {
	Contract *DexabiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DexabiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DexabiTransactorSession struct {
	Contract     *DexabiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexabiRaw is an auto generated low-level Go binding around an Ethereum contract.
type DexabiRaw struct {
	Contract *Dexabi // Generic contract binding to access the raw methods on
}

// DexabiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DexabiCallerRaw struct {
	Contract *DexabiCaller // Generic read-only contract binding to access the raw methods on
}

// DexabiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DexabiTransactorRaw struct {
	Contract *DexabiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDexabi creates a new instance of Dexabi, bound to a specific deployed contract.
func NewDexabi(address common.Address, backend bind.ContractBackend) (*Dexabi, error) {
	contract, err := bindDexabi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dexabi{DexabiCaller: DexabiCaller{contract: contract}, DexabiTransactor: DexabiTransactor{contract: contract}, DexabiFilterer: DexabiFilterer{contract: contract}}, nil
}

// NewDexabiCaller creates a new read-only instance of Dexabi, bound to a specific deployed contract.
func NewDexabiCaller(address common.Address, caller bind.ContractCaller) (*DexabiCaller, error) {
	contract, err := bindDexabi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DexabiCaller{contract: contract}, nil
}

// NewDexabiTransactor creates a new write-only instance of Dexabi, bound to a specific deployed contract.
func NewDexabiTransactor(address common.Address, transactor bind.ContractTransactor) (*DexabiTransactor, error) {
	contract, err := bindDexabi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DexabiTransactor{contract: contract}, nil
}

// NewDexabiFilterer creates a new log filterer instance of Dexabi, bound to a specific deployed contract.
func NewDexabiFilterer(address common.Address, filterer bind.ContractFilterer) (*DexabiFilterer, error) {
	contract, err := bindDexabi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DexabiFilterer{contract: contract}, nil
}

// bindDexabi binds a generic wrapper to an already deployed contract.
func bindDexabi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DexabiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dexabi *DexabiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dexabi.Contract.DexabiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dexabi *DexabiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dexabi.Contract.DexabiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dexabi *DexabiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dexabi.Contract.DexabiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dexabi *DexabiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dexabi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dexabi *DexabiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dexabi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dexabi *DexabiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dexabi.Contract.contract.Transact(opts, method, params...)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_Dexabi *DexabiCaller) Parameters(opts *bind.CallOpts) (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	var out []interface{}
	err := _Dexabi.contract.Call(opts, &out, "parameters")

	outstruct := new(struct {
		Factory     common.Address
		Token0      common.Address
		Token1      common.Address
		Fee         *big.Int
		TickSpacing *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Factory = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token0 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TickSpacing = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_Dexabi *DexabiSession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _Dexabi.Contract.Parameters(&_Dexabi.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_Dexabi *DexabiCallerSession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _Dexabi.Contract.Parameters(&_Dexabi.CallOpts)
}
