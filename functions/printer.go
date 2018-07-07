package functions

func Printer(c <- chan int){
	for {
		a := <-c
		println(a)
	}
}
