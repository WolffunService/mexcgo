package types

// SideType define side type of order
type SideType string

// OrderType define order type
type OrderType string

const (
	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"

	OrderTypeLimit  OrderType = "LIMIT"
	OrderTypeMarket OrderType = "MARKET"
	//LIMIT_MAKER (Limit maker order)
	//IMMEDIATE_OR_CANCEL (Immediate or cancel order)
	FILL_OR_KILL OrderType = "FILL_OR_KILL" // (Fill or kill order)
)

type OrderStatus string

const (
	NEW                OrderStatus = "NEW"                // Uncompleted
	FILLED             OrderStatus = "FILLED"             //Filled
	PARTIALLY_FILLED   OrderStatus = "PARTIALLY_FILLED"   //Partially filled
	CANCELED           OrderStatus = "CANCELED"           // Canceled
	PARTIALLY_CANCELED OrderStatus = "PARTIALLY_CANCELED" //Partially canceled
)
