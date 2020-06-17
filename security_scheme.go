package gopenapi

// SecurityScheme Object
type SecurityScheme struct {
	Type             string      `yaml:"bearerFormat"`
	Description      string      `yaml:"description"`
	Name             string      `yaml:"name"`             // TODO: Applies To "apiKey" then REQUIRED
	In               string      `yaml:"in"`               // TODO: Applies To "apiKey" then REQUIRED
	Scheme           string      `yaml:"scheme"`           // TODO: Applies To "http" then REQUIRED
	BearerFormat     string      `yaml:"bearerFormat"`     //
	Flows            *OAuthFlows `yaml:"flows"`            // TODO: Applies To "http" then REQUIRED
	OpenIDConnectURL string      `yaml:"openIdConnectUrl"` // TODO: Applies To "openIdConnect" then REQUIRED
}
