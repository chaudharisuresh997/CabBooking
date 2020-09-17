package main

import (
	"fmt"
	"workspace/service"
	"workspace/model"
	"log"
)

func main() {
	fmt.Println("Shiv")
	manager:=service.Manager{}
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
}

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
