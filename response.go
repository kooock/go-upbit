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

type Candle struct {
	Market string `json:"market"`
	CandleDateTimeUTC string `json:"candle_date_time_utc"`
	CandleDateTimeKST string `json:"candle_date_time_kst"`
	OpeningPrice float64 `json:"opening_price"`
	HighPrice float64 `json:"high_price"`
	LowPrice float64 `json:"low_price"`
	TradePrice float64 `json:"trade_price"`
	Timestamp int64 `json:"timestamp"`
	CandleAccTradePrice float64 `json:"candle_acc_trade_price"`
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"`
	Unit int32 `json:"unit"`
}

type OrderbookUnit struct {
	AskPrice float64 `json:"ask_price"`
	BidPrice float64 `json:"bid_price"`
	AskSize float64 `json:"ask_size"`
	BidSize float64 `json:"bid_size"`
}

type Orderbook struct {
	Market string `json:"market"`
	Timestamp int64 `json:"timestamp"`
	TotalAskSize float64 `json:"total_ask_size"`
	TotalBidSize float64 `json:"total_bid_size"`
	OrderbookUnits []OrderbookUnit `json:"orderbook_units"`
}
