package config

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func Config() {
	goth.UseProviders(
		GoogleConfig(),
	)
}

func GoogleConfig() *google.Provider {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURL := os.Getenv("GOOGLE_REDIRECT_URL")

	return google.New(
		clientID,
		clientSecret,
		redirectURL,
		"email",
		"profile",
	)
}
