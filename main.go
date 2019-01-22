package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const (
	Apikey = "7qPl7fYevNrFsIDBJo8fObZ0C3hCRoIFqsXD4YbccH9SQ37ETKKKHiTZdGaL"
	Url    = "https://www.worldtradingdata.com/api/v1/stock?symbol="
)

type Responsestock struct {
	Symbol               string `json:"symbol"`
	Name                 string `json:"name"`
	Currency             string `json:"currency"`
	Price                string `json:"price"`
	Priceopen            string `json:"price_open"`
	Dayhigh              string `json:"day_high"`
	Daylow               string `json:"day_low"`
	week_high            string `json:"52_week_high" `
	week_low             string `json:"52_week_low" `
	Day_change           string `json:"day_change" `
	Change_pct           string `json:"change_pct"`
	Close_yesterday      string `json:"close_yesterday"`
	Market_cap           string `json:"market_cap"`
	Volume               string `json:"volume"`
	Volume_avg           string `json:"volume_avg"`
	Shares               string `json:"shares"`
	Stock_exchange_long  string `json:"stock_exchange_long"`
	Stock_exchange_short string `json:"stock_exchange_short"`
	Timezone             string `json:"timezone"`
	Timezone_name        string `json:"timezone_name"`
	Gmt_offset           string `json:"gmt_offset"`
	Last_trade_time      string `json:"last_trade_time"`
}

type StockExchangeResponse struct {
	Symbol          string `json:"symbol"`
	Name            string `json:"name"`
	Price           string `json:"price"`
	Close_yesterday string `json:"close_yesterday"`
	Currency        string `json:"currency"`
	Market_cap      string `json:"market_cap"`
	Volume          string `json:"volume"`
	Timezone        string `json:"timezone"`
	Timezone_name   string `json:"timezone_name"`
	Gmt_offset      string `json:"gmt_offset"`
	Last_trade_time string `json:"last_trade_time"`
}
type Stockresponse struct {
	Symbols_requested int             `json:"symbols_requested"`
	Symbols_returned  int             `json:"symbols_returned"`
	Data              []Responsestock `json:"data"`
}

func GetStockHandler(w http.ResponseWriter, req *http.Request) {
	var stockexange []string
	vars := mux.Vars(req)
	data := vars["symbol"]
	fmt.Println(data)
	keys, ok := req.URL.Query()["stock_exchange"]
	if !ok || len(keys[0]) < 1 {
		fmt.Println("Url Param 'key' is missing")
		data = "AMEX"
	} else {
		stockexange = strings.Split(keys[0], ",")
	}
	url := Url + data + "&api_token=" + Apikey
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in the URL:Response Error", err)
	}
	respbody, err := ioutil.ReadAll(resp.Body)
	if data == "AMEX" ||  {
		w.Write(respbody)
	}
	var response Stockresponse
	err = json.Unmarshal(respbody, &response)
	if err != nil {
		fmt.Print("Cannot Unmarshal", err)
	}
	Exchangeresposne := make(map[string]interface{})
	for _, k := range stockexange {
		var Stockres []StockExchangeResponse
		for _, v := range response.Data {

			if v.Stock_exchange_short == k {
				obj := StockExchangeResponse{}
				obj.Symbol = v.Symbol
				obj.Name = v.Name
				obj.Close_yesterday = v.Close_yesterday
				obj.Currency = v.Currency
				obj.Gmt_offset = v.Gmt_offset
				obj.Last_trade_time = v.Last_trade_time
				obj.Market_cap = v.Market_cap
				obj.Timezone = v.Timezone
				obj.Timezone_name = v.Timezone_name
				obj.Price = v.Price
				obj.Volume = v.Volume
				_, exist := Exchangeresposne[v.Stock_exchange_short]

				if exist {
					Stockres = append(Stockres, obj)

					fmt.Println("stockresponse", len(Stockres))
					Exchangeresposne[v.Stock_exchange_short] = Stockres

				} else {
					Stockres = append(Stockres, obj)
					Exchangeresposne[v.Stock_exchange_short] = Stockres

				}

			}
		}

	}
	SEresponse, err := json.Marshal(Exchangeresposne)
	w.Write(SEresponse)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/stock/{symbol}", GetStockHandler)
	http.ListenAndServe(":8080", r)
}
