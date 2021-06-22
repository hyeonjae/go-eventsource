package eventbus

import (
	"math/rand"
	"time"

	"github.com/Shopify/sarama"
	eventsource "github.com/hyeonjae/go-eventsource"
)

type EventBus interface {
	Send()
	Receive()
}

type Kafka struct {
	kafkaCfg *sarama.Config
	consumer sarama.Consumer
	producer sarama.SyncProducer
}

func New(cfg *eventsource.Config) (*Kafka, error) {
	rand.Seed(time.Now().UnixNano())

	var kafkaCfg = sarama.NewConfig()
	kafkaCfg.Version = sarama.V2_2_0_0
	kafkaCfg.Producer.Return.Successes = true
	kafkaCfg.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	kafkaCfg.Producer.Flush.Frequency = 200 * time.Millisecond // Flush batches every 200ms
	kafkaCfg.Net.DialTimeout = 300 * time.Millisecond
	kafkaCfg.Net.ReadTimeout = 300 * time.Millisecond
	kafkaCfg.Net.WriteTimeout = 300 * time.Millisecond

	producer, err := sarama.NewSyncProducer(cfg.Kafka.Addresses, kafkaCfg)
	if err != nil {
		return nil, err
	}

	consumer, err := sarama.NewConsumer(cfg.Kafka.Addresses, kafkaCfg)
	if err != nil {
		return nil, err
	}

	return &Kafka{
		kafkaCfg: kafkaCfg,
		producer: producer,
		consumer: consumer,
	}, nil
}

func (k Kafka) Send(topic string, message sarama.Encoder) (int32, int64, error) {
	msg := sarama.ProducerMessage{
		Topic:     topic,
		Value:     message,
		Timestamp: time.Now(),
	}
	return k.producer.SendMessage(&msg)
}

func (k Kafka) Receive(topic string) (<-chan *sarama.ConsumerMessage, <-chan *sarama.ConsumerError, error) {
	partitions, _ := k.consumer.Partitions(topic)
	consumer, err := k.consumer.ConsumePartition(topic, partitions[0], sarama.OffsetNewest)
	return consumer.Messages(), consumer.Errors(), err
}
