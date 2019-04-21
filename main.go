package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// no push notifications! that means that we need to either send sms or something of the sort to the user dabout the state of
	// his request

	// create new connection

	conn, err := NewConnection()
	if err != nil {
		panic(err)
	}

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.FormValue("sessionId")
		phoneNumber := r.FormValue("phoneNumber")[1:]
		networkCode := r.FormValue("networkCode")
		text := r.FormValue("text")
		// check if the user has a history, if not create one
		_, err := conn.memory.Get(phoneNumber)
		if err != nil {
			// if error, create a new one
			txHistory := TransactionHistory{}
			conn.memory.Write(phoneNumber, txHistory)
		}
		fmt.Println(sessionId, phoneNumber, networkCode, text)
		if text == "" {
			w.Write([]byte(`CON 
			Welcome to crypto USSD Portal, what would you like to do?
			1. Send Money
			2. Withdraw
			3. Buy airtime
			4. Pay goods
			5. Pay bill
			6. Get a loan PRWD by MKR
			7. My Account
			`))
			return
		}
		textArray := strings.Split(text, "*")

		switch textArray[0] {
		case "1":
			// Send Money
			resp, err := conn.SendMoney(textArray, sessionId, phoneNumber, networkCode, text)
			if err != nil {
				// handle it
				w.Write([]byte(err.Error()))
			}
			w.Write([]byte(resp))
		case "2":
			// Withdraw Cash
			resp, err := conn.Withdraw(textArray, sessionId, phoneNumber, networkCode, text)
			if err != nil {
				// handle it
				w.Write([]byte(err.Error()))
			}
			w.Write([]byte(resp))
		case "3":
			w.Write([]byte("END Not Implemented, sorry!"))
		case "4":
			w.Write([]byte("END Not Implemented, sorry!"))
		case "5":
			w.Write([]byte("END Not Implemented, sorry!"))
		case "6":
			resp, err := conn.MakerDao(textArray, sessionId, phoneNumber, networkCode, text)
			if err != nil {
				// handle it
				w.Write([]byte(err.Error()))
			}
			w.Write([]byte(resp))
		case "7":
			// My Account
			resp, err := conn.MyAccount(textArray, sessionId, phoneNumber, networkCode, text)
			if err != nil {
				// handle it
				w.Write([]byte(err.Error()))
			}
			w.Write([]byte(resp))
		case "8":
			w.Write([]byte("END Not Implemented, sorry!"))
		default:
			w.Write([]byte("END Not Implemented, sorry!"))
		}

	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("Listening on port", port)
	http.ListenAndServe(":"+port, r)
}
