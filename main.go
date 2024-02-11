package main

import (
	//ffneuralnet "github.com/JoshuaTrevor/Neat-Go/FFNeuralNet"

	ffneuralnet "github.com/JoshuaTrevor/Neat-Go/FFNeuralNet"
	snake "github.com/JoshuaTrevor/Neat-Go/Snake"
	testexample "github.com/JoshuaTrevor/Neat-Go/TestExample"
)

func main() {
	// Current state, it should work but it doesn't learn :)
	// Will be hard to debug without creating cases to explore behaviour, so may as well make test cases.
	// TODO hitlist: Check it still works with a less complicated game. Like teach it to multiply numbers with random inputs.
		// IE. make fake game where inputs are random numbers 1-100, then it has to multiply input 1 by input 2, then divide by input three.
			// Fitness is distance from actual answer
		// Write some tests of basic things I want to be sure of. IE. 
			// a test to ensure feeding all zeros to an arbitrary NN results in all zero outputs.
			// a test to ensure that different inputs result in different outputs
			// a test to ensure non-deterministic output
			// a test to train an arbitrary implementation to a minimum fitness score within 1 second. (ie. the math game?)

	trainSnake()
}

// Single evaluation pass
func playSnake() {
	snake.Evaluate(ffneuralnet.InitNeuralNet([]int{10, 100, 4}))	
}

func trainSnake() {
	ffneuralnet.Train(snake.Evaluate, false)
}



func example () {
	testexample.Train()
}
