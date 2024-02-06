package testexample

import (
	ffneuralnet "github.com/JoshuaTrevor/Neat-Go/FFNeuralNet"
)


func Train() {
	fromScratch := false
	ffneuralnet.Train(evaluate, fromScratch)
}



func evaluate(neuralNet *ffneuralnet.FFNeuralNet) float64 {
	outputs := neuralNet.Feed([]float64{1, 1, 1, 1})

	score := outputs[0] - outputs[1]
	return score
}
