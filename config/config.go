package config

/*
 * "Inspired" from teacher's init project
 */

import (
	"fmt"
	"os"
)

var conf Config

type Config struct {
	Scheme           string
	Host             string
	Port             string
	MongoURI         string
	MongoDB          string
	MongoCollImage   string
	MongoCollUser    string
	GoogleCredFile   string
	GoogleBucketName string
}

func Init(config Config) Config {
	conf = config
	return conf
}

func Get() (Config, error) {
	if len(conf.MongoURI) == 0 {
		return conf, fmt.Errorf("Config not initialized")
	}
	return conf, nil
}

// unused
func LoadConfig() Config {
	return Config{
		Port:             getConfig("PORT"),
		MongoURI:         getConfig("MONGO_URI"),
		MongoDB:          getConfig("MONGO_DB"),
		MongoCollImage:   getConfig("MONGO_COLL_IMAGE"),
		MongoCollUser:    getConfig("MONGO_COLL_USER"),
		GoogleCredFile:   getConfig("GOOGLE_APPLICATION_CREDENTIALS"),
		GoogleBucketName: getConfig("GOOGLE_APPLICATION_BUCKET"),
	}
}

func getConfig(key string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		panic(fmt.Sprintf("Key %s cannot empty", key))
	}
	return val
}
