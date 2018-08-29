package websocket


type ResponseMessage struct {
	Code string `json:"cd"`
	Market string  `json:"mk"`
	Timestamp int64 `json:"tms"`
	TotalAskSize float64 `json:"tas"`
	TotalBidSize float64 `json:"tbs"`
	OrderBook []UnitOrder `json:"obu"`
	StreamType string  `json:"st"`
}

type UnitOrder struct {
	AskPrice float64 `json:"askPrice"`
	BidPrice float64 `json:"bidPrice"`
	AskSize float64 `json:"askSize"`
	BidSize float64 `json:"bidSize"`
}
