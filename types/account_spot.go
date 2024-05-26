package types

import "github.com/msw-x/moon/ujson"

type AccountInfo struct {
	AccountType      string    `json:"accountType"`
	Balances         []Balance `json:"balances"`
	BuyerCommission  string    `json:"buyerCommission"`
	CanDeposit       bool      `json:"canDeposit"`
	CanTrade         bool      `json:"canTrade"`
	CanWithdraw      bool      `json:"canWithdraw"`
	MakerCommission  string    `json:"makerCommission"`
	Permissions      []string  `json:"permissions"`
	SellerCommission string    `json:"sellerCommission"`
	TakerCommission  string    `json:"takerCommission"`
	UpdateTime       string    `json:"updateTime"`
}

type Balance struct {
	Asset  string        `json:"asset"`
	Free   ujson.Float64 `json:"free"`
	Locked ujson.Float64 `json:"locked"`
}

type ExchangeInfoQuery struct {
	Symbol  string
	Symbols string
}

type ExchangeInfoResponse struct {
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	RateLimits      []interface{} `json:"rateLimits"`
	ServerTime      int64         `json:"serverTime"`
	Symbols         []struct {
		BaseAsset                  string        `json:"baseAsset"`
		BaseAssetPrecision         int           `json:"baseAssetPrecision"`
		BaseCommissionPrecision    int           `json:"baseCommissionPrecision"`
		BaseSizePrecision          string        `json:"baseSizePrecision"`
		Filters                    []interface{} `json:"filters"`
		FullName                   string        `json:"fullName"`
		IsMarginTradingAllowed     bool          `json:"isMarginTradingAllowed"`
		IsSpotTradingAllowed       bool          `json:"isSpotTradingAllowed"`
		MakerCommission            string        `json:"makerCommission"`
		MaxQuoteAmount             string        `json:"maxQuoteAmount"`
		MaxQuoteAmountMarket       string        `json:"maxQuoteAmountMarket"`
		OrderTypes                 []string      `json:"orderTypes"`
		Permissions                []string      `json:"permissions"`
		QuoteAmountPrecision       string        `json:"quoteAmountPrecision"`
		QuoteAmountPrecisionMarket string        `json:"quoteAmountPrecisionMarket"`
		QuoteAsset                 string        `json:"quoteAsset"`
		QuoteAssetPrecision        int           `json:"quoteAssetPrecision"`
		QuoteCommissionPrecision   int           `json:"quoteCommissionPrecision"`
		QuotePrecision             int           `json:"quotePrecision"`
		Status                     string        `json:"status"`
		Symbol                     string        `json:"symbol"`
		TakerCommission            string        `json:"takerCommission"`
	} `json:"symbols"`
	Timezone string `json:"timezone"`
}
