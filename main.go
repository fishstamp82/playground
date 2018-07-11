package main

import (
	"net/http"
	"playground/functions"
	"playground/producers"
	"github.com/julienschmidt/httprouter"
)

func main() {

	channelToKafkaProducer := make(chan int, 250)
	router := httprouter.New()

	router.GET("/publish/:bucket_id", functions.CreateBucketListener(channelToKafkaProducer))
	go producers.KafkaProducer(channelToKafkaProducer)

	println("Running server on localhost:8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
