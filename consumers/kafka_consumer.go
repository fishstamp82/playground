package consumers

import "time"

func KafkaConsumer(host string)  {
	for {
		println("Kafka Consumer Sleeping for 5 seconds")
		time.Sleep(5 * 1000000000 )
	}
}

