package main

import (
	"context"
	"encoding/json"
	"workspace/model"
	"workspace/service"

	//"errors"
	"log"
	"net/http"

	//"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func DecodeAddCabRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request service.Car
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func AddCabEndpoint(svc service.BookingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(service.Car)
		v, err := svc.AddCab(req)
		if err != nil {
			return model.Reply{Msg: "", Error: err.Error()}, err
		}
		return v, nil
	}
}

func DecodeBookCabRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func BookCabEndpoint(svc service.BookingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.Request)
		reply, err := svc.Book(req)
		return reply, err
	}
}
func methodControl(method string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			h.ServeHTTP(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
func main() {

	var manager service.Manager
	//	var cars []*service.Car
	//	manager.Cars = cars
	svc := service.BookingService{BookingManager: manager}
	cabAddHeandler := httptransport.NewServer(
		AddCabEndpoint(svc),
		DecodeAddCabRequest,
		encodeResponse,
	)

	cabBookHandler := httptransport.NewServer(
		BookCabEndpoint(svc),
		DecodeBookCabRequest,
		encodeResponse,
	)

	http.Handle("/addcab", methodControl("POST", cabAddHeandler))
	http.Handle("/bookcab", methodControl("POST", cabBookHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
