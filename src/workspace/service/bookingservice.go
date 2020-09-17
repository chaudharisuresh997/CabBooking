package service

import (
	"errors"
	"fmt"
	"workspace/model"
)

type IBookingService interface {
	Book(model.Request) (model.Reply, error)
	AddCab(Vehicle) (model.Reply, error)
}
type BookingService struct {
	BookingManager
}

func (bookingService BookingService) Book(r model.Request) (model.Reply, error) {
	carObj, response := bookingService.BookingManager.BookCabCommand(r)
	if len(response.Error) > 0 {
		return response, errors.New(response.Error)
	}
	response.Msg = response.Msg + "" + fmt.Sprintf("Car Number : %d", carObj.CarNumber)
	return response, nil

}

func (bookingService BookingService) AddCab(vehicle Vehicle) (model.Reply, error) {
	reply := bookingService.BookingManager.AddCabCommand(vehicle)

	return reply, nil
}
