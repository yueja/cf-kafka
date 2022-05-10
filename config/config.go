package config

import "github.com/Shopify/sarama"

// DefaultConfig 默认配置
func DefaultConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	return config
}

// DefaultProducerConfig 默认配置
func DefaultProducerConfig() *sarama.Config {
	kfkConfig := DefaultConfig()
	kfkConfig.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	kfkConfig.Producer.Return.Successes = true
	kfkConfig.Producer.MaxMessageBytes = 10000000 //最大消息支持10M，默认10M
	return kfkConfig
}

// DefaultConsumerConfig 默认配置
func DefaultConsumerConfig() *sarama.Config {
	config := DefaultConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Fetch.Max = 10000000
	return config
}
