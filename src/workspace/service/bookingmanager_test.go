package service

import (
	"reflect"
	"testing"
	"workspace/model"
)

// func TestManager_AddCab(t *testing.T) {
// 	type args struct {
// 		vehicle Vehicle
// 	}
// 	tests := []struct {
// 		name string
// 		m    Manager
// 		args args
// 		want model.Reply
// 	}{}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.m.AddCab(tt.args.vehicle); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Manager.AddCab() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestGetNearestCar(t *testing.T) {
	Cars := []*Car{
		&Car{X: 100, Y: 100, CarNumber: 100},
		&Car{X: 300, Y: 300, CarNumber: 300},
		&Car{X: 1, Y: 1, CarNumber: 1},
	}
	request := model.Request{X: 1, Y: 1}
	type args struct {
		request model.Request
		Cars    []*Car
	}
	tests := []struct {
		name string
		args args
		want *Car
		//want1 model.Reply
	}{
		{name: "Smallest Distance", args: args{request: request, Cars: Cars}, want: &Car{X: 1, Y: 1, CarNumber: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := GetNearestCar(tt.args.request, tt.args.Cars)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNearestCar() got = %v, want %v", got, tt.want)
			}
			// if !reflect.DeepEqual(got1, tt.want1) {
			// 	t.Errorf("GetNearestCar() got1 = %v, want %v", got1, tt.want1)
			// }
		})
	}
}
