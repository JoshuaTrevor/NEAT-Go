package main

import (
	//ffneuralnet "github.com/JoshuaTrevor/Neat-Go/FFNeuralNet"

	snake "github.com/JoshuaTrevor/Neat-Go/Snake"
	testexample "github.com/JoshuaTrevor/Neat-Go/TestExample"
)

func main() {
	//neuralNet := ffneuralnet.InitNeuralNet([]int{1, 2, 3})
	//neuralNet.MutateConnections()
	//example()
	//snake.DoTest()
	playSnake()
}

func playSnake() {
	game := snake.NewGame()
	for !game.Snake.ShouldDie() {
		game.Move(snake.LEFT)
	}
}


func example () {
	testexample.Train()
}
