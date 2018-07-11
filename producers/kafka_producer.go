package producers

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"fmt"
)

// BucketCreated is a shit
type BucketCreated struct {
	eventType string
	bucketID int
}

// KafkaProducer produces and event to kafka based on the value supplied by the channel its listening on
func KafkaProducer(channelToKafkaProducer <-chan int)  {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	for {
		val := <-channelToKafkaProducer
		go func() {
			for e := range p.Events() {
				switch ev := e.(type) {
				case *kafka.Message:
					if ev.TopicPartition.Error != nil {
						fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
					} else {
						fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
					}
				}
			}
			
		}()
		// Produce messages to topic (asynchronously)
		topic := "bucket_service"
		bucketCreated := BucketCreated{eventType: "bucket_created", bucketID: val}
		fmt.Println( bucketCreated )

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value: []byte(fmt.Sprintf("%v", bucketCreated )),
		}, nil)
		

		// Wait for message deliveries
		p.Flush(1000)
	}
}
