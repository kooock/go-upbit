package upbit

import (
	"log"
	"encoding/json"
	"fmt"
	"strconv"
)

func GetTicker(client *UpbitClient,market string, minutes int,count int) ([]Candle,error) {

	url := fmt.Sprintf("https://api.upbit.com/v1/candles/minutes/%v",minutes)
	queries := map[string]string{}
	queries["market"] = market
	queries["count"] = strconv.Itoa(count)
	resp, err := client.SendRequest("GET",url,queries)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	respObj := make([]Candle,count)

	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return respObj, nil
}

func GetOrderbook(client *UpbitClient, market string) (*Orderbook,error) {
	url := "https://api.upbit.com/v1/orderbook"
	respObj := make([]*Orderbook,0)


	queries := map[string]string{}
	queries["markets"] = market
	resp, err := client.SendRequest("GET",url,queries)
	if err != nil {
		log.Fatalln(err)
		return respObj[0], err
	}


	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		log.Fatalln(err)
		return respObj[0], err
	}

	return respObj[0], nil
}
