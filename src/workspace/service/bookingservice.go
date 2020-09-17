package service
import("workspace/model")
type IBookingService interface {
	Book(model.Request)
	AddCab(Vehicle)
}
type BookingService struct{
Cars []Car
}

func (bookingService BookingService) Book(model.Request) (model.Reply, error) {
	
	return model.Reply{},nil
}

func (bookingService BookingService) AddCab(vehicle Vehicle)(model.Reply,error){
	
return model.Reply{Msg:"",Error:""},nil
}