package config

type JwtKey struct {
	PubPath string `required:"true" envconfig:"JWT_PUBLIC"`
	PriPath string `required:"true" envconfig:"JWT_PRIVATE"`
}
