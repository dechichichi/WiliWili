package utils

import (
	"github.com/segmentio/kafka-go"
)

var KafkaClientGlobal *kafka.Writer

func InitKafkaClient(addr string, topic string) {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(addr),
		Topic:        topic,
		RequiredAcks: kafka.RequireOne,
		Balancer:     &kafka.Hash{},
		BatchSize:    1,
	}
	KafkaClientGlobal = writer
}
