package go_eventsource

type Config struct {
	Kafka Kafka
	Mongo Mongo
}

type Kafka struct {
	Addresses []string
}

type Mongo struct {
	Addresses []string
	Username string
	Password string
	Database string
}

func NewConfig() *Config {
	return &Config{
		Kafka: Kafka{
			Addresses: []string{
				":9092",
			},
		},
		Mongo: Mongo{
			Addresses: []string{
				"mongodb://localhost:27017",
			},
			Username: "dev",
			Password: "eventsource!@#",
			Database: "eventsource",
		},
	}
}
