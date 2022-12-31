package helpers

import (
	"encoding/json"
	"net/http"
)

func ReadFromReqBody(r *http.Request, reslut interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(reslut)
	if err != nil {
		panic(err)
	}
}

func WriteToReqBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		panic(err)
	}
}
