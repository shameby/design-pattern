package main

import "fmt"

func main() {
	cs := NewCommonMessage(ViaSMS())
	ce := NewCommonMessage(ViaEmail())
	us := NewUrgencyMessage(ViaSMS())
	ue := NewUrgencyMessage(ViaEmail())

	cs.SendMsg("hello", "wrh")
	ce.SendMsg("hello", "wrh")
	us.SendMsg("hello", "wrh")
	ue.SendMsg("hello", "wrh")
}

type MessageImplementer interface {
	Send(text, to string)
}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}

type CommonMessage struct {
	method MessageImplementer
}

func (m *CommonMessage) SendMsg(text, to string) {
	m.method.Send(text, to)
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
	return &UrgencyMessage{
		method: method,
	}
}

type UrgencyMessage struct {
	method MessageImplementer
}

func (m *UrgencyMessage) SendMsg(text, to string) {
	m.method.Send(fmt.Sprintf("[urgency] %s", text), to)
}

type MessageSMS struct {
}

func ViaSMS() MessageImplementer {
	return &MessageSMS{}
}

func (*MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s by SMS \n", text, to)
}

type MessageEmail struct {
}

func ViaEmail() MessageImplementer {
	return &MessageEmail{}
}

func (*MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s by Email \n", text, to)
}