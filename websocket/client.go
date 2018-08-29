package websocket

import (
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
	"encoding/json"
	. "time"
	"log"
)
/*
[{"ticket":"test"},{"format":"SIMPLE"},{"type":"trade","codes":["KRW-BTC","BTC-BCC"]},{"format":"SIMPLE"}]
*/

type Field interface {}

type TicketField struct {
	Ticket string `json:"ticket"`
}

type FormatField struct {
	Format string `json:"format"`
}

type TypeField struct {
	Type string `json:"type"`
	Code []string `json:"codes"`
}

type WebSocketClient struct {
	Url string
	Connection *websocket.Conn
	Message *ResponseMessage
}

func CreateWebSocketClient() *WebSocketClient{
	u := url.URL{Scheme: "wss", Host: "api.upbit.com", Path: "/websocket/v1"}
	return &WebSocketClient{Url:u.String(),Message: &ResponseMessage{}}
}

func (webSocketClient * WebSocketClient)Connect() (err error){
	webSocketClient.Connection, _, err = websocket.DefaultDialer.Dial(webSocketClient.Url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	return err
}

func CreateAndConnectWebSocket()(*WebSocketClient,error){
	client := CreateWebSocketClient()
	err := client.Connect()
	return client, err
}

func (webSocketClient *WebSocketClient)RunStream(ticket string, format string, infoType string, codes []string){


	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	messageRequested := []Field{TicketField{Ticket:ticket},FormatField{Format:format},TypeField{Type:infoType,Code:codes}}

	messageSended, err := json.Marshal(messageRequested)
	if err != nil {
		panic(err)
	}
	webSocketClient.Connection.WriteMessage(websocket.TextMessage,messageSended)

	for {
		Sleep(Millisecond)
		_, messageReceived, err := webSocketClient.Connection.ReadMessage()
		if err != nil {
			println("read:", err)
			continue
		}

		err = json.Unmarshal(messageReceived,webSocketClient.Message)
		if err != nil {
			println(err)
			continue
		}

	}
}

func (websocketClient *WebSocketClient)Close() error{
	err := websocketClient.Connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}
	err = websocketClient.Connection.Close()
	if err != nil {
		return err
	}
	return nil
}