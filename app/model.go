package app

type CoinMarketCap struct {
	Status *Status `json:"status"`
	Data   *[]Coin `json:"data"`
}

type Status struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int    `json:"elapsed"`
	CreditCount  int    `json:"credit_count"`
	Notice       string `json:"notice"`
	TotalCount   int    `json:"total_count"`
}

type Coin struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	Slug              string    `json:"slug"`
	NumMarketPairs    int       `json:"num_market_pairs"`
	Tag               []string  `json:"tags"`
	MaxSupply         float64   `json:"max_supply"`
	CirculatingSupply float64   `json:"circulating_supply"`
	TotalSupply       float64   `json:"total_supply"`
	Platform          *Platform `json:"platform"`
	CmcRank           int       `json:"cmc_rank"`
	LastUpdated       string    `json:"last_updated"`
	Quote             *Quote    `json:"quote"`
}

type Platform struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Slug         string `json:"slug"`
	TokenAddress string `json:"token_address"`
}

type Quote struct {
	USD *MetaData `json:"USD"`
}

type MetaData struct {
	Price            float64 `json:"price"`
	Volume24h        float64 `json:"volume_24h"`
	PercentChange1h  float64 `json:"percent_change_1h"`
	PercentChange24h float64 `json:"percent_change_24h"`
	PercentChange7d  float64 `json:"percent_change_7d"`
	PercentChange30d float64 `json:"percent_change_30d"`
	PercentChange60d float64 `json:"percent_change_60d"`
	PercentChange90d float64 `json:"percent_change_90d"`
	MarketCap        float64 `json:"market_cap"`
	LastUpdated      string  `json:"last_updated"`
}
