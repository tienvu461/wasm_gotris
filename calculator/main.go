package calculator

import (
	"fmt"
	"syscall/js"
)

func calculateSquare(x int) int {
	return x * x
}

func calculateCube(x int) int {
	return x * x * x
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

func calculateCubeWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		x := args[0].Int()
		return calculateCube(x)
	})
}

func main() {
	js.Global().Set("calculateCube", calculateCubeWrapper())
	js.Global().Set("calculateSquare", calculateSquareWrapper())

	fmt.Println("Welcome to calculator")
	// make sure the program doesn't exit
	<-make(chan bool)
}
