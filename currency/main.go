package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	ratesURL = "https://openexchangerates.org/api/latest.json?app_id="
)

type Response struct {
	Rates struct {
		TRY float32 `json:"TRY"`
		XAU float32 `json:"XAU"`
		EUR float32 `json:"EUR"`
	} `json:"rates"`
}

func main() {

	apiKey := os.Getenv("EXCHANGE_API")
	if apiKey == "" {
		log.Fatal("Please provide EXCHANGE_API key in environment. To get the key please visit https://openexchangerates.org")
	}

	ratesURL = fmt.Sprintf("%v%v", ratesURL, apiKey)

	resp, err := http.Get(ratesURL)
	if err != nil {
		log.Fatalf("Error in request %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Response code %v", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error while reading body %v", err)
	}
	var rates Response
	if err = json.Unmarshal(respBody, &rates); err != nil {
		log.Fatalf("error while unmarshalling")
	}

	euroInTRY := (1 / rates.Rates.EUR) * rates.Rates.TRY
	xauInTRY := (1 / rates.Rates.XAU)

	fmt.Printf("USD: %v\n", rates.Rates.TRY)
	fmt.Printf("EUR: %v\n", euroInTRY)
	fmt.Printf("GOLD: %v\n", xauInTRY)

}
