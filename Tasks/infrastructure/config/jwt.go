package config

type JwtKey struct {
	PubPath string `json:"pub_path" env:"JWT_PUBLIC"`
	PriPath string `json:"pri_path" env:"JWT_PRIVATE"`
}
