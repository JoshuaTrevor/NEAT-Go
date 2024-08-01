package ffneuralnet_test

import (
	"testing"

	ffneuralnet "github.com/JoshuaTrevor/Neat-Go/FFNeuralNet"
)

func TestChangingInputChangesOutput(t *testing.T) {
	neuralNet := ffneuralnet.InitNeuralNet([]int{2, 10, 2})
	
	emptyInput := []float64{0, 0}
	output1 := neuralNet.Feed(emptyInput)

	nonEmptyInput := []float64{1, 1}
	output2 := neuralNet.Feed(nonEmptyInput)

	if output1[0] == output2[0] && output1[1] == output2[1] {
		t.Fatalf("Different inputs should result in different outputs, but they were the same!")
	}
}

func TestSameInputSameOutput(t *testing.T) {
	neuralNet := ffneuralnet.InitNeuralNet([]int{2, 10, 2})
	
	emptyInput := []float64{0, 0}
	output1 := neuralNet.Feed(emptyInput)
	output2 := neuralNet.Feed(emptyInput)

	if output1[0] != output2[0] || output1[1] != output2[1] {
		t.Fatalf("The same input should have the same output with no training step in between!")
	}
}

func TestMutationChangesOutput(t *testing.T) {
	neuralNet := ffneuralnet.InitNeuralNet([]int{2, 10, 2})
	
	emptyInput := []float64{0, 0}
	output1 := neuralNet.Feed(emptyInput)
	neuralNet.MutateConnections()
	output2 := neuralNet.Feed(emptyInput)

	if output1[0] == output2[0] && output1[1] == output2[1] {
		t.Fatalf("After mutation, the same input should not produce the same output!")
	}
}
