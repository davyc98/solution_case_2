package movie

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"log"
)

func SearchMovie(url string, resp interface{}) (executionTime float64, err error) {

	start := time.Now()
	response, err := http.Get(url)
	since := time.Since(start)

	executionTime = float64(since.Microseconds() / 1000)
	if err != nil {
		log.Println(err.Error())
		return executionTime, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
		return executionTime, err
	}
	response.Body.Close()

	json.Unmarshal(body, &resp)

	return executionTime, err
}
