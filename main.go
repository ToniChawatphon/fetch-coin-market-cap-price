package main

import (
	"github.com/3TS-services/toni-coin-marketcap-api/app"
)

func main() {
	app.Init()

	// Run Mannually
	app.Run.FetchCoinsMarketValue()

	// Run Agent
	app.Agent.Schedule(4, app.Hour, app.Agent.Service.FetchCoinsMarketValue)
	app.Agent.Start()
}
