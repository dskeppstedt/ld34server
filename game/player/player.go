package player

import (
	"math"
)

type Player struct {
	Pos   Position
	Angle float32
}

type Position struct {
	X float32
	Y float32
}

func (this *Player) Update() {
	this.move(5)
}
func (this *Player) AngleInc(amt float32) {
	this.Angle += amt
}

func (this *Player) move(speed float32) {
	this.Pos.X += speed * float32(math.Cos(float64(this.Angle)*math.Pi/180))
	this.Pos.Y += speed * float32(math.Sin(float64(this.Angle)*math.Pi/180))
}

type PlayerManager struct {
	Players []Player
}

func (this *PlayerManager) Update() {
	for _, player := range this.Players {
		player.Update()
	}
}

func (this *PlayerManager) CreatePlayer() Player {
	pos := &Position{700, 398}
	p := &Player{Pos: *pos}
	this.Players = append(this.Players, *p)
	return *p
}
