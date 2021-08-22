package pusdafil

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func searchMovie(url string, resp interface{}) (err error, executionTime float64) {

	start := time.Now()
	response, err := http.Get(url)
	since := time.Since(start)

	executionTime = float64(since.Microseconds() / 1000)
	if err != nil {
		log.Println(err.Error())
		return err, executionTime
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
		return err, executionTime
	}
	response.Body.Close()

	json.Unmarshal(body, &resp)

	return err, executionTime
}
