package shop

import (
	"context"
	"github.com/erfansahebi/lamia_gateway/config"
	"github.com/erfansahebi/lamia_shared/go/log"
	shopProto "github.com/erfansahebi/lamia_shared/go/proto/shop"
	"google.golang.org/grpc"
)

type ShopServiceInterface interface {
	Configuration() *config.ServiceConfiguration

	Client() shopProto.ShopServiceClient
}

type shopService struct {
	ctx                 context.Context
	serverConfiguration *config.Config

	configuration *config.ServiceConfiguration
	client        shopProto.ShopServiceClient
}

func NewShopService(ctx context.Context, serverConfiguration *config.Config) ShopServiceInterface {
	return &shopService{
		ctx:                 ctx,
		serverConfiguration: serverConfiguration,
	}
}

func (s *shopService) Configuration() *config.ServiceConfiguration {
	if err := s.initConfiguration(); err != nil {
		log.WithError(err).Fatalf(s.ctx, "error in load configurations in shop service")
		panic(err)
	}

	return s.configuration
}

func (s *shopService) initConfiguration() error {
	if s.configuration != nil {
		return nil
	}

	s.configuration = &config.ServiceConfiguration{
		Host: s.serverConfiguration.Services.Shop.Host,
		Port: s.serverConfiguration.Services.Shop.Port,
	}

	return nil
}

func (s *shopService) Client() shopProto.ShopServiceClient {
	if err := s.initClient(); err != nil {
		log.WithError(err).Fatalf(s.ctx, "error in init shop service")
		panic(err)
	}

	return s.client
}

func (s *shopService) initClient() error {
	if s.client != nil {
		return nil
	}

	cc, err := grpc.Dial(s.Configuration().URL(), grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Fatalf(s.ctx, "error in create client connection with shop service")
		return err
	}

	s.client = shopProto.NewShopServiceClient(cc)

	return nil
}
