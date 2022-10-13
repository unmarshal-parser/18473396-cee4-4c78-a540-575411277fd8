// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"math/big"
	"time"
)

var (
	_ = decimal.Decimal{}
	_ = big.NewInt
	_ = datatypes.JSON{}
	_ = time.Time{}
)

func GetEthWithdrawnEventHash() string {
	return "0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b"
}

type EthWithdrawnEvent struct {
	Target string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:2bf9061f-73b0-42da-9be7-6876a30a05f0,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:2bf9061f-73b0-42da-9be7-6876a30a05f0,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:2bf9061f-73b0-42da-9be7-6876a30a05f0,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetSwapTargetAddedEventHash() string {
	return "0xb907822409611d127ab6a64611591b98e03a6a85ade4f258bae26b7c1efdfeaf"
}

type SwapTargetAddedEvent struct {
	Target string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:e4b1c76c-546b-431a-be81-f7070ca5d9c8,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:e4b1c76c-546b-431a-be81-f7070ca5d9c8,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:e4b1c76c-546b-431a-be81-f7070ca5d9c8,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetSwapTargetRemovedEventHash() string {
	return "0x393b8be3e26787f19285ecd039dfd80bc6507828750f4d50367e6efe2524695c"
}

type SwapTargetRemovedEvent struct {
	Target string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:cd3f69c7-88cd-45ac-a84a-cb842741e18b,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:cd3f69c7-88cd-45ac-a84a-cb842741e18b,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:cd3f69c7-88cd-45ac-a84a-cb842741e18b,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetTokenWithdrawnEventHash() string {
	return "0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620"
}

type TokenWithdrawnEvent struct {
	Token  string
	Target string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:ab645d2f-f33d-461b-8c16-a6c3bf2f36e6,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:ab645d2f-f33d-461b-8c16-a6c3bf2f36e6,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:ab645d2f-f33d-461b-8c16-a6c3bf2f36e6,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToEthWithPermitMethodHash() string {
	return "b3093838"
}

type FillQuoteTokenToEthWithPermitMethod struct {
	SellTokenAddress         string
	Target                   string
	SwapCallData             []byte
	SellAmount               decimal.Decimal `gorm:"type:numeric"`
	FeePercentageBasisPoints decimal.Decimal `gorm:"type:numeric"`
	PermitData               datatypes.JSON

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	TokenPriceFeePercentageBasisPoints      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeePercentageBasisPoints float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:64f8071e-5fb4-415a-864b-f9ddbc753446,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:64f8071e-5fb4-415a-864b-f9ddbc753446,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToTokenMethodHash() string {
	return "55e4b7be"
}

type FillQuoteTokenToTokenMethod struct {
	SellTokenAddress string
	BuyTokenAddress  string
	Target           string
	SwapCallData     []byte
	SellAmount       decimal.Decimal `gorm:"type:numeric"`
	FeeAmount        decimal.Decimal `gorm:"type:numeric"`

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	TokenPriceFeeAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeeAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:bc9310c4-751f-4b82-b3aa-6558f9cde7ab,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:bc9310c4-751f-4b82-b3aa-6558f9cde7ab,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToTokenWithPermitMethodHash() string {
	return "b0480bbd"
}

type FillQuoteTokenToTokenWithPermitMethod struct {
	SellTokenAddress string
	BuyTokenAddress  string
	Target           string
	SwapCallData     []byte
	SellAmount       decimal.Decimal `gorm:"type:numeric"`
	FeeAmount        decimal.Decimal `gorm:"type:numeric"`
	PermitData       datatypes.JSON

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	TokenPriceFeeAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeeAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:a75b379b-05a0-4cac-af09-c6870613d0f4,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:a75b379b-05a0-4cac-af09-c6870613d0f4,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetWithdrawTokenMethodHash() string {
	return "01e33667"
}

type WithdrawTokenMethod struct {
	Token  string
	To     string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:d7b07f55-9e5a-4d04-88ce-22c2e3f5ad77,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:d7b07f55-9e5a-4d04-88ce-22c2e3f5ad77,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteEthToTokenMethodHash() string {
	return "3c2b9a7d"
}

type FillQuoteEthToTokenMethod struct {
	BuyTokenAddress string
	Target          string
	SwapCallData    []byte
	FeeAmount       decimal.Decimal `gorm:"type:numeric"`

	TokenPriceFeeAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeeAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:a7d22484-073a-47d6-8bc7-c76610eebc6e,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:a7d22484-073a-47d6-8bc7-c76610eebc6e,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToEthMethodHash() string {
	return "999b6464"
}

type FillQuoteTokenToEthMethod struct {
	SellTokenAddress         string
	Target                   string
	SwapCallData             []byte
	SellAmount               decimal.Decimal `gorm:"type:numeric"`
	FeePercentageBasisPoints decimal.Decimal `gorm:"type:numeric"`

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	TokenPriceFeePercentageBasisPoints      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeePercentageBasisPoints float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:7b8c4e37-5f13-4dfd-b101-7bd2109a9a56,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:7b8c4e37-5f13-4dfd-b101-7bd2109a9a56,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetTransferOwnershipMethodHash() string {
	return "f2fde38b"
}

type TransferOwnershipMethod struct {
	NewOwner string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:2d7e4540-4b38-41e2-9ec9-f577cf32b9f5,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:2d7e4540-4b38-41e2-9ec9-f577cf32b9f5,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetUpdateSwapTargetsMethodHash() string {
	return "97bbda0e"
}

type UpdateSwapTargetsMethod struct {
	Target string
	Add    bool

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:b4bfb6ad-722b-4503-b185-7f5137044c48,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:b4bfb6ad-722b-4503-b185-7f5137044c48,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetWithdrawEthMethodHash() string {
	return "1b9a91a4"
}

type WithdrawEthMethod struct {
	To     string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:4b98a186-acd6-480c-aeac-3e20580deced,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:4b98a186-acd6-480c-aeac-3e20580deced,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

type LastSyncedBlock struct {
	Contract    string `gorm:"primaryKey"`
	ChainID     string `gorm:"primaryKey"`
	SyncType    string `gorm:"primaryKey"`
	BlockNumber uint64
}

// Plugin Models

type TokenDetails struct {
	ID      int
	Address string `gorm:"uniqueIndex:address_and_chain"`
	Symbol  string
	ChainID string `gorm:"uniqueIndex:address_and_chain"`
	Decimal int
}

// Config

type PostgresConfig struct {
	ConnectionString string `mapstructure:"connection_string"`
	TablePrefix      string `mapstructure:"table_prefix"`
	CreateBatchSize  int    `mapstructure:"create_batch_size"`
}

type IndexerConfig struct {
	EthEndpoint       string `mapstructure:"eth_endpoint"`
	ContractAddress   string `mapstructure:"contract_address"`
	StartBlock        int    `mapstructure:"start_block"`
	ApiKey            string `mapstructure:"api_key"`
	PostgresConfig    `mapstructure:"postgres_config"`
	LagToHighestBlock int `mapstructure:"lag_to_highest_block"`
	StepSize          int `mapstructure:"step_size"`
	ParallelCalls     int `mapstructure:"parallel_calls_for_logs"`
}

func (i *IndexerConfig) AssignDefaults() {
	if i.PostgresConfig.CreateBatchSize == 0 {
		i.PostgresConfig.CreateBatchSize = 100
	}
	if i.StepSize == 0 {
		i.StepSize = 50
	}
	if i.LagToHighestBlock == 0 {
		i.LagToHighestBlock = 10
	}
}
