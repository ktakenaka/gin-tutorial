package config

import (
	"context"
	"sync"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	once            sync.Once
	googleOauth2Cnf *oauth2.Config
	googleOIDCCnf   *oidc.Config
	googleProvider  *oidc.Provider
)

const (
	jwksURL   = "https://www.googleapis.com/oauth2/v3/certs"
	issuerURL = "https://accounts.google.com"
)

// TODO: refactor considering to use pkg

// InitGoogleOIDCCnf initializes OIDC configuration
func InitGoogleOIDCCnf(redirectURL, clientID, clientSecret string) {
	once.Do(func() {
		googleOauth2Cnf = &oauth2.Config{
			RedirectURL:  redirectURL,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
			Endpoint:     google.Endpoint,
		}
		googleOIDCCnf = &oidc.Config{
			ClientID: clientID,
		}
		provider, err := oidc.NewProvider(context.TODO(), issuerURL)
		if err != nil {
			panic(err)
		}
		googleProvider = provider
	})
}

// GetGoogleOauthConfig returns oauth2.conf for Google
func GetGoogleOauthConfig() *oauth2.Config {
	return googleOauth2Cnf
}

// GetGoogleVerifier returns oidc.IDTokenVerifier for Google
func GetGoogleVerifier(ctx context.Context) *oidc.IDTokenVerifier {
	keySet := oidc.NewRemoteKeySet(ctx, jwksURL)
	verifier := oidc.NewVerifier(issuerURL, keySet, googleOIDCCnf)
	return verifier
}

// GetGoogleProvider returns oidc.Provider for Google
func GetGoogleProvider() *oidc.Provider {
	return googleProvider
}
