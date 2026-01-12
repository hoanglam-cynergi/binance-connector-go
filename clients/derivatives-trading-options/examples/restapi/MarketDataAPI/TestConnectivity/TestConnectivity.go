package main

import (
	"context"
	"log"

	client "github.com/hoanglam-cynergi/binance-connector-go/clients/derivativestradingoptions"
	"github.com/hoanglam-cynergi/binance-connector-go/common/common"
)

func main() {
	TestConnectivity()
}

func TestConnectivity() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DerivativesTradingOptionsRestApiProdUrl),
		common.WithApiKey("Your API Key"),
	)
	apiClient := client.NewBinanceDerivativesTradingOptionsClient(
		client.WithRestAPI(configuration),
	)
	_, err := apiClient.RestApi.MarketDataAPI.TestConnectivity(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}
}
