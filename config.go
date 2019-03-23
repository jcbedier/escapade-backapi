package main

import "os"

type MyConfig struct {
	MongoLogin string
	MongoPass  string
	MongoDB    string
	MongoHost  string
}

func NewConfiguration() *MyConfig {
	return &MyConfig{
		MongoLogin: os.Getenv("ENV_MOGO_LOGIN"),
		MongoPass:  os.Getenv("ENV_MOGO_PASS"),
		MongoDB:    os.Getenv("ENV_MOGO_DB"),
	}
}
