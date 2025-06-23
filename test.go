package main

import (
	"GeneReport_platform/internal/contracts/sharingPlatformContract"
	"GeneReport_platform/internal/contracts/testCon"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func test() {
	var err error
	var dataHashBytes32s [][32]byte
	var tempDataHashBytes32 [32]byte
	userAddressHex := "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"
	copy(tempDataHashBytes32[:], common.Hex2Bytes("7c6cd50ffe512475c5bfb9140f45f6d4df79c5d793d28a8f02bb3a913a707679"))
	dataHashBytes32s = append(dataHashBytes32s, tempDataHashBytes32)

	newGeneSharingAddress, _, receipt, err := sharingPlatformContract.GetContractIns().CreateAllFromThirdParty(sharingPlatformContract.NewAdminTransactor(), common.HexToAddress(userAddressHex), dataHashBytes32s)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(newGeneSharingAddress.Hex())
	fmt.Println(receipt.TransactionHash)
}

func testBytes32() {
	DataHashBytes32 := [32]byte(common.Hex2Bytes("7c6cd50ffe512475c5bfb9140f45f6d4df79c5d793d28a8f02bb3a913a707679"))
	tx, receipt, err := testCon.GetContractIns().StoreData(testCon.NewAdminTransactor(), DataHashBytes32, big.NewInt(12332))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(tx)
	fmt.Println(receipt)
}

func main() {
	test()
	testBytes32()
}
