package main

import (
	"crypto/ecdsa"
	"io/ioutil"

	"github.com/Dev43/crypto-ussd/memory"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type connection struct {
	client          *ethclient.Client
	keyedTransactor *bind.TransactOpts
	memory          *memory.InMemorySessionStore
}

func NewConnection() (*connection, error) {
	key, err := getKey()
	if err != nil {
		return nil, err
	}
	auth := bind.NewKeyedTransactor(key)
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		return nil, err
	}
	mem := memory.NewInMemorySessionStore()
	return &connection{
		client:          client,
		keyedTransactor: auth,
		memory:          mem,
	}, nil
}

func getKey() (*ecdsa.PrivateKey, error) {
	data, err := ioutil.ReadFile("./keystore/key.txt")
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(string(data))
	if err != nil {
		return nil, err
	}

	return (privateKey), nil
}

func (conn *connection) SendMoney(textArray []string, sessionID, phoneNumber, networkCode, text string) (string, error) {
	var msg string
	var err error
	if len(textArray) == 1 {
		msg = `CON Please send the total amount to send`
		return msg, nil
	}

	switch len(textArray) {
	case 1:
		msg = `CON Please write their phone number`
	case 2:
		msg = `CON Please choose what token to send
		1. DAI
		2. ECT
		3. ETH
		`

	case 3:
		msg = `CON Please send the total amount to send`
	case 4:
		msg = `CON Please confirm what you are sending:`
	}
	return msg, err
	// // to send money
	// // first validate all the inputs
	// //
	// addr := common.HexToAddress("0xC28614fEcD3109EFf192DD3cABc7ac9b82C7eD11")
	// telcoInteractor, err := telco.NewTelco(addr, conn.client)
	// if err != nil {
	// 	return "", err
	// }
	// a, _ := telcoInteractor.TelcoCaller.Owner(&bind.CallOpts{})
	// fmt.Println(a.String())
	// telcoInteractor.AuthTransfer(conn.keyedTransactor, addr, "aa", big.NewInt(1), "11", []byte("hi"))
	// return "", nil
}
