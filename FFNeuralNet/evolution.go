package ffneuralnet

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	config "github.com/JoshuaTrevor/Neat-Go/Config"
)

type ScoredNeuralNet struct {
	NeuralNet *FFNeuralNet
	Score     float64
}

type FitnessFunc func(*FFNeuralNet) float64

func Train(evaluate FitnessFunc, fromScratch bool) {
	var generation []*FFNeuralNet
	var trainingTimeSeconds int = 0
	start := time.Now().Unix()
	if fromScratch {
		generation = InitGeneration() // Completely random starter generation
	} else {
		storedNeuralNet, err := Load()
		if err != nil {
			panic(err)
		}
		trainingTimeSeconds = storedNeuralNet.TrainingDurationSeconds
		generation = storedNeuralNet.NeuralNet.SpawnGeneration()
	}
	conf := config.GetConfig()
	topSpeciesNumFloat := conf.PreservationRate * float32(conf.GenerationPopulation)
	topSpeciesNum := int(math.Ceil(float64(topSpeciesNumFloat)))

	var bestBrain *FFNeuralNet
	for i := 0; i < 100000; i++ { // how many generations to loop for
		evaluationQueue := make(chan *FFNeuralNet, conf.GenerationPopulation)
		evaluatedSpecies := make(chan ScoredNeuralNet, conf.GenerationPopulation)
		for _, species := range generation {
			evaluationQueue <- species
		}
		generationStartTime := time.Now().Unix()
		for j := 0; j < conf.Workers; j++ {
			go EvaluateSpeciesConcurrent(evaluationQueue, evaluatedSpecies, evaluate)
		}

		scoredNetworks := []ScoredNeuralNet{}
		for i := 0; i < conf.GenerationPopulation; i++ {
			scoredNetworks = append(scoredNetworks, <-evaluatedSpecies)
		}
		generationScoreTime := time.Now().Unix() - generationStartTime

		// Sort once at the end, then grab top X into a new slice. Memory inefficient but decently fast.
		sort.Slice(scoredNetworks, func(i, j int) bool {
			return scoredNetworks[i].Score > scoredNetworks[j].Score //descending
		})

		fmt.Println("Generation", i, "had mean score of", mean(scoredNetworks), "was scored in", generationScoreTime, "seconds.")
		//fmt.Println("Best score of generation", i, "was", scoredNetworks[0].Score)

		topSpecies := []*FFNeuralNet{}
		for i := 0; i < topSpeciesNum; i++ {
			topSpecies = append(topSpecies, scoredNetworks[i].NeuralNet)
		}
		bestBrain = topSpecies[0]
		generation = PadGeneration(topSpecies)

		if i%10 == 0 {
			bestBrain.Save(trainingTimeSeconds + int(time.Now().Unix()) - int(start))
		}
	}
	bestBrain.Save(trainingTimeSeconds + int(time.Now().Unix()) - int(start))
}

func mean(scoredNNs []ScoredNeuralNet) float64 {
	sum := 0.0
	for _, scoredNN := range scoredNNs {
		sum += scoredNN.Score
	}
	return sum / float64(len(scoredNNs))
}

// In future consider whether this should be a constantly running service instead, would need to keep same channels since the start I guess.
func EvaluateSpeciesConcurrent(evaluationQueue chan *FFNeuralNet, evaluatedSpecies chan ScoredNeuralNet, evaluate FitnessFunc) {
	for neuralNet := range evaluationQueue {
		scoredNN := ScoredNeuralNet{NeuralNet: neuralNet, Score: evaluate(neuralNet)}
		evaluatedSpecies <- scoredNN
	}
}

func PadGeneration(initialSpecies []*FFNeuralNet) []*FFNeuralNet {
	conf := config.GetConfig()

	for len(initialSpecies) < conf.GenerationPopulation {
		parentIdx := rand.Intn(len(initialSpecies))
		child := initialSpecies[parentIdx].DeepCopy()
		child.MutateConnections() // Potential optimisation: Two operations could be combined into SpawnChild, which does the mutation during deep copy.
		initialSpecies = append(initialSpecies, child)
	}

	// 'initialSpecies' is now a full sized generation ready for evaluation, based on the initial species.
	return initialSpecies
}
