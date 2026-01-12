package spot

import (
	BinanceSpotRestApi "github.com/hoanglam-cynergi/binance-connector-go/clients/spot/src/restapi"
	BinanceSpotWebsocketApi "github.com/hoanglam-cynergi/binance-connector-go/clients/spot/src/websocketapi"
	BinanceSpotWebsocketStreams "github.com/hoanglam-cynergi/binance-connector-go/clients/spot/src/websocketstreams"
	"github.com/hoanglam-cynergi/binance-connector-go/common/common"
)

type BinanceSpotClient struct {
	RestApi          *BinanceSpotRestApi.RestAPIClient
	WebsocketAPI     *BinanceSpotWebsocketApi.WebsocketAPIClient
	WebsocketStreams *BinanceSpotWebsocketStreams.WebsocketStreamsClient
}

type Option func(*BinanceSpotClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceSpotClient) {
		c.RestApi = BinanceSpotRestApi.NewRestAPIClient(cfg)
	}
}

func WithWebsocketAPI(cfg *common.ConfigurationWebsocketApi) Option {
	return func(c *BinanceSpotClient) {
		c.WebsocketAPI = BinanceSpotWebsocketApi.NewWebsocketAPIClient(cfg)
	}
}

func WithWebsocketStreams(cfg *common.ConfigurationWebsocketStreams) Option {
	return func(c *BinanceSpotClient) {
		c.WebsocketStreams = BinanceSpotWebsocketStreams.NewWebsocketStreamsClient(cfg)
	}
}

func NewBinanceSpotClient(opts ...Option) *BinanceSpotClient {
	client := &BinanceSpotClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
