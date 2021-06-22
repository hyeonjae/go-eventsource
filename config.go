package go_eventsource

type Config struct {
	Kafka *Kafka
	Mongo *Mongo
	MySQL *MySQL
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

type MySQL struct {
	Address  string
	Username string
	Password string
	Database string
}

func NewConfig() *Config {
	return &Config{
		Kafka: &Kafka{
			Addresses: []string{
				":9092",
			},
		},
		Mongo: &Mongo{
			Addresses: []string{
				"mongodb://localhost:27017",
			},
			Username: "dev",
			Password: "eventsource!@#",
			Database: "eventsource",
		},
		MySQL: &MySQL{
			Address: "localhost:3306",
			Username: "dev",
			Password: "eventsource",
			Database: "eventsource",
		},
	}
}
