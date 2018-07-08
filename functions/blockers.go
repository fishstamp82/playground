package functions

func Send(ch chan<- int, d int) {
	for i := 0; i < d; i++ {
		ch <- i
	}}

func Recieve(ch <-chan int) {
	for  {
		i := <-ch
		if i%100000 == 0 {
			println(i)
		}
	}

}