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
	Port               string
	MongoURI           string
	MongoDB            string
	GoogleCredFile     string
	GoogleBucketName   string
	KafkaBrokers       string
	KafkaConsumerGroup string
	KafkaTopic         string
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

func LoadConfig() Config {
	conf := Config{
		Port:               getConfig("PORT", "8090"),
		MongoURI:           getConfig("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:            getConfig("MONGO_DB", "gosocial_db"),
		GoogleCredFile:     getConfig("GOOGLE_APPLICATION_CREDENTIALS", "gcs.json"),
		GoogleBucketName:   getConfig("GOOGLE_APPLICATION_BUCKET", "ct-go-social"),
		KafkaBrokers:       getConfig("KAFKA_BROKERS", "localhost:9093"),
		KafkaConsumerGroup: getConfig("KAFKA_CONSUMER_GROUP", "group_notify"),
		KafkaTopic:         getConfig("KAFKA_TOPIC", "like_event"),
	}
	return conf
}

func getConfig(key string, defaultVal string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		fmt.Printf("Env %s is empty, fallback to default value %s", key, defaultVal)
		val = defaultVal
	}
	return val
}
