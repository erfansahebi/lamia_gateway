package auth

import (
	"context"
	"github.com/erfansahebi/lamia_gateway/config"
	"github.com/erfansahebi/lamia_shared/log"
	authProto "github.com/erfansahebi/lamia_shared/services/auth"
	"google.golang.org/grpc"
)

type AuthServiceInterface interface {
	Configuration() *config.ServiceConfiguration

	Client() authProto.AuthServiceClient
}

type authService struct {
	serverConfiguration *config.Config

	configuration *config.ServiceConfiguration
	client        authProto.AuthServiceClient
}

func NewAuthService(serverConfiguration *config.Config) AuthServiceInterface {
	return &authService{
		serverConfiguration: serverConfiguration,
	}
}

func (s *authService) Configuration() *config.ServiceConfiguration {
	if err := s.initConfiguration(); err != nil {
		log.WithError(err).Fatalf(context.Background(), "error in load configurations in auth service")
		panic(err)
	}

	return s.configuration
}

func (s *authService) initConfiguration() error {
	if s.configuration != nil {
		return nil
	}

	s.configuration = &config.ServiceConfiguration{
		Host: s.serverConfiguration.Services.Auth.Host,
		Port: s.serverConfiguration.Services.Auth.Port,
	}

	return nil
}

func (s *authService) Client() authProto.AuthServiceClient {
	if err := s.initClient(); err != nil {
		log.WithError(err).Fatalf(context.Background(), "error in init auth service")
		panic(err)
	}

	return s.client
}

func (s *authService) initClient() error {
	if s.client != nil {
		return nil
	}

	cc, err := grpc.Dial(s.Configuration().URL(), grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Fatalf(context.Background(), "error in create client connection with auth service")
		return err
	}

	s.client = authProto.NewAuthServiceClient(cc)

	return nil
}
