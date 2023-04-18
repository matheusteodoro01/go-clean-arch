package providers

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaMessagerBroker struct {
	Producer *kafka.Producer
}

func NewKafkaMessagerProducer() *KafkaMessagerBroker {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err == nil {
		panic(err)
	}
	return &KafkaMessagerBroker{Producer: producer}
}

func (producer *KafkaMessagerBroker) Send(msg []byte, topic string) {
	defer producer.Producer.Close()
	producer.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, nil)

	producer.Producer.Flush(15 * 1000)
}
