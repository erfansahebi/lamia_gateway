package auth

import (
	"github.com/erfansahebi/lamia_gateway/config"
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
		return err
	}

	s.client = authProto.NewAuthServiceClient(cc)

	return nil
}
