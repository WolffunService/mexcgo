package examples

import (
	"github.com/WolffunService/mexcgo/mexc"
	"github.com/WolffunService/mexcgo/types"
	"os"
	"testing"
)

var c *mexc.Client

var apiKey = os.Getenv("MEXC_API_KEY")
var apiSecret = os.Getenv("MEXC_SECRET_KEY")

func TestClient(t *testing.T) {
	c = mexc.NewClient()
	res := c.GetOrderBook(types.OrderBookQuery{
		Symbol: "THGUSDT",
		Limit:  1,
	})

	if !res.Ok() {
		t.Fatal(res.Error)
	}
	t.Log(res.Data)
}

func TestAccount(t *testing.T) {

	c = mexc.NewClient().WithAuth(apiKey, apiSecret)
	res := c.GetAccountInfo()

	if !res.Ok() {
		t.Fatal(res.Error)
	}
	t.Log(res.Data)
}

func TestExchangeInfo(t *testing.T) {
	c = mexc.NewClient()
	res := c.ExchangeInfo(types.ExchangeInfoQuery{
		Symbol: "THGUSDT",
	})

	if !res.Ok() {
		t.Fatal(res.Error)
	}
	t.Log(res.Data)
}

func TestPlaceOrder(t *testing.T) {
	c = mexc.NewClient().WithAuth(apiKey, apiSecret)
	res := c.NewOrder(types.NewOrder{
		Symbol:   "THGUSDT",
		Side:     types.SideTypeSell,
		Type:     types.OrderTypeLimit,
		Quantity: 6000,
		Price:    0.0888,
	})

	if !res.Ok() {
		t.Fatal(res.Error)
	}
	t.Log(res.Data)
}

func TestGetOpenOrders(t *testing.T) {
	c = mexc.NewClient().WithAuth(apiKey, apiSecret)
	res := c.GetOpenOrders(types.OpenOrderQuery{
		Symbol: "THGUSDT",
	})

	if !res.Ok() {
		t.Fatal(res.Error)
	}
	t.Log(res.Data)
}

func TestGetAllOrders(t *testing.T) {
	c = mexc.NewClient().WithAuth(apiKey, apiSecret)
	res := c.GetAllOrders(types.AllOrderQuery{Symbol: "THGUSDT"})

	if !res.Ok() {
		t.Fatal(res.Error)
	}
	t.Log(res.Data)
}

func TestCancelOrder(t *testing.T) {
	c = mexc.NewClient().WithAuth(apiKey, apiSecret)
	orderId := "C02__423626447278653441018"
	res := c.CancelOrder(types.CancelOrder{Symbol: "THGUSDT", OrderId: orderId})

	if !res.Ok() {
		t.Fatal(res.Error)
	}
	t.Log(res.Data)
}
