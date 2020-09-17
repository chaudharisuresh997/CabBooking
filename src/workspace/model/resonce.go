package model

type Request struct {
	Color string
	X     int
	Y int
	FinalX int
	FinalY int
}
type Reply struct {
	Msg   string
	Error string
}