package config

import "flag"

type Config struct {
	Host      string
	Port      string
	StaticDir string
}

func NewConfig() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.Host, "host", "localhost", "usage <localhost>'")
	flag.StringVar(&cfg.Port, "port", "8080", "usage :<port>")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()
	return cfg
}

// go run cmd/app/main.go  --port 8900
