package main

import (
	"context"
	"log"

	client "github.com/hoanglam-cynergi/binance-connector-go/clients/margintrading"
	"github.com/hoanglam-cynergi/binance-connector-go/common/common"
)

func main() {
	CloseUserDataStream()
}

func CloseUserDataStream() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.MarginTradingRestApiProdUrl),
		common.WithApiKey("Your API Key"),
	)
	apiClient := client.NewBinanceMarginTradingClient(
		client.WithRestAPI(configuration),
	)
	_, err := apiClient.RestApi.RiskDataStreamAPI.CloseUserDataStream(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}
}
