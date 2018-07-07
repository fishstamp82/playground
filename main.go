package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"playground/functions"
)


func main() {
	channel := make(chan int)
	go playground.Printer(channel)

	router := httprouter.New()
	router.GET("/listen/:digit", playground.CreateListener(channel))
	router.GET("/index", playground.Index)

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}