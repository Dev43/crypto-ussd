package main

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"

	"github.com/Dev43/crypto-ussd/telco"
	"github.com/Dev43/crypto-ussd/user"

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
	userInteractor  *user.User
}

var telcoAddress = common.HexToAddress("0xf7116a103306f0CF026851e2f78Ef73C4C81b278")
var userAddress = common.HexToAddress("0xf65DF84ee55FCb9e6c628ef4C7a96f7713933Ee0")

var tokenMap = map[string]common.Address{
	"DAI": common.HexToAddress("0x40f54c622ea1c3B91Eb1eEEE89bDb0A96AD336B5"),
	"ETH": common.HexToAddress("0x40f54c622ea1c3B91Eb1eEEE89bDb0A96AD336B5"),
	"ECT": common.HexToAddress("0x40f54c622ea1c3B91Eb1eEEE89bDb0A96AD336B5"),
}

var tokenArray = []string{"DAI", "ETH", "ECT"}

func NewConnection() (*connection, error) {
	key, err := getKey()
	if err != nil {
		return nil, err
	}
	auth := bind.NewKeyedTransactor(key)

	auth.GasLimit = 6000000
	gp, _ := new(big.Int).SetString("10000000000", 10)
	auth.GasPrice = gp
	// client, err := ethclient.Dial("https://goerli.infura.io/v3/1eecb15771324b71961a05dc3398ebd4")
	url := os.Getenv("BLOCKCHAIN_URL")
	if url == "" {
		url = "http://localhost:8545"
	}
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	mem := memory.NewInMemorySessionStore()

	telcoInteractor, err := telco.NewTelco(telcoAddress, client)
	if err != nil {
		return nil, err
	}
	userInteractor, err := user.NewUser(userAddress, client)
	if err != nil {
		return nil, err
	}
	return &connection{
		client:          client,
		keyedTransactor: auth,
		memory:          mem,
		telcoInteractor: telcoInteractor,
		userInteractor:  userInteractor,
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
	transferKey := sessionID + "-transfer"
	// comes here when 1
	if len(textArray) == 1 {
		msg = `CON Please send the total amount to send`
		return msg, nil
	}
	//                 1*  777 * 8989  * 1 * hi * 1
	// comes here when 1*{amount}*{phone}*{Coinchoice#}*{password}*{confirm = 0 && 1}
	switch len(textArray) {
	case 2:
		// we have an amount
		amount := textArray[1]

		a, ok := new(big.Int).SetString(amount, 10)
		if !ok {
			//
		}
		tx := TransferInformation{
			Amount: a,
			State:  "init",
		}
		conn.memory.Write(transferKey, tx)
		msg = `CON Please write their phone number`
	case 3:
		// we have the phone number here
		number := textArray[2]
		fmt.Println(number)

		und, err := conn.memory.Get(transferKey)
		if err != nil {

		}

		tx := und.(TransferInformation)
		tx.ToPhoneNumber = number
		conn.memory.Write(transferKey, tx)

		addr, err := conn.telcoInteractor.TelcoCaller.GetFromPhoneNumberToAddress(&bind.CallOpts{}, phoneNumber)
		if err != nil {
			msg = "error getting phone number to address"
			break
		}
		msg = "CON Please choose what token to send\n"
		for index, tokenName := range tokenArray {
			tokenAddress := tokenMap[tokenName]
			bal, err := conn.telcoInteractor.TelcoCaller.GetUserBalance(&bind.CallOpts{}, addr, tokenAddress)
			if err != nil {
				msg = "error getting balance to address"
				break
			}
			msg += fmt.Sprintf("%d. %s: %s \n", index+1, tokenName, bal.String())
		}

	case 4:
		// we get their choice
		index := textArray[3]
		i, err := strconv.ParseInt(index, 10, 64)
		if len(tokenArray) <= int(i) || int(i) <= 0 {
			msg = "END Wrong token chosen"
			break
		}
		und, err := conn.memory.Get(transferKey)
		if err != nil {

		}
		tx := und.(TransferInformation)
		tokenName := tokenArray[i-1]
		tx.Token = tokenName
		conn.memory.Write(transferKey, tx)

		pLeft, err := conn.userInteractor.UserCaller.TotalPasswordsLeft(&bind.CallOpts{})
		fmt.Println(pLeft.String)
		msg = fmt.Sprintf("%s \n%s %s", `CON Please send a password`, "Passwords left:", pLeft.String()) // can add (you have x left)
	case 5:
		// we get the password
		password := textArray[4]

		und, err := conn.memory.Get(transferKey)
		if err != nil {

		}
		isValid, err := conn.userInteractor.UserCaller.IsPasswordValid(&bind.CallOpts{}, []byte(password))
		if err != nil || !isValid {
			msg = `END Password is not valid`
			break
		}
		tx := und.(TransferInformation)
		tx.Password = password
		conn.memory.Write(transferKey, tx)
		msg = fmt.Sprintf("%s \n %s: %s\n %s: %s\n %s: %s\n %s: %s\n 1.Confirm\n 2.Deny\n",
			`CON Please confirm what you are sending:`,
			"Token", tx.Token,
			"Amount", tx.Amount.String(),
			"To", tx.ToPhoneNumber,
			"Password", tx.Password,
		)
	case 6:
		// we get the confirmation
		conf := textArray[5]
		i, err := strconv.ParseInt(conf, 10, 64)
		if i == 2 {
			msg = "END Operation aborted"
			break
		}
		if i != 1 {
			msg = "END Invalid choice"
			break
		}
		und, err := conn.memory.Get(transferKey)
		if err != nil {

		}
		tx := und.(TransferInformation)
		// actually send the transaction
		fmt.Println(tx)
		bTx, err := conn.telcoInteractor.AuthTransfer(conn.keyedTransactor, tokenMap[tx.Token], tx.ToPhoneNumber, tx.Amount, phoneNumber, []byte(tx.Password))
		if err != nil {
			fmt.Println(err)
			msg = "END Problem sending the transaction"
			break
		}
		tx.State = "pending"
		tx.TxHash = bTx.Hash().String()
		conn.memory.Write(transferKey, tx)

		// create a new user transaction
		c, _ := conn.memory.Get(phoneNumber)
		txHistory := c.(TransactionHistory)
		txHistory.Transactions = append(txHistory.Transactions, tx)
		conn.memory.Write(phoneNumber, txHistory)

		// check it's ok
		msg = `END Sent!`
	}
	return msg, err
}

func (conn *connection) Withdraw(textArray []string, sessionID, phoneNumber, networkCode, text string) (string, error) {
	var msg string
	var err error
	withdrawKey := sessionID + "-withdraw"
	// comes here when 1
	if len(textArray) == 1 {
		w := WithdrawRequest{}
		conn.memory.Write(withdrawKey, w)

		addr, err := conn.telcoInteractor.TelcoCaller.GetFromPhoneNumberToAddress(&bind.CallOpts{}, phoneNumber)
		if err != nil {
			msg = "error getting phone number to address"
		}
		msg = "CON Please choose what token to withdraw\n"
		for index, tokenName := range tokenArray {
			tokenAddress := tokenMap[tokenName]
			bal, err := conn.telcoInteractor.TelcoCaller.GetUserBalance(&bind.CallOpts{}, addr, tokenAddress)
			if err != nil {
				msg = "error getting balance to address"
				break
			}
			msg += fmt.Sprintf("%d. %s: %s \n", index+1, tokenName, bal.String())
		}
	}
	// comes here when 2*{choice#}
	switch len(textArray) {
	case 2:
		// we have a tokenID
		index := textArray[1]
		i, err := strconv.ParseInt(index, 10, 64)
		if len(tokenArray) <= int(i) || int(i) <= 0 {
			msg = "END Wrong token chosen"
			break
		}
		und, err := conn.memory.Get(withdrawKey)
		if err != nil {

		}
		w := und.(WithdrawRequest)
		tokenName := tokenArray[i-1]
		w.Token = tokenName
		conn.memory.Write(withdrawKey, w)

		msg = fmt.Sprintf("%s", `CON How many tokens would you like to withdraw?`)

	case 3:
		// we have the token
		amount := textArray[2]

		a, ok := new(big.Int).SetString(amount, 10)
		if !ok {
			//
		}
		addr, err := conn.telcoInteractor.TelcoCaller.GetFromPhoneNumberToAddress(&bind.CallOpts{}, phoneNumber)
		if err != nil {
			msg = "error getting phone number to address"
			break
		}
		und, err := conn.memory.Get(withdrawKey)
		if err != nil {

		}
		w := und.(WithdrawRequest)
		w.Amount = a
		w.State = "init"
		w.ToAddress = addr.String()
		conn.memory.Write(withdrawKey, w)
		msg = `CON Please provide a valid password`
		// 2*{choice#}*{amount}*{password}*stop
	case 4:
		// we get the password
		password := textArray[3]

		und, err := conn.memory.Get(withdrawKey)
		if err != nil {

		}
		isValid, err := conn.userInteractor.UserCaller.IsPasswordValid(&bind.CallOpts{}, []byte(password))
		if err != nil || !isValid {
			msg = `END Password is not valid`
			break
		}
		w := und.(WithdrawRequest)
		w.Password = password
		conn.memory.Write(withdrawKey, w)
		msg = fmt.Sprintf("%s \n %s: %s\n %s: %s\n %s: %s\n %s: %s\n 1.Confirm\n 2.Deny\n",
			`CON Please confirm what you are sending:`,
			"Token", w.Token,
			"Amount", w.Amount.String(),
			"To", "Yourself",
			"Password", w.Password,
		)
	case 5:
		// we get the confirmation
		conf := textArray[4]
		i, err := strconv.ParseInt(conf, 10, 64)
		if i == 2 {
			msg = "END Operation aborted"
			break
		}
		if i != 1 {
			msg = "END Invalid choice"
			break
		}
		und, err := conn.memory.Get(withdrawKey)
		if err != nil {

		}
		w := und.(WithdrawRequest)
		// actually send the transaction
		bTx, err := conn.telcoInteractor.UserInDirectWithdraw(conn.keyedTransactor, tokenMap[w.Token], w.Amount, phoneNumber, []byte(w.Password))
		if err != nil {
			fmt.Println(err)
			msg = "END Problem sending the transaction"
			break
		}
		w.State = "pending"
		w.TxHash = bTx.Hash().String()
		conn.memory.Write(withdrawKey, w)

		// check it's ok
		msg = `END Withdral Sent!`
	}
	return msg, err
}

func (conn *connection) MyAccount(textArray []string, sessionID, phoneNumber, networkCode, text string) (string, error) {
	var msg string
	var err error
	// comes here when 1
	if len(textArray) == 1 {
		msg = `CON What would you like to access?
		1. My Account Balances
		2. My Account Address
		3. My Phone Number
		4. My Transaction History`
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
			msg = fmt.Sprintf("%s %s", `END You phone number is`, phoneNumber)
		case "4":
			// Tx History
			c, _ := conn.memory.Get(phoneNumber)
			txHistory := c.(TransactionHistory)
			if len(txHistory.Transactions) == 0 {
				msg = "No transaction history"
			}
			msg = "END Here are your transactions:\n"
			for _, tx := range txHistory.Transactions {
				msg += fmt.Sprintf("%s: %s\n %s: %s\n %s: %s\n %s: %s\n ------------------- \n",
					"Token", tx.Token,
					"Amount", tx.Amount.String(),
					"To", tx.ToPhoneNumber,
					"Password", tx.Password,
				)
			}
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
