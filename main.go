package main

import (
	//ffneuralnet "github.com/JoshuaTrevor/Neat-Go/FFNeuralNet"

	ffneuralnet "github.com/JoshuaTrevor/Neat-Go/FFNeuralNet"
	multiplication "github.com/JoshuaTrevor/Neat-Go/Multiplication"
	snake "github.com/JoshuaTrevor/Neat-Go/Snake"
	testexample "github.com/JoshuaTrevor/Neat-Go/TestExample"
)

func main() {
	trainSnake()
	//playMultiplication()
}

func playMultiplication() {
	ffneuralnet.Train(multiplication.Evaluate, true)
}

// Single evaluation pass
func playSnake() {
	snake.Evaluate(ffneuralnet.InitNeuralNet([]int{10, 100, 4}))
}

func trainSnake() {
	ffneuralnet.Train(snake.Evaluate, false)
}

func example() {
	testexample.Train()
}
