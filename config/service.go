package config

import "fmt"

type ServiceConfiguration struct {
	Host string
	Port string
}

func (sc *ServiceConfiguration) URL() string {
	return fmt.Sprintf("%s:%s", sc.Host, sc.Port)
}
