package services

import (
	"context"
	"github.com/erfansahebi/lamia_gateway/config"
	"github.com/erfansahebi/lamia_gateway/services/auth"
	"github.com/erfansahebi/lamia_gateway/services/shop"
	"github.com/erfansahebi/lamia_shared/go/log"
)

type ServiceContainerInterface interface {
	ServerConfiguration() *config.Config

	Auth() auth.AuthServiceInterface
	Shop() shop.ShopServiceInterface
}

type serviceContainer struct {
	ctx                 context.Context
	serverConfiguration *config.Config

	authService auth.AuthServiceInterface
	shopService shop.ShopServiceInterface
}

func NewServiceContainer(ctx context.Context, serverConfiguration *config.Config) ServiceContainerInterface {
	return &serviceContainer{
		ctx:                 ctx,
		serverConfiguration: serverConfiguration,
	}
}

func (s *serviceContainer) ServerConfiguration() *config.Config {
	return s.serverConfiguration
}

func (s *serviceContainer) Auth() auth.AuthServiceInterface {
	if err := s.initAuth(); err != nil {
		log.WithError(err).Fatalf(s.ctx, "error in load auth service")
		panic(err)
	}

	return s.authService
}

func (s *serviceContainer) initAuth() error {
	if s.authService != nil {
		return nil
	}

	s.authService = auth.NewAuthService(s.ctx, s.serverConfiguration)

	return nil
}

func (s *serviceContainer) Shop() shop.ShopServiceInterface {
	if err := s.initShop(); err != nil {
		log.WithError(err).Fatalf(s.ctx, "error in load shop service")
		panic(err)
	}

	return s.shopService
}

func (s *serviceContainer) initShop() error {
	if s.shopService != nil {
		return nil
	}

	s.shopService = shop.NewShopService(s.ctx, s.serverConfiguration)

	return nil
}
