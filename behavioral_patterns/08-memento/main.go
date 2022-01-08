package main

import "fmt"

func main() {
	game := &Game{50, 50}
	m := game.Save()
	game.Play(10, 10)
	game.Status()
	game.Load(m)
	game.Status()
}

type Memento interface {
}

type gameMemento struct {
	hp, mp int64
}

type Game struct {
	hp, mp int64
}

func (g *Game) Play(mpDelta, hpDelta int64) {
	g.mp += mpDelta
	g.hp += hpDelta
}

func (g *Game) Save() Memento {
	return &gameMemento{
		hp: g.hp,
		mp: g.mp,
	}
}

func (g *Game) Load(m Memento) {
	gm := m.(*gameMemento)
	g.mp = gm.mp
	g.hp = gm.hp
}

func (g *Game) Status() {
	fmt.Printf("Current HP: %d, MP: %d. \n", g.hp, g.mp)
}


