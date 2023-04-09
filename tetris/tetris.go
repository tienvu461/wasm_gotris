/*
Copyright © 2023 tienvu461@gmail.com
*/
package tetris

import (
	"fmt"
	"time"
)

const B_HEIGHT = 20
const B_WIDTH = 15

type gameState int

const (
	G_INIT gameState = iota
	G_PLAY
	G_OVER
)

type game struct {
	board     [][]int
	position  vector
	block     block
	state     gameState
	score     int
	FallSpeed *time.Timer
}

func (g *game) genBlock() {
	g.block = randBlock()
	g.position = vector{1, B_WIDTH / 2}
}

func (g *game) GetScore() int {
	return g.score
}

func (g *game) resetFallSpeed() {
	g.FallSpeed.Reset(700 * time.Millisecond)
}

func (g *game) Start() {
	g.state = G_PLAY
	g.genBlock()
	g.resetFallSpeed()
}

func (g *game) Quit() {
	g.state = G_OVER
}

func (g *game) colision() bool {
	for _, v := range g.block.shape {
		pos := g.blockOnBoardByPosition(v)
		if pos.x < 0 || pos.x >= B_WIDTH {
			return true
		}
		if pos.y < 0 || pos.y >= B_HEIGHT {
			return true
		}
		if g.board[pos.y][pos.x] > 0 {
			return true
		}
	}
	return false
}

func (g *game) moveIfPosible(v vector) bool {
	g.position.x += v.x
	g.position.y += v.y
	if g.colision() {
		g.position.x -= v.x
		g.position.y -= v.y
		return false
	}
	fmt.Println(g.position)
	return true
}

func (g *game) MoveLeft(unit ...int) {
	default_unit := 1
	if len(unit) > 0 {
		default_unit = unit[0]
	}
	g.moveIfPosible(vector{0, -(default_unit)})
}

func (g *game) MoveRight(unit ...int) {
	default_unit := 1
	if len(unit) > 0 {
		default_unit = unit[0]
	}
	g.moveIfPosible(vector{0, default_unit})
}

func (g *game) MoveDown(unit ...int) {
	default_unit := 1
	if len(unit) > 0 {
		default_unit = unit[0]
	}
	g.moveIfPosible(vector{default_unit, 0})
}

func (g *game) MoveUp(unit ...int) {
	default_unit := 1
	if len(unit) > 0 {
		default_unit = unit[0]
	}
	g.moveIfPosible(vector{-(default_unit), 0})
}

func (g *game) SpeedUp() {
	g.FallSpeed.Reset(50 * time.Millisecond)
}

func (g *game) Rotate() {
	g.block.rotate()
	// TODO: handle exception rotate will crash on border
	if g.colision() {
		// g.block.rotateBack()
		ymin, _, xmin, xmax := g.block.ShapeMinMax()
		switch {
		case xmax+g.position.x > B_WIDTH-1:
			fmt.Printf("xmax = %d\n", xmax)
			g.MoveLeft(xmax + g.position.x - B_WIDTH + 1)
		case g.position.x+xmin < 0:
			g.MoveRight(g.position.x - xmin)
		case g.position.y+ymin < 0:
			fmt.Printf("ymin = %d\n", ymin)
			g.MoveDown(g.position.y + ymin + 2)
		default:
			g.block.rotateBack()
		}
	}
}

func (g *game) Fall() {
	for {
		if !g.moveIfPosible(vector{1, 0}) {
			g.FallSpeed.Reset(1 * time.Millisecond)
			return
		}
	}
}
func (g *game) lockBlocks() {
	g.board = g.GetBoard()
}
func (g *game) clearLine() {
	line := make([]int, B_WIDTH)
	for i := 0; i < B_WIDTH; i++ {
		line[i] = 0
	}
	clearLine := [][]int{line}
	for y := 0; y < B_HEIGHT; y++ {
		for x := 0; x < B_WIDTH; x++ {
			if g.board[y][x] == 0 {
				break
			} else if x == B_WIDTH-1 {
				newBoard := append(clearLine, g.board[:y]...)
				g.board = append(newBoard, g.board[y+1:]...)
				g.score += 100
			}
		}
	}
}
func (g *game) GameLoop() {
	if !g.moveIfPosible(vector{1, 0}) {
		g.lockBlocks()
		g.clearLine()
		g.genBlock()
		if g.colision() {
			g.FallSpeed.Stop()
			g.state = G_OVER
			return
		}
	}
	g.resetFallSpeed()
}

func (g *game) blockOnBoardByPosition(v vector) vector {
	px := g.position.x + v.x
	py := g.position.y + v.y
	return vector{py, px}
}

// return current board
func (g *game) GetBoard() [][]int {
	cBoard := make([][]int, len(g.board))
	for y := 0; y < len(g.board); y++ {
		cBoard[y] = make([]int, len(g.board[y]))
		copy(cBoard[y], g.board[y])
		// for x := 0; x < len(g.board[y]); x++ {
		// 	cBoard[y][x] = g.board[y][x]
		// }
	}

	for _, v := range g.block.shape {
		pos := g.blockOnBoardByPosition(v)
		cBoard[pos.y][pos.x] = g.block.color
	}
	return cBoard
}

func (g *game) GetState() gameState {
	return g.state
}

func (g *game) init() {
	// initialize 2d array
	g.board = make([][]int, B_HEIGHT)
	for y := 0; y < B_HEIGHT; y++ {
		g.board[y] = make([]int, B_WIDTH)
		for x := 0; x < B_WIDTH; x++ {
			g.board[y][x] = 0
		}
	}
	g.position = vector{0, B_WIDTH / 2}
	// g.block = blocks[0]
	g.FallSpeed = time.NewTimer(time.Duration(1000 * time.Second))
	g.FallSpeed.Stop()
	g.state = G_INIT
}

func NewGame() *game {
	g := &game{}
	g.init()
	return g
}
