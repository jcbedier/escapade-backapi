package main

import "os"

type Config struct {
	MongoLogin string
	MongoPass  string
	MongoDB    string
	MongoHost  string
}

func NewConfiguration() *Config {
	return &Config{
		MongoLogin: os.Getenv("ENV_MOGO_LOGIN"),
		MongoPass:  os.Getenv("ENV_MOGO_PASS"),
		MongoDB:    os.Getenv("ENV_MOGO_DB"),
	}
}
