package main

import (
	"fmt"
	"strings"
)

func main() {
	m := GetMediatorIns()
	m.CD.ReadData("music,image")
}

var mediator *Mediator

func GetMediatorIns() *Mediator {
	if mediator == nil {
		mediator = &Mediator{
			&CDDriver{},
			&CPU{},
			&VideoCard{},
			&SoundCard{},
		}
	}
	return mediator
}

type Mediator struct {
	CD    *CDDriver
	CPU   *CPU
	Video *VideoCard
	Sound *SoundCard
}

func (m *Mediator) changed(i interface{}) {
	switch inst := i.(type) {
	case *CDDriver:
		m.CPU.Process(inst.Data)
	case *CPU:
		m.Video.DisPlay(inst.Video)
		m.Sound.Play(inst.Sound)
	}
}

type CDDriver struct {
	Data string
}

func (c *CDDriver) ReadData(data string) {
	c.Data = data
	fmt.Printf("CDDriver: reading data %s \n", c.Data)
	GetMediatorIns().changed(c)
}

type CPU struct {
	Video string
	Sound string
}

func (c *CPU) Process(data string) {
	sp := strings.Split(data, ",")
	c.Sound = sp[0]
	c.Video = sp[1]
	fmt.Printf("CPU: split data with sound: %s, video: %s. \n", c.Sound, c.Video)
	GetMediatorIns().changed(c)
}

type VideoCard struct {
	Data string
}

func (v *VideoCard) DisPlay(data string) {
	v.Data = data
	fmt.Printf("VideoCard play: %s \n", data)
	GetMediatorIns().changed(v)
}

type SoundCard struct {
	Data string
}

func (s *SoundCard) Play(data string) {
	s.Data = data
	fmt.Printf("SoundCard play: %s \n", data)
	GetMediatorIns().changed(s)
}
