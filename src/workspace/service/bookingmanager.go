package service

import (
	"log"
	"math"
	"strings"
	"workspace/model"
)

var Cars []*Car
var bookedCars []int

type Manager struct {
}
type BookingManager interface {
	BookCabCommand(model.Request) (*Car, model.Reply)
	AddCabCommand(Vehicle) model.Reply
}

const (
	Booking_Full    string = "Vehicle not available"
	Booking_Success string = "Booking Successfully done"
)

func (m Manager) BookCabCommand(request model.Request) (*Car, model.Reply) {
	var reply model.Reply
	var car *Car
	if len(request.Color) > 0 {
		car, reply = BookByColor(request, Cars)
		return car, reply
	} else {
		return GetNearestCar(request, Cars)
	}
	reply.Msg = Booking_Full

	return &Car{}, reply
}

func (m Manager) AddCabCommand(vehicle Vehicle) model.Reply {
	car, _ := vehicle.(Car)
	Cars = append(Cars, &car)
	return model.Reply{Msg: "Cab registerd successfully", Error: ""}
}
func BookByColor(request model.Request, cars []*Car) (car *Car, reply model.Reply) {
	for i := 0; i < len(cars); i++ {
		if !cars[i].IsBooked() && strings.EqualFold(cars[i].Color, Red) {
			cars[i].Book(true)
			return cars[i], model.Reply{}
		}
	}
	return &Car{}, model.Reply{Error: Booking_Full}
}
func GetNearestCar(request model.Request, allCars []*Car) (*Car, model.Reply) {
	log.Printf("Check all cars")
	minDistance := math.MaxFloat64
	var minCar *Car
	for i := 0; i < len(allCars); i++ {
		log.Printf("searching cars")
		if allCars[i].IsBooked() != true {
			//Calculate distance
			xTerm := math.Pow(float64(request.X-allCars[i].X), 2)
			yTerm := math.Pow(float64(request.Y-allCars[i].Y), 2)
			log.Printf("Smallest distance is now  and the car is %v,%v", minDistance, allCars[i])
			//minDistance = math.Min(minDistance, math.Sqrt(xTerm+yTerm))
			if minDistance > math.Sqrt(xTerm+yTerm) {
				//get minimum car
				minCar = allCars[i]
				minDistance = math.Sqrt(xTerm + yTerm)
			}
			log.Printf("got free car")

		}
	}
	if minCar != nil {
		log.Printf("Booked Car is ==> %v,%v", minDistance, minCar)
		//		bookedCars = append(bookedCars, minCar.CarNumber)
		minCar.Book(true)
		//minCar.Book(true)
		return minCar, model.Reply{Msg: "", Error: ""}
	}
	return nil, model.Reply{Error: Booking_Full}
}
