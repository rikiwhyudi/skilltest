package main

import (
	"e-market/pkg/mysql"
	"fmt"
	"math/rand"
)

// Dice represents a dice.
type Dice struct {
	topSideVal int
}

// GetTopSideVal returns the top side value of the dice.
func (d *Dice) GetTopSideVal() int {
	return d.topSideVal
}

// Roll rolls the dice.
func (d *Dice) Roll() *Dice {
	d.topSideVal = rand.Intn(6) + 1
	return d
}

// SetTopSideVal sets the top side value of the dice.
func (d *Dice) SetTopSideVal(topSideVal int) *Dice {
	d.topSideVal = topSideVal
	return d
}

// Player represents a player.
type Player struct {
	diceInCup []*Dice
	name      string
	position  int
	point     int
}

// GetDiceInCup returns the dice in the cup of the player.
func (p *Player) GetDiceInCup() []*Dice {
	return p.diceInCup
}

// GetName returns the name of the player.
func (p *Player) GetName() string {
	return p.name
}

// GetPosition returns the position of the player.
func (p *Player) GetPosition() int {
	return p.position
}

// AddPoint adds the given point to the player's point.
func (p *Player) AddPoint(point int) {
	p.point += point
}

// GetPoint returns the player's point.
func (p *Player) GetPoint() int {
	return p.point
}

// Play rolls the dice in the player's cup.
func (p *Player) Play() {
	for _, dice := range p.diceInCup {
		dice.Roll()
	}
}

// RemoveDice removes the dice at the given index from the player's cup.
func (p *Player) RemoveDice(index int) {
	p.diceInCup = append(p.diceInCup[:index], p.diceInCup[index+1:]...)
}

// InsertDice inserts the given dice to the player's cup.
func (p *Player) InsertDice(dice *Dice) {
	p.diceInCup = append(p.diceInCup, dice)
}

// Game represents a game.
type Game struct {
	players               []*Player
	round                 int
	numberOfPlayer        int
	numberOfDicePerPlayer int
}

const (
	removedWhenDiceTop = 6
	moveWhenDiceTop    = 1
)

// NewGame creates a new game with the given number of players and dices per player.
func NewGame(numberOfPlayer, numberOfDicePerPlayer int) *Game {
	game := &Game{
		round:                 0,
		numberOfPlayer:        numberOfPlayer,
		numberOfDicePerPlayer: numberOfDicePerPlayer,
	}

	// Create players and dices
	for i := 0; i < game.numberOfPlayer; i++ {
		player := &Player{
			diceInCup: make([]*Dice, game.numberOfDicePerPlayer),
			position:  i,
			name:      string('A' + i),
		}
		for j := 0; j < game.numberOfDicePerPlayer; j++ {
			player.diceInCup[j] = &Dice{}
		}
		game.players = append(game.players, player)
	}

	return game
}

// DisplayRound displays the current round.
func (g *Game) DisplayRound() *Game {
	fmt.Printf("Round %d\n", g.round)
	return g
}

// PlayRound plays a round of the game.
func (g *Game) PlayRound() *Game {
	// Display the current round
	g.DisplayRound()

	// Play each player
	for _, player := range g.players {
		player.Play()
	}

	// Remove dice with top value of 6
	for _, player := range g.players {
		for i := 0; i < len(player.diceInCup); i++ {
			if player.diceInCup[i].GetTopSideVal() == removedWhenDiceTop {
				player.RemoveDice(i)
				i--
			}
		}
	}

	// Move the player with dice top value of 1
	for _, player := range g.players {
		for _, dice := range player.diceInCup {
			if dice.GetTopSideVal() == moveWhenDiceTop {
				player.position++
				break
			}
		}
	}

	// Award points to the player at the end of the round
	for _, player := range g.players {
		if player.position == g.numberOfPlayer {
			player.AddPoint(1)
		}
	}

	// Remove player with no dice left
	for i := 0; i < len(g.players); i++ {
		if len(g.players[i].diceInCup) == 0 {
			g.players = append(g.players[:i], g.players[i+1:]...)
			i--
		}
	}

	// Increment the round
	g.round++

	return g
}

func main() {
	// Migrataions database models
	mysql.DatabaseInit()

	// Create a new game with 2 players and 2 dices per player
	game := NewGame(3, 4)

	// Play 3 rounds of the game
	for i := 0; i < 3; i++ {
		game.PlayRound()
	}

	// Find the player with the highest score
	var winner *Player
	for _, player := range game.players {
		if winner == nil || player.GetPoint() > winner.GetPoint() {
			winner = player
		}
	}

	// Display the final score of each player and the winner
	for _, player := range game.players {
		fmt.Printf("Player %s: %d point(s)\n", player.GetName(), player.GetPoint())
	}
	fmt.Printf("The winner is Player %s with %d point(s)!\n", winner.GetName(), winner.GetPoint())
}
