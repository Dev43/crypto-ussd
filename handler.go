package main

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"

	"github.com/Dev43/crypto-ussd/telco"

	"github.com/Dev43/crypto-ussd/memory"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type connection struct {
	client          *ethclient.Client
	keyedTransactor *bind.TransactOpts
	memory          *memory.InMemorySessionStore
	telcoInteractor *telco.Telco
}

var telcoAddress = common.HexToAddress("0x6437B6F32F9f4365c93FDE230E5E5E31197e56f8")
var userAddress = common.HexToAddress("0xA9EB6d4624cAbc500208b66DD43F8937FBd32EFD")

var tokenMap = map[string]common.Address{
	"DAI": common.HexToAddress("0xD15Bc67136E3ED8B9AAc9A396468eD3aBD9E5656"),
	"ETH": common.HexToAddress("0xD15Bc67136E3ED8B9AAc9A396468eD3aBD9E5656"),
	"ECT": common.HexToAddress("0xD15Bc67136E3ED8B9AAc9A396468eD3aBD9E5656"),
}

func NewConnection() (*connection, error) {
	key, err := getKey()
	if err != nil {
		return nil, err
	}
	auth := bind.NewKeyedTransactor(key)
	// client, err := ethclient.Dial("https://goerli.infura.io/v3/1eecb15771324b71961a05dc3398ebd4")
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		return nil, err
	}
	mem := memory.NewInMemorySessionStore()

	telcoInteractor, err := telco.NewTelco(telcoAddress, client)
	if err != nil {
		return nil, err
	}
	return &connection{
		client:          client,
		keyedTransactor: auth,
		memory:          mem,
		telcoInteractor: telcoInteractor,
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

var choices = []string{"", "DAI", "ECT", "ETH"}

func (conn *connection) SendMoney(textArray []string, sessionID, phoneNumber, networkCode, text string) (string, error) {
	var msg string
	var err error
	// comes here when 1
	if len(textArray) == 1 {
		msg = `CON Please send the total amount to send`
		return msg, nil
	}
	// comes here when 1*{amount}*{phone}*{Coinchoice#}*{token symbol #}*{confirm = 0 && 1}
	switch len(textArray) {
	case 2:
		// we have an amount
		amount := textArray[1]
		fmt.Println(amount)
		msg = `CON Please write their phone number`
	case 3:
		// we have the phone number here
		number := textArray[2]
		fmt.Println(number)

		msg = `CON Please choose what token to send
		1. DAI
		2. ECT
		3. ETH
		`
	case 4:
		// we get their choice
		index := textArray[3]
		fmt.Println(index)

		// make sure it's within the bounds

		msg = `CON Please send a password` // can add (you have x left)
	case 5:
		// we get the password
		password := textArray[4]
		fmt.Println(password)

		msg = `CON Please confirm what you are sending:`
	case 6:
		// we get the confirmation
		conf := textArray[5]
		fmt.Println(conf)

		// check it's ok
		msg = `END Sent!`
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

func (conn *connection) MyAccount(textArray []string, sessionID, phoneNumber, networkCode, text string) (string, error) {
	var msg string
	var err error
	// comes here when 1
	if len(textArray) == 1 {
		msg = `CON What would you like to access?
		1. My Account Balances
		2. My Account Address
		3. My Phone Number`
		return msg, nil
	}
	// comes here when 7*{choice #}*1
	switch len(textArray) {
	case 2:
		// we have an choice
		choice := textArray[1]
		switch choice {
		case "1":
			/// account balance
			addr, err := conn.telcoInteractor.TelcoCaller.GetFromPhoneNumberToAddress(&bind.CallOpts{}, phoneNumber)
			if err != nil {
				msg = "error getting phone number to address"
				break
			}
			for tokenName, tokenAddress := range tokenMap {
				bal, err := conn.telcoInteractor.TelcoCaller.GetUserBalance(&bind.CallOpts{}, addr, tokenAddress)
				if err != nil {
					msg = "error getting balance to address"
					break
				}
				msg += fmt.Sprintf("%s: %s \n", tokenName, bal.String())
			}

		case "2":
		// address
		addr, err := conn.telcoInteractor.TelcoCaller.GetFromPhoneNumberToAddress(&bind.CallOpts{}, phoneNumber)
		if err != nil {
			msg = "error getting phone number to address"
			break
		}
		msg = fmt.Sprintf("%s %s", "Your address is", addr.String())
		case "3":
			// Phone number
			msg = fmt.Sprintf("%s %s", `ENV You phone number is`, phoneNumber)
		}
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
