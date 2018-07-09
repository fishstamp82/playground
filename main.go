package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"playground/functions"
	"runtime"
	"playground/producers"
	"playground/consumers"
)


func main() {
	runtime.GOMAXPROCS(1)
	channel := make(chan int, 4)
	channel_to_kafka_producer := make(chan []byte)

	go functions.Printer(channel)
	go producers.KafkaProducer( channel_to_kafka_producer )
	go consumers.KafkaConsumer( "localhost" )


	router := httprouter.New()
	router.GET("/listen/:digit", functions.CreateListener(channel))
	router.GET("/kafka/:digit", functions.CreateKafkaProducerListener(channel))

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}