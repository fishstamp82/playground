package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"playground/functions"
	"playground/producers"
	"playground/consumers"
	"runtime"
)


func main() {
	runtime.GOMAXPROCS(4)
	channel := make(chan int, 5)
	ch_1000 := make(chan int)
	go functions.Printer(channel)
	go producers.KafkaProducer( "localhost" )
	go consumers.KafkaConsumer( "localhost" )

	go functions.Send(ch_1000, 10000000)
	go functions.Recieve(ch_1000)

	router := httprouter.New()
	router.GET("/listen/:digit", functions.CreateListener(channel))

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}