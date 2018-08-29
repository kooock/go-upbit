package upbit

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"strconv"
)

func (upbitClient *UpbitClient)OrderToMarket(market string, side string, volume string, price string, ordType string,identifier string) (*Order,error){
	method := "POST"
	url := "https://api.upbit.com/v1/orders"
	queries := map[string]string{}
	if (market == "default") || (market == ""){
		return nil,errors.New("you have to input market id")
	}
	queries["market"] = market

	if (side == "default") || (side == ""){
		return nil,errors.New("you have to input side as ask or bid at least one")
	}
	queries["side"] = side

	if (volume == "default") || (volume == ""){
		return nil,errors.New("you have to input volume")
	}
	queries["volume"] = volume

	if (price == "default") || (price == ""){
		return nil,errors.New("you have to input unit price")
	}
	queries["price"] = price

	if (ordType != "default") && (ordType != ""){
		queries["ord_type"] = ordType
	}

	responseObject := &Order{}
	responseJson,err := upbitClient.SendRequest(method,url,queries)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseJson, &responseObject)
	if err != nil {
		return nil, err
	}
	return responseObject,nil

}

func (upbitClient *UpbitClient)CancelOrder(uuid string) (*Order,error){
	method := "DELETE"
	url := "https://api.upbit.com/v1/orders"
	if uuid == "default" || uuid == ""{
		return nil, errors.New("You have to input order uuid")
	}
	queries := map[string]string{
		"uuid" : uuid,
	}
	responseObject := &Order{}
	responseJson,err := upbitClient.SendRequest(method,url,queries)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseJson, &responseObject)
	if err != nil {
		return nil, err
	}
	return responseObject,nil
}

func (upbitClient *UpbitClient)OrderBuy(market string, volume float64, price float64, ordType string,identifier string) (*Order,error) {
	volumeString := fmt.Sprintf("%f",volume)
	priceString := fmt.Sprintf("%f",price)
	return upbitClient.OrderToMarket(market, "bid", volumeString, priceString, ordType, identifier)
}

func (upbitClient *UpbitClient)OrderSell(market string, volume float64, price float64, ordType string,identifier string) (*Order,error) {
	volumeString := fmt.Sprintf("%f",volume)
	priceString := fmt.Sprintf("%f",price)
	return upbitClient.OrderToMarket(market, "ask", volumeString, priceString, ordType, identifier)
}


func (upbitClient *UpbitClient)OrderBuyMarketPrice(market string, volume float64, ordType string,identifier string) (*Order,error) {
	orderbook,err := GetOrderbook(upbitClient,market)
	if err != nil {
		return nil,err
	}

	bidPrice := orderbook.OrderbookUnits[0].BidPrice
	bidSize := orderbook.OrderbookUnits[0].BidSize

	if volume > bidSize {
		volume = bidSize
	}
	return upbitClient.OrderBuy(market,volume,bidPrice,"limit",identifier)
}

func (upbitClient *UpbitClient)OrderBuyMarketPriceByAccountWeight(market string, weight float64, ordType string,identifier string) (*Order,error) {

	accounts, err := upbitClient.CheckAccounts()
	if err != nil {
		return nil,err
	}

	balance := float64(0)
	currency := strings.Split(market, "-")[0]
	for _,account := range accounts{
		if account.Currency == currency{
			balance,err = strconv.ParseFloat(account.Balance, 64)
			if err != nil {
				return nil, err
			}
		}
	}

	orderbook,err := GetOrderbook(upbitClient,market)
	if err != nil {
		return nil,err
	}

	askPrice := orderbook.OrderbookUnits[0].AskPrice
	askSize := orderbook.OrderbookUnits[0].AskSize

	weightBalance := balance * weight
	if currency == "KRW"{
		weightBalance = float64(int64(weightBalance * 0.9995))
	}else{
		weightBalance = float64(int64(weightBalance * 0.9995 * float64(100000000))) / float64(100000000)
	}
	orderSize := weightBalance / askPrice
	orderSize = float64(int64(orderSize * float64(100000000))) / float64(100000000)

	if orderSize > askSize {
		orderSize = askSize
	}


	return upbitClient.OrderBuy(market,orderSize,askPrice,"limit",identifier)
}


func (upbitClient *UpbitClient)OrderSellMarketPrice(market string, volume float64, ordType string,identifier string) (*Order,error) {
	orderbook,err := GetOrderbook(upbitClient,market)
	if err != nil {
		return nil,err
	}

	bidPrice := orderbook.OrderbookUnits[0].BidPrice
	bidSize := orderbook.OrderbookUnits[0].BidSize

	if volume > bidSize {
		volume = bidSize
	}
	return upbitClient.OrderSell(market,volume,bidPrice,"limit",identifier)
}
