package tetris

import (
	"fmt"
	"math/rand"
	"syscall/js"
	"time"
)

const FPS = 60 * time.Millisecond

type browserKeyCode int

const (
	BKC_Space      browserKeyCode = 32
	BKC_ArrowLeft  browserKeyCode = 38
	BKC_ArrowUp    browserKeyCode = 39
	BKC_ArrowRight browserKeyCode = 40
	BKC_ArrowDown  browserKeyCode = 41
	BKC_KeyQ       browserKeyCode = 81
	BKC_KeyS       browserKeyCode = 83
)

func calculateSquare(x int) int {
	return x * x
}

func calculateSquareWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		x := args[0].Int()
		return calculateSquare(x)
	})
}

func (g *game) move(key browserKeyCode) bool {
	fmt.Println(key)
	switch {
	case key == BKC_ArrowUp:
		g.Rotate()
	case key == BKC_ArrowDown:
		g.SpeedUp()
	case key == BKC_ArrowLeft:
		g.MoveLeft()
	case key == BKC_ArrowRight:
		g.MoveRight()
	case key == BKC_Space:
		g.Fall()
	case key == BKC_KeyS:
		g.Start()
	case key == BKC_KeyQ:
		g.Quit()
	default:
		return false
	}
	return true
}

func (g *game) MoveWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		x := args[0].Int()
		return g.move(browserKeyCode(x))
	})
}

func (g *game) GetBoardWrapper() js.Func {
	fmt.Println("getting board")
	board := g.GetBoard()
	//TODO: Optimize
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		x := args[0].Int()
		arr := make([]interface{}, len(board[x]))
		// fmt.Println(board[x])
		for i := 0; i < len(arr); i++ {
			arr[i] = board[x][i]
		}
		return arr
	})
}

func (g *game) GetStateWrapper() js.Func {
	fmt.Println("getting state")
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		return g.GetState()
	})
}

func GameInit() {

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to GoTris")
	// make sure the program doesn't exit

	ticker := time.NewTimer(time.Duration(FPS))
	game := NewGame()
	js.Global().Set("move", game.MoveWrapper())
	js.Global().Set("getBoard", game.GetBoardWrapper())
	js.Global().Set("getState", game.GetStateWrapper())

	for game.GetState() != G_OVER {
		select {
		case <-ticker.C:
			ticker.Reset(FPS)
		case <-game.FallSpeed.C:
			game.GameLoop()
		}
	}
	<-make(chan bool)
}
