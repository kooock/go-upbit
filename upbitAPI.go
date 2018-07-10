package upbit

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"encoding/json"
)


type UpbitClient struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}



func NewUpbitClient() (*UpbitClient, error){

	goPath := os.Getenv("GOPATH")

	configPath := path.Join(goPath,"src/github.com/kooock/upbit-go/client.json")
	println(configPath)
	config,err := os.Open(configPath)
	defer config.Close()
	if err != nil {
		return nil, err
	}

	byteValue, _ := ioutil.ReadAll(config)
	upbitClient := &UpbitClient{}
	json.Unmarshal(byteValue, upbitClient)
	println(upbitClient.AccessKey)
	println(upbitClient.SecretKey)
	return upbitClient, nil
}

func CreateRequest(method string,url string, queries map[string]string, tokenString string) (*http.Request,error){
	req, err := http.NewRequest(method,url,nil)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	authKey := fmt.Sprintf("Bearer %v",tokenString)

	if queries != nil{
		q := req.URL.Query()
		for key,value := range queries {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}
	req.Header.Add("Authorization", authKey)
	return req, err
}



func (upbitClient *UpbitClient)SendRequest(method string, url string, queries map[string]string) ([]byte,error){
	jwtToken := NewJWTtoken(queries, upbitClient.AccessKey)
	token, err := jwtToken.CreateTokenString(upbitClient.SecretKey)
	if err != nil {
		log.Fatalln(err)
		return nil,err
	}
	request, err := CreateRequest(method,url,queries,token)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
		return nil,err
	}
	defer resp.Body.Close()
	// 결과 출력
	bytes, _ := ioutil.ReadAll(resp.Body)
	return bytes,nil
}

