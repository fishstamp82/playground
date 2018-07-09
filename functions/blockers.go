package functions

import "time"

func Send(ch chan<- int, d int) {
	for i := 0; i < d; i++ {
		ch <- i
	}}

func Recieve(ch <-chan int) {
	for  {
		i :=<-ch
		time.Sleep(1e9)
		println(i)
	}

}