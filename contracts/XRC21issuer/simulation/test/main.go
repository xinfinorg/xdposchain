package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/XRC21issuer"
	"github.com/ethereum/go-ethereum/contracts/XRC21issuer/simulation"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

var (
	XRC21TokenAddr = common.HexToAddress("0x80430A33EaB86890a346bCf64F86CFeAC73287f3")
)

func airDropTokenToAccountNoXDC() {
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}
	nonce, _ := client.NonceAt(context.Background(), simulation.MainAddr, nil)
	mainAccount := bind.NewKeyedTransactor(simulation.MainKey)
	mainAccount.Nonce = big.NewInt(int64(nonce))
	mainAccount.Value = big.NewInt(0)      // in wei
	mainAccount.GasLimit = uint64(4000000) // in units
	mainAccount.GasPrice = big.NewInt(21000)
	XRC21Instance, _ := XRC21issuer.NewXRC21(mainAccount, XRC21TokenAddr, client)
	XRC21IssuerInstance, _ := XRC21issuer.NewXRC21Issuer(mainAccount, common.XRC21IssuerSMC, client)
	// air drop token
	remainFee, _ := XRC21IssuerInstance.GetTokenCapacity(XRC21TokenAddr)
	tx, err := XRC21Instance.Transfer(simulation.AirdropAddr, simulation.AirDropAmount)
	if err != nil {
		log.Fatal("can't air drop to ", err)
	}
	// check balance after transferAmount
	fmt.Println("wait 10s to airdrop success ", tx.Hash().Hex())
	time.Sleep(10 * time.Second)

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil || receipt == nil {
		log.Fatal("can't transaction's receipt ", err, "hash", tx.Hash().Hex())
	}
	remainFee = big.NewInt(0).Sub(remainFee, big.NewInt(0).SetUint64(receipt.GasUsed))
	//check balance fee
	balanceIssuerFee, err := XRC21IssuerInstance.GetTokenCapacity(XRC21TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	if err != nil {
		log.Fatal("can't execute transferAmount in tr21:", err)
	}
}
func testTransferXRC21TokenWithAccountNoXDC() {
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}

	// access to address which received token XRC20 but dont have tomo
	nonce, _ := client.NonceAt(context.Background(), simulation.AirdropAddr, nil)
	airDropAccount := bind.NewKeyedTransactor(simulation.AirdropKey)
	airDropAccount.Nonce = big.NewInt(int64(nonce))
	airDropAccount.Value = big.NewInt(0)      // in wei
	airDropAccount.GasLimit = uint64(4000000) // in units
	airDropAccount.GasPrice = big.NewInt(21000)
	XRC21Instance, _ := XRC21issuer.NewXRC21(airDropAccount, XRC21TokenAddr, client)
	XRC21IssuerInstance, _ := XRC21issuer.NewXRC21Issuer(airDropAccount, common.XRC21IssuerSMC, client)

	remainFee, _ := XRC21IssuerInstance.GetTokenCapacity(XRC21TokenAddr)
	airDropBalanceBefore, err := XRC21Instance.BalanceOf(simulation.AirdropAddr)
	receiverBalanceBefore, err := XRC21Instance.BalanceOf(simulation.ReceiverAddr)
	// execute transferAmount XRC to other address
	tx, err := XRC21Instance.Transfer(simulation.ReceiverAddr, simulation.TransferAmount)
	if err != nil {
		log.Fatal("can't execute transferAmount in tr21:", err)
	}

	// check balance after transferAmount
	fmt.Println("wait 10s to transferAmount success ")
	time.Sleep(10 * time.Second)

	balance, err := XRC21Instance.BalanceOf(simulation.ReceiverAddr)
	wantedBalance := big.NewInt(0).Add(receiverBalanceBefore, simulation.TransferAmount)
	if err != nil || balance.Cmp(wantedBalance) != 0 {
		log.Fatal("check balance after fail receiverAmount in tr21: ", err, "get", balance, "wanted", wantedBalance)
	}

	remainAirDrop := big.NewInt(0).Sub(airDropBalanceBefore, simulation.TransferAmount)
	remainAirDrop = remainAirDrop.Sub(remainAirDrop, simulation.Fee)
	// check balance XRC21 again
	balance, err = XRC21Instance.BalanceOf(simulation.AirdropAddr)
	if err != nil || balance.Cmp(remainAirDrop) != 0 {
		log.Fatal("check balance after fail transferAmount in tr21: ", err, "get", balance, "wanted", remainAirDrop)
	}

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal("can't transaction's receipt ", err, "hash", tx.Hash().Hex())
	}
	remainFee = big.NewInt(0).Sub(remainFee, big.NewInt(0).SetUint64(receipt.GasUsed))
	//check balance fee
	balanceIssuerFee, err := XRC21IssuerInstance.GetTokenCapacity(XRC21TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	//check XRC21 SMC balance
	balance, err = client.BalanceAt(context.Background(), common.XRC21IssuerSMC, nil)
	if err != nil || balance.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
}
func testTransferXRC21Fail() {
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}
	nonce, _ := client.NonceAt(context.Background(), simulation.AirdropAddr, nil)
	airDropAccount := bind.NewKeyedTransactor(simulation.AirdropKey)
	airDropAccount.Nonce = big.NewInt(int64(nonce))
	airDropAccount.Value = big.NewInt(0)      // in wei
	airDropAccount.GasLimit = uint64(4000000) // in units
	airDropAccount.GasPrice = big.NewInt(21000)
	XRC21Instance, _ := XRC21issuer.NewXRC21(airDropAccount, XRC21TokenAddr, client)
	XRC21IssuerInstance, _ := XRC21issuer.NewXRC21Issuer(airDropAccount, common.XRC21IssuerSMC, client)
	balanceIssuerFee, err := XRC21IssuerInstance.GetTokenCapacity(XRC21TokenAddr)

	minFee, err := XRC21Instance.MinFee()
	if err != nil {
		log.Fatal("can't get minFee of XRC21 smart contract:", err)
	}
	ownerBalance, err := XRC21Instance.BalanceOf(simulation.MainAddr)
	remainFee, err := XRC21IssuerInstance.GetTokenCapacity(XRC21TokenAddr)
	airDropBalanceBefore, err := XRC21Instance.BalanceOf(simulation.AirdropAddr)

	tx, err := XRC21Instance.Transfer(common.Address{}, big.NewInt(1))
	if err != nil {
		log.Fatal("can't execute test transfer to zero address in tr21:", err)
	}
	fmt.Println("wait 10s to transfer to zero address")
	time.Sleep(10 * time.Second)

	fmt.Println("airDropBalanceBefore", airDropBalanceBefore)
	// check balance XRC21 again
	airDropBalanceBefore = big.NewInt(0).Sub(airDropBalanceBefore, minFee)
	balance, err := XRC21Instance.BalanceOf(simulation.AirdropAddr)
	if err != nil || balance.Cmp(airDropBalanceBefore) != 0 {
		log.Fatal("check balance after fail transferAmount in tr21: ", err, "get", balance, "wanted", airDropBalanceBefore)
	}

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal("can't transaction's receipt ", err, "hash", tx.Hash().Hex())
	}
	ownerBalance = big.NewInt(0).Add(ownerBalance, minFee)
	//check balance fee
	balance, err = XRC21Instance.BalanceOf(simulation.MainAddr)
	if err != nil || balance.Cmp(ownerBalance) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}

	remainFee = big.NewInt(0).Sub(remainFee, big.NewInt(0).SetUint64(receipt.GasUsed))
	//check balance fee
	balanceIssuerFee, err = XRC21IssuerInstance.GetTokenCapacity(XRC21TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	//check XRC21 SMC balance
	balance, err = client.BalanceAt(context.Background(), common.XRC21IssuerSMC, nil)
	if err != nil || balance.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}

}
func main() {
	fmt.Println("========================")
	fmt.Println("airdropAddr", simulation.AirdropAddr.Hex())
	fmt.Println("receiverAddr", simulation.ReceiverAddr.Hex())
	fmt.Println("========================")

	start := time.Now()
	for i := 0; i < 10000000; i++ {
		airDropTokenToAccountNoXDC()
		fmt.Println("Finish airdrop token to a account")
		testTransferXRC21TokenWithAccountNoXDC()
		fmt.Println("Finish transfer XRC21 token with a account no tomo")
		testTransferXRC21Fail()
		fmt.Println("Finish testing ! Success transferAmount token XRC20 with a account no tomo")
	}
	fmt.Println(common.PrettyDuration(time.Since(start)))
}
