package service
import("workspace/model"
"strings")
type Manager struct {
	Cars []Car
}
type BookingManager interface {
	Book(model.Request)
	AddCab(Vehicle)
}
const(Booking_Full string="Vehicle not available"
Booking_Success string="Booking Successfully done")


func (m Manager) Book(request model.Request) (Car,model.Reply) {
	var reply model.Reply
	var car Car
	if(len(request.Color)>0){
		car,reply=BookByColor(request,m.Cars);
		return car,reply
	}else{
		for i := 0; i < len(m.Cars); i++ {
			if !m.Cars[i].IsBooked() {
				return m.Cars[i],model.Reply{Msg:"", Error:""}
			}
		}
	}
	reply.Msg=Booking_Full

	return Car{},reply
}

func (m Manager) AddCab(vehicle Vehicle)(model.Reply){
	car,_:=vehicle.(Car)
	m.Cars=append(m.Cars,car)
return model.Reply{Msg:"",Error:""}
}
func BookByColor(request model.Request,cars []Car)(car Car,reply model.Reply){
	for i := 0; i < len(cars); i++ {
		if(!cars[i].IsBooked() && strings.EqualFold(cars[i].Color,Red)){
			cars[i].Book(true)
			return cars[i],model.Reply{}
		}
	}
return Car{},model.Reply{Msg:Booking_Full}
}