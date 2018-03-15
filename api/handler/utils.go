package handler

import (
	"api/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// SendJSONWithStatus Encode data into JSON.
func SendJSONWithStatus(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=\"UTF-8\"")
	w.WriteHeader(status)

	if status == http.StatusNoContent {
		return
	}

	if err := json.NewEncoder(w).Encode(&data); err != nil {
		panic(err)
	}
}

func SendJSONOk(w http.ResponseWriter, data interface{}) {
	SendJSONWithStatus(w, data, http.StatusOK)
}

func SendJSONNotFound(w http.ResponseWriter, data interface{}) {
	SendJSONWithStatus(w, data, http.StatusNotFound)
}

func SendJSONError(w http.ResponseWriter, message string, code int) {
	if message == "" {
		message = http.StatusText(code)
	}

	err := model.Error{
		Message: message,
		Code:    code,
	}

	SendJSONWithStatus(w, err, code)
}

func ParseJSON(r *http.Request, data interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(&data)
}

func GetParams(key string, r *http.Request) string {
	return mux.Vars(r)[key]
}
