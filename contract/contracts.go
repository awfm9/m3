// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// AssertiveABI is the input ABI used to generate the binding from.
const AssertiveABI = `[]`

// AssertiveBin is the compiled bytecode used for deploying new contracts.
const AssertiveBin = `0x606060405260068060106000396000f3606060405200`

// DeployAssertive deploys a new Ethereum contract, binding an instance of Assertive to it.
func DeployAssertive(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Assertive, error) {
	parsed, err := abi.JSON(strings.NewReader(AssertiveABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AssertiveBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Assertive{AssertiveCaller: AssertiveCaller{contract: contract}, AssertiveTransactor: AssertiveTransactor{contract: contract}}, nil
}

// Assertive is an auto generated Go binding around an Ethereum contract.
type Assertive struct {
	AssertiveCaller     // Read-only binding to the contract
	AssertiveTransactor // Write-only binding to the contract
}

// AssertiveCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssertiveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertiveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssertiveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertiveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssertiveSession struct {
	Contract     *Assertive        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssertiveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssertiveCallerSession struct {
	Contract *AssertiveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AssertiveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssertiveTransactorSession struct {
	Contract     *AssertiveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AssertiveRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssertiveRaw struct {
	Contract *Assertive // Generic contract binding to access the raw methods on
}

// AssertiveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssertiveCallerRaw struct {
	Contract *AssertiveCaller // Generic read-only contract binding to access the raw methods on
}

// AssertiveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssertiveTransactorRaw struct {
	Contract *AssertiveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssertive creates a new instance of Assertive, bound to a specific deployed contract.
func NewAssertive(address common.Address, backend bind.ContractBackend) (*Assertive, error) {
	contract, err := bindAssertive(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &Assertive{AssertiveCaller: AssertiveCaller{contract: contract}, AssertiveTransactor: AssertiveTransactor{contract: contract}}, nil
}

// NewAssertiveCaller creates a new read-only instance of Assertive, bound to a specific deployed contract.
func NewAssertiveCaller(address common.Address, caller bind.ContractCaller) (*AssertiveCaller, error) {
	contract, err := bindAssertive(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AssertiveCaller{contract: contract}, nil
}

// NewAssertiveTransactor creates a new write-only instance of Assertive, bound to a specific deployed contract.
func NewAssertiveTransactor(address common.Address, transactor bind.ContractTransactor) (*AssertiveTransactor, error) {
	contract, err := bindAssertive(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AssertiveTransactor{contract: contract}, nil
}

// bindAssertive binds a generic wrapper to an already deployed contract.
func bindAssertive(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssertiveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Assertive *AssertiveRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Assertive.Contract.AssertiveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Assertive *AssertiveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assertive.Contract.AssertiveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Assertive *AssertiveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Assertive.Contract.AssertiveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Assertive *AssertiveCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Assertive.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Assertive *AssertiveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assertive.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Assertive *AssertiveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Assertive.Contract.contract.Transact(opts, method, params...)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = `[{"constant":false,"inputs":[{"name":"spender","type":"address"},{"name":"value","type":"uint256"}],"name":"approve","outputs":[{"name":"ok","type":"bool"}],"type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"from","type":"address"},{"name":"to","type":"address"},{"name":"value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"ok","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"who","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"to","type":"address"},{"name":"value","type":"uint256"}],"name":"transfer","outputs":[{"name":"ok","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"owner","type":"address"},{"name":"spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"}]`

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(ok bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(ok bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(ok bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(ok bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(ok bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(ok bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(ok bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(ok bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(ok bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20BaseABI is the input ABI used to generate the binding from.
const ERC20BaseABI = `[{"constant":false,"inputs":[{"name":"spender","type":"address"},{"name":"value","type":"uint256"}],"name":"approve","outputs":[{"name":"ok","type":"bool"}],"type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"supply","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"from","type":"address"},{"name":"to","type":"address"},{"name":"value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"ok","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"who","type":"address"}],"name":"balanceOf","outputs":[{"name":"value","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"to","type":"address"},{"name":"value","type":"uint256"}],"name":"transfer","outputs":[{"name":"ok","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"owner","type":"address"},{"name":"spender","type":"address"}],"name":"allowance","outputs":[{"name":"_allowance","type":"uint256"}],"type":"function"},{"inputs":[{"name":"initial_balance","type":"uint256"}],"type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"}]`

// ERC20BaseBin is the compiled bytecode used for deploying new contracts.
const ERC20BaseBin = `0x6060604052604051602080610382833950608060405251600160a060020a033316600090815260208190526040902081905560028190555061033d806100456000396000f3606060405236156100565760e060020a6000350463095ea7b3811461005857806318160ddd146100cd57806323b872dd146100e357806370a0823114610115578063a9059cbb14610135578063dd62ed3e14610164575b005b61019860043560243533600160a060020a03908116600081815260016020908152604080832094871680845294825280832086905580518681529051929493927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b6002545b60408051918252519081900360200190f35b610198600435602435604435600160a060020a038316600090815260208190526040812054829010156101ac57610002565b600160a060020a03600435166000908152602081905260409020546100d1565b61019860043560243533600160a060020a0316600090815260208190526040812054829010156102a157610002565b6100d1600435602435600160a060020a038281166000908152600160209081526040808320938516835292905220546100c7565b604080519115158252519081900360200190f35b600160a060020a0384811660009081526001602090815260408083203390941683529290522054829010156101e057610002565b600160a060020a03831660009081526020819052604090205461020b90835b818101829010156100c7565b151561021657610002565b600160a060020a038481166000818152600160209081526040808320338616845282528083208054889003905583835282825280832080548890039055938716808352918490208054870190558351868152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35060019392505050565b600160a060020a0383166000908152602081905260409020546102c490836101ff565b15156102cf57610002565b33600160a060020a0390811660008181526020818152604080832080548890039055938716808352918490208054870190558351868152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35060016100c756`

// DeployERC20Base deploys a new Ethereum contract, binding an instance of ERC20Base to it.
func DeployERC20Base(auth *bind.TransactOpts, backend bind.ContractBackend, initial_balance *big.Int) (common.Address, *types.Transaction, *ERC20Base, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20BaseBin), backend, initial_balance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Base{ERC20BaseCaller: ERC20BaseCaller{contract: contract}, ERC20BaseTransactor: ERC20BaseTransactor{contract: contract}}, nil
}

// ERC20Base is an auto generated Go binding around an Ethereum contract.
type ERC20Base struct {
	ERC20BaseCaller     // Read-only binding to the contract
	ERC20BaseTransactor // Write-only binding to the contract
}

// ERC20BaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BaseSession struct {
	Contract     *ERC20Base        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BaseCallerSession struct {
	Contract *ERC20BaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ERC20BaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BaseTransactorSession struct {
	Contract     *ERC20BaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ERC20BaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BaseRaw struct {
	Contract *ERC20Base // Generic contract binding to access the raw methods on
}

// ERC20BaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BaseCallerRaw struct {
	Contract *ERC20BaseCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BaseTransactorRaw struct {
	Contract *ERC20BaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Base creates a new instance of ERC20Base, bound to a specific deployed contract.
func NewERC20Base(address common.Address, backend bind.ContractBackend) (*ERC20Base, error) {
	contract, err := bindERC20Base(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &ERC20Base{ERC20BaseCaller: ERC20BaseCaller{contract: contract}, ERC20BaseTransactor: ERC20BaseTransactor{contract: contract}}, nil
}

// NewERC20BaseCaller creates a new read-only instance of ERC20Base, bound to a specific deployed contract.
func NewERC20BaseCaller(address common.Address, caller bind.ContractCaller) (*ERC20BaseCaller, error) {
	contract, err := bindERC20Base(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BaseCaller{contract: contract}, nil
}

// NewERC20BaseTransactor creates a new write-only instance of ERC20Base, bound to a specific deployed contract.
func NewERC20BaseTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BaseTransactor, error) {
	contract, err := bindERC20Base(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC20BaseTransactor{contract: contract}, nil
}

// bindERC20Base binds a generic wrapper to an already deployed contract.
func bindERC20Base(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Base *ERC20BaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Base.Contract.ERC20BaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Base *ERC20BaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Base.Contract.ERC20BaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Base *ERC20BaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Base.Contract.ERC20BaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Base *ERC20BaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Base.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Base *ERC20BaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Base.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Base *ERC20BaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Base.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_ERC20Base *ERC20BaseCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Base.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_ERC20Base *ERC20BaseSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Base.Contract.Allowance(&_ERC20Base.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_ERC20Base *ERC20BaseCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Base.Contract.Allowance(&_ERC20Base.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_ERC20Base *ERC20BaseCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Base.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_ERC20Base *ERC20BaseSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Base.Contract.BalanceOf(&_ERC20Base.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_ERC20Base *ERC20BaseCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Base.Contract.BalanceOf(&_ERC20Base.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_ERC20Base *ERC20BaseCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Base.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_ERC20Base *ERC20BaseSession) TotalSupply() (*big.Int, error) {
	return _ERC20Base.Contract.TotalSupply(&_ERC20Base.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_ERC20Base *ERC20BaseCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Base.Contract.TotalSupply(&_ERC20Base.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(ok bool)
func (_ERC20Base *ERC20BaseTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Base.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(ok bool)
func (_ERC20Base *ERC20BaseSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Base.Contract.Approve(&_ERC20Base.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(ok bool)
func (_ERC20Base *ERC20BaseTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Base.Contract.Approve(&_ERC20Base.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(ok bool)
func (_ERC20Base *ERC20BaseTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Base.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(ok bool)
func (_ERC20Base *ERC20BaseSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Base.Contract.Transfer(&_ERC20Base.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(ok bool)
func (_ERC20Base *ERC20BaseTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Base.Contract.Transfer(&_ERC20Base.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(ok bool)
func (_ERC20Base *ERC20BaseTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Base.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(ok bool)
func (_ERC20Base *ERC20BaseSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Base.Contract.TransferFrom(&_ERC20Base.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(ok bool)
func (_ERC20Base *ERC20BaseTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Base.Contract.TransferFrom(&_ERC20Base.TransactOpts, from, to, value)
}

// EventfulMarketABI is the input ABI used to generate the binding from.
const EventfulMarketABI = `[{"anonymous":false,"inputs":[{"indexed":false,"name":"id","type":"uint256"}],"name":"ItemUpdate","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"sell_how_much","type":"uint256"},{"indexed":true,"name":"sell_which_token","type":"address"},{"indexed":false,"name":"buy_how_much","type":"uint256"},{"indexed":true,"name":"buy_which_token","type":"address"}],"name":"Trade","type":"event"}]`

// EventfulMarketBin is the compiled bytecode used for deploying new contracts.
const EventfulMarketBin = `0x606060405260068060106000396000f3606060405200`

// DeployEventfulMarket deploys a new Ethereum contract, binding an instance of EventfulMarket to it.
func DeployEventfulMarket(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EventfulMarket, error) {
	parsed, err := abi.JSON(strings.NewReader(EventfulMarketABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EventfulMarketBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EventfulMarket{EventfulMarketCaller: EventfulMarketCaller{contract: contract}, EventfulMarketTransactor: EventfulMarketTransactor{contract: contract}}, nil
}

// EventfulMarket is an auto generated Go binding around an Ethereum contract.
type EventfulMarket struct {
	EventfulMarketCaller     // Read-only binding to the contract
	EventfulMarketTransactor // Write-only binding to the contract
}

// EventfulMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventfulMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventfulMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventfulMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventfulMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventfulMarketSession struct {
	Contract     *EventfulMarket   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventfulMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventfulMarketCallerSession struct {
	Contract *EventfulMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EventfulMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventfulMarketTransactorSession struct {
	Contract     *EventfulMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EventfulMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventfulMarketRaw struct {
	Contract *EventfulMarket // Generic contract binding to access the raw methods on
}

// EventfulMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventfulMarketCallerRaw struct {
	Contract *EventfulMarketCaller // Generic read-only contract binding to access the raw methods on
}

// EventfulMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventfulMarketTransactorRaw struct {
	Contract *EventfulMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventfulMarket creates a new instance of EventfulMarket, bound to a specific deployed contract.
func NewEventfulMarket(address common.Address, backend bind.ContractBackend) (*EventfulMarket, error) {
	contract, err := bindEventfulMarket(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &EventfulMarket{EventfulMarketCaller: EventfulMarketCaller{contract: contract}, EventfulMarketTransactor: EventfulMarketTransactor{contract: contract}}, nil
}

// NewEventfulMarketCaller creates a new read-only instance of EventfulMarket, bound to a specific deployed contract.
func NewEventfulMarketCaller(address common.Address, caller bind.ContractCaller) (*EventfulMarketCaller, error) {
	contract, err := bindEventfulMarket(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &EventfulMarketCaller{contract: contract}, nil
}

// NewEventfulMarketTransactor creates a new write-only instance of EventfulMarket, bound to a specific deployed contract.
func NewEventfulMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*EventfulMarketTransactor, error) {
	contract, err := bindEventfulMarket(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &EventfulMarketTransactor{contract: contract}, nil
}

// bindEventfulMarket binds a generic wrapper to an already deployed contract.
func bindEventfulMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventfulMarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventfulMarket *EventfulMarketRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EventfulMarket.Contract.EventfulMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventfulMarket *EventfulMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventfulMarket.Contract.EventfulMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventfulMarket *EventfulMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventfulMarket.Contract.EventfulMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventfulMarket *EventfulMarketCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EventfulMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventfulMarket *EventfulMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventfulMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventfulMarket *EventfulMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventfulMarket.Contract.contract.Transact(opts, method, params...)
}

// FallbackFailerABI is the input ABI used to generate the binding from.
const FallbackFailerABI = `[]`

// FallbackFailerBin is the compiled bytecode used for deploying new contracts.
const FallbackFailerBin = `0x606060405260108060106000396000f360606040523615600a575b6000600256`

// DeployFallbackFailer deploys a new Ethereum contract, binding an instance of FallbackFailer to it.
func DeployFallbackFailer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FallbackFailer, error) {
	parsed, err := abi.JSON(strings.NewReader(FallbackFailerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FallbackFailerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FallbackFailer{FallbackFailerCaller: FallbackFailerCaller{contract: contract}, FallbackFailerTransactor: FallbackFailerTransactor{contract: contract}}, nil
}

// FallbackFailer is an auto generated Go binding around an Ethereum contract.
type FallbackFailer struct {
	FallbackFailerCaller     // Read-only binding to the contract
	FallbackFailerTransactor // Write-only binding to the contract
}

// FallbackFailerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FallbackFailerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FallbackFailerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FallbackFailerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FallbackFailerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FallbackFailerSession struct {
	Contract     *FallbackFailer   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FallbackFailerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FallbackFailerCallerSession struct {
	Contract *FallbackFailerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FallbackFailerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FallbackFailerTransactorSession struct {
	Contract     *FallbackFailerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FallbackFailerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FallbackFailerRaw struct {
	Contract *FallbackFailer // Generic contract binding to access the raw methods on
}

// FallbackFailerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FallbackFailerCallerRaw struct {
	Contract *FallbackFailerCaller // Generic read-only contract binding to access the raw methods on
}

// FallbackFailerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FallbackFailerTransactorRaw struct {
	Contract *FallbackFailerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFallbackFailer creates a new instance of FallbackFailer, bound to a specific deployed contract.
func NewFallbackFailer(address common.Address, backend bind.ContractBackend) (*FallbackFailer, error) {
	contract, err := bindFallbackFailer(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &FallbackFailer{FallbackFailerCaller: FallbackFailerCaller{contract: contract}, FallbackFailerTransactor: FallbackFailerTransactor{contract: contract}}, nil
}

// NewFallbackFailerCaller creates a new read-only instance of FallbackFailer, bound to a specific deployed contract.
func NewFallbackFailerCaller(address common.Address, caller bind.ContractCaller) (*FallbackFailerCaller, error) {
	contract, err := bindFallbackFailer(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &FallbackFailerCaller{contract: contract}, nil
}

// NewFallbackFailerTransactor creates a new write-only instance of FallbackFailer, bound to a specific deployed contract.
func NewFallbackFailerTransactor(address common.Address, transactor bind.ContractTransactor) (*FallbackFailerTransactor, error) {
	contract, err := bindFallbackFailer(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &FallbackFailerTransactor{contract: contract}, nil
}

// bindFallbackFailer binds a generic wrapper to an already deployed contract.
func bindFallbackFailer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FallbackFailerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FallbackFailer *FallbackFailerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FallbackFailer.Contract.FallbackFailerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FallbackFailer *FallbackFailerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FallbackFailer.Contract.FallbackFailerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FallbackFailer *FallbackFailerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FallbackFailer.Contract.FallbackFailerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FallbackFailer *FallbackFailerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FallbackFailer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FallbackFailer *FallbackFailerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FallbackFailer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FallbackFailer *FallbackFailerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FallbackFailer.Contract.contract.Transact(opts, method, params...)
}

// SimpleMarketABI is the input ABI used to generate the binding from.
const SimpleMarketABI = `[{"constant":true,"inputs":[],"name":"last_offer_id","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"id","type":"uint256"}],"name":"cancel","outputs":[{"name":"_success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"id","type":"uint256"}],"name":"getOffer","outputs":[{"name":"","type":"uint256"},{"name":"","type":"address"},{"name":"","type":"uint256"},{"name":"","type":"address"}],"type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"offers","outputs":[{"name":"sell_how_much","type":"uint256"},{"name":"sell_which_token","type":"address"},{"name":"buy_how_much","type":"uint256"},{"name":"buy_which_token","type":"address"},{"name":"owner","type":"address"},{"name":"active","type":"bool"}],"type":"function"},{"constant":false,"inputs":[{"name":"id","type":"uint256"},{"name":"quantity","type":"uint256"}],"name":"buyPartial","outputs":[{"name":"_success","type":"bool"}],"type":"function"},{"constant":false,"inputs":[{"name":"id","type":"uint256"}],"name":"buy","outputs":[{"name":"_success","type":"bool"}],"type":"function"},{"constant":false,"inputs":[{"name":"sell_how_much","type":"uint256"},{"name":"sell_which_token","type":"address"},{"name":"buy_how_much","type":"uint256"},{"name":"buy_which_token","type":"address"}],"name":"offer","outputs":[{"name":"id","type":"uint256"}],"type":"function"},{"anonymous":false,"inputs":[{"indexed":false,"name":"id","type":"uint256"}],"name":"ItemUpdate","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"sell_how_much","type":"uint256"},{"indexed":true,"name":"sell_which_token","type":"address"},{"indexed":false,"name":"buy_how_much","type":"uint256"},{"indexed":true,"name":"buy_which_token","type":"address"}],"name":"Trade","type":"event"}]`

// SimpleMarketBin is the compiled bytecode used for deploying new contracts.
const SimpleMarketBin = `0x60606040526108a2806100126000396000f3606060405236156100615760e060020a6000350463232cae0b811461006957806340e58ee5146100725780634579268a146100a85780638a72ea6a146100ff578063a5d0bab114610150578063d96a094a14610181578063f09ea2a6146101ad575b610203610002565b61020560005481565b61021760043560008181526001602052604081206004810154829061026c9060a060020a900460ff165b8015156107be57610002565b6004356000908152600160208181526040805193819020805460028201549482015460039290920154908652600160a060020a03918216938601939093528482019390935291166060830152519081900360800190f35b61022b600435600160208190526000918252604090912080546002820154600483015493830154600393909301549193600160a060020a0393841693919282169181169060a060020a900460ff1686565b6102176004356024356000828152600160205260408120600481015482906103879060a060020a900460ff1661009c565b610217600435600081815260016020526040812060048101546104c09060a060020a900460ff1661009c565b6102056004356024356044356064356000600060c0604051908101604052806000815260200160008152602001600081526020016000815260200160008152602001600081526020015061061d6000881161009c565b005b60408051918252519081900360200190f35b604080519115158252519081900360200190f35b60408051968752600160a060020a039586166020880152868101949094529184166060860152909216608084015290151560a0830152519081900360c00190f35b600482015461028a90600160a060020a03908116339091161461009c565b60408051835460018501547fa9059cbb00000000000000000000000000000000000000000000000000000000835233600160a060020a0390811660048501526024840192909252925192169163a9059cbb9160448181019260209290919082900301816000876161da5a03f11561000257505060405151915061030e90508161009c565b60008481526001602081815260408084208481559283018054600160a060020a031990811690915560028401949094556003830180549094169093556004919091018054600160a860020a0319169055815186815291516000805160206108828339815191529281900390910190a15060019392505050565b600085815260016020526040902054849010156103a757600092506104b8565b6000858152600160205260409020548414156103f2576004820154600183015460038401548454600286015461044494600160a060020a0390811694929392811692339291166104f0565b506000848152600160205260408120805460029190910154850204908111156104b85760048201546001830154600384015461056992600160a060020a039081169288929082169133918791166104f0565b60008581526001602081815260408084208481559283018054600160a060020a031990811690915560028401949094556003830180549094169093556004919091018054600160a860020a0319169055815187815291516000805160206108828339815191529281900390910190a1600192505b505092915050565b600481015481546001830154600284015460038501546105a594600160a060020a03908116949381169233929091165b6000600082600160a060020a03166323b872dd868a876040518460e060020a0281526004018084600160a060020a0316815260200183600160a060020a0316815260200182815260200193505050506020604051808303816000876161da5a03f1156100025750506040515192506107c190508261009c565b8154849003825560028201805482900390556040805186815290516000805160206108828339815191529181900360200190a1600192506104b8565b60008381526001602081815260408084208481559283018054600160a060020a031990811690915560028401949094556003830180549094169093556004919091018054600160a860020a0319169055815185815291516000805160206108828339815191529281900390910190a150600192915050565b610633600160a060020a0387166000141561009c565b61063f6000861161009c565b610655600160a060020a0385166000141561009c565b85600160a060020a03166323b872dd33308a6040518460e060020a0281526004018084600160a060020a0316815260200183600160a060020a0316815260200182815260200193505050506020604051808303816000876161da5a03f1156100025750506040515192506106ca90508261009c565b868152600160a060020a03868116602083015260408201869052848116606083015233166080820152600160a082015261070b600080546001019081905590565b60008181526001602081815260409283902085518155918201805486830151600160a060020a03199182161790915585840151600284015560038301805460608801519083161790556004929092018054608087015160a088015160a060020a02919094169390931774ff00000000000000000000000000000000000000001916929092179091558151838152915192955060008051602061088283398151915292918290030190a15050949350505050565b50565b85600160a060020a031663a9059cbb86896040518360e060020a0281526004018083600160a060020a03168152602001828152602001925050506020604051808303816000876161da5a03f11561000257505060405151915061082590508161009c565b82600160a060020a031686600160a060020a03167fa5ca35f5c7b1c108bbc4c25279f619f720805890f993005d9f00ef1e32663f9b8987604051808381526020018281526020019250505060405180910390a3505050505050505056de857d2761836ca6234345c7f7f4c783271ed7d1aedf9268b3fe32800d186fde`

// DeploySimpleMarket deploys a new Ethereum contract, binding an instance of SimpleMarket to it.
func DeploySimpleMarket(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleMarket, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleMarketABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleMarketBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleMarket{SimpleMarketCaller: SimpleMarketCaller{contract: contract}, SimpleMarketTransactor: SimpleMarketTransactor{contract: contract}}, nil
}

// SimpleMarket is an auto generated Go binding around an Ethereum contract.
type SimpleMarket struct {
	SimpleMarketCaller     // Read-only binding to the contract
	SimpleMarketTransactor // Write-only binding to the contract
}

// SimpleMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleMarketSession struct {
	Contract     *SimpleMarket     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleMarketCallerSession struct {
	Contract *SimpleMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SimpleMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleMarketTransactorSession struct {
	Contract     *SimpleMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SimpleMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleMarketRaw struct {
	Contract *SimpleMarket // Generic contract binding to access the raw methods on
}

// SimpleMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleMarketCallerRaw struct {
	Contract *SimpleMarketCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleMarketTransactorRaw struct {
	Contract *SimpleMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleMarket creates a new instance of SimpleMarket, bound to a specific deployed contract.
func NewSimpleMarket(address common.Address, backend bind.ContractBackend) (*SimpleMarket, error) {
	contract, err := bindSimpleMarket(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &SimpleMarket{SimpleMarketCaller: SimpleMarketCaller{contract: contract}, SimpleMarketTransactor: SimpleMarketTransactor{contract: contract}}, nil
}

// NewSimpleMarketCaller creates a new read-only instance of SimpleMarket, bound to a specific deployed contract.
func NewSimpleMarketCaller(address common.Address, caller bind.ContractCaller) (*SimpleMarketCaller, error) {
	contract, err := bindSimpleMarket(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleMarketCaller{contract: contract}, nil
}

// NewSimpleMarketTransactor creates a new write-only instance of SimpleMarket, bound to a specific deployed contract.
func NewSimpleMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleMarketTransactor, error) {
	contract, err := bindSimpleMarket(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SimpleMarketTransactor{contract: contract}, nil
}

// bindSimpleMarket binds a generic wrapper to an already deployed contract.
func bindSimpleMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleMarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleMarket *SimpleMarketRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleMarket.Contract.SimpleMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleMarket *SimpleMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleMarket.Contract.SimpleMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleMarket *SimpleMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleMarket.Contract.SimpleMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleMarket *SimpleMarketCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleMarket *SimpleMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleMarket *SimpleMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleMarket.Contract.contract.Transact(opts, method, params...)
}

// GetOffer is a free data retrieval call binding the contract method 0x4579268a.
//
// Solidity: function getOffer(id uint256) constant returns(uint256, address, uint256, address)
func (_SimpleMarket *SimpleMarketCaller) GetOffer(opts *bind.CallOpts, id *big.Int) (*big.Int, common.Address, *big.Int, common.Address, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _SimpleMarket.contract.Call(opts, out, "getOffer", id)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetOffer is a free data retrieval call binding the contract method 0x4579268a.
//
// Solidity: function getOffer(id uint256) constant returns(uint256, address, uint256, address)
func (_SimpleMarket *SimpleMarketSession) GetOffer(id *big.Int) (*big.Int, common.Address, *big.Int, common.Address, error) {
	return _SimpleMarket.Contract.GetOffer(&_SimpleMarket.CallOpts, id)
}

// GetOffer is a free data retrieval call binding the contract method 0x4579268a.
//
// Solidity: function getOffer(id uint256) constant returns(uint256, address, uint256, address)
func (_SimpleMarket *SimpleMarketCallerSession) GetOffer(id *big.Int) (*big.Int, common.Address, *big.Int, common.Address, error) {
	return _SimpleMarket.Contract.GetOffer(&_SimpleMarket.CallOpts, id)
}

// Last_offer_id is a free data retrieval call binding the contract method 0x232cae0b.
//
// Solidity: function last_offer_id() constant returns(uint256)
func (_SimpleMarket *SimpleMarketCaller) Last_offer_id(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleMarket.contract.Call(opts, out, "last_offer_id")
	return *ret0, err
}

// Last_offer_id is a free data retrieval call binding the contract method 0x232cae0b.
//
// Solidity: function last_offer_id() constant returns(uint256)
func (_SimpleMarket *SimpleMarketSession) Last_offer_id() (*big.Int, error) {
	return _SimpleMarket.Contract.Last_offer_id(&_SimpleMarket.CallOpts)
}

// Last_offer_id is a free data retrieval call binding the contract method 0x232cae0b.
//
// Solidity: function last_offer_id() constant returns(uint256)
func (_SimpleMarket *SimpleMarketCallerSession) Last_offer_id() (*big.Int, error) {
	return _SimpleMarket.Contract.Last_offer_id(&_SimpleMarket.CallOpts)
}

// Offers is a free data retrieval call binding the contract method 0x8a72ea6a.
//
// Solidity: function offers( uint256) constant returns(sell_how_much uint256, sell_which_token address, buy_how_much uint256, buy_which_token address, owner address, active bool)
func (_SimpleMarket *SimpleMarketCaller) Offers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Sell_how_much    *big.Int
	Sell_which_token common.Address
	Buy_how_much     *big.Int
	Buy_which_token  common.Address
	Owner            common.Address
	Active           bool
}, error) {
	ret := new(struct {
		Sell_how_much    *big.Int
		Sell_which_token common.Address
		Buy_how_much     *big.Int
		Buy_which_token  common.Address
		Owner            common.Address
		Active           bool
	})
	out := ret
	err := _SimpleMarket.contract.Call(opts, out, "offers", arg0)
	return *ret, err
}

// Offers is a free data retrieval call binding the contract method 0x8a72ea6a.
//
// Solidity: function offers( uint256) constant returns(sell_how_much uint256, sell_which_token address, buy_how_much uint256, buy_which_token address, owner address, active bool)
func (_SimpleMarket *SimpleMarketSession) Offers(arg0 *big.Int) (struct {
	Sell_how_much    *big.Int
	Sell_which_token common.Address
	Buy_how_much     *big.Int
	Buy_which_token  common.Address
	Owner            common.Address
	Active           bool
}, error) {
	return _SimpleMarket.Contract.Offers(&_SimpleMarket.CallOpts, arg0)
}

// Offers is a free data retrieval call binding the contract method 0x8a72ea6a.
//
// Solidity: function offers( uint256) constant returns(sell_how_much uint256, sell_which_token address, buy_how_much uint256, buy_which_token address, owner address, active bool)
func (_SimpleMarket *SimpleMarketCallerSession) Offers(arg0 *big.Int) (struct {
	Sell_how_much    *big.Int
	Sell_which_token common.Address
	Buy_how_much     *big.Int
	Buy_which_token  common.Address
	Owner            common.Address
	Active           bool
}, error) {
	return _SimpleMarket.Contract.Offers(&_SimpleMarket.CallOpts, arg0)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(id uint256) returns(_success bool)
func (_SimpleMarket *SimpleMarketTransactor) Buy(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _SimpleMarket.contract.Transact(opts, "buy", id)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(id uint256) returns(_success bool)
func (_SimpleMarket *SimpleMarketSession) Buy(id *big.Int) (*types.Transaction, error) {
	return _SimpleMarket.Contract.Buy(&_SimpleMarket.TransactOpts, id)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(id uint256) returns(_success bool)
func (_SimpleMarket *SimpleMarketTransactorSession) Buy(id *big.Int) (*types.Transaction, error) {
	return _SimpleMarket.Contract.Buy(&_SimpleMarket.TransactOpts, id)
}

// BuyPartial is a paid mutator transaction binding the contract method 0xa5d0bab1.
//
// Solidity: function buyPartial(id uint256, quantity uint256) returns(_success bool)
func (_SimpleMarket *SimpleMarketTransactor) BuyPartial(opts *bind.TransactOpts, id *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _SimpleMarket.contract.Transact(opts, "buyPartial", id, quantity)
}

// BuyPartial is a paid mutator transaction binding the contract method 0xa5d0bab1.
//
// Solidity: function buyPartial(id uint256, quantity uint256) returns(_success bool)
func (_SimpleMarket *SimpleMarketSession) BuyPartial(id *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _SimpleMarket.Contract.BuyPartial(&_SimpleMarket.TransactOpts, id, quantity)
}

// BuyPartial is a paid mutator transaction binding the contract method 0xa5d0bab1.
//
// Solidity: function buyPartial(id uint256, quantity uint256) returns(_success bool)
func (_SimpleMarket *SimpleMarketTransactorSession) BuyPartial(id *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _SimpleMarket.Contract.BuyPartial(&_SimpleMarket.TransactOpts, id, quantity)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(id uint256) returns(_success bool)
func (_SimpleMarket *SimpleMarketTransactor) Cancel(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _SimpleMarket.contract.Transact(opts, "cancel", id)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(id uint256) returns(_success bool)
func (_SimpleMarket *SimpleMarketSession) Cancel(id *big.Int) (*types.Transaction, error) {
	return _SimpleMarket.Contract.Cancel(&_SimpleMarket.TransactOpts, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(id uint256) returns(_success bool)
func (_SimpleMarket *SimpleMarketTransactorSession) Cancel(id *big.Int) (*types.Transaction, error) {
	return _SimpleMarket.Contract.Cancel(&_SimpleMarket.TransactOpts, id)
}

// Offer is a paid mutator transaction binding the contract method 0xf09ea2a6.
//
// Solidity: function offer(sell_how_much uint256, sell_which_token address, buy_how_much uint256, buy_which_token address) returns(id uint256)
func (_SimpleMarket *SimpleMarketTransactor) Offer(opts *bind.TransactOpts, sell_how_much *big.Int, sell_which_token common.Address, buy_how_much *big.Int, buy_which_token common.Address) (*types.Transaction, error) {
	return _SimpleMarket.contract.Transact(opts, "offer", sell_how_much, sell_which_token, buy_how_much, buy_which_token)
}

// Offer is a paid mutator transaction binding the contract method 0xf09ea2a6.
//
// Solidity: function offer(sell_how_much uint256, sell_which_token address, buy_how_much uint256, buy_which_token address) returns(id uint256)
func (_SimpleMarket *SimpleMarketSession) Offer(sell_how_much *big.Int, sell_which_token common.Address, buy_how_much *big.Int, buy_which_token common.Address) (*types.Transaction, error) {
	return _SimpleMarket.Contract.Offer(&_SimpleMarket.TransactOpts, sell_how_much, sell_which_token, buy_how_much, buy_which_token)
}

// Offer is a paid mutator transaction binding the contract method 0xf09ea2a6.
//
// Solidity: function offer(sell_how_much uint256, sell_which_token address, buy_how_much uint256, buy_which_token address) returns(id uint256)
func (_SimpleMarket *SimpleMarketTransactorSession) Offer(sell_how_much *big.Int, sell_which_token common.Address, buy_how_much *big.Int, buy_which_token common.Address) (*types.Transaction, error) {
	return _SimpleMarket.Contract.Offer(&_SimpleMarket.TransactOpts, sell_how_much, sell_which_token, buy_how_much, buy_which_token)
}

// TradeProxyABI is the input ABI used to generate the binding from.
const TradeProxyABI = `[{"constant":false,"inputs":[{"name":"market","type":"address"},{"name":"firstId","type":"uint256"},{"name":"secondId","type":"uint256"}],"name":"trade","outputs":[],"type":"function"}]`

// TradeProxyBin is the compiled bytecode used for deploying new contracts.
const TradeProxyBin = `0x6060604052610503806100126000396000f36060604052361561001f5760e060020a6000350463b44b96e18114610027575b610135610002565b61013560043560243560443560006000600060006000600060006000600060008c600160a060020a0316634579268a8d6040518260e060020a028152600401808281526020019150506080604051808303816000876161da5a03f115610002575050506040518051906020018051906020018051906020018051906020015099509950995099508c600160a060020a0316634579268a8c6040518260e060020a028152600401808281526020019150506080604051808303816000876161da5a03f1156100025750505060405180519060200180519060200180519060200180519060200150955095509550955082600160a060020a031689600160a060020a03161415156101e957610002565b005b82600160a060020a031663095ea7b38e866040518360e060020a0281526004018083600160a060020a03168152602001828152602001925050506020604051808303816000876161da5a03f1156100025750505060405180519060200150508c600160a060020a031663d96a094a8c6040518260e060020a028152600401808281526020019150506020604051808303816000876161da5a03f115610002575050505b50505050505050505050505050565b84600160a060020a031687600160a060020a031614151561020957610002565b86600160a060020a031663dd62ed3e33306040518360e060020a0281526004018083600160a060020a0316815260200182600160a060020a03168152602001925050506020604051808303816000876161da5a03f115610002575050604051519250508181141561027957610002565b818811156103355786600160a060020a031663095ea7b38e846040518360e060020a0281526004018083600160a060020a03168152602001828152602001925050506020604051808303816000876161da5a03f1156100025750505060405180519060200150508c600160a060020a031663a5d0bab18d8c8b8602046040518360e060020a02815260040180838152602001828152602001925050506020604051808303816000876161da5a03f11561000257506103d8915050565b86600160a060020a031663095ea7b38e8a6040518360e060020a0281526004018083600160a060020a03168152602001828152602001925050506020604051808303816000876161da5a03f1156100025750505060405180519060200150508c600160a060020a031663d96a094a8d6040518260e060020a028152600401808281526020019150506020604051808303816000876161da5a03f115610002575050505b82600160a060020a031663dd62ed3e33306040518360e060020a0281526004018083600160a060020a0316815260200182600160a060020a03168152602001925050506020604051808303816000876161da5a03f1156100025750506040515191821415905061044757610002565b808411156101375782600160a060020a031663095ea7b38e836040518360e060020a0281526004018083600160a060020a03168152602001828152602001925050506020604051808303816000876161da5a03f1156100025750505060405180519060200150508c600160a060020a031663a5d0bab18c88878502046040518360e060020a02815260040180838152602001828152602001925050506020604051808303816000876161da5a03f11561000257506101da91505056`

// DeployTradeProxy deploys a new Ethereum contract, binding an instance of TradeProxy to it.
func DeployTradeProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TradeProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(TradeProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TradeProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TradeProxy{TradeProxyCaller: TradeProxyCaller{contract: contract}, TradeProxyTransactor: TradeProxyTransactor{contract: contract}}, nil
}

// TradeProxy is an auto generated Go binding around an Ethereum contract.
type TradeProxy struct {
	TradeProxyCaller     // Read-only binding to the contract
	TradeProxyTransactor // Write-only binding to the contract
}

// TradeProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradeProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradeProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradeProxySession struct {
	Contract     *TradeProxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradeProxyCallerSession struct {
	Contract *TradeProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TradeProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradeProxyTransactorSession struct {
	Contract     *TradeProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TradeProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradeProxyRaw struct {
	Contract *TradeProxy // Generic contract binding to access the raw methods on
}

// TradeProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradeProxyCallerRaw struct {
	Contract *TradeProxyCaller // Generic read-only contract binding to access the raw methods on
}

// TradeProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradeProxyTransactorRaw struct {
	Contract *TradeProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTradeProxy creates a new instance of TradeProxy, bound to a specific deployed contract.
func NewTradeProxy(address common.Address, backend bind.ContractBackend) (*TradeProxy, error) {
	contract, err := bindTradeProxy(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &TradeProxy{TradeProxyCaller: TradeProxyCaller{contract: contract}, TradeProxyTransactor: TradeProxyTransactor{contract: contract}}, nil
}

// NewTradeProxyCaller creates a new read-only instance of TradeProxy, bound to a specific deployed contract.
func NewTradeProxyCaller(address common.Address, caller bind.ContractCaller) (*TradeProxyCaller, error) {
	contract, err := bindTradeProxy(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TradeProxyCaller{contract: contract}, nil
}

// NewTradeProxyTransactor creates a new write-only instance of TradeProxy, bound to a specific deployed contract.
func NewTradeProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*TradeProxyTransactor, error) {
	contract, err := bindTradeProxy(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TradeProxyTransactor{contract: contract}, nil
}

// bindTradeProxy binds a generic wrapper to an already deployed contract.
func bindTradeProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TradeProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradeProxy *TradeProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TradeProxy.Contract.TradeProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradeProxy *TradeProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradeProxy.Contract.TradeProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradeProxy *TradeProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradeProxy.Contract.TradeProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradeProxy *TradeProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TradeProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradeProxy *TradeProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradeProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradeProxy *TradeProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradeProxy.Contract.contract.Transact(opts, method, params...)
}

// Trade is a paid mutator transaction binding the contract method 0xb44b96e1.
//
// Solidity: function trade(market address, firstId uint256, secondId uint256) returns()
func (_TradeProxy *TradeProxyTransactor) Trade(opts *bind.TransactOpts, market common.Address, firstId *big.Int, secondId *big.Int) (*types.Transaction, error) {
	return _TradeProxy.contract.Transact(opts, "trade", market, firstId, secondId)
}

// Trade is a paid mutator transaction binding the contract method 0xb44b96e1.
//
// Solidity: function trade(market address, firstId uint256, secondId uint256) returns()
func (_TradeProxy *TradeProxySession) Trade(market common.Address, firstId *big.Int, secondId *big.Int) (*types.Transaction, error) {
	return _TradeProxy.Contract.Trade(&_TradeProxy.TransactOpts, market, firstId, secondId)
}

// Trade is a paid mutator transaction binding the contract method 0xb44b96e1.
//
// Solidity: function trade(market address, firstId uint256, secondId uint256) returns()
func (_TradeProxy *TradeProxyTransactorSession) Trade(market common.Address, firstId *big.Int, secondId *big.Int) (*types.Transaction, error) {
	return _TradeProxy.Contract.Trade(&_TradeProxy.TransactOpts, market, firstId, secondId)
}
