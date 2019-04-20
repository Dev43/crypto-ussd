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
		phoneNumber := r.FormValue("phoneNumber")
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
			6. ATM Withdrawal
			7. My Account
			`))
			return
		}
		textArray := strings.Split(text, "*")

		switch textArray[0] {
		case "1":
			// Send Money
			resp, err := conn.SendMoney(textArray, sessionId, phoneNumber[1:], networkCode, text)
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
			// Buy airtime
			// resp, err := SendMoney(sessionId, phoneNumber, networkCode, text)
			// if err != nil {
			// 	// handle it
			// }
			// w.Write([]byte(resp))
		case "4":
			// Pay goods
			// resp, err := SendMoney(sessionId, phoneNumber, networkCode, text)
			// if err != nil {
			// 	// handle it
			// }
			// w.Write([]byte(resp))
		case "5":
			// Pay bill
			// resp, err := SendMoney(sessionId, phoneNumber, networkCode, text)
			// if err != nil {
			// 	// handle it
			// }
			// w.Write([]byte(resp))
		case "6":
			// ATM withdrawal
			// resp, err := SendMoney(sessionId, phoneNumber, networkCode, text)
			// if err != nil {
			// 	// handle it
			// }
			// w.Write([]byte(resp))
		case "7":
			// My Account
			resp, err := conn.MyAccount(textArray, sessionId, phoneNumber, networkCode, text)
			if err != nil {
				// handle it
				w.Write([]byte(err.Error()))

			}
			w.Write([]byte(resp))
		case "8":
			// Create a loan!
		}

	})
	// // FIRST REGISTER TOKEN
	// registerToken("0x396764f15ed1467883A9a5B7D42AcFb788CD1826")
	// THEN OPEN CHANNEL
	// openChannel("0x0bae0289AAA26845224F528F9B9DefE69e01606E", "0x396764f15ed1467883A9a5B7D42AcFb788CD1826", 150)
	// NOW TOP UP THE CHANNEL
	// topUpChannel("0x0bae0289AAA26845224F528F9B9DefE69e01606E", "0x396764f15ed1467883A9a5B7D42AcFb788CD1826", 15000000)
	// now make a payment
	// makePayment("0x0bae0289AAA26845224F528F9B9DefE69e01606E", "0x396764f15ed1467883A9a5B7D42AcFb788CD1826", 100, 3)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("Listening on port", port)
	http.ListenAndServe(":"+port, r)
}

/*
2019-04-19T21:40:01.446686+00:00 app[web.1]: ATUid_d74e8947568bba6721afbb8b51205b9c +27777777777 99999 1
2019-04-19T21:40:01.446772+00:00 app[web.1]: 2019/04/19 21:40:01 [7be3e2c2-b9ff-4ab8-b4e5-b4dacdc4eff2] "POST http://ancient-ravine-46133.herokuapp.com/ HTTP/1.1" from 134.213.151.223 - 200 69B in 1.32171ms
2019-04-19T21:40:04.444290+00:00 heroku[router]: at=info method=POST path="/" host=ancient-ravine-46133.herokuapp.com request_id=f31652e7-bd8a-4561-8f25-b56b40bac9e6 fwd="134.213.151.223" dyno=web.1 connect=0ms service=1ms status=200 bytes=205 protocol=https
2019-04-19T21:40:04.443644+00:00 app[web.1]: ATUid_d74e8947568bba6721afbb8b51205b9c +27777777777 99999 1*2
2019-04-19T21:40:04.443667+00:00 app[web.1]: 2019/04/19 21:40:04 [f31652e7-bd8a-4561-8f25-b56b40bac9e6] "POST http://ancient-ravine-46133.herokuapp.com/ HTTP/1.1" from 134.213.151.223 - 200 69B in 310.642µs
2019-04-19T21:40:12.387964+00:00 app[web.1]: ATUid_d74e8947568bba6721afbb8b51205b9c +27777777777 99999 1*2*3
2019-04-19T21:40:12.387996+00:00 app[web.1]: 2019/04/19 21:40:12 [bd9cca8a-ff82-44b2-9bf9-fc493c8ac539] "POST http://ancient-ravine-46133.herokuapp.com/ HTTP/1.1" from 134.213.151.223 - 200 69B in 354.129µs
2019-04-19T21:40:12.388910+00:00 heroku[router]: at=info method=POST path="/" host=ancient-ravine-46133.herokuapp.com request_id=bd9cca8a-ff82-44b2-9bf9-fc493c8ac539 fwd="134.213.151.223" dyno=web.1 connect=0ms service=1ms status=200 bytes=205 protocol=https
2019-04-19T21:40:14.977871+00:00 heroku[router]: at=info method=POST path="/" host=ancient-ravine-46133.herokuapp.com request_id=58845bf4-09bb-4563-97a6-74a65c27ba74 fwd="134.213.151.223" dyno=web.1 connect=0ms service=1ms status=200 bytes=205 protocol=https
2019-04-19T21:40:14.977056+00:00 app[web.1]: ATUid_d74e8947568bba6721afbb8b51205b9c +27777777777 99999 1*2*3*4
2019-04-19T21:40:14.977079+00:00 app[web.1]: 2019/04/19 21:40:14 [58845bf4-09bb-4563-97a6-74a65c27ba74] "POST http://ancient-ravine-46133.herokuapp.com/ HTTP/1.1" from 134.213.151.223 - 200 69B in 151.287µs
2019-04-19T21:40:26.430557+00:00 app[web.1]: ATUid_d74e8947568bba6721afbb8b51205b9c +27777777777 99999 1*2*3*4*5
2019-04-19T21:40:26.430576+00:00 app[web.1]: 2019/04/19 21:40:26 [84c7078f-2fc3-4070-bfb0-6391c85287da] "POST http://ancient-ravine-46133.herokuapp.com/ HTTP/1.1" from 134.213.151.223 - 200 69B in 340.312µs
2019-04-19T21:40:26.431449+00:00 heroku[router]: at=info method=POST path="/" host=ancient-ravine-46133.herokuapp.com request_id=84c7078f-2fc3-4070-bfb0-6391c85287da fwd="134.213.151.223" dyno=web.1 connect=0ms service=1ms status=200 bytes=205 protocol=https
2019-04-19T21:40:28.905944+00:00 heroku[router]: at=info method=POST path="/" host=ancient-ravine-46133.herokuapp.com request_id=e4eee290-4802-48cd-b183-4549532d6c52 fwd="134.213.151.223" dyno=web.1 connect=0ms service=1ms status=200 bytes=205 protocol=https
2019-04-19T21:40:28.904832+00:00 app[web.1]: ATUid_d74e8947568bba6721afbb8b51205b9c +27777777777 99999 1*2*3*4*5*6
2019-04-19T21:40:28.905166+00:00 app[web.1]: 2019/04/19 21:40:28 [e4eee290-4802-48cd-b183-4549532d6c52] "POST http://ancient-ravine-46133.herokuapp.com/ HTTP/1.1" from 134.213.151.223 - 200 69B in 470.674µs */

// user goes and needs a private key
// user opens account, adds in all of their one time passwords in it
// from there they are in the smart contract -- as their phone number
// they need to use their one time passwords for the service provider to actually be able to interact with the system
