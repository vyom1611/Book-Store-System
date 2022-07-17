package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Parsing request body to json
func ParseBody(r *http.Request, x interface{}) {
	//Reading the data from the request and unmarshalling it into json format, while saving it to a body slice
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			//If error exists, then this stops the function
			return
		}
	}
}