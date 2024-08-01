package snake

import (
	"math"

	ffneuralnet "github.com/JoshuaTrevor/Neat-Go/FFNeuralNet"
)

func Evaluate(neuralNet *ffneuralnet.FFNeuralNet) float64 {
	fitnessSum := float64(0)
	fitnessCount := float64(0)

	for i := 0; i < 5; i++ {
		fitnessSum += evaluateSingleRun(neuralNet)
		fitnessCount += 1
	}
	return fitnessSum / fitnessCount
}

func getManhattanDistanceToApple(snakeGame *SnakeGame) float64 {
	return math.Abs(float64(snakeGame.Food.x-snakeGame.Snake.head.x)) + math.Abs(float64(snakeGame.Food.y-snakeGame.Snake.head.y))
}

func evaluateSingleRun(neuralNet *ffneuralnet.FFNeuralNet) float64 {
	snakeGame := NewGame()
	moves := 0
	constructiveMoves := 0
	destructiveMoves := 0
	movesSinceApple := 0
	snakeLength := len(snakeGame.Snake.body)
	abort := false

	for !snakeGame.Snake.ShouldDie() && !abort {
		//fmt.Printf("Snake game %+v, movesSinceApple: %v, moves: %v\n", snakeGame, movesSinceApple, moves)
		vision := snakeGame.getVision()
		nn_output := neuralNet.Feed(vision)

		// Convert to direction
		maxVal := float64(-1000)
		maxValIndex := -1
		//fmt.Println(nn_output)
		for idx, val := range nn_output {
			if val > maxVal {
				maxVal = val
				maxValIndex = idx
			}
		}

		prevManhattan := getManhattanDistanceToApple(snakeGame)
		snakeGame.Move(maxValIndex)
		movesSinceApple++

		// Check if it ate an apple
		if len(snakeGame.Snake.body) > snakeLength {
			movesSinceApple = 0
			snakeLength = len(snakeGame.Snake.body)
		}

		// Reward the snake for moving closer to the apple
		newManhattan := getManhattanDistanceToApple(snakeGame)
		if newManhattan > prevManhattan {
			destructiveMoves++
		} else {
			constructiveMoves++
		}

		// Reward the snake for not dying straight away, kill the snake if it's making no progress
		if moves < 20 {
			moves++
		} else if float64(movesSinceApple) > 10*10*math.Min(0.4+float64(len(snakeGame.Snake.body))*0.1, 1) {
			abort = true
		}
	}

	fitness := (len(snakeGame.Snake.body)-1)*250 + constructiveMoves - destructiveMoves
	return float64(fitness)
}
