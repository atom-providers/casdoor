package casdoor

import (
	"github.com/atom-providers/cert"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config Config
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}
	return container.Container.Provide(func(cert *cert.Cert) *casdoorsdk.Client {
		certificate := config.Certificate
		if certificate == "" {
			certificate = cert.Cert
		}

		config := casdoorsdk.AuthConfig{
			Endpoint:         config.Endpoint,
			ClientId:         config.ClientId,
			ClientSecret:     config.ClientSecret,
			Certificate:      certificate,
			OrganizationName: config.OrganizationName,
			ApplicationName:  config.ApplicationName,
		}

		return casdoorsdk.NewClientWithConf(&config)
	}, o.DiOptions()...)
}
