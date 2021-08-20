package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Market struct {
	Name                  string          `json:"name"`
	Underlying            string          `json:"underlying"`
	BaseCurrency          string          `json:"baseCurrency"`
	QuoteCurrency         string          `json:"quoteCurrency"`
	Enabled               bool            `json:"enabled"`
	Ask                   decimal.Decimal `json:"ask"`
	Bid                   decimal.Decimal `json:"bid"`
	Last                  decimal.Decimal `json:"last"`
	PostOnly              bool            `json:"postOnly"`
	PriceIncrement        decimal.Decimal `json:"priceIncrement"`
	SizeIncrement         decimal.Decimal `json:"sizeIncrement"`
	Restricted            bool            `json:"restricted"`
	MinProvideSize        decimal.Decimal `json:"minProvideSize"`
	VolumeUSD24h          decimal.Decimal `json:"volumeUsd24h"`
	Type                  string          `json:"type"`
	QuoteVolume24h        decimal.Decimal `json:"quoteVolume24h"`
	HighLeverageFeeExempt bool            `json:"highLeverageFeeExempt"`
	Change1h              decimal.Decimal `json:"change1h"`
	Change24h             decimal.Decimal `json:"change24h"`
	ChangeBod             decimal.Decimal `json:"changeBod"`
}

type Trade struct {
	ID          int64           `json:"id"`
	Liquidation bool            `json:"liquidation"`
	Price       decimal.Decimal `json:"price"`
	Side        string          `json:"side"`
	Size        decimal.Decimal `json:"size"`
	Time        time.Time       `json:"time"`
}

type HistoricalPrice struct {
	StartTime time.Time       `json:"startTime"`
	Open      decimal.Decimal `json:"open"`
	Close     decimal.Decimal `json:"close"`
	High      decimal.Decimal `json:"high"`
	Low       decimal.Decimal `json:"low"`
	Volume    decimal.Decimal `json:"volume"`
}

type Ticker struct {
	Bid     decimal.Decimal `json:"bid"`
	Ask     decimal.Decimal `json:"ask"`
	BidSize decimal.Decimal `json:"bidSize"`
	AskSize decimal.Decimal `json:"askSize"`
	Last    decimal.Decimal `json:"last"`
	Time    FTXTime         `json:"time"`
}

type GetTradesParams struct {
	Limit     *int `json:"limit"`
	StartTime *int `json:"start_time"`
	EndTime   *int `json:"end_time"`
}

type GetHistoricalPricesParams struct {
	Resolution Resolution `json:"resolution"`
	Limit      *int       `json:"limit"`
	StartTime  *int       `json:"start_time"`
	EndTime    *int       `json:"end_time"`
}
