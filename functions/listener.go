package functions

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"fmt"
	"os"
	"io"
)

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

func CreateKafkaProducerListener(c chan<- []byte) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, d httprouter.Params) {
		val, err := strconv.Atoi(d.ByName("digit"))
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		c <- []byte(val) 
		io.WriteString(w, "Sent Message to kafka!\n")
	}
}


