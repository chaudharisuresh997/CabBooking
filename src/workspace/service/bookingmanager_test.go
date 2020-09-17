package service

import (
	"reflect"
	"testing"
	"workspace/model"
)

func TestManager_AddCab(t *testing.T) {
	type args struct {
		vehicle Vehicle
	}
	tests := []struct {
		name string
		m    Manager
		args args
		want model.Reply
	}{
	
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.AddCab(tt.args.vehicle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.AddCab() = %v, want %v", got, tt.want)
			}
		})
	}
}
