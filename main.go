package main

import (
	"fmt"

	ffneuralnet "github.com/JoshuaTrevor/Neat-Go/FFNeuralNet"
)

func main() {
	ffneuralnet.InitNeuralNet([]int{1, 2, 3})
	fmt.Println("done")
}
