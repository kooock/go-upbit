package upbit

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (upbitClient *UpbitClient)CheckAccounts() ([]Account,error){
	method := "GET"
	url := "https://api.upbit.com/v1/accounts"

	responseObject := make([]Account,0)
	responseJson, err := upbitClient.SendRequest(method,url,nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseJson, &responseObject)
	if err != nil {
		return nil, err
	}
	return responseObject,nil
}

func (upbitClient *UpbitClient)CheckOrderChance(market string) (*OrderChance,error){
	method := "GET"
	url := "https://api.upbit.com/v1/orders/chance"
	if (market == "default") || (market == "") {
		return nil,errors.New("You have to input market")
	}
	queries := map[string]string{
		"market" : market,
	}
	responseObject := &OrderChance{}
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


func (upbitClient *UpbitClient)CheckOrder(uuid string,identifier string) (*Order, error){
	method := "GET"
	url := "https://api.upbit.com/v1/order"
	queries := map[string]string{}

	if ((uuid == "default") || (uuid == "")) && ((identifier == "default") || identifier == ""){
		return nil, errors.New("You have to input uuid or identifier At least one")
	}


	if (uuid != "default") && (uuid != ""){
		queries["uuid"] = uuid
	}
	if (identifier != "default") && (identifier != ""){
		queries["identifier"] = identifier
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

func (upbitClient *UpbitClient)CheckOrderList(market string,state string, page int, orderBy string) ([]Order, error){
	method := "GET"
	url := "https://api.upbit.com/v1/orders"
	queries := map[string]string{}

	if (market != "default") && (market != ""){
		queries["market"] = market
	}

	if state == "default" || state == ""{
		return nil,errors.New("You have to input state value")
	}
	queries["state"] = state

	if page != 0{
		queries["page"] = fmt.Sprintf("%v",page)
	}

	if (orderBy != "default") && (orderBy != ""){
		queries["order_by"] = orderBy
	}

	responseObject := make([]Order,0)
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

