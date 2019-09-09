package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/contracts/XRC21issuer"
	"github.com/ethereum/go-ethereum/contracts/XRC21issuer/simulation"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

func main() {
	fmt.Println("========================")
	fmt.Println("mainAddr", simulation.MainAddr.Hex())
	fmt.Println("airdropAddr", simulation.AirdropAddr.Hex())
	fmt.Println("receiverAddr", simulation.ReceiverAddr.Hex())
	fmt.Println("========================")
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}
	nonce, _ := client.NonceAt(context.Background(), simulation.MainAddr, nil)

	// init XRC21 issuer
	auth := bind.NewKeyedTransactor(simulation.MainKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(4000000) // in units
	auth.GasPrice = big.NewInt(210000000000000)
	XRC21IssuerAddr, XRC21Issuer, err := XRC21issuer.DeployXRC21Issuer(auth, client, simulation.MinApply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("main address", simulation.MainAddr.Hex(), "nonce", nonce)
	fmt.Println("===> XRC21 issuer address", XRC21IssuerAddr.Hex())

	auth.Nonce = big.NewInt(int64(nonce + 1))

	// init XRC20
	XRC21TokenAddr, XRC21Token, err := XRC21issuer.DeployXRC21(auth, client, "TEST", "TOMO", 18, simulation.Cap, simulation.Fee)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("===>  XRC21 token address", XRC21TokenAddr.Hex(), "cap", simulation.Cap)

	fmt.Println("wait 10s to execute init smart contract")
	time.Sleep(10 * time.Second)

	XRC21Issuer.TransactOpts.Nonce = big.NewInt(int64(nonce + 2))
	XRC21Issuer.TransactOpts.Value = simulation.MinApply
	XRC21Issuer.TransactOpts.GasPrice = big.NewInt(21000)
	XRC21Token.TransactOpts.GasPrice = big.NewInt(21000)
	XRC21Token.TransactOpts.GasLimit = uint64(4000000)
	auth.GasPrice = big.NewInt(21000)
	// get balance init XRC21 smart contract
	balance, err := XRC21Token.BalanceOf(simulation.MainAddr)
	if err != nil || balance.Cmp(simulation.Cap) != 0 {
		log.Fatal(err, "\tget\t", balance, "\twant\t", simulation.Cap)
	}
	fmt.Println("balance", balance, "mainAddr", simulation.MainAddr.Hex())

	// add XRC20 list token XRC21 issuer
	_, err = XRC21Issuer.Apply(XRC21TokenAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("wait 10s to add token to list issuer")
	time.Sleep(10 * time.Second)

	//check XRC21 SMC balance
	balance, err = client.BalanceAt(context.Background(), XRC21IssuerAddr, nil)
	if err != nil || balance.Cmp(simulation.MinApply) != 0 {
		log.Fatal("can't get balance  in XRC21Issuer SMC: ", err, "got", balance, "wanted", simulation.MinApply)
	}

	//check balance fee
	balanceIssuerFee, err := XRC21Issuer.GetTokenCapacity(XRC21TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(simulation.MinApply) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", simulation.MinApply)
	}
}
