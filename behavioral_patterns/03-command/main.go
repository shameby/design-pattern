package main

import "fmt"

func main() {
	mb := &MotherBoard{}
	box := &Box{
		button1: &StartCommand{mb: mb},
		button2: &RebootCommand{mb: mb},
	}
	box.PressButton1()
	box.PressButton2()
}

type MotherBoard struct{}

func (m *MotherBoard) Start() {
	fmt.Println("system starting")
}

func (m *MotherBoard) Reboot() {
	fmt.Println("system rebooting")
}

type Command interface {
	Execute()
}

type StartCommand struct {
	mb *MotherBoard
}

func (s *StartCommand) Execute() {
	s.mb.Start()
}

type RebootCommand struct {
	mb *MotherBoard
}

func (r *RebootCommand) Execute() {
	r.mb.Reboot()
}

type Box struct {
	button1 Command
	button2 Command
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}
