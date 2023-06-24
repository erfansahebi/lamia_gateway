package services

import (
	"context"
	"github.com/erfansahebi/lamia_gateway/config"
	"github.com/erfansahebi/lamia_gateway/services/auth"
	"github.com/erfansahebi/lamia_shared/log"
)

type ServiceContainerInterface interface {
	ServerConfiguration() *config.Config

	Auth() auth.AuthServiceInterface
}

type serviceContainer struct {
	serverConfiguration *config.Config

	authService auth.AuthServiceInterface
}

func NewServiceContainer(serverConfiguration *config.Config) ServiceContainerInterface {
	return &serviceContainer{
		serverConfiguration: serverConfiguration,
	}
}

func (s *serviceContainer) ServerConfiguration() *config.Config {
	return s.serverConfiguration
}

func (s *serviceContainer) Auth() auth.AuthServiceInterface {
	if err := s.initAuth(); err != nil {
		log.WithError(err).Fatalf(context.Background(), "error in load auth service")
		panic(err)
	}

	return s.authService
}

func (s *serviceContainer) initAuth() error {
	if s.authService != nil {
		return nil
	}

	s.authService = auth.NewAuthService(s.serverConfiguration)

	return nil
}
