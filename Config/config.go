package config

// For now just hold hardcoded config struct, but in future read from JSON?

type Config struct {
	Workers int
	MutateRate float32
	MutateAmount float32
	PreservationRate float32
	GenerationPopulation int
	Dimensions []int // Ordered list of layer sizes
}

// If I do make this into a JSON fetch, I'm going to have to rework the fetches in FFNeuralNet to not call this method every single mutate iteration...
func GetConfig() Config {
	return Config{
		Workers: 15,
		MutateRate: 0.1,
		MutateAmount: 0.1,
		PreservationRate: 0.12,
		GenerationPopulation: 25000,
		Dimensions: []int{10, 100, 4},
	}
}
