package config

import "fmt"

type Redis struct {
	Host string `json:"host" env:"REDIS_HOST"`
	Port int    `json:"port" env:"REDIS_PORT"`
}

// URL return redis connection URL.
func (c *Redis) URL() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)
}
