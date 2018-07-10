package upbit

import (
	"encoding/json"
	"errors"
	"fmt"
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
	volumeString := fmt.Sprintf("%v",volume)
	priceString := fmt.Sprintf("%v",price)
	return upbitClient.OrderToMarket(market, "bid", volumeString, priceString, ordType, identifier)
}

func (upbitClient *UpbitClient)OrderSell(market string, volume float64, price float64, ordType string,identifier string) (*Order,error) {
	volumeString := fmt.Sprintf("%v",volume)
	priceString := fmt.Sprintf("%v",price)
	return upbitClient.OrderToMarket(market, "ask", volumeString, priceString, ordType, identifier)
}

