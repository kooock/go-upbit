package upbit

type Account struct{
	Currency string `json:"currency"`
	Balance string `json:"balance"`
	Locked string `json:"locked"`
	AvgKrwBuyPrice string `json:"avg_krw_buy_price"`
	Modified bool `json:"modified"`
}

type MarketOrder struct {
	Currency string `json:"currency"`
	PriceUnit float64 `json:"price_unit"`
	MinTotal float64 `json:"price_unit"`
}


type Market struct {
	Id string `json:"id"`
	Name string `json:"name"`
	OrderTypes []string `json:"order_types"`
	OrderSides []string `json:"order_sides"`
	Bid MarketOrder `json:"bid"`
	Ask MarketOrder `json:"ask"`
	MaxTotal string `json:"max_total"`
	State string `json:"state"`
}

type OrderChance struct {
	BidFee string `json:"bid_fee"`
	AskFee string `json:"ask_fee"`
	Market Market `json:"market"`
	BidAccount Account `json:"bid_account"`
	AskAccount Account `json:"ask_account"`
}

type Transaction struct {
	Market string `json:"market"`
	Uuid string `json:"uuid"`
	Price string `json:"price"`
	Volume string `json:"volume"`
	Funds string `json:"funds"`
	AskFee string `json:"ask_fee"`
	BidFee string `json:"bid_fee"`
	CreateAt string `json:"created_at"`
	Side string `json:"side"`
}

type Order struct {
	Uuid string `json:"uuid"`
	Side string `json:"side"`
	OrdType string `json:"ord_type"`
	Price string `json:"price"`
	AvgPrice string `json:"avg_price"`
	State string `json:"state"`
	Market string `json:"market"`
	Create string `json:"created_at"`
	Volume string `json:"volume"`
	RemainingVolume string `json:"remaining_volume"`
	ReservedFee string `json:"reserved_fee"`
	RemainingFee string `json:"remaining_fee"`
	PaidFee string `json:"paid_fee"`
	Locked string `json:"locked"`
	ExecutedVolume string `json:"executed_volume"`
	TradesCount int64 `json:"trades_count"`
	Trades []Transaction `json:"trades"`
}