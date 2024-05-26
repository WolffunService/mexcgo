package types

import (
	"github.com/msw-x/moon/ujson"
)

type OrderBookQuery struct {
	Symbol string
	Limit  int //default is 100
}

type Orderbook struct {
	Asks         [][]ujson.Float64 `json:"asks"`
	Bids         [][]ujson.Float64 `json:"bids"`
	LastUpdateId int64             `json:"lastUpdateId"`
	//Timestamp    int64             `json:"timestamp"`
}

//place new order
type NewOrder struct {
	Symbol   string
	Side     SideType
	Type     OrderType
	Quantity float64
	Price    float64
}

type NewOrderResponse struct {
	OrderId      string        `json:"orderId"`
	OrderListId  int           `json:"orderListId"`
	OrigQty      ujson.Float64 `json:"origQty"`
	Price        ujson.Float64 `json:"price"`
	Side         SideType      `json:"side"`
	Symbol       string        `json:"symbol"`
	TransactTime int64         `json:"transactTime"`
	Type         OrderType     `json:"type"`
}

//cancel order
type CancelOrder struct {
	Symbol  string
	OrderId string
}

type CancelOrderResponse struct {
	Symbol              string        `json:"symbol"`
	OrderId             string        `json:"orderId"`
	OrderListId         int           `json:"orderListId"`
	OrigQty             ujson.Float64 `json:"origQty"`
	Price               ujson.Float64 `json:"price"`
	ExecutedQty         ujson.Float64 `json:"executedQty"`
	CummulativeQuoteQty ujson.Float64 `json:"cummulativeQuoteQty"` //Cummulative quote quantity
	Side                SideType      `json:"side"`
	Status              int64         `json:"transactTime"`
	Type                OrderType     `json:"type"`
}

type OpenOrderQuery struct {
	Symbol string
}

type QueryOrder struct {
	Symbol  string
	OrderId string
}

type AllOrderQuery struct {
	Symbol     string
	StartTime  *int64
	EndTime    *int64
	Limit      *int   // 	Default 500; max 1000;
	RecvWindow *int64 //	NO
	Timestamp  *int64 //YES
}

type OrderResponse struct {
	Symbol              string        `json:"symbol"`
	ClientOrderId       string        `json:"clientOrderId"`
	CummulativeQuoteQty string        `json:"cummulativeQuoteQty"`
	ExecutedQty         ujson.Float64 `json:"executedQty"`
	IcebergQty          ujson.Float64 `json:"icebergQty"`
	IsWorking           bool          `json:"isWorking"`
	OrderId             string        `json:"orderId"`
	OrderListId         int           `json:"orderListId"`
	OrigQty             ujson.Float64 `json:"origQty"`
	OrigQuoteOrderQty   ujson.Float64 `json:"origQuoteOrderQty"`
	Price               ujson.Float64 `json:"price"`
	Side                SideType      `json:"side"`
	Status              OrderStatus   `json:"status"`
	StopPrice           string        `json:"stopPrice"`
	Time                int64         `json:"time"`
	TimeInForce         string        `json:"timeInForce"`
	Type                OrderType     `json:"type"`
	UpdateTime          string        `json:"updateTime"`
}
