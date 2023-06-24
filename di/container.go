package di

import (
	"context"
	"github.com/erfansahebi/lamia_gateway/config"
	"github.com/erfansahebi/lamia_gateway/services"
	"github.com/erfansahebi/lamia_shared/log"
)

type DIContainerInterface interface {
	Config() *config.Config

	Services() services.ServiceContainerInterface
}

type diContainer struct {
	ctx           context.Context
	configuration *config.Config

	services services.ServiceContainerInterface
}

func NewDIContainer(ctx context.Context, config *config.Config) DIContainerInterface {
	return &diContainer{
		ctx:           ctx,
		configuration: config,
	}
}

func (d *diContainer) Config() *config.Config {
	return d.configuration
}

func (d *diContainer) Services() services.ServiceContainerInterface {
	if err := d.initServices(); err != nil {
		log.WithError(err).Fatalf(d.ctx, "error in load service")
		panic(err)
	}

	return d.services
}

func (d *diContainer) initServices() error {
	if d.services != nil {
		return nil
	}

	d.services = services.NewServiceContainer(d.configuration)

	return nil
}
