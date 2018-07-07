package playground

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"time"
	"fmt"
	"os"
	"io"
	"strconv"
)



func Printer(c <- chan int){
	for {
		a := <-c
		println(a)
		time.Sleep(1)
	}
}


func CreateListener(c chan<- int) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, d httprouter.Params) {
		val, err := strconv.Atoi(d.ByName("digit"))
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		c <- val
		io.WriteString(w, "hello, world!\n")
	}
}



func Index(w http.ResponseWriter, r *http.Request, d httprouter.Params) {
	println("asdf")
	println("asdf")
	println("asdf")
	println("asdf")
	println("asdf")
	println("asdf")
	println("asdf")
	println("asdf")
	println("asdf")
	println("asdf")
}