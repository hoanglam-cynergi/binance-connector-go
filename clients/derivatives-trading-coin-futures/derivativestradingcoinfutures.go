package derivativestradingcoinfutures

import (
	BinanceDerivativesTradingCoinFuturesRestApi "github.com/hoanglam-cynergi/binance-connector-go/clients/derivativestradingcoinfutures/src/restapi"
	BinanceDerivativesTradingCoinFuturesWebsocketApi "github.com/hoanglam-cynergi/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketapi"
	BinanceDerivativesTradingCoinFuturesWebsocketStreams "github.com/hoanglam-cynergi/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams"
	"github.com/hoanglam-cynergi/binance-connector-go/common/common"
)

type BinanceDerivativesTradingCoinFuturesClient struct {
	RestApi          *BinanceDerivativesTradingCoinFuturesRestApi.RestAPIClient
	WebsocketAPI     *BinanceDerivativesTradingCoinFuturesWebsocketApi.WebsocketAPIClient
	WebsocketStreams *BinanceDerivativesTradingCoinFuturesWebsocketStreams.WebsocketStreamsClient
}

type Option func(*BinanceDerivativesTradingCoinFuturesClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceDerivativesTradingCoinFuturesClient) {
		c.RestApi = BinanceDerivativesTradingCoinFuturesRestApi.NewRestAPIClient(cfg)
	}
}

func WithWebsocketAPI(cfg *common.ConfigurationWebsocketApi) Option {
	return func(c *BinanceDerivativesTradingCoinFuturesClient) {
		c.WebsocketAPI = BinanceDerivativesTradingCoinFuturesWebsocketApi.NewWebsocketAPIClient(cfg)
	}
}

func WithWebsocketStreams(cfg *common.ConfigurationWebsocketStreams) Option {
	return func(c *BinanceDerivativesTradingCoinFuturesClient) {
		c.WebsocketStreams = BinanceDerivativesTradingCoinFuturesWebsocketStreams.NewWebsocketStreamsClient(cfg)
	}
}

func NewBinanceDerivativesTradingCoinFuturesClient(opts ...Option) *BinanceDerivativesTradingCoinFuturesClient {
	client := &BinanceDerivativesTradingCoinFuturesClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
