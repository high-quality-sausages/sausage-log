package drivers

import (
	"sync"
	"testing"

	"github.com/Shopify/sarama"
)

func TestProducer(t *testing.T) {
	producer := NewProducer()
	t.Log(producer)

	defer producer.Close()

	// 定义一个生产消息，包括Topic、消息内容、
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test"
	msg.Key = sarama.StringEncoder("miles")
	msg.Value = sarama.StringEncoder("hello world...")

	// 发送消息
	pid, offset, err := producer.SendMessage(msg)
	if err != nil {
		t.Error("err:", err)
	}
	t.Log(pid)
	t.Log(offset)
}

func TestConsumer(t *testing.T) {
	var wg sync.WaitGroup
	consumer := NewConsumer()
	t.Log(consumer)

	defer consumer.Close()

	t.Log("connnect success...")
	defer consumer.Close()
	partitions, err := consumer.Partitions("revolution")
	if err != nil {
		t.Error("geet partitions failed, err:", err)
		return
	}


	for _, p := range partitions {
		partitionConsumer, err := consumer.ConsumePartition("test", p, sarama.OffsetOldest)
		if err != nil {
			t.Log("partitionConsumer err:", err)
			continue
		}
		wg.Add(1)
		go func() {
			for m := range partitionConsumer.Messages() {
				t.Logf("key: %s, text: %s, offset: %d\n", string(m.Key), string(m.Value), m.Offset)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
