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
