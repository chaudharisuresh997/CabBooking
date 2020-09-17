package service

import (
	"log"
	"strings"
	"workspace/model"
)

var Cars []*Car

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
		log.Printf("Check all cars")
		for i := 0; i < len(Cars); i++ {
			log.Printf("searching cars")
			if !Cars[i].IsBooked() {
				log.Printf("got free car")
				return Cars[i], model.Reply{Msg: "", Error: ""}
			}
		}
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
