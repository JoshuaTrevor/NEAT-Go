package ffneuralnet

import (
	"math"
	"math/rand"
	"sort"

	config "github.com/JoshuaTrevor/Neat-Go/Config"
)

type ScoredNeuralNet struct {
	NeuralNet *FFNeuralNet
	Score float64	
}

type FitnessFunc func(*FFNeuralNet) float64


func Train(evaluate FitnessFunc, fromScratch bool) {
	var generation []*FFNeuralNet
	if fromScratch {
		generation = InitGeneration() // Completely random starter generation
	} else {
		generation = Load().SpawnGeneration() // Spawn generation based off single 'Adam' species
	}
	conf := config.GetConfig()
	topSpeciesNumFloat := conf.PreservationRate * float32(conf.GenerationPopulation)
	topSpeciesNum := int(math.Ceil(float64(topSpeciesNumFloat)))

	for i := 0; i < 100; i++ { // how many generations to loop for
		evaluationQueue := make(chan *FFNeuralNet, conf.GenerationPopulation)
		evaluatedSpecies := make(chan ScoredNeuralNet, conf.GenerationPopulation)
		for _, species := range generation {
			evaluationQueue <- species
		}
		for j := 0; j < conf.Workers; j++ {
			go EvaluateSpeciesConcurrent(evaluationQueue, evaluatedSpecies, evaluate)
		}

				
		scoredNetworks := []ScoredNeuralNet{}
		for i := 0; i < conf.GenerationPopulation; i++ {
			scoredNetworks = append(scoredNetworks, <-evaluatedSpecies)
		}

		// Sort once at the end, then grab top X into a new slice. Memory inefficient but decently fast.
		sort.Slice(scoredNetworks, func(i, j int) bool {
			return scoredNetworks[i].Score > scoredNetworks[j].Score //descending
		})

		topSpecies := []*FFNeuralNet{}
		for i := 0; i < topSpeciesNum; i++ {
			if i == 0 {
			}
			topSpecies = append(topSpecies, scoredNetworks[i].NeuralNet)
		}
	
		generation = PadGeneration(topSpecies)
	}
}

// Accept input channel and output channel as args
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
