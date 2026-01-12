package main

import (
	"context"
	"log"

	client "github.com/hoanglam-cynergi/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/hoanglam-cynergi/binance-connector-go/common/common"
)

func main() {
	TestConnectivity()
}

func TestConnectivity() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DerivativesTradingUsdsFuturesRestApiProdUrl),
		common.WithApiKey("Your API Key"),
	)
	apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
		client.WithRestAPI(configuration),
	)
	_, err := apiClient.RestApi.MarketDataAPI.TestConnectivity(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}
}
