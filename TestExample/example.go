package testexample

import (
	ffneuralnet "github.com/JoshuaTrevor/Neat-Go/FFNeuralNet"
)


func Train() {
	fromScratch := false
	ffneuralnet.Train(evaluate, fromScratch)
}



// Trivial fitness function that should maximise the first weight and minimise the second.
	// This is the most simplistic way to verify that the NN is capable of learning
func evaluate(neuralNet *ffneuralnet.FFNeuralNet) float64 {
	outputs := neuralNet.Feed([]float64{1, 1, 1, 1})

	score := outputs[0] - outputs[1]
	return score
}
