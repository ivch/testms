package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	client "github.com/ivch/testms/client/service"
)

//ErrorResponse represents error message struct
type ErrorResponse struct {
	Text string
}

//New create new http router for client service
func New(svc client.Service) *mux.Router {
	r := mux.NewRouter()

	r.Methods("GET", "HEAD").Path("/v1/port/{id:[A-Z]+}").HandlerFunc(getEndpoint(svc))
	// more http endpoints

	return r
}

func getEndpoint(svc client.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id := vars["id"]

		if len(id) != 5 {
			encodeHTTPError(errors.New("id should be 5 characters long"), w)
			return
		}

		res, err := svc.Get(id)
		if err != nil {
			encodeHTTPError(err, w)
			return
		}

		if err = encodeHTTPResponse(w, res); err != nil {
			encodeHTTPError(err, w)
		}
	}
}

func encodeHTTPError(err error, w http.ResponseWriter) {
	//there should be added custom error type so we could send http status code regarding to error level
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(ErrorResponse{
		Text: err.Error(),
	})
}

func encodeHTTPResponse(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(response)
}
