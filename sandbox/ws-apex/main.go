package main

import (
	"context"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type Coin struct {
	Symbol            string `json:"s"`
	LastPrice         string `json:"p"`
	DailyChange       string `json:"pr"`
	DailyHighestPrice string `json:"h"`
	DailyLowestPrice  string `json:"l"`
	MarkPrice         string `json:"mp"`
	IndexPrice        string `json:"xp"`
	TurnOver          string `json:"to"`
	TradingVolume     string `json:"v"`
	FundingRate       string `json:"fr"`
	OpenInterest      string `json:"o"`
	TradeCount        string `json:"tc"`
	Timestamp         int64  `json:"ts"`
}

func main() {
	u := url.URL{Scheme: "wss", Host: "quote.omni.apex.exchange", Path: "/realtime_public"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}

	defer c.Close()
	c.WriteJSON(map[string]any{"op": "subscribe", "args": []string{"instrumentInfo.all"}})
	_, cancel := context.WithCancel(context.Background())

	defer cancel()

	for {
		var msg struct {
			Topic string `json:"topic"`
			Data  []Coin `json:"data"`
			Type  string `json:"type"`
			TS    int64  `json:"ts"`
		}

		err := c.ReadJSON(&msg)
		if err != nil {
			log.Print("reading error:", err)
		}

		for _, value := range msg.Data {
			log.Printf("Websocket message")
			log.Println(value)
		}
	}
}
