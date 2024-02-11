package ffneuralnet

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	config "github.com/JoshuaTrevor/Neat-Go/Config"
)

type StoredFFNeuralNet struct {
	NeuralNet               FFNeuralNet `json:"Neural Net"`
	TrainingDurationSeconds int         `json:"Training Duration Seconds"`
}

type FFNeuralNet []Layer

type Layer []Node

type Node struct {
	InputSum    float64      `json:"Input Sum"`
	Connections []Connection `json:"Connections"`
}
type Connection struct {
	Target int     `json:"Target"`
	Weight float64 `json:"Weight"`
}

// Run many times for the first generation only, subsequent neural nets will always be produced by mutation
func InitNeuralNet(layerSizes []int) *FFNeuralNet {
	layers := make(FFNeuralNet, len(layerSizes))

	for i := 0; i < len(layers); i++ {
		layer := make(Layer, layerSizes[i])
		if i != len(layers)-1 {
			initConnections(&layer, layerSizes[i+1])
		}
		layers[i] = layer
	}

	return &layers
}

func initConnections(layer *Layer, nextLayerSize int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(*layer); i++ {
		for j := 0; j < nextLayerSize; j++ {
			randomWeightCon := Connection{Target: j, Weight: rand.Float64()}
			(*layer)[i].Connections = append((*layer)[i].Connections, randomWeightCon)
		}
	}
}

func (neuralNet *FFNeuralNet) MutateConnections() {
	conf := config.GetConfig()
	rand.Seed(time.Now().UnixNano())

	for layerIdx := 0; layerIdx < len(*neuralNet); layerIdx++ {
		for nodeIdx := 0; nodeIdx < len((*neuralNet)[layerIdx]); nodeIdx++ {
			for connectionIdx := 0; connectionIdx < len((*neuralNet)[layerIdx][nodeIdx].Connections); connectionIdx++ {
				if rand.Float32() < conf.MutateRate {
					mutateDiff := (conf.MutateAmount * 2 * rand.Float32()) - conf.MutateAmount
					(*neuralNet)[layerIdx][nodeIdx].Connections[connectionIdx].Weight += float64(mutateDiff)
				}
			}
		}
	}
}

func activation(input float64) float64 {
	return (1 / (1 + math.Pow(math.E, (-1*input))))
}

// For every layer except the last layer, propagate inputs via weights
func (neuralNet *FFNeuralNet) Feed(inputs []float64) []float64 {
	layerCount := len(*neuralNet)

	// Send through inputs
	for nodeIdx := 0; nodeIdx < len((*neuralNet)[0]); nodeIdx++ {
		(*neuralNet)[0][nodeIdx].InputSum = inputs[nodeIdx]
	}

	// propagate inputs forward
	for layerIdx := 0; layerIdx < layerCount-1; layerIdx++ {
		for nodeIdx := 0; nodeIdx < len((*neuralNet)[layerIdx]); nodeIdx++ {
			for connectionIdx := 0; connectionIdx < len((*neuralNet)[layerIdx][nodeIdx].Connections); connectionIdx++ {
				connection := (*neuralNet)[layerIdx][nodeIdx].Connections[connectionIdx]
				// Multiple the connected node in the next layer by this node's input (after activation function) * con weight
				(*neuralNet)[layerIdx+1][connection.Target].InputSum += activation((*neuralNet)[layerIdx][nodeIdx].InputSum) * connection.Weight
			}
		}
	}

	// Now the outputs should be the input energy to the last layer.
	result := []float64{}
	for nodeIdx := 0; nodeIdx < len((*neuralNet)[layerCount-1]); nodeIdx++ {
		result = append(result, activation((*neuralNet)[layerCount-1][nodeIdx].InputSum))
	}

	// Reset input sums to zero. Should think of a better system than this. Best way would be to use a mapping of node to inputSum, but maybe this isn't too bad
	neuralNet.ResetInputs()

	return result

}

func (neuralNet *FFNeuralNet) DeepCopy() *FFNeuralNet {
	conf := config.GetConfig()

	// Lazy implementation: reuse code to init random neuralnet, then just replace connection weights
	newNeuralNet := InitNeuralNet(conf.Dimensions)
	for layerIdx := 0; layerIdx < len(*neuralNet); layerIdx++ {
		for nodeIdx := 0; nodeIdx < len((*neuralNet)[layerIdx]); nodeIdx++ {
			for connectionIdx := 0; connectionIdx < len((*neuralNet)[layerIdx][nodeIdx].Connections); connectionIdx++ {
				(*newNeuralNet)[layerIdx][nodeIdx].Connections[connectionIdx].Weight = (*neuralNet)[layerIdx][nodeIdx].Connections[connectionIdx].Weight
			}
		}
	}

	return newNeuralNet
}

// Spawn an entire generation from scratch, with completely random species (maximum entropy)
func InitGeneration() []*FFNeuralNet {
	conf := config.GetConfig()
	generation := []*FFNeuralNet{}

	for speciesIdx := 0; speciesIdx < conf.GenerationPopulation; speciesIdx++ {
		generation = append(generation, InitNeuralNet(conf.Dimensions))
	}

	return generation
}

// Reset 'InputSum' to 0. Must be done before cloning to avoid copies preserving input sum!
func (neuralNet *FFNeuralNet) ResetInputs() {
	// Start at one since resetting input layer is not needed
	for layerIdx := 1; layerIdx < len(*neuralNet); layerIdx++ {
		for nodeIdx := 0; nodeIdx < len((*neuralNet)[layerIdx]); nodeIdx++ {
			(*neuralNet)[layerIdx][nodeIdx].InputSum = 0
		}
	}

}

// Spawn an entire generation from a single NeuralNet, useful when trying to train an existing neural network loaded from file.
func (neuralNet *FFNeuralNet) SpawnGeneration() []*FFNeuralNet {
	conf := config.GetConfig()
	generation := []*FFNeuralNet{}

	generation = append(generation, neuralNet)
	for speciesIdx := 1; speciesIdx < conf.GenerationPopulation; speciesIdx++ {
		generation = append(generation, neuralNet.DeepCopy())
	}
	return generation
}

// TODO Measure moves/sec in training simulations. Captures speed better and would also be easy to put in java also for comparison.
