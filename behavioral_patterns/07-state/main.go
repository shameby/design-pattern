package main

import "fmt"

func main() {
	ctx := NewDayCtx()
	for i := 0; i < 21; i++ {
		ctx.Today()
		ctx.Next()
	}
}

func NewDayCtx() *DayContext {
	return &DayContext{
		today: &Sunday{},
	}
}

type WeekDay interface {
	Today()
	Next(*DayContext)
}

type DayContext struct {
	today WeekDay
}

func (d *DayContext) Today() {
	d.today.Today()
}

func (d *DayContext) Next() {
	d.today.Next(d)
}

type Sunday struct {}

func (*Sunday) Today() {
	fmt.Println("Sunday")
}

func (*Sunday) Next(ctx *DayContext) {
	ctx.today = &Monday{}
}

type Monday struct {}

func (*Monday) Today() {
	fmt.Println("Monday")
}

func (*Monday) Next(ctx *DayContext) {
	ctx.today = &Tuesday{}
}

type Tuesday struct {}

func (*Tuesday) Today() {
	fmt.Println("Tuesday")
}

func (*Tuesday) Next(ctx *DayContext) {
	ctx.today = &Wednesday{}
}

type Wednesday struct {}

func (*Wednesday) Today() {
	fmt.Println("Wednesday")
}

func (*Wednesday) Next(ctx *DayContext) {
	ctx.today = &Thursday{}
}

type Thursday struct {}

func (*Thursday) Today() {
	fmt.Println("Thursday")
}

func (*Thursday) Next(ctx *DayContext) {
	ctx.today = &Friday{}
}

type Friday struct {}

func (*Friday) Today() {
	fmt.Println("Friday")
}

func (*Friday) Next(ctx *DayContext) {
	ctx.today = &Saturday{}
}

type Saturday struct {}

func (*Saturday) Today() {
	fmt.Println("Saturday")
}

func (*Saturday) Next(ctx *DayContext) {
	ctx.today = &Sunday{}
}
