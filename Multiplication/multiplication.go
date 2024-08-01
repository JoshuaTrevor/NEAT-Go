package multiplication

import (
	"math/rand"
	"time"

	config "github.com/JoshuaTrevor/Neat-Go/Config"
	ffneuralnet "github.com/JoshuaTrevor/Neat-Go/FFNeuralNet"
)


func Evaluate(neuralNet *ffneuralnet.FFNeuralNet) float64 {
	sum := 0.0
	for i := 0; i < 100; i++ {
		sum += EvaluateSingle(neuralNet) 
	}
	return sum / 100
}

func EvaluateSingle(neuralNet *ffneuralnet.FFNeuralNet) float64 { 
	conf := config.GetConfig()
	rand.Seed(time.Now().UnixNano())
	inputs := []float64{}
	for i := 0; i < conf.Dimensions[0]; i++ {
		inputs = append(inputs, rand.Float64())
	}
	output := neuralNet.Feed(inputs)
	return 1 - output[0]
}
