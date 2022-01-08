package main

import "fmt"

func main() {
	cPay, bPay := NewPayment("wrh", "", 10000, Cash{}), NewPayment("wrh", "62220231000", 56999, Bank{})
	cPay.Pay()
	bPay.Pay()
}

type PaymentStrategy interface {
	Pay(*PaymentContext)
}

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

type PaymentContext struct {
	Name, CardId string
	Money        int64
}

type Cash struct {
}

func (Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash \n", ctx.Money, ctx.Name)
}

type Bank struct {
}

func (Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s \n", ctx.Money, ctx.Name, ctx.CardId)
}

func NewPayment(name, cardId string, money int64, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardId: cardId,
			Money:  money,
		},
		strategy: strategy,
	}
}
