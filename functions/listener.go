package functions

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// CreateBucketListener is designed to
func CreateBucketListener(channelToKafkaProducer chan<- int) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, d httprouter.Params) {

		strDigit := d.ByName("bucket_id")
		intDigit, err := strconv.Atoi(strDigit)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		io.WriteString(w, fmt.Sprintf("Sending bucket_id %d to kafka producer", intDigit))
		channelToKafkaProducer <- intDigit

		
	}
}
