package main

import (
	"fmt"
	"log"

	tronWallet "github.com/ranjbar-dev/tron-wallet"
	"github.com/ranjbar-dev/tron-wallet/enums"
)

// https://nile.tronscan.org/
var JST_TEST tronWallet.Token = tronWallet.Token{ContractAddress: "TF17BgPaZYbz8oxbjhriubPDsA7ArKoLX3"}
var WINK_TEST tronWallet.Token = tronWallet.Token{ContractAddress: "TU2T8vpHZhCNY8fXGVaHyeZrKm8s6HEXWe"}

func main() {

	// w := tronWallet.GenerateTronWallet(enums.NILE_NODE)
	// fmt.Println(w.PrivateKey)

	// address: TQg43vM2rNNsyKBVxagtcRFBz1LnWfcUNe
	// load private key from .keys file.
	privateKey := ""
	// https://nileex.io/join/getJoinPage
	w, err := tronWallet.CreateTronWallet(enums.NILE_NODE, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Wallet Address: %s\n", w.AddressBase58)
	// fmt.Println(w.PublicKey)
	// fmt.Println(w.PrivateKey)

	b, err := w.Balance()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance of TRX: %v\n", b)

	tb, err := w.BalanceTRC20(&JST_TEST)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance of JST: %v\n", tb)

	tb, err = w.BalanceTRC20(&WINK_TEST)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance of WINK: %v\n", tb)

	// w2 := tronWallet.GenerateTronWallet(enums.NILE_NODE)
	// fmt.Println(w2.PrivateKey)

	// address: TEKxk75VrMo7aHMHMfETMvnAopxkASoN7h
	// load private key from .keys file.
	privateKey2 := ""
	// https://nileex.io/join/getJoinPage
	w2, err := tronWallet.CreateTronWallet(enums.NILE_NODE, privateKey2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Wallet2 Address: %s\n", w2.AddressBase58)
	// fmt.Println(w.PublicKey)
	// fmt.Println(w.PrivateKey)

	b2, err := w2.Balance()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance of TRX: %v\n", b2)

	txId, err := w.Transfer(w2.AddressBase58, 1000000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sending TX: %s\n", txId)

}
