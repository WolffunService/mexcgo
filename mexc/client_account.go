package mexc

import "mexcgo/types"

type defaultStruct struct {
}

var empty = func() defaultStruct {
	return defaultStruct{}
}

func (o *Client) GetAccountInfo() Response[types.AccountInfo] {
	return Get(o, "account", empty(), forward[types.AccountInfo])
}

func (o *Client) ExchangeInfo(req types.ExchangeInfoQuery) Response[types.ExchangeInfoResponse] {
	return GetPub(o, "exchangeInfo", req, forward[types.ExchangeInfoResponse])
}
