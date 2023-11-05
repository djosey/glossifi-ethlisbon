// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package index

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
	_ = abi.ConvertType
)

// IndexMetaData contains all meta data concerning the Index contract.
var IndexMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getIPFSCID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ipfsCID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040518060600160405280602e815260200161061d602e91395f90816100369190610276565b50610345565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806100b757607f821691505b6020821081036100ca576100c9610073565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261012c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826100f1565b61013686836100f1565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61017a6101756101708461014e565b610157565b61014e565b9050919050565b5f819050919050565b61019383610160565b6101a761019f82610181565b8484546100fd565b825550505050565b5f90565b6101bb6101af565b6101c681848461018a565b505050565b5b818110156101e9576101de5f826101b3565b6001810190506101cc565b5050565b601f82111561022e576101ff816100d0565b610208846100e2565b81016020851015610217578190505b61022b610223856100e2565b8301826101cb565b50505b505050565b5f82821c905092915050565b5f61024e5f1984600802610233565b1980831691505092915050565b5f610266838361023f565b9150826002028217905092915050565b61027f8261003c565b67ffffffffffffffff81111561029857610297610046565b5b6102a282546100a0565b6102ad8282856101ed565b5f60209050601f8311600181146102de575f84156102cc578287015190505b6102d6858261025b565b86555061033d565b601f1984166102ec866100d0565b5f5b82811015610313578489015182556001820191506020850194506020810190506102ee565b86831015610330578489015161032c601f89168261023f565b8355505b6001600288020188555050505b505050505050565b6102cb806103525f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c806349fc73dd14610038578063d5ac21c114610056575b5f80fd5b610040610074565b60405161004d9190610218565b60405180910390f35b61005e6100ff565b60405161006b9190610218565b60405180910390f35b5f805461008090610265565b80601f01602080910402602001604051908101604052809291908181526020018280546100ac90610265565b80156100f75780601f106100ce576101008083540402835291602001916100f7565b820191905f5260205f20905b8154815290600101906020018083116100da57829003601f168201915b505050505081565b60605f805461010d90610265565b80601f016020809104026020016040519081016040528092919081815260200182805461013990610265565b80156101845780601f1061015b57610100808354040283529160200191610184565b820191905f5260205f20905b81548152906001019060200180831161016757829003601f168201915b5050505050905090565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156101c55780820151818401526020810190506101aa565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6101ea8261018e565b6101f48185610198565b93506102048185602086016101a8565b61020d816101d0565b840191505092915050565b5f6020820190508181035f83015261023081846101e0565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061027c57607f821691505b60208210810361028f5761028e610238565b5b5091905056fea26469706673582212200b50918e16eed99587328ba7ec27f412d6608c179358eba505edbbf09595c16364736f6c63430008140033516d553761754a436b71774565743654745478443937786e3571344a745a4d6a704a6b787445785331476a79554b",
}

// IndexABI is the input ABI used to generate the binding from.
// Deprecated: Use IndexMetaData.ABI instead.
var IndexABI = IndexMetaData.ABI

// IndexBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IndexMetaData.Bin instead.
var IndexBin = IndexMetaData.Bin

// DeployIndex deploys a new Ethereum contract, binding an instance of Index to it.
func DeployIndex(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Index, error) {
	parsed, err := IndexMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IndexBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Index{IndexCaller: IndexCaller{contract: contract}, IndexTransactor: IndexTransactor{contract: contract}, IndexFilterer: IndexFilterer{contract: contract}}, nil
}

// Index is an auto generated Go binding around an Ethereum contract.
type Index struct {
	IndexCaller     // Read-only binding to the contract
	IndexTransactor // Write-only binding to the contract
	IndexFilterer   // Log filterer for contract events
}

// IndexCaller is an auto generated read-only Go binding around an Ethereum contract.
type IndexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IndexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IndexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IndexSession struct {
	Contract     *Index            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IndexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IndexCallerSession struct {
	Contract *IndexCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IndexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IndexTransactorSession struct {
	Contract     *IndexTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IndexRaw is an auto generated low-level Go binding around an Ethereum contract.
type IndexRaw struct {
	Contract *Index // Generic contract binding to access the raw methods on
}

// IndexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IndexCallerRaw struct {
	Contract *IndexCaller // Generic read-only contract binding to access the raw methods on
}

// IndexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IndexTransactorRaw struct {
	Contract *IndexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIndex creates a new instance of Index, bound to a specific deployed contract.
func NewIndex(address common.Address, backend bind.ContractBackend) (*Index, error) {
	contract, err := bindIndex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Index{IndexCaller: IndexCaller{contract: contract}, IndexTransactor: IndexTransactor{contract: contract}, IndexFilterer: IndexFilterer{contract: contract}}, nil
}

// NewIndexCaller creates a new read-only instance of Index, bound to a specific deployed contract.
func NewIndexCaller(address common.Address, caller bind.ContractCaller) (*IndexCaller, error) {
	contract, err := bindIndex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IndexCaller{contract: contract}, nil
}

// NewIndexTransactor creates a new write-only instance of Index, bound to a specific deployed contract.
func NewIndexTransactor(address common.Address, transactor bind.ContractTransactor) (*IndexTransactor, error) {
	contract, err := bindIndex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IndexTransactor{contract: contract}, nil
}

// NewIndexFilterer creates a new log filterer instance of Index, bound to a specific deployed contract.
func NewIndexFilterer(address common.Address, filterer bind.ContractFilterer) (*IndexFilterer, error) {
	contract, err := bindIndex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IndexFilterer{contract: contract}, nil
}

// bindIndex binds a generic wrapper to an already deployed contract.
func bindIndex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IndexMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Index *IndexRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Index.Contract.IndexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Index *IndexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Index.Contract.IndexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Index *IndexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Index.Contract.IndexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Index *IndexCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Index.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Index *IndexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Index.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Index *IndexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Index.Contract.contract.Transact(opts, method, params...)
}

// GetIPFSCID is a free data retrieval call binding the contract method 0xd5ac21c1.
//
// Solidity: function getIPFSCID() view returns(string)
func (_Index *IndexCaller) GetIPFSCID(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Index.contract.Call(opts, &out, "getIPFSCID")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetIPFSCID is a free data retrieval call binding the contract method 0xd5ac21c1.
//
// Solidity: function getIPFSCID() view returns(string)
func (_Index *IndexSession) GetIPFSCID() (string, error) {
	return _Index.Contract.GetIPFSCID(&_Index.CallOpts)
}

// GetIPFSCID is a free data retrieval call binding the contract method 0xd5ac21c1.
//
// Solidity: function getIPFSCID() view returns(string)
func (_Index *IndexCallerSession) GetIPFSCID() (string, error) {
	return _Index.Contract.GetIPFSCID(&_Index.CallOpts)
}

// IpfsCID is a free data retrieval call binding the contract method 0x49fc73dd.
//
// Solidity: function ipfsCID() view returns(string)
func (_Index *IndexCaller) IpfsCID(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Index.contract.Call(opts, &out, "ipfsCID")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// IpfsCID is a free data retrieval call binding the contract method 0x49fc73dd.
//
// Solidity: function ipfsCID() view returns(string)
func (_Index *IndexSession) IpfsCID() (string, error) {
	return _Index.Contract.IpfsCID(&_Index.CallOpts)
}

// IpfsCID is a free data retrieval call binding the contract method 0x49fc73dd.
//
// Solidity: function ipfsCID() view returns(string)
func (_Index *IndexCallerSession) IpfsCID() (string, error) {
	return _Index.Contract.IpfsCID(&_Index.CallOpts)
}
