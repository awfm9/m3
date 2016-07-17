// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package binding

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// M3WalletABI is the input ABI used to generate the binding from.
const M3WalletABI = `[{"constant":false,"inputs":[{"name":"market","type":"address"},{"name":"first_id","type":"uint256"},{"name":"first_selling","type":"uint256"},{"name":"second_id","type":"uint256"},{"name":"second_selling","type":"uint256"}],"name":"atomicTradePair","outputs":[{"name":"","type":"bool"}],"type":"function"},{"constant":false,"inputs":[{"name":"amount","type":"uint256"}],"name":"withdraw","outputs":[{"name":"","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"token","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[{"name":"market","type":"address"},{"name":"order_id","type":"uint256"}],"name":"getBuyToken","outputs":[{"name":"","type":"address"}],"type":"function"},{"constant":true,"inputs":[{"name":"market","type":"address"},{"name":"order_id","type":"uint256"},{"name":"selling_amount","type":"uint256"}],"name":"getBuyingAmount","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"inputs":[],"type":"constructor"}]`

// M3WalletBin is the compiled bytecode used for deploying new contracts.
const M3WalletBin = `0x606060405260008054600160a060020a031916331790556103f8806100246000396000f3606060405260e060020a60003504630616191a81146100475780632e1a7d4d1461006d57806370a08231146100935780638832cf69146100f3578063beb481911461015c575b005b6101d2600435602435604435606435608435600060006000600060006102158a8a6100fd565b6101d260043560006103c88230600160a060020a03163110155b8015156103f557610002565b6101e6600435600081600160a060020a03166370a08231306040518260e060020a0281526004018082600160a060020a031681526020019150506020604051808303816000876161da5a03f1156100025750506040515191506103c39050565b6101f86004356024355b6000600060006000600086600160a060020a0316634579268a876040518260e060020a028152600401808281526020019150506080604051808303816000876161da5a03f1156100025750506040516060015198975050505050505050565b6101e66004356024356044355b6000600060006000600087600160a060020a0316634579268a886040518260e060020a028152600401808281526020019150506080604051808303816000876161da5a03f115610002575050604080518051910151979097029690960498975050505050505050565b604080519115158252519081900360200190f35b60408051918252519081900360200190f35b60408051600160a060020a03929092168252519081900360200190f35b93506102218a886100fd565b925061028484600160a060020a031663095ea7b38c8b6040518360e060020a0281526004018083600160a060020a03168152602001828152602001925050506020604051808303816000876161da5a03f115610002575050604051519050610087565b6102e583600160a060020a031663095ea7b38c896040518360e060020a0281526004018083600160a060020a03168152602001828152602001925050506020604051808303816000876161da5a03f115610002575050604051519050610087565b6102f08a8a8a610169565b91506102fd8a8888610169565b90506103578a600160a060020a031663a5d0bab18b856040518360e060020a02815260040180838152602001828152602001925050506020604051808303816000876161da5a03f115610002575050604051519050610087565b6103af8a600160a060020a031663a5d0bab189846040518360e060020a02815260040180838152602001828152602001925050506020604051808303816000876161da5a03f115610002575050604051519050610087565b5060019998505050505050505050565b5060015b919050565b600080546040516103bf92600160a060020a03929092169190859082818181858883f19350505050610087565b5056`

// DeployM3Wallet deploys a new Ethereum contract, binding an instance of M3Wallet to it.
func DeployM3Wallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *M3Wallet, error) {
	parsed, err := abi.JSON(strings.NewReader(M3WalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(M3WalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &M3Wallet{M3WalletCaller: M3WalletCaller{contract: contract}, M3WalletTransactor: M3WalletTransactor{contract: contract}}, nil
}

// M3Wallet is an auto generated Go binding around an Ethereum contract.
type M3Wallet struct {
	M3WalletCaller     // Read-only binding to the contract
	M3WalletTransactor // Write-only binding to the contract
}

// M3WalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type M3WalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// M3WalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type M3WalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// M3WalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type M3WalletSession struct {
	Contract     *M3Wallet         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// M3WalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type M3WalletCallerSession struct {
	Contract *M3WalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// M3WalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type M3WalletTransactorSession struct {
	Contract     *M3WalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// M3WalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type M3WalletRaw struct {
	Contract *M3Wallet // Generic contract binding to access the raw methods on
}

// M3WalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type M3WalletCallerRaw struct {
	Contract *M3WalletCaller // Generic read-only contract binding to access the raw methods on
}

// M3WalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type M3WalletTransactorRaw struct {
	Contract *M3WalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewM3Wallet creates a new instance of M3Wallet, bound to a specific deployed contract.
func NewM3Wallet(address common.Address, backend bind.ContractBackend) (*M3Wallet, error) {
	contract, err := bindM3Wallet(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &M3Wallet{M3WalletCaller: M3WalletCaller{contract: contract}, M3WalletTransactor: M3WalletTransactor{contract: contract}}, nil
}

// NewM3WalletCaller creates a new read-only instance of M3Wallet, bound to a specific deployed contract.
func NewM3WalletCaller(address common.Address, caller bind.ContractCaller) (*M3WalletCaller, error) {
	contract, err := bindM3Wallet(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &M3WalletCaller{contract: contract}, nil
}

// NewM3WalletTransactor creates a new write-only instance of M3Wallet, bound to a specific deployed contract.
func NewM3WalletTransactor(address common.Address, transactor bind.ContractTransactor) (*M3WalletTransactor, error) {
	contract, err := bindM3Wallet(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &M3WalletTransactor{contract: contract}, nil
}

// bindM3Wallet binds a generic wrapper to an already deployed contract.
func bindM3Wallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(M3WalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_M3Wallet *M3WalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _M3Wallet.Contract.M3WalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_M3Wallet *M3WalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _M3Wallet.Contract.M3WalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_M3Wallet *M3WalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _M3Wallet.Contract.M3WalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_M3Wallet *M3WalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _M3Wallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_M3Wallet *M3WalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _M3Wallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_M3Wallet *M3WalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _M3Wallet.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_M3Wallet *M3WalletCaller) BalanceOf(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _M3Wallet.contract.Call(opts, out, "balanceOf", token)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_M3Wallet *M3WalletSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _M3Wallet.Contract.BalanceOf(&_M3Wallet.CallOpts, token)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_M3Wallet *M3WalletCallerSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _M3Wallet.Contract.BalanceOf(&_M3Wallet.CallOpts, token)
}

// GetBuyToken is a free data retrieval call binding the contract method 0x8832cf69.
//
// Solidity: function getBuyToken(market address, order_id uint256) constant returns(address)
func (_M3Wallet *M3WalletCaller) GetBuyToken(opts *bind.CallOpts, market common.Address, order_id *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _M3Wallet.contract.Call(opts, out, "getBuyToken", market, order_id)
	return *ret0, err
}

// GetBuyToken is a free data retrieval call binding the contract method 0x8832cf69.
//
// Solidity: function getBuyToken(market address, order_id uint256) constant returns(address)
func (_M3Wallet *M3WalletSession) GetBuyToken(market common.Address, order_id *big.Int) (common.Address, error) {
	return _M3Wallet.Contract.GetBuyToken(&_M3Wallet.CallOpts, market, order_id)
}

// GetBuyToken is a free data retrieval call binding the contract method 0x8832cf69.
//
// Solidity: function getBuyToken(market address, order_id uint256) constant returns(address)
func (_M3Wallet *M3WalletCallerSession) GetBuyToken(market common.Address, order_id *big.Int) (common.Address, error) {
	return _M3Wallet.Contract.GetBuyToken(&_M3Wallet.CallOpts, market, order_id)
}

// GetBuyingAmount is a free data retrieval call binding the contract method 0xbeb48191.
//
// Solidity: function getBuyingAmount(market address, order_id uint256, selling_amount uint256) constant returns(uint256)
func (_M3Wallet *M3WalletCaller) GetBuyingAmount(opts *bind.CallOpts, market common.Address, order_id *big.Int, selling_amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _M3Wallet.contract.Call(opts, out, "getBuyingAmount", market, order_id, selling_amount)
	return *ret0, err
}

// GetBuyingAmount is a free data retrieval call binding the contract method 0xbeb48191.
//
// Solidity: function getBuyingAmount(market address, order_id uint256, selling_amount uint256) constant returns(uint256)
func (_M3Wallet *M3WalletSession) GetBuyingAmount(market common.Address, order_id *big.Int, selling_amount *big.Int) (*big.Int, error) {
	return _M3Wallet.Contract.GetBuyingAmount(&_M3Wallet.CallOpts, market, order_id, selling_amount)
}

// GetBuyingAmount is a free data retrieval call binding the contract method 0xbeb48191.
//
// Solidity: function getBuyingAmount(market address, order_id uint256, selling_amount uint256) constant returns(uint256)
func (_M3Wallet *M3WalletCallerSession) GetBuyingAmount(market common.Address, order_id *big.Int, selling_amount *big.Int) (*big.Int, error) {
	return _M3Wallet.Contract.GetBuyingAmount(&_M3Wallet.CallOpts, market, order_id, selling_amount)
}

// AtomicTradePair is a paid mutator transaction binding the contract method 0x0616191a.
//
// Solidity: function atomicTradePair(market address, first_id uint256, first_selling uint256, second_id uint256, second_selling uint256) returns(bool)
func (_M3Wallet *M3WalletTransactor) AtomicTradePair(opts *bind.TransactOpts, market common.Address, first_id *big.Int, first_selling *big.Int, second_id *big.Int, second_selling *big.Int) (*types.Transaction, error) {
	return _M3Wallet.contract.Transact(opts, "atomicTradePair", market, first_id, first_selling, second_id, second_selling)
}

// AtomicTradePair is a paid mutator transaction binding the contract method 0x0616191a.
//
// Solidity: function atomicTradePair(market address, first_id uint256, first_selling uint256, second_id uint256, second_selling uint256) returns(bool)
func (_M3Wallet *M3WalletSession) AtomicTradePair(market common.Address, first_id *big.Int, first_selling *big.Int, second_id *big.Int, second_selling *big.Int) (*types.Transaction, error) {
	return _M3Wallet.Contract.AtomicTradePair(&_M3Wallet.TransactOpts, market, first_id, first_selling, second_id, second_selling)
}

// AtomicTradePair is a paid mutator transaction binding the contract method 0x0616191a.
//
// Solidity: function atomicTradePair(market address, first_id uint256, first_selling uint256, second_id uint256, second_selling uint256) returns(bool)
func (_M3Wallet *M3WalletTransactorSession) AtomicTradePair(market common.Address, first_id *big.Int, first_selling *big.Int, second_id *big.Int, second_selling *big.Int) (*types.Transaction, error) {
	return _M3Wallet.Contract.AtomicTradePair(&_M3Wallet.TransactOpts, market, first_id, first_selling, second_id, second_selling)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns(bool)
func (_M3Wallet *M3WalletTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _M3Wallet.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns(bool)
func (_M3Wallet *M3WalletSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _M3Wallet.Contract.Withdraw(&_M3Wallet.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns(bool)
func (_M3Wallet *M3WalletTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _M3Wallet.Contract.Withdraw(&_M3Wallet.TransactOpts, amount)
}
