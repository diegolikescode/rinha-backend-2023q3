package config

import "rinha-backend-2023q3/pkg"

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Cache struct {
	Host string
	Port string
}

type Server struct {
	Port     string
	UseSonic bool
	Prefork  bool
}

type Profiling struct {
	Enabled bool
	Cpu     string
	Mem     string
}

type Config struct {
	Database  Database
	Cache     Cache
	Server    Server
	Profiling Profiling
}

func NewConfig() *Config {
	return &Config{
		Database{
			Host:     pkg.GetEnvOrDieTrying("DATABASE_HOST"),
			Port:     pkg.GetEnvOrDieTrying("DATABASE_PORT"),
			User:     pkg.GetEnvOrDieTrying("DATABASE_USER"),
			Password: pkg.GetEnvOrDieTrying("DATABASE_PASSWORD"),
			Name:     pkg.GetEnvOrDieTrying("DATABASE_NAME"),
		},
		Cache{
			Host: pkg.GetEnvOrDieTrying("CACHE_HOST"),
			Port: pkg.GetEnvOrDieTrying("CACHE_PORT"),
		},
		Server{
			Port:     pkg.GetEnvOrDieTrying("SERVER_PORT"),
			UseSonic: pkg.GetEnvOrDieTrying("ENABLE_SONIC_JSON") == "1",
			Prefork:  pkg.GetEnvOrDieTrying("ENABLE_PREFORK") == "1",
		},
		Profiling{
			Enabled: pkg.GetEnvOrDieTrying("ENABLE_PROFILING") == "1",
			Cpu:     pkg.GetEnvOrDieTrying("CPU_PROFILE"),
			Mem:     pkg.GetEnvOrDieTrying("MEM_PROFILE"),
		},
	}
}
