package arcade

import (
	"fmt"

	"github.com/chr-ras/advent-of-code-2019/13-care-package/intcode"
	q "github.com/enriquebris/goconcurrentqueue"
	"github.com/gosuri/uilive"
)

// RunGame runs a game of bricks.
func RunGame(program []int64, gameFinished chan struct{}) {
	inputQueue := q.NewFIFO()
	outputQueue := q.NewFIFO()
	finalMemory := make(chan []int64)
	gameLoopFinished := make(chan struct{})

	go intcode.ExecuteProgram(program, finalMemory, inputQueue, outputQueue, 1024)
	go gameLoop(inputQueue, outputQueue, gameLoopFinished)

	<-finalMemory

	outputQueue.Enqueue(int64(99))

	<-gameLoopFinished

	gameFinished <- struct{}{}
}

func gameLoop(inputQueue q.Queue, outputQueue q.Queue, gameLoopFinished chan struct{}) {
	writer := uilive.New()
	writer.Start()

	gameState := prepareGameState()
	score := int64(0)
	paddlePos := int64(-1)
	ballPos := int64(-1)

	initialRenderComplete := false
	initialAction := false
	renderRounds := 0
	paddleMovement := int64(0)

	for {
		xElement, _ := outputQueue.DequeueOrWaitForNextElement()
		x := xElement.(int64)
		if x == 99 {
			gameLoopFinished <- struct{}{}
			return
		}

		yElement, _ := outputQueue.DequeueOrWaitForNextElement()
		typeElement, _ := outputQueue.DequeueOrWaitForNextElement()
		y := yElement.(int64)
		output := typeElement.(int64)
		ballMoved := false

		if x == -1 && y == 0 {
			score = output
		} else {
			newCellState := cellState(output)
			gameState[y][x] = newCellState

			if newCellState == paddle {
				paddlePos = x
				paddleMoved = true
			}

			if newCellState == ball {
				ballPos = x
				ballMoved = true
			}

			renderRounds++
			if renderRounds == 1170 {
				initialRenderComplete = true
				initialAction = true
			}

			paddleMovement = int64(0)
			if initialRenderComplete && (ballMoved || initialAction) {

				if initialAction {
					initialAction = false
				}

				if ballPos > paddlePos {
					paddleMovement = 1
				} else if ballPos < paddlePos {
					paddleMovement = -1
				}

				inputQueue.Enqueue(paddleMovement)

			}
		}

		render(gameState, score, writer)
	}
}

func prepareGameState() [][]cellState {
	gameState := make([][]cellState, 26)
	for i := range gameState {
		gameState[i] = make([]cellState, 45)
	}

	return gameState
}

func render(gameState [][]cellState, score int64, writer *uilive.Writer) {
	output := fmt.Sprintf("Score: %    d\n", score)
	for _, row := range gameState {
		for _, cell := range row {
			cellOutput := ""
			switch cell {
			case empty:
				cellOutput = " "
			case wall:
				cellOutput = "░"
			case block:
				cellOutput = "#"
			case paddle:
				cellOutput = "="
			case ball:
				cellOutput = "O"
			default:
				panic(fmt.Errorf("Unexpected cell type %v", cell))
			}

			output += cellOutput
		}

		output += "\n"
	}

	fmt.Fprintf(writer, output)
}

type cellState int64

const (
	empty cellState = iota
	wall
	block
	paddle
	ball
)
