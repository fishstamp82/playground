package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"playground/functions"
	"playground/producers"
	"playground/consumers"
)


func main() {
	channel := make(chan int)
	go functions.Printer(channel)
	go producers.KafkaProducer( "localhost" )
	go consumers.KafkaConsumer( "localhost" )

	router := httprouter.New()
	router.GET("/listen/:digit", functions.CreateListener(channel))

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}