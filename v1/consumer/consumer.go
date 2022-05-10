package consumer

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/yueja/cf-kafka/config"
	"github.com/yueja/cf-kafka/errors"
	"go-common/library/log"
	"sync"
)

// Consumer 消费者结构
type Consumer struct {
	Address []string
	Config  *sarama.Config
	Topic   string
	Group   string
	Handler sarama.ConsumerGroupHandler
	Version sarama.KafkaVersion
}

// NewConsumer 创建消费者
func NewConsumer(addresses []string, topic string, group string, h sarama.ConsumerGroupHandler) *Consumer {
	if len(addresses) == 0 {
		log.Error("address err:%+v", errors.ErrNoKafkaServerAddress)
		panic(errors.ErrNoKafkaServerAddress)
	}

	if topic == "" {
		log.Error("topic err:%+v", errors.ErrNoCanConsumerTopic)
		panic(errors.ErrNoCanConsumerTopic)
	}

	if group == "" {
		log.Error("group err:%+v", errors.ErrNoKafkaConsumerGroup)
		panic(errors.ErrNoKafkaConsumerGroup)
	}

	consumer := &Consumer{
		Address: addresses,
		Config:  config.DefaultConsumerConfig(),
		Topic:   topic,
		Group:   group,
		Handler: h,
	}
	return consumer
}

// Consume 进行消费 ctx context.Context, wg *sync.WaitGroup
func (c *Consumer) Consume(superCtx context.Context, superWg *sync.WaitGroup) {
	defer superWg.Done()
	ctx, cancel := context.WithCancel(superCtx)
	cg, err := sarama.NewConsumerGroup(c.Address, c.Group, c.Config)
	if err != nil {
		log.Error("consume err:%+v", err)
		panic(err)
	}
	consumerWg := &sync.WaitGroup{}
	consumerWg.Add(1)
	go func() {
		defer func() {
			consumerWg.Done()
			if e := recover(); e != nil {
				log.Error("consume recover e:%+v", e)
			}
		}()
		for {
			// 触发reBalance时重连
			if err := cg.Consume(ctx, []string{c.Topic}, c.Handler); err != nil {
				log.Error("cg.consume  consume:%+v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				log.Error("terminating：拉取下来的信息处理完成")
				return
			}
		}
	}()

	<-superCtx.Done()
	log.Info("terminating: %s context cancelled", c.Topic)
	// 有中断信号，就关闭当前的上下文，通知当前消费者停止拉取
	cancel()
	// 等待消费者处理完已拉取下来的信息
	consumerWg.Wait()
	if err = cg.Close(); err != nil {
		log.Error("cg.Close:%+v", err)
	}
	log.Info("terminating: %s context cancelled:%+v", c.Topic)
}
