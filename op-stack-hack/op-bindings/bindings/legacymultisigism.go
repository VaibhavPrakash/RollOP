// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// LegacyMultisigIsmMetaData contains all meta data concerning the LegacyMultisigIsm contract.
var LegacyMultisigIsmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"CommitmentUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorCount\",\"type\":\"uint256\"}],\"name\":\"ValidatorEnrolled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorCount\",\"type\":\"uint256\"}],\"name\":\"ValidatorUnenrolled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"commitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"enrollValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"_domains\",\"type\":\"uint32[]\"},{\"internalType\":\"address[][]\",\"name\":\"_validators\",\"type\":\"address[][]\"}],\"name\":\"enrollValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isEnrolled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"moduleType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"_domains\",\"type\":\"uint32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_thresholds\",\"type\":\"uint8[]\"}],\"name\":\"setThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"unenrollValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"validatorsAndThreshold\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6122628061007e6000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063715018a611610097578063f211b73211610066578063f211b73214610247578063f2fde38b1461026a578063f7e83aee1461027d578063fe9f5ab11461029057600080fd5b8063715018a6146101e45780638da5cb5b146101ec5780639a6e9a2414610214578063e7d5d3a11461023457600080fd5b8063414c676f116100d3578063414c676f146101845780634422cc2a146101975780635d94f2b9146101aa5780636465e69f146101ca57600080fd5b8063020addf114610105578063174769dc1461012b5780632e0ed234146101405780634098991214610161575b600080fd5b610118610113366004611a8b565b6102a3565b6040519081526020015b60405180910390f35b61013e610139366004611aeb565b6102c7565b005b61015361014e366004611b99565b610439565b604051610122929190611c2c565b61017461016f366004611c75565b61047d565b6040519015158152602001610122565b61013e610192366004611aeb565b6104ab565b61013e6101a5366004611cb9565b61058b565b6101186101b8366004611a8b565b60036020526000908152604090205481565b6101d2600381565b60405160ff9091168152602001610122565b61013e61069c565b60005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610122565b610227610222366004611a8b565b6106b0565b6040516101229190611ce3565b61013e610242366004611c75565b610781565b6101d2610255366004611a8b565b60016020526000908152604090205460ff1681565b61013e610278366004611cf6565b610907565b61017461028b366004611d11565b6109be565b61013e61029e366004611c75565b610aaf565b63ffffffff811660009081526002602052604081206102c190610aca565b92915050565b6102cf610ad4565b8281811461033e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f216c656e6774680000000000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b60005b818110156104315736600085858481811061035e5761035e611d71565b90506020028101906103709190611da0565b90925090508060005b818110156103e9576103d78a8a8781811061039657610396611d71565b90506020020160208101906103ab9190611a8b565b8585848181106103bd576103bd611d71565b90506020020160208101906103d29190611cf6565b610b55565b6103e2600182611e37565b9050610379565b506104198989868181106103ff576103ff611d71565b90506020020160208101906104149190611a8b565b610cb7565b5050505060018161042a9190611e37565b9050610341565b505050505050565b60606000806104488585610d7c565b90506000610455826106b0565b63ffffffff9092166000908152600160205260409020549193505060ff1690505b9250929050565b63ffffffff808316600090815260026020526040812090916104a39082908590610d9f16565b949350505050565b6104b3610ad4565b8281811461051d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f216c656e677468000000000000000000000000000000000000000000000000006044820152606401610335565b60005b818110156104315761057986868381811061053d5761053d611d71565b90506020020160208101906105529190611a8b565b85858481811061056457610564611d71565b90506020020160208101906101a59190611e4f565b610584600182611e37565b9050610520565b610593610ad4565b60008160ff161180156105b157506105aa826102a3565b8160ff1611155b610617576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600660248201527f2172616e676500000000000000000000000000000000000000000000000000006044820152606401610335565b63ffffffff821660008181526001602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff861690811790915591519182527ff25cfff98c95cf069df801752174d854732576e4b283bc4299386f65676e386a910160405180910390a261069782610cb7565b505050565b6106a4610ad4565b6106ae6000610dd1565b565b63ffffffff811660009081526002602052604081206060916106d182610aca565b905060008167ffffffffffffffff8111156106ee576106ee611e99565b604051908082528060200260200182016040528015610717578160200160208202803683370190505b50905060005b828110156107785761072f8482610e46565b82828151811061074157610741611d71565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101528061077081611ec8565b91505061071d565b50949350505050565b610789610ad4565b63ffffffff80831660009081526002602052604090206107ab918390610e5216565b610811576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f21656e726f6c6c656400000000000000000000000000000000000000000000006044820152606401610335565b600061081c836102a3565b63ffffffff841660009081526001602052604090205490915060ff168110156108a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f76696f6c617465732071756f72756d207468726573686f6c64000000000000006044820152606401610335565b6108aa83610cb7565b508173ffffffffffffffffffffffffffffffffffffffff168363ffffffff167fa4c7a7b783c9afd72ed0b93a7e67ca063acdaa9c3b3268bd43abe1199bdad27c836040516108fa91815260200190565b60405180910390a3505050565b61090f610ad4565b73ffffffffffffffffffffffffffffffffffffffff81166109b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610335565b6109bb81610dd1565b50565b60006109cc85858585610e74565b610a32576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f216d65726b6c65000000000000000000000000000000000000000000000000006044820152606401610335565b610a3e85858585610eee565b610aa4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600560248201527f21736967730000000000000000000000000000000000000000000000000000006044820152606401610335565b506001949350505050565b610ab7610ad4565b610ac18282610b55565b61069782610cb7565b60006102c1825490565b60005473ffffffffffffffffffffffffffffffffffffffff1633146106ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610335565b73ffffffffffffffffffffffffffffffffffffffff8116610bd2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f7a65726f206164647265737300000000000000000000000000000000000000006044820152606401610335565b63ffffffff8083166000908152600260205260409020610bf491839061116616565b610c5a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f616c726561647920656e726f6c6c6564000000000000000000000000000000006044820152606401610335565b8073ffffffffffffffffffffffffffffffffffffffff168263ffffffff167f890997ec79a2b993cbc6e69433ed0ffa6303de16e3c21e315ba91c21ab5ec15a610ca2856102a3565b60405190815260200160405180910390a35050565b600080610cc3836106b0565b63ffffffff8416600090815260016020908152604080832054905193945060ff1692610cf3918491869101611f00565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152828252805160209182012063ffffffff8916600081815260038452849020829055845290830181905292507f6f17fabbcaf1b0e8cfbc153d8884a27de46c3076c2de8ac2aef0094a134909e3910160405180910390a1949350505050565b6000610d8c600960058486611f78565b610d9591611fa2565b60e01c9392505050565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260018301602052604081205415155b9392505050565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000610dca8383611188565b6000610dca8373ffffffffffffffffffffffffffffffffffffffff84166111b2565b600080610ed8610eb985858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506112a592505050565b610ec388886112b0565b610ecd87876112d4565b63ffffffff166112e4565b9050610ee48686611395565b1495945050505050565b600080610efb86866113ad565b9050600080610f0a8686610d7c565b9050600083610f198a8a6113d2565b604051602001610f2b93929190611fea565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012063ffffffff8516600090815260039093529120549091508114610fe0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f21636f6d6d69746d656e740000000000000000000000000000000000000000006044820152606401610335565b61100782610fee8b8b6113f7565b610ff88c8c611395565b6110028d8d611407565b611417565b92505050600061101788886114cb565b90506000805b8460ff16811015611156576000611073856110398d8d866114ed565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061153392505050565b90505b83831080156110bb575061108b8b8b8561154f565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b156110d0576110c983611ec8565b9250611076565b838310611139576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f217468726573686f6c64000000000000000000000000000000000000000000006044820152606401610335565b61114283611ec8565b9250508061114f90611ec8565b905061101d565b5060019998505050505050505050565b6000610dca8373ffffffffffffffffffffffffffffffffffffffff84166115af565b600082600001828154811061119f5761119f611d71565b9060005260206000200154905092915050565b6000818152600183016020526040812054801561129b5760006111d6600183612029565b85549091506000906111ea90600190612029565b905081811461124f57600086600001828154811061120a5761120a611d71565b906000526020600020015490508087600001848154811061122d5761122d611d71565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061126057611260612040565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506102c1565b60009150506102c1565b805160209091012090565b6112b8611a53565b6112c761044460448486611f78565b810190610dca919061206f565b6000610d8c600560018486611f78565b8260005b602081101561138d57600183821c16600085836020811061130b5761130b611d71565b602002015190508160010361134b576040805160208101839052908101859052606001604051602081830303815290604052805190602001209350611378565b60408051602081018690529081018290526060016040516020818303038152906040528051906020012093505b5050808061138590611ec8565b9150506112e8565b509392505050565b60006113a46020828486611f78565b610dca91612115565b60006113bf6104456104448486611f78565b6113c891612151565b60f81c9392505050565b36600083836113e186866115fe565b6113ec928290611f78565b915091509250929050565b60006113a4604460248486611f78565b6000610d8c602460208486611f78565b6000806114248686611625565b60408051602080820184905281830188905260e087901b7fffffffff00000000000000000000000000000000000000000000000000000000166060830152825160448184030181526064830184528051908201207f19457468657265756d205369676e6564204d6573736167653a0a333200000000608484015260a0808401919091528351808403909101815260c090920190925280519101209091509695505050505050565b600060206114d984846115fe565b6114e39084612029565b610dca9190612197565b366000806114fc6041856121d2565b61150890610445611e37565b90506000611517604183611e37565b90506115258183888a611f78565b935093505050935093915050565b600080600061154285856116a5565b9150915061138d816116e7565b60008061155d8360206121d2565b61156786866115fe565b6115719190611e37565b61157c90600c611e37565b9050600061158b826014611e37565b905061159981838789611f78565b6115a29161220f565b60601c9695505050505050565b60008181526001830160205260408120546115f6575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556102c1565b5060006102c1565b6000604161160c84846113ad565b60ff1661161991906121d2565b610dca90610445611e37565b6040805160e084901b7fffffffff0000000000000000000000000000000000000000000000000000000016602080830191909152602482018490527f48595045524c414e45000000000000000000000000000000000000000000000060448301528251602d818403018152604d9092019092528051910120600090610dca565b60008082516041036116db5760208301516040840151606085015160001a6116cf8782858561193b565b94509450505050610476565b50600090506002610476565b60008160048111156116fb576116fb611e6a565b036117035750565b600181600481111561171757611717611e6a565b0361177e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610335565b600281600481111561179257611792611e6a565b036117f9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610335565b600381600481111561180d5761180d611e6a565b0361189a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610335565b60048160048111156118ae576118ae611e6a565b036109bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610335565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156119725750600090506003611a4a565b8460ff16601b1415801561198a57508460ff16601c14155b1561199b5750600090506004611a4a565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156119ef573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116611a4357600060019250925050611a4a565b9150600090505b94509492505050565b6040518061040001604052806020906020820280368337509192915050565b803563ffffffff81168114611a8657600080fd5b919050565b600060208284031215611a9d57600080fd5b610dca82611a72565b60008083601f840112611ab857600080fd5b50813567ffffffffffffffff811115611ad057600080fd5b6020830191508360208260051b850101111561047657600080fd5b60008060008060408587031215611b0157600080fd5b843567ffffffffffffffff80821115611b1957600080fd5b611b2588838901611aa6565b90965094506020870135915080821115611b3e57600080fd5b50611b4b87828801611aa6565b95989497509550505050565b60008083601f840112611b6957600080fd5b50813567ffffffffffffffff811115611b8157600080fd5b60208301915083602082850101111561047657600080fd5b60008060208385031215611bac57600080fd5b823567ffffffffffffffff811115611bc357600080fd5b611bcf85828601611b57565b90969095509350505050565b600081518084526020808501945080840160005b83811015611c2157815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101611bef565b509495945050505050565b604081526000611c3f6040830185611bdb565b905060ff831660208301529392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611a8657600080fd5b60008060408385031215611c8857600080fd5b611c9183611a72565b9150611c9f60208401611c51565b90509250929050565b803560ff81168114611a8657600080fd5b60008060408385031215611ccc57600080fd5b611cd583611a72565b9150611c9f60208401611ca8565b602081526000610dca6020830184611bdb565b600060208284031215611d0857600080fd5b610dca82611c51565b60008060008060408587031215611d2757600080fd5b843567ffffffffffffffff80821115611d3f57600080fd5b611d4b88838901611b57565b90965094506020870135915080821115611d6457600080fd5b50611b4b87828801611b57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611dd557600080fd5b83018035915067ffffffffffffffff821115611df057600080fd5b6020019150600581901b360382131561047657600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115611e4a57611e4a611e08565b500190565b600060208284031215611e6157600080fd5b610dca82611ca8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611ef957611ef9611e08565b5060010190565b7fff000000000000000000000000000000000000000000000000000000000000008360f81b168152600060018083018451602080870160005b83811015611f6a57815173ffffffffffffffffffffffffffffffffffffffff16855293820193908201908501611f39565b509298975050505050505050565b60008085851115611f8857600080fd5b83861115611f9557600080fd5b5050820193919092039150565b7fffffffff000000000000000000000000000000000000000000000000000000008135818116916004851015611fe25780818660040360031b1b83161692505b505092915050565b7fff000000000000000000000000000000000000000000000000000000000000008460f81b168152818360018301376000910160010190815292915050565b60008282101561203b5761203b611e08565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600061040080838503121561208357600080fd5b83601f84011261209257600080fd5b60405181810181811067ffffffffffffffff821117156120db577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040529083019080858311156120f057600080fd5b845b8381101561210a5780358252602091820191016120f2565b509095945050505050565b803560208310156102c1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602084900360031b1b1692915050565b7fff000000000000000000000000000000000000000000000000000000000000008135818116916001851015611fe25760019490940360031b84901b1690921692915050565b6000826121cd577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561220a5761220a611e08565b500290565b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008135818116916014851015611fe25760149490940360031b84901b169092169291505056fea164736f6c634300080f000a",
}

// LegacyMultisigIsmABI is the input ABI used to generate the binding from.
// Deprecated: Use LegacyMultisigIsmMetaData.ABI instead.
var LegacyMultisigIsmABI = LegacyMultisigIsmMetaData.ABI

// LegacyMultisigIsmBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LegacyMultisigIsmMetaData.Bin instead.
var LegacyMultisigIsmBin = LegacyMultisigIsmMetaData.Bin

// DeployLegacyMultisigIsm deploys a new Ethereum contract, binding an instance of LegacyMultisigIsm to it.
func DeployLegacyMultisigIsm(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LegacyMultisigIsm, error) {
	parsed, err := LegacyMultisigIsmMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LegacyMultisigIsmBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LegacyMultisigIsm{LegacyMultisigIsmCaller: LegacyMultisigIsmCaller{contract: contract}, LegacyMultisigIsmTransactor: LegacyMultisigIsmTransactor{contract: contract}, LegacyMultisigIsmFilterer: LegacyMultisigIsmFilterer{contract: contract}}, nil
}

// LegacyMultisigIsm is an auto generated Go binding around an Ethereum contract.
type LegacyMultisigIsm struct {
	LegacyMultisigIsmCaller     // Read-only binding to the contract
	LegacyMultisigIsmTransactor // Write-only binding to the contract
	LegacyMultisigIsmFilterer   // Log filterer for contract events
}

// LegacyMultisigIsmCaller is an auto generated read-only Go binding around an Ethereum contract.
type LegacyMultisigIsmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyMultisigIsmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LegacyMultisigIsmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyMultisigIsmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LegacyMultisigIsmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyMultisigIsmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LegacyMultisigIsmSession struct {
	Contract     *LegacyMultisigIsm // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LegacyMultisigIsmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LegacyMultisigIsmCallerSession struct {
	Contract *LegacyMultisigIsmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LegacyMultisigIsmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LegacyMultisigIsmTransactorSession struct {
	Contract     *LegacyMultisigIsmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LegacyMultisigIsmRaw is an auto generated low-level Go binding around an Ethereum contract.
type LegacyMultisigIsmRaw struct {
	Contract *LegacyMultisigIsm // Generic contract binding to access the raw methods on
}

// LegacyMultisigIsmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LegacyMultisigIsmCallerRaw struct {
	Contract *LegacyMultisigIsmCaller // Generic read-only contract binding to access the raw methods on
}

// LegacyMultisigIsmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LegacyMultisigIsmTransactorRaw struct {
	Contract *LegacyMultisigIsmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLegacyMultisigIsm creates a new instance of LegacyMultisigIsm, bound to a specific deployed contract.
func NewLegacyMultisigIsm(address common.Address, backend bind.ContractBackend) (*LegacyMultisigIsm, error) {
	contract, err := bindLegacyMultisigIsm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigIsm{LegacyMultisigIsmCaller: LegacyMultisigIsmCaller{contract: contract}, LegacyMultisigIsmTransactor: LegacyMultisigIsmTransactor{contract: contract}, LegacyMultisigIsmFilterer: LegacyMultisigIsmFilterer{contract: contract}}, nil
}

// NewLegacyMultisigIsmCaller creates a new read-only instance of LegacyMultisigIsm, bound to a specific deployed contract.
func NewLegacyMultisigIsmCaller(address common.Address, caller bind.ContractCaller) (*LegacyMultisigIsmCaller, error) {
	contract, err := bindLegacyMultisigIsm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigIsmCaller{contract: contract}, nil
}

// NewLegacyMultisigIsmTransactor creates a new write-only instance of LegacyMultisigIsm, bound to a specific deployed contract.
func NewLegacyMultisigIsmTransactor(address common.Address, transactor bind.ContractTransactor) (*LegacyMultisigIsmTransactor, error) {
	contract, err := bindLegacyMultisigIsm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigIsmTransactor{contract: contract}, nil
}

// NewLegacyMultisigIsmFilterer creates a new log filterer instance of LegacyMultisigIsm, bound to a specific deployed contract.
func NewLegacyMultisigIsmFilterer(address common.Address, filterer bind.ContractFilterer) (*LegacyMultisigIsmFilterer, error) {
	contract, err := bindLegacyMultisigIsm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigIsmFilterer{contract: contract}, nil
}

// bindLegacyMultisigIsm binds a generic wrapper to an already deployed contract.
func bindLegacyMultisigIsm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LegacyMultisigIsmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LegacyMultisigIsm *LegacyMultisigIsmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LegacyMultisigIsm.Contract.LegacyMultisigIsmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LegacyMultisigIsm *LegacyMultisigIsmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.LegacyMultisigIsmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LegacyMultisigIsm *LegacyMultisigIsmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.LegacyMultisigIsmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LegacyMultisigIsm *LegacyMultisigIsmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LegacyMultisigIsm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LegacyMultisigIsm *LegacyMultisigIsmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LegacyMultisigIsm *LegacyMultisigIsmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.contract.Transact(opts, method, params...)
}

// Commitment is a free data retrieval call binding the contract method 0x5d94f2b9.
//
// Solidity: function commitment(uint32 ) view returns(bytes32)
func (_LegacyMultisigIsm *LegacyMultisigIsmCaller) Commitment(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _LegacyMultisigIsm.contract.Call(opts, &out, "commitment", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Commitment is a free data retrieval call binding the contract method 0x5d94f2b9.
//
// Solidity: function commitment(uint32 ) view returns(bytes32)
func (_LegacyMultisigIsm *LegacyMultisigIsmSession) Commitment(arg0 uint32) ([32]byte, error) {
	return _LegacyMultisigIsm.Contract.Commitment(&_LegacyMultisigIsm.CallOpts, arg0)
}

// Commitment is a free data retrieval call binding the contract method 0x5d94f2b9.
//
// Solidity: function commitment(uint32 ) view returns(bytes32)
func (_LegacyMultisigIsm *LegacyMultisigIsmCallerSession) Commitment(arg0 uint32) ([32]byte, error) {
	return _LegacyMultisigIsm.Contract.Commitment(&_LegacyMultisigIsm.CallOpts, arg0)
}

// IsEnrolled is a free data retrieval call binding the contract method 0x40989912.
//
// Solidity: function isEnrolled(uint32 _domain, address _address) view returns(bool)
func (_LegacyMultisigIsm *LegacyMultisigIsmCaller) IsEnrolled(opts *bind.CallOpts, _domain uint32, _address common.Address) (bool, error) {
	var out []interface{}
	err := _LegacyMultisigIsm.contract.Call(opts, &out, "isEnrolled", _domain, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEnrolled is a free data retrieval call binding the contract method 0x40989912.
//
// Solidity: function isEnrolled(uint32 _domain, address _address) view returns(bool)
func (_LegacyMultisigIsm *LegacyMultisigIsmSession) IsEnrolled(_domain uint32, _address common.Address) (bool, error) {
	return _LegacyMultisigIsm.Contract.IsEnrolled(&_LegacyMultisigIsm.CallOpts, _domain, _address)
}

// IsEnrolled is a free data retrieval call binding the contract method 0x40989912.
//
// Solidity: function isEnrolled(uint32 _domain, address _address) view returns(bool)
func (_LegacyMultisigIsm *LegacyMultisigIsmCallerSession) IsEnrolled(_domain uint32, _address common.Address) (bool, error) {
	return _LegacyMultisigIsm.Contract.IsEnrolled(&_LegacyMultisigIsm.CallOpts, _domain, _address)
}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_LegacyMultisigIsm *LegacyMultisigIsmCaller) ModuleType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LegacyMultisigIsm.contract.Call(opts, &out, "moduleType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_LegacyMultisigIsm *LegacyMultisigIsmSession) ModuleType() (uint8, error) {
	return _LegacyMultisigIsm.Contract.ModuleType(&_LegacyMultisigIsm.CallOpts)
}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_LegacyMultisigIsm *LegacyMultisigIsmCallerSession) ModuleType() (uint8, error) {
	return _LegacyMultisigIsm.Contract.ModuleType(&_LegacyMultisigIsm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LegacyMultisigIsm *LegacyMultisigIsmCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LegacyMultisigIsm.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LegacyMultisigIsm *LegacyMultisigIsmSession) Owner() (common.Address, error) {
	return _LegacyMultisigIsm.Contract.Owner(&_LegacyMultisigIsm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LegacyMultisigIsm *LegacyMultisigIsmCallerSession) Owner() (common.Address, error) {
	return _LegacyMultisigIsm.Contract.Owner(&_LegacyMultisigIsm.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0xf211b732.
//
// Solidity: function threshold(uint32 ) view returns(uint8)
func (_LegacyMultisigIsm *LegacyMultisigIsmCaller) Threshold(opts *bind.CallOpts, arg0 uint32) (uint8, error) {
	var out []interface{}
	err := _LegacyMultisigIsm.contract.Call(opts, &out, "threshold", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0xf211b732.
//
// Solidity: function threshold(uint32 ) view returns(uint8)
func (_LegacyMultisigIsm *LegacyMultisigIsmSession) Threshold(arg0 uint32) (uint8, error) {
	return _LegacyMultisigIsm.Contract.Threshold(&_LegacyMultisigIsm.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0xf211b732.
//
// Solidity: function threshold(uint32 ) view returns(uint8)
func (_LegacyMultisigIsm *LegacyMultisigIsmCallerSession) Threshold(arg0 uint32) (uint8, error) {
	return _LegacyMultisigIsm.Contract.Threshold(&_LegacyMultisigIsm.CallOpts, arg0)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x020addf1.
//
// Solidity: function validatorCount(uint32 _domain) view returns(uint256)
func (_LegacyMultisigIsm *LegacyMultisigIsmCaller) ValidatorCount(opts *bind.CallOpts, _domain uint32) (*big.Int, error) {
	var out []interface{}
	err := _LegacyMultisigIsm.contract.Call(opts, &out, "validatorCount", _domain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCount is a free data retrieval call binding the contract method 0x020addf1.
//
// Solidity: function validatorCount(uint32 _domain) view returns(uint256)
func (_LegacyMultisigIsm *LegacyMultisigIsmSession) ValidatorCount(_domain uint32) (*big.Int, error) {
	return _LegacyMultisigIsm.Contract.ValidatorCount(&_LegacyMultisigIsm.CallOpts, _domain)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x020addf1.
//
// Solidity: function validatorCount(uint32 _domain) view returns(uint256)
func (_LegacyMultisigIsm *LegacyMultisigIsmCallerSession) ValidatorCount(_domain uint32) (*big.Int, error) {
	return _LegacyMultisigIsm.Contract.ValidatorCount(&_LegacyMultisigIsm.CallOpts, _domain)
}

// Validators is a free data retrieval call binding the contract method 0x9a6e9a24.
//
// Solidity: function validators(uint32 _domain) view returns(address[])
func (_LegacyMultisigIsm *LegacyMultisigIsmCaller) Validators(opts *bind.CallOpts, _domain uint32) ([]common.Address, error) {
	var out []interface{}
	err := _LegacyMultisigIsm.contract.Call(opts, &out, "validators", _domain)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0x9a6e9a24.
//
// Solidity: function validators(uint32 _domain) view returns(address[])
func (_LegacyMultisigIsm *LegacyMultisigIsmSession) Validators(_domain uint32) ([]common.Address, error) {
	return _LegacyMultisigIsm.Contract.Validators(&_LegacyMultisigIsm.CallOpts, _domain)
}

// Validators is a free data retrieval call binding the contract method 0x9a6e9a24.
//
// Solidity: function validators(uint32 _domain) view returns(address[])
func (_LegacyMultisigIsm *LegacyMultisigIsmCallerSession) Validators(_domain uint32) ([]common.Address, error) {
	return _LegacyMultisigIsm.Contract.Validators(&_LegacyMultisigIsm.CallOpts, _domain)
}

// ValidatorsAndThreshold is a free data retrieval call binding the contract method 0x2e0ed234.
//
// Solidity: function validatorsAndThreshold(bytes _message) view returns(address[], uint8)
func (_LegacyMultisigIsm *LegacyMultisigIsmCaller) ValidatorsAndThreshold(opts *bind.CallOpts, _message []byte) ([]common.Address, uint8, error) {
	var out []interface{}
	err := _LegacyMultisigIsm.contract.Call(opts, &out, "validatorsAndThreshold", _message)

	if err != nil {
		return *new([]common.Address), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// ValidatorsAndThreshold is a free data retrieval call binding the contract method 0x2e0ed234.
//
// Solidity: function validatorsAndThreshold(bytes _message) view returns(address[], uint8)
func (_LegacyMultisigIsm *LegacyMultisigIsmSession) ValidatorsAndThreshold(_message []byte) ([]common.Address, uint8, error) {
	return _LegacyMultisigIsm.Contract.ValidatorsAndThreshold(&_LegacyMultisigIsm.CallOpts, _message)
}

// ValidatorsAndThreshold is a free data retrieval call binding the contract method 0x2e0ed234.
//
// Solidity: function validatorsAndThreshold(bytes _message) view returns(address[], uint8)
func (_LegacyMultisigIsm *LegacyMultisigIsmCallerSession) ValidatorsAndThreshold(_message []byte) ([]common.Address, uint8, error) {
	return _LegacyMultisigIsm.Contract.ValidatorsAndThreshold(&_LegacyMultisigIsm.CallOpts, _message)
}

// Verify is a free data retrieval call binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes _metadata, bytes _message) view returns(bool)
func (_LegacyMultisigIsm *LegacyMultisigIsmCaller) Verify(opts *bind.CallOpts, _metadata []byte, _message []byte) (bool, error) {
	var out []interface{}
	err := _LegacyMultisigIsm.contract.Call(opts, &out, "verify", _metadata, _message)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes _metadata, bytes _message) view returns(bool)
func (_LegacyMultisigIsm *LegacyMultisigIsmSession) Verify(_metadata []byte, _message []byte) (bool, error) {
	return _LegacyMultisigIsm.Contract.Verify(&_LegacyMultisigIsm.CallOpts, _metadata, _message)
}

// Verify is a free data retrieval call binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes _metadata, bytes _message) view returns(bool)
func (_LegacyMultisigIsm *LegacyMultisigIsmCallerSession) Verify(_metadata []byte, _message []byte) (bool, error) {
	return _LegacyMultisigIsm.Contract.Verify(&_LegacyMultisigIsm.CallOpts, _metadata, _message)
}

// EnrollValidator is a paid mutator transaction binding the contract method 0xfe9f5ab1.
//
// Solidity: function enrollValidator(uint32 _domain, address _validator) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmTransactor) EnrollValidator(opts *bind.TransactOpts, _domain uint32, _validator common.Address) (*types.Transaction, error) {
	return _LegacyMultisigIsm.contract.Transact(opts, "enrollValidator", _domain, _validator)
}

// EnrollValidator is a paid mutator transaction binding the contract method 0xfe9f5ab1.
//
// Solidity: function enrollValidator(uint32 _domain, address _validator) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmSession) EnrollValidator(_domain uint32, _validator common.Address) (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.EnrollValidator(&_LegacyMultisigIsm.TransactOpts, _domain, _validator)
}

// EnrollValidator is a paid mutator transaction binding the contract method 0xfe9f5ab1.
//
// Solidity: function enrollValidator(uint32 _domain, address _validator) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmTransactorSession) EnrollValidator(_domain uint32, _validator common.Address) (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.EnrollValidator(&_LegacyMultisigIsm.TransactOpts, _domain, _validator)
}

// EnrollValidators is a paid mutator transaction binding the contract method 0x174769dc.
//
// Solidity: function enrollValidators(uint32[] _domains, address[][] _validators) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmTransactor) EnrollValidators(opts *bind.TransactOpts, _domains []uint32, _validators [][]common.Address) (*types.Transaction, error) {
	return _LegacyMultisigIsm.contract.Transact(opts, "enrollValidators", _domains, _validators)
}

// EnrollValidators is a paid mutator transaction binding the contract method 0x174769dc.
//
// Solidity: function enrollValidators(uint32[] _domains, address[][] _validators) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmSession) EnrollValidators(_domains []uint32, _validators [][]common.Address) (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.EnrollValidators(&_LegacyMultisigIsm.TransactOpts, _domains, _validators)
}

// EnrollValidators is a paid mutator transaction binding the contract method 0x174769dc.
//
// Solidity: function enrollValidators(uint32[] _domains, address[][] _validators) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmTransactorSession) EnrollValidators(_domains []uint32, _validators [][]common.Address) (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.EnrollValidators(&_LegacyMultisigIsm.TransactOpts, _domains, _validators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyMultisigIsm.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmSession) RenounceOwnership() (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.RenounceOwnership(&_LegacyMultisigIsm.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.RenounceOwnership(&_LegacyMultisigIsm.TransactOpts)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x4422cc2a.
//
// Solidity: function setThreshold(uint32 _domain, uint8 _threshold) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmTransactor) SetThreshold(opts *bind.TransactOpts, _domain uint32, _threshold uint8) (*types.Transaction, error) {
	return _LegacyMultisigIsm.contract.Transact(opts, "setThreshold", _domain, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x4422cc2a.
//
// Solidity: function setThreshold(uint32 _domain, uint8 _threshold) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmSession) SetThreshold(_domain uint32, _threshold uint8) (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.SetThreshold(&_LegacyMultisigIsm.TransactOpts, _domain, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x4422cc2a.
//
// Solidity: function setThreshold(uint32 _domain, uint8 _threshold) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmTransactorSession) SetThreshold(_domain uint32, _threshold uint8) (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.SetThreshold(&_LegacyMultisigIsm.TransactOpts, _domain, _threshold)
}

// SetThresholds is a paid mutator transaction binding the contract method 0x414c676f.
//
// Solidity: function setThresholds(uint32[] _domains, uint8[] _thresholds) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmTransactor) SetThresholds(opts *bind.TransactOpts, _domains []uint32, _thresholds []uint8) (*types.Transaction, error) {
	return _LegacyMultisigIsm.contract.Transact(opts, "setThresholds", _domains, _thresholds)
}

// SetThresholds is a paid mutator transaction binding the contract method 0x414c676f.
//
// Solidity: function setThresholds(uint32[] _domains, uint8[] _thresholds) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmSession) SetThresholds(_domains []uint32, _thresholds []uint8) (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.SetThresholds(&_LegacyMultisigIsm.TransactOpts, _domains, _thresholds)
}

// SetThresholds is a paid mutator transaction binding the contract method 0x414c676f.
//
// Solidity: function setThresholds(uint32[] _domains, uint8[] _thresholds) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmTransactorSession) SetThresholds(_domains []uint32, _thresholds []uint8) (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.SetThresholds(&_LegacyMultisigIsm.TransactOpts, _domains, _thresholds)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LegacyMultisigIsm.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.TransferOwnership(&_LegacyMultisigIsm.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.TransferOwnership(&_LegacyMultisigIsm.TransactOpts, newOwner)
}

// UnenrollValidator is a paid mutator transaction binding the contract method 0xe7d5d3a1.
//
// Solidity: function unenrollValidator(uint32 _domain, address _validator) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmTransactor) UnenrollValidator(opts *bind.TransactOpts, _domain uint32, _validator common.Address) (*types.Transaction, error) {
	return _LegacyMultisigIsm.contract.Transact(opts, "unenrollValidator", _domain, _validator)
}

// UnenrollValidator is a paid mutator transaction binding the contract method 0xe7d5d3a1.
//
// Solidity: function unenrollValidator(uint32 _domain, address _validator) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmSession) UnenrollValidator(_domain uint32, _validator common.Address) (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.UnenrollValidator(&_LegacyMultisigIsm.TransactOpts, _domain, _validator)
}

// UnenrollValidator is a paid mutator transaction binding the contract method 0xe7d5d3a1.
//
// Solidity: function unenrollValidator(uint32 _domain, address _validator) returns()
func (_LegacyMultisigIsm *LegacyMultisigIsmTransactorSession) UnenrollValidator(_domain uint32, _validator common.Address) (*types.Transaction, error) {
	return _LegacyMultisigIsm.Contract.UnenrollValidator(&_LegacyMultisigIsm.TransactOpts, _domain, _validator)
}

// LegacyMultisigIsmCommitmentUpdatedIterator is returned from FilterCommitmentUpdated and is used to iterate over the raw logs and unpacked data for CommitmentUpdated events raised by the LegacyMultisigIsm contract.
type LegacyMultisigIsmCommitmentUpdatedIterator struct {
	Event *LegacyMultisigIsmCommitmentUpdated // Event containing the contract specifics and raw log

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
func (it *LegacyMultisigIsmCommitmentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyMultisigIsmCommitmentUpdated)
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
		it.Event = new(LegacyMultisigIsmCommitmentUpdated)
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
func (it *LegacyMultisigIsmCommitmentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyMultisigIsmCommitmentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyMultisigIsmCommitmentUpdated represents a CommitmentUpdated event raised by the LegacyMultisigIsm contract.
type LegacyMultisigIsmCommitmentUpdated struct {
	Domain     uint32
	Commitment [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommitmentUpdated is a free log retrieval operation binding the contract event 0x6f17fabbcaf1b0e8cfbc153d8884a27de46c3076c2de8ac2aef0094a134909e3.
//
// Solidity: event CommitmentUpdated(uint32 domain, bytes32 commitment)
func (_LegacyMultisigIsm *LegacyMultisigIsmFilterer) FilterCommitmentUpdated(opts *bind.FilterOpts) (*LegacyMultisigIsmCommitmentUpdatedIterator, error) {

	logs, sub, err := _LegacyMultisigIsm.contract.FilterLogs(opts, "CommitmentUpdated")
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigIsmCommitmentUpdatedIterator{contract: _LegacyMultisigIsm.contract, event: "CommitmentUpdated", logs: logs, sub: sub}, nil
}

// WatchCommitmentUpdated is a free log subscription operation binding the contract event 0x6f17fabbcaf1b0e8cfbc153d8884a27de46c3076c2de8ac2aef0094a134909e3.
//
// Solidity: event CommitmentUpdated(uint32 domain, bytes32 commitment)
func (_LegacyMultisigIsm *LegacyMultisigIsmFilterer) WatchCommitmentUpdated(opts *bind.WatchOpts, sink chan<- *LegacyMultisigIsmCommitmentUpdated) (event.Subscription, error) {

	logs, sub, err := _LegacyMultisigIsm.contract.WatchLogs(opts, "CommitmentUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyMultisigIsmCommitmentUpdated)
				if err := _LegacyMultisigIsm.contract.UnpackLog(event, "CommitmentUpdated", log); err != nil {
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

// ParseCommitmentUpdated is a log parse operation binding the contract event 0x6f17fabbcaf1b0e8cfbc153d8884a27de46c3076c2de8ac2aef0094a134909e3.
//
// Solidity: event CommitmentUpdated(uint32 domain, bytes32 commitment)
func (_LegacyMultisigIsm *LegacyMultisigIsmFilterer) ParseCommitmentUpdated(log types.Log) (*LegacyMultisigIsmCommitmentUpdated, error) {
	event := new(LegacyMultisigIsmCommitmentUpdated)
	if err := _LegacyMultisigIsm.contract.UnpackLog(event, "CommitmentUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyMultisigIsmOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LegacyMultisigIsm contract.
type LegacyMultisigIsmOwnershipTransferredIterator struct {
	Event *LegacyMultisigIsmOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LegacyMultisigIsmOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyMultisigIsmOwnershipTransferred)
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
		it.Event = new(LegacyMultisigIsmOwnershipTransferred)
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
func (it *LegacyMultisigIsmOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyMultisigIsmOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyMultisigIsmOwnershipTransferred represents a OwnershipTransferred event raised by the LegacyMultisigIsm contract.
type LegacyMultisigIsmOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LegacyMultisigIsm *LegacyMultisigIsmFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LegacyMultisigIsmOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LegacyMultisigIsm.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigIsmOwnershipTransferredIterator{contract: _LegacyMultisigIsm.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LegacyMultisigIsm *LegacyMultisigIsmFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LegacyMultisigIsmOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LegacyMultisigIsm.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyMultisigIsmOwnershipTransferred)
				if err := _LegacyMultisigIsm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LegacyMultisigIsm *LegacyMultisigIsmFilterer) ParseOwnershipTransferred(log types.Log) (*LegacyMultisigIsmOwnershipTransferred, error) {
	event := new(LegacyMultisigIsmOwnershipTransferred)
	if err := _LegacyMultisigIsm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyMultisigIsmThresholdSetIterator is returned from FilterThresholdSet and is used to iterate over the raw logs and unpacked data for ThresholdSet events raised by the LegacyMultisigIsm contract.
type LegacyMultisigIsmThresholdSetIterator struct {
	Event *LegacyMultisigIsmThresholdSet // Event containing the contract specifics and raw log

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
func (it *LegacyMultisigIsmThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyMultisigIsmThresholdSet)
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
		it.Event = new(LegacyMultisigIsmThresholdSet)
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
func (it *LegacyMultisigIsmThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyMultisigIsmThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyMultisigIsmThresholdSet represents a ThresholdSet event raised by the LegacyMultisigIsm contract.
type LegacyMultisigIsmThresholdSet struct {
	Domain    uint32
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdSet is a free log retrieval operation binding the contract event 0xf25cfff98c95cf069df801752174d854732576e4b283bc4299386f65676e386a.
//
// Solidity: event ThresholdSet(uint32 indexed domain, uint8 threshold)
func (_LegacyMultisigIsm *LegacyMultisigIsmFilterer) FilterThresholdSet(opts *bind.FilterOpts, domain []uint32) (*LegacyMultisigIsmThresholdSetIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _LegacyMultisigIsm.contract.FilterLogs(opts, "ThresholdSet", domainRule)
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigIsmThresholdSetIterator{contract: _LegacyMultisigIsm.contract, event: "ThresholdSet", logs: logs, sub: sub}, nil
}

// WatchThresholdSet is a free log subscription operation binding the contract event 0xf25cfff98c95cf069df801752174d854732576e4b283bc4299386f65676e386a.
//
// Solidity: event ThresholdSet(uint32 indexed domain, uint8 threshold)
func (_LegacyMultisigIsm *LegacyMultisigIsmFilterer) WatchThresholdSet(opts *bind.WatchOpts, sink chan<- *LegacyMultisigIsmThresholdSet, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _LegacyMultisigIsm.contract.WatchLogs(opts, "ThresholdSet", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyMultisigIsmThresholdSet)
				if err := _LegacyMultisigIsm.contract.UnpackLog(event, "ThresholdSet", log); err != nil {
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

// ParseThresholdSet is a log parse operation binding the contract event 0xf25cfff98c95cf069df801752174d854732576e4b283bc4299386f65676e386a.
//
// Solidity: event ThresholdSet(uint32 indexed domain, uint8 threshold)
func (_LegacyMultisigIsm *LegacyMultisigIsmFilterer) ParseThresholdSet(log types.Log) (*LegacyMultisigIsmThresholdSet, error) {
	event := new(LegacyMultisigIsmThresholdSet)
	if err := _LegacyMultisigIsm.contract.UnpackLog(event, "ThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyMultisigIsmValidatorEnrolledIterator is returned from FilterValidatorEnrolled and is used to iterate over the raw logs and unpacked data for ValidatorEnrolled events raised by the LegacyMultisigIsm contract.
type LegacyMultisigIsmValidatorEnrolledIterator struct {
	Event *LegacyMultisigIsmValidatorEnrolled // Event containing the contract specifics and raw log

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
func (it *LegacyMultisigIsmValidatorEnrolledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyMultisigIsmValidatorEnrolled)
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
		it.Event = new(LegacyMultisigIsmValidatorEnrolled)
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
func (it *LegacyMultisigIsmValidatorEnrolledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyMultisigIsmValidatorEnrolledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyMultisigIsmValidatorEnrolled represents a ValidatorEnrolled event raised by the LegacyMultisigIsm contract.
type LegacyMultisigIsmValidatorEnrolled struct {
	Domain         uint32
	Validator      common.Address
	ValidatorCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorEnrolled is a free log retrieval operation binding the contract event 0x890997ec79a2b993cbc6e69433ed0ffa6303de16e3c21e315ba91c21ab5ec15a.
//
// Solidity: event ValidatorEnrolled(uint32 indexed domain, address indexed validator, uint256 validatorCount)
func (_LegacyMultisigIsm *LegacyMultisigIsmFilterer) FilterValidatorEnrolled(opts *bind.FilterOpts, domain []uint32, validator []common.Address) (*LegacyMultisigIsmValidatorEnrolledIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LegacyMultisigIsm.contract.FilterLogs(opts, "ValidatorEnrolled", domainRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigIsmValidatorEnrolledIterator{contract: _LegacyMultisigIsm.contract, event: "ValidatorEnrolled", logs: logs, sub: sub}, nil
}

// WatchValidatorEnrolled is a free log subscription operation binding the contract event 0x890997ec79a2b993cbc6e69433ed0ffa6303de16e3c21e315ba91c21ab5ec15a.
//
// Solidity: event ValidatorEnrolled(uint32 indexed domain, address indexed validator, uint256 validatorCount)
func (_LegacyMultisigIsm *LegacyMultisigIsmFilterer) WatchValidatorEnrolled(opts *bind.WatchOpts, sink chan<- *LegacyMultisigIsmValidatorEnrolled, domain []uint32, validator []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LegacyMultisigIsm.contract.WatchLogs(opts, "ValidatorEnrolled", domainRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyMultisigIsmValidatorEnrolled)
				if err := _LegacyMultisigIsm.contract.UnpackLog(event, "ValidatorEnrolled", log); err != nil {
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

// ParseValidatorEnrolled is a log parse operation binding the contract event 0x890997ec79a2b993cbc6e69433ed0ffa6303de16e3c21e315ba91c21ab5ec15a.
//
// Solidity: event ValidatorEnrolled(uint32 indexed domain, address indexed validator, uint256 validatorCount)
func (_LegacyMultisigIsm *LegacyMultisigIsmFilterer) ParseValidatorEnrolled(log types.Log) (*LegacyMultisigIsmValidatorEnrolled, error) {
	event := new(LegacyMultisigIsmValidatorEnrolled)
	if err := _LegacyMultisigIsm.contract.UnpackLog(event, "ValidatorEnrolled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyMultisigIsmValidatorUnenrolledIterator is returned from FilterValidatorUnenrolled and is used to iterate over the raw logs and unpacked data for ValidatorUnenrolled events raised by the LegacyMultisigIsm contract.
type LegacyMultisigIsmValidatorUnenrolledIterator struct {
	Event *LegacyMultisigIsmValidatorUnenrolled // Event containing the contract specifics and raw log

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
func (it *LegacyMultisigIsmValidatorUnenrolledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyMultisigIsmValidatorUnenrolled)
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
		it.Event = new(LegacyMultisigIsmValidatorUnenrolled)
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
func (it *LegacyMultisigIsmValidatorUnenrolledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyMultisigIsmValidatorUnenrolledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyMultisigIsmValidatorUnenrolled represents a ValidatorUnenrolled event raised by the LegacyMultisigIsm contract.
type LegacyMultisigIsmValidatorUnenrolled struct {
	Domain         uint32
	Validator      common.Address
	ValidatorCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorUnenrolled is a free log retrieval operation binding the contract event 0xa4c7a7b783c9afd72ed0b93a7e67ca063acdaa9c3b3268bd43abe1199bdad27c.
//
// Solidity: event ValidatorUnenrolled(uint32 indexed domain, address indexed validator, uint256 validatorCount)
func (_LegacyMultisigIsm *LegacyMultisigIsmFilterer) FilterValidatorUnenrolled(opts *bind.FilterOpts, domain []uint32, validator []common.Address) (*LegacyMultisigIsmValidatorUnenrolledIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LegacyMultisigIsm.contract.FilterLogs(opts, "ValidatorUnenrolled", domainRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigIsmValidatorUnenrolledIterator{contract: _LegacyMultisigIsm.contract, event: "ValidatorUnenrolled", logs: logs, sub: sub}, nil
}

// WatchValidatorUnenrolled is a free log subscription operation binding the contract event 0xa4c7a7b783c9afd72ed0b93a7e67ca063acdaa9c3b3268bd43abe1199bdad27c.
//
// Solidity: event ValidatorUnenrolled(uint32 indexed domain, address indexed validator, uint256 validatorCount)
func (_LegacyMultisigIsm *LegacyMultisigIsmFilterer) WatchValidatorUnenrolled(opts *bind.WatchOpts, sink chan<- *LegacyMultisigIsmValidatorUnenrolled, domain []uint32, validator []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LegacyMultisigIsm.contract.WatchLogs(opts, "ValidatorUnenrolled", domainRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyMultisigIsmValidatorUnenrolled)
				if err := _LegacyMultisigIsm.contract.UnpackLog(event, "ValidatorUnenrolled", log); err != nil {
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

// ParseValidatorUnenrolled is a log parse operation binding the contract event 0xa4c7a7b783c9afd72ed0b93a7e67ca063acdaa9c3b3268bd43abe1199bdad27c.
//
// Solidity: event ValidatorUnenrolled(uint32 indexed domain, address indexed validator, uint256 validatorCount)
func (_LegacyMultisigIsm *LegacyMultisigIsmFilterer) ParseValidatorUnenrolled(log types.Log) (*LegacyMultisigIsmValidatorUnenrolled, error) {
	event := new(LegacyMultisigIsmValidatorUnenrolled)
	if err := _LegacyMultisigIsm.contract.UnpackLog(event, "ValidatorUnenrolled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
