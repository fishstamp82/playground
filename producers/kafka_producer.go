package producers

import "time"

func KafkaProducer(host string)  {
	for {
		println("Kafka Producer Sleeping for 5 seconds")
		time.Sleep(5 * 1000000000 )
	}
}
