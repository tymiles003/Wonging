package wonging

import (
	"fmt"
)

type Player struct {
	id uint8
	//current table the player is sitting at
	table *Table
	//current casino
	casino *Casino
	//current hand, can be split into multiple hands
	hands []*Hand
	//how much is the player betting
	currentBet float64
	//if bought insurance for dealer getting blackjack
	isInsured bool
	//if already doubled
	isDoubled bool
	//how much money does the player have
	totalCash float64

	//TODO: implement in phase 2, for group counting
	//groupId uint8

	//TODO: implement in phase 2, for simulation of getting caught
	//strikes uint8
}

func (p *Player) Initialize(id uint8, c *Casino, t *Table, h *Hand) *Player {
	p.id = id
	p.casino = c
	p.table = t
	p.hands = []*Hand{h}
	p.currentBet = 0
	p.totalCash = DEFAULTPLAYERSTARTINGCASH
	p.isInsured = false
	p.isDoubled = false
	return p
}

func (p *Player) bet(money float64) {
	if money <= 0 || (p.totalCash-money) < 0 {
		fmt.Printf("No more money to make that bet, current cash: %f", p.totalCash)
	} else {
		p.currentBet += money
		p.totalCash -= money
	}
}
func (p *Player) win(money float64) {
	if money < 0 {
		fmt.Println("Use lose() instead!")
	} else {
		p.totalCash += money
		p.currentBet = 0
	}
}
func (p *Player) lose(money float64) {
	if p.totalCash >= money {
		p.totalCash -= money
	} else {
		p.totalCash = 0
	}
}

func (p *Player) changeTable(table *Table) {
	p.table = table
}

func (p *Player) acceptCard(c *Card, handIndex uint8) {
	if handIndex > 0 {
		p.hands[handIndex].AddCard(c)
	} else {
		p.hands[0].AddCard(c)
	}
}
func (p *Player) isNatural() bool {
	return (len(p.hands) == 1 && p.hands[0].isBlackJack())
}

func (p *Player) printPlayer() {
	fmt.Println("Player %d, sitting at table %d, currently betting %f, total cash: %f", p.id, p.table.id, p.currentBet, p.totalCash)
}

//player actions
func (p *Player) hit() {
	if p.currentBet != 0 {

	}
}

func (p *Player) stand() {
	if p.currentBet != 0 {

	}
}

func (p *Player) double() {
	if p.currentBet != 0 && !p.isDoubled {

	}
}

func (p *Player) split() {
	if p.currentBet != 0 {

	}
}

func (p *Player) buyInsurance() {
	if p.currentBet != 0 && !p.isInsured {
		p.bet(p.currentBet / 2)
		p.isInsured = true
	}
}

func (p *Player) isBroke() bool {
	return p.currentBet+p.totalCash == 0
}
