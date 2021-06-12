package app

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

type API struct {
	client *http.Client
	Coins  *CoinMarketCap
}

func (a *API) Connect() *http.Client {
	a.client = &http.Client{}
	return a.client
}

func (a *API) GetMarketPrice(uri string) (*http.Response, *http.Request) {
	client := a.Connect()
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", Setting.ApiKey)
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		log.Printf("%v Error sending request to server %v\n", res.StatusCode, err.Error())
		os.Exit(1)
	}

	log.Println("- status", res.Status)

	return res, req
}
