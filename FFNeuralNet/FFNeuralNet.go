package ffneuralnet

import (
	"fmt"
	"math/rand"
)

type FFNeuralNet []Layer

type Layer []Node

type Node []Connection

type Connection struct {
	Target int
	Weight float64
}

// Run many times for the first generation only, subsequent neural nets will always be produced by mutation
func InitNeuralNet(layerSizes []int) {
	layers := make(FFNeuralNet, len(layerSizes))

	for i := 0; i < len(layers); i++ {
		layer := make(Layer, layerSizes[i])
		// Don't connect the last layer to anything
		if i == len(layers) - 1 {
			continue
		}

		initConnections(&layer, layerSizes[i+1])

		layers[i] = layer
	}

	fmt.Println(layers)

}

func initConnections(layer *Layer, nextLayerSize int) {
	for i := 0; i < len(*layer); i++ {
		for j := 0; j < nextLayerSize; j++ {
			randomWeightCon :=Connection{Target: j, Weight: rand.Float64()}
			(*layer)[i] = append((*layer)[i], randomWeightCon)
		}
	}
}

func mutateConnections(neuralNet *FFNeuralNet) {
	// Do I want a config package? Or a config JSON? Either way need a config struct, so probably a config package!
	// Then hardcode it at start and can always change it later very easily.
	// Anyway just do basic mutation like last time -> mutate random * by random amount within bounds.
	// Then only missing activation function and feeding before I can do the basic fitness function of 
	// what is the score of node 0 minus node 1. Hopefully should instantly trend to [1, 0]
	// Then I can get into the scalability (and then... the... snake implementation...)
}

