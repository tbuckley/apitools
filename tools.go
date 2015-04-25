package apitools

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

// WriteJSON writes an object as JSON to the ResponseWriter. It will write
// an error if there is any issue marshaling or writing the data. It also
// sets the application/json content-type.
func WriteJSON(w http.ResponseWriter, obj interface{}) {
	data, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// WriteResponseJSON(w http.ResponseWriter, inteface{})
func WriteResponseJSON(w http.ResponseWriter, obj interface{}) {
	response := Response{Data: obj}
	WriteJSON(w, response)
}

func WriteErrorJSON(w http.ResponseWriter, obj interface{}, statusCode int) {
	response := Response{Error: obj}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	WriteJSON(w, response)
}

// ReadJSON(Reader)
func ReadJSON(r io.Reader, obj interface{}) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, obj)
}
