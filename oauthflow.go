package gopenapi

// OAuthFlow Object.
type OAuthFlow struct {
	AuthorizationURL string            `yaml:"authorizationUrl"` // TODO: Applies To "oauth2 ("implicit", "authorizationCode")" when REQUIRED
	TokenURL         string            `yaml:"tokenUrl"`         // TODO: Applies To "oauth2 ("password", "clientCredentials", "authorizationCode")" when REQUIRED
	RefreshURL       string            `yaml:"refreshUrl"`       //
	Scopes           map[string]string `yaml:"tokenUrl"`         // TODO: Applies To "oauth2" when REQUIRED
}
