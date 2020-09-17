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

/*	manager:=service.Manager{}
	car:=service.Car{
		CarNumber:1,
		X:1,
		Y:1,
		FinalX:5,
		FinalY:5,
	}
	reply:=manager.AddCab(car)
	log.Printf("Booked status %v",reply)
	request:=model.Request{
		X:1,
		Y:1,
		FinalX:7,
		FinalY:7,
		Color:"pink",
	}
	manager.AddCab(car)
	manager.Book(request)
*/
// class User{
// int x;
// int y;
// int fx;

// public void requestCab(int x,int y){
// Cab b=carpool.getFreeCab(x,y);
// }
// }

// public class CarpoolManager{
// List<Car> cars=new ArrayList<>();
// public void addCar(Car c){
// cars.add(c);
// }
// public Car getClosest(int x,int y){
// int mindiff=Integer.MAX_VALUE;
// Car minddistanceCar=null;
// for(Car car:cars){
// if(!car.isAvailable()) continue;
// int distance=Math.abs(car.x-x)+Math.abs(car.y-y);
// if(mindiff>distance){
// minddistanceCar=car;
// }

// }
// return minddistanceCar;

// }
// public Car getFreeCab(x,y){
// if(!pq.isEmpty()){
// return pq.poll();
// }
// else{return null;}
// }

// }
