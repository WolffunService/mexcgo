package mexc

import (
	"mexcgo/types"
)

func (o *Client) GetOrderBook(mq types.OrderBookQuery) Response[types.Orderbook] {
	return GetPub(o, "depth", mq, forward[types.Orderbook])
}

func (o *Client) NewOrderTest(mq types.NewOrder) Response[types.NewOrderResponse] {
	//mq.Timestamp = time.Now().UnixMilli()
	return Post(o, "order/test", mq, forward[types.NewOrderResponse])
}

func (o *Client) NewOrder(mq types.NewOrder) Response[types.NewOrderResponse] {
	return Post(o, "order", mq, forward[types.NewOrderResponse])
}

func (o *Client) CancelOrder(co types.CancelOrder) Response[types.CancelOrderResponse] {
	return Delete(o, "order", co, forward[types.CancelOrderResponse])
}

func (o *Client) GetOpenOrders(mq types.OpenOrderQuery) Response[[]types.OrderResponse] {
	return Get(o, "openOrders", mq, forward[[]types.OrderResponse])
}

func (o *Client) GetAllOrders(req types.AllOrderQuery) Response[[]types.OrderResponse] {
	return Get(o, "allOrders", req, forward[[]types.OrderResponse])
}

func (o *Client) QueryOrder(req types.QueryOrder) Response[types.OrderResponse] {
	return Get(o, "order", req, forward[types.OrderResponse])
}
