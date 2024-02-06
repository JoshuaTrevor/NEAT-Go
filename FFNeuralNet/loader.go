package ffneuralnet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (neuralNet *FFNeuralNet) Save(trainingDurationSeconds int) error {
	annotatedNN := StoredFFNeuralNet{
		NeuralNet: *neuralNet,
		TrainingDurationSeconds: trainingDurationSeconds,
	}
	marshalled, err := json.MarshalIndent(annotatedNN, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("BestBrain.json", marshalled, 0644)
	if err == nil {
		fmt.Println("File saved")
	}
	return err
}

func Load() (*StoredFFNeuralNet, error) {
	var storedNN StoredFFNeuralNet

	fileBytes, err := ioutil.ReadFile("BestBrain.json")
	if err != nil {
		return &storedNN, err
	}

	err = json.Unmarshal(fileBytes, &storedNN)
	if err == nil {
		fmt.Println("File loaded")
	}
	return &storedNN, err	
}
