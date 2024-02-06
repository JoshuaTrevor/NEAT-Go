package main

import (
	//ffneuralnet "github.com/JoshuaTrevor/Neat-Go/FFNeuralNet"
	testexample "github.com/JoshuaTrevor/Neat-Go/TestExample"
)

func main() {
	//neuralNet := ffneuralnet.InitNeuralNet([]int{1, 2, 3})
	//neuralNet.MutateConnections()
	example()
}

func example () {
	testexample.Train()
}
