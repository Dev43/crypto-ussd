package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func doHttpCall(method string, url string, body []byte) ([]byte, error) {
	defaultURL := "http://localhost:5001/api/v1/"
	fullURL := defaultURL + url
	b := bytes.NewReader(body)
	req, err := http.NewRequest(method, fullURL, b)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	fmt.Println(res.Body)
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func registerToken(tokenAddress string) error {
	resp, err := doHttpCall("PUT", fmt.Sprintf("%s/%s", "tokens", tokenAddress), nil)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

func openChannel(partnerAddress string, tokenAddress string, totalDeposit int) (*OpenChannelResponse, error) {
	// default partner address 0x0bae0289AAA26845224F528F9B9DefE69e01606E
	body := OpenChannelRequest{
		PartnerAddress: partnerAddress,
		TokenAddress:   tokenAddress,
		TotalDeposit:   totalDeposit,
		SettleTimeout:  500,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	fmt.Println(b)
	resp, err := doHttpCall("PUT", fmt.Sprintf("%s", "channels"), b)
	if err != nil {
		return nil, err
	}
	data := OpenChannelResponse{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	// need to save the channel here somewhere
	fmt.Println(string(resp))
	return &data, nil
}

func getChannelInformation() ([]ChannelInformation, error) {
	// default partner address 0x0bae0289AAA26845224F528F9B9DefE69e01606E
	resp, err := doHttpCall("GET", fmt.Sprintf("%s", "channels"), nil)
	if err != nil {
		return nil, err
	}
	data := []ChannelInformation{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	// need to save the channel here somewhere
	fmt.Println(string(resp))
	return data, nil
}

func topUpChannel(partnerAddress string, tokenAddress string, totalDeposit int) (*TopUpChannelResponse, error) {
	// default partner address 0x0bae0289AAA26845224F528F9B9DefE69e01606E
	body := TopUpChannelRequest{
		TotalDeposit: totalDeposit,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := doHttpCall("PATCH", fmt.Sprintf("%s/%s/%s", "channels", tokenAddress, partnerAddress), b)
	if err != nil {
		return nil, err
	}
	data := TopUpChannelResponse{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	// need to save the channel here somewhere
	fmt.Println(string(resp))
	return &data, nil
}

func closeChannel(partnerAddress string, tokenAddress string, totalDeposit int) (*CloseChannelRequest, error) {
	// default partner address 0x0bae0289AAA26845224F528F9B9DefE69e01606E
	body := CloseChannelRequest{
		State: "closed",
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := doHttpCall("PATCH", fmt.Sprintf("%s/%s/%s", "channels", tokenAddress, partnerAddress), b)
	if err != nil {
		return nil, err
	}
	data := CloseChannelRequest{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	// need to save the channel here somewhere
	fmt.Println(string(resp))
	return &data, nil
}

func makePayment(targetAddress string, tokenAddress string, amount int, identifier int) (*PaymentResponse, error) {
	// default partner address 0x0bae0289AAA26845224F528F9B9DefE69e01606E
	body := PaymentRequest{
		Amount:     amount,
		Identifier: identifier,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(b))
	resp, err := doHttpCall("POST", fmt.Sprintf("%s/%s/%s", "payments", tokenAddress, targetAddress), b)
	if err != nil {
		return nil, err
	}
	data := PaymentResponse{}
	err = json.Unmarshal(resp, &data)

	if err != nil {
		return nil, err
	}
	// need to save the channel here somewhere
	fmt.Println(string(resp))
	return &data, nil
}
