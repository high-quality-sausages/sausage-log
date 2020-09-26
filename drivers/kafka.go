package drivers

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func NewProducer() sarama.SyncProducer {
	config := sarama.NewConfig()

	// WaitForAll waits for all in-sync replicas to commit before responding.
	config.Producer.RequiredAcks = sarama.WaitForAll

	// NewRandomPartitioner returns a Partitioner which chooses a random partition each time.
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	config.Producer.Return.Successes = true

	// 新建一个同步生产者
	producer, err := sarama.NewSyncProducer([]string{"kafka-master:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return nil
	}

	return producer
}

func NewConsumer() sarama.Consumer {
	consumer, err := sarama.NewConsumer([]string{"kafka-master:9092"}, nil)
	if err != nil {
		fmt.Println("consumer connect error:", err)
		return nil
	}

	return consumer
	// 	fmt.Println("connnect success...")
	// 	defer consumer.Close()
	// 	partitions, err := consumer.Partitions("revolution")
	// 	if err != nil {
	// 		fmt.Println("geet partitions failed, err:", err)
	// 		return nil
	// 	}

	// 	for _, p := range partitions {
	// 		partitionConsumer, err := consumer.ConsumePartition("revolution", p, sarama.OffsetOldest)
	// 		if err != nil {
	// 			fmt.Println("partitionConsumer err:", err)
	// 			continue
	// 		}
	// 		wg.Add(1)
	// 		go func() {
	// 			for m := range partitionConsumer.Messages() {
	// 				fmt.Printf("key: %s, text: %s, offset: %d\n", string(m.Key), string(m.Value), m.Offset)
	// 			}
	// 			wg.Done()
	// 		}()
	// 	}
	// 	wg.Wait()
}
