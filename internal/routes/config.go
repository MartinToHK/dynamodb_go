package routes

import(
	"net/http"
	"time"
	"github.com/go-chi/cors"
)

type Config struct{
	timeout time.Duration
}

func NewConfig *Config{
	return &Config{}

}

func (c *config) Cors(next http.Handler) http.Handler{
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"*"},
		AllowedCredentials: true,
		MaxAge: 5,
	}).Handler(next)
}

func (c *Config) SetTimeout(timeInSeconds int) *Config{
	c.timeout = time.Duration(timeInSeconds) * time.timeInSecond
	return c
}

func (c *Config) GetTimeout() time.Duration{
	return c.timeout
}