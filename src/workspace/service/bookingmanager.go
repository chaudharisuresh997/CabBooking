package service
import("workspace/model"
"strings")
type Manager struct {
}
type BookingManager interface {
	Book(model.Request)
	AddCab(Vehicle)
}
const(Booking_Full string="Vehicle not available"
Booking_Success string="Booking Successfully done")

var cars []Car

func (m Manager) Book(request model.Request) (Car,model.Reply) {
	var reply model.Reply
	var car Car
	if(len(request.Color)>0){
		car,reply=BookByColor(request);
		return car,reply
	}else{
		// for i := 0; i < len(cars); i++ {
		// 	if !cars[i].IsBooked() {
		// 		return model.Reply{Msg:"", Error:""}
		// 	}
		// }
	}
	reply.Msg=Booking_Full

	return Car{},reply
}

func (m Manager) AddCab(vehicle Vehicle)(model.Reply){
	car,_:=vehicle.(Car)
	cars=append(cars,car)
return model.Reply{Msg:"",Error:""}
}
func BookByColor(request model.Request)(car Car,reply model.Reply){
	for i := 0; i < len(cars); i++ {
		if(!cars[i].IsBooked() && strings.EqualFold(cars[i].Color,Red)){
			cars[i].Book(true)
			return cars[i],model.Reply{}
		}
	}
return Car{},model.Reply{Msg:Booking_Full}
}