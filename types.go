package main

import "math/big"

type OpenChannelRequest struct {
	PartnerAddress string `json:"partner_address"`
	TokenAddress   string `json:"token_address"`
	TotalDeposit   int    `json:"total_deposit"`
	SettleTimeout  int    `json:"settle_timeout"`
}

type OpenChannelResponse struct {
	TokenNetworkIdentifier string `json:"token_network_identifier"`
	ChannelIdentifier      int    `json:"channel_identifier"`
	PartnerAddress         string `json:"partner_address"`
	TokenAddress           string `json:"token_address"`
	Balance                int    `json:"balance"`
	TotalDeposit           int    `json:"total_deposit"`
	State                  string `json:"state"`
	SettleTimeout          int    `json:"settle_timeout"`
	RevealTimeout          int    `json:"reveal_timeout"`
}

type TopUpChannelRequest struct {
	TotalDeposit int `json:"total_deposit"`
}
type TopUpChannelResponse struct {
	TokenNetworkIdentifier string `json:"token_network_identifier"`
	ChannelIdentifier      int    `json:"channel_identifier"`
	PartnerAddress         string `json:"partner_address"`
	TokenAddress           string `json:"token_address"`
	Balance                int    `json:"balance"`
	TotalDeposit           int    `json:"total_deposit"`
	State                  string `json:"state"`
	SettleTimeout          int    `json:"settle_timeout"`
	RevealTimeout          int    `json:"reveal_timeout"`
}

type CloseChannelRequest struct {
	State string `json:"state"`
}
type CloseChannelResponse struct {
	TokenNetworkIdentifier string `json:"token_network_identifier"`
	ChannelIdentifier      int    `json:"channel_identifier"`
	PartnerAddress         string `json:"partner_address"`
	TokenAddress           string `json:"token_address"`
	Balance                int    `json:"balance"`
	TotalDeposit           int    `json:"total_deposit"`
	State                  string `json:"state"`
	SettleTimeout          int    `json:"settle_timeout"`
	RevealTimeout          int    `json:"reveal_timeout"`
}

type PaymentRequest struct {
	Amount     int `json:"amount"`
	Identifier int `json:"identifier"`
}

type PaymentResponse struct {
	InitiatorAddress string `json:"initiator_address"`
	TargetAddress    string `json:"target_address"`
	TokenAddress     string `json:"token_address"`
	Amount           int    `json:"amount"`
	Identifier       int    `json:"identifier"`
}

type ChannelInformation struct {
	TokenNetworkIdentifier string `json:"token_network_identifier"`
	ChannelIdentifier      int    `json:"channel_identifier"`
	PartnerAddress         string `json:"partner_address"`
	TokenAddress           string `json:"token_address"`
	Balance                int    `json:"balance"`
	TotalDeposit           int    `json:"total_deposit"`
	State                  string `json:"state"`
	SettleTimeout          int    `json:"settle_timeout"`
	RevealTimeout          int    `json:"reveal_timeout"`
}

type TransferInformation struct {
	Amount        *big.Int
	ToPhoneNumber string
	Token         string
	Password      string
	TxHash        string
	State         string
}

type TransactionHistory struct {
	Transactions []TransferInformation
}

type WithdrawRequest struct {
	Amount    *big.Int
	Token     string
	ToAddress string
	Password  string
	TxHash    string
	State     string
}
