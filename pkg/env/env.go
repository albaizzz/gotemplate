package env

import (
	"os"
)

const (
	//Env environment key
	Env = "ENV"
	//EnvDevelopment environment development
	EnvDevelopment = "dev"
	//EnvStaging environment staging
	EnvStaging = "staging"
	//EnvProduction environment production
	EnvProduction = "prod"
)

//Get gets environment by key
func Get(key string) string {
	return os.Getenv(key)
}
