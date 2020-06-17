package gopenapi

// OAuthFlows Object
type OAuthFlows struct {
	Implicit          *OAuthFlow `yaml:"implicit"`
	Password          *OAuthFlow `yaml:"password"`
	ClientCredentials *OAuthFlow `yaml:"clientCredentials"`
	AuthorizationCode *OAuthFlow `yaml:"authorizationCode"`
}
