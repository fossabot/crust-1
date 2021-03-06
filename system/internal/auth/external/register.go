package external

import (
	"context"
	"fmt"

	"github.com/crusttech/go-oidc"
)

func RegisterNewOpenIdClient(ctx context.Context, eas *externalAuthSettings, name, url string) (eap *externalAuthProvider, err error) {
	var (
		provider    *oidc.Provider
		client      *oidc.Client
		redirectUrl = fmt.Sprintf(eas.redirectUrl, "openid-connect."+name)
	)

	if provider, err = oidc.NewProvider(ctx, url); err != nil {
		return
	}

	client, err = provider.RegisterClient(ctx, &oidc.ClientRegistration{
		Name:          "Crust",
		RedirectURIs:  []string{redirectUrl},
		ResponseTypes: []string{"token id_token", "code"},
	})

	if err != nil {
		return
	}

	eap = &externalAuthProvider{
		redirectUrl: redirectUrl,
		key:         client.ID,
		secret:      client.Secret,
		issuerUrl:   url,
	}

	return
}
