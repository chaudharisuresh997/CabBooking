package service

const (
	Red   string = "red"
	Blue  string = "blue"
	Black string = "black"
	Pink  string = "pink"
)

type Car struct {
	CarNumber int
	isbooked  bool
	X         int
	Y         int
	FinalX    int
	FinalY    int
	Color     string
}

type Vehicle interface {
	IsBooked() bool
	Move()
}

func (car Car) IsBooked() bool {
	return car.isbooked
}

func (car *Car) Book(flag bool) {
	car.isbooked = flag
}

func (car Car) Move() {
}

// public void drive(int x,int y,int finalx,int finaly){
// this.occupy();
// if(x>finalx){
// int temp=finalx;
// int finalx=x;
// int x=temp;
// }
// if(x<finalx){
// x++
// }
// this.release();
// }
// public void occupy(){
// this.isAvailable=false;
// }
