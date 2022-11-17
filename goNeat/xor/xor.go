package xor

import (
	"context"
	"math"

	"github.com/yaricom/goNEAT/v3/experiment"
	"github.com/yaricom/goNEAT/v3/neat"
	"github.com/yaricom/goNEAT/v3/neat/genetics"
)

const fitnessThreshold = 15.5

type xorGenerationEvaluator struct {
}

func NewXorGenerationEvaluator(outputPath string) experiment.GenerationEvaluator {
	return &xorGenerationEvaluator{}
}

func (e *xorGenerationEvaluator) GenerationEvaluate(ctx context.Context, pop *genetics.Population, epoch *experiment.Generation) error {
	options, ok := neat.FromContext(ctx)
	if !ok {
		return neat.ErrNEATOptionsNotFound
	}
	// Evaluate each organism on a test
	for _, org := range pop.Organisms {
		res, err := e.orgEvaluate(org)
		if err != nil {
			return err
		}

		if res && (epoch.Champion == nil || org.Fitness > epoch.Champion.Fitness) {
			epoch.Solved = true
			epoch.WinnerNodes = len(org.Genotype.Nodes)
			epoch.WinnerGenes = org.Genotype.Extrons()
			epoch.WinnerEvals = options.PopSize*epoch.Id + org.Genotype.Id
			epoch.Champion = org
		}
	}

	// Fill statistics about current epoch
	epoch.FillPopulationStatistics(pop)

	return nil
}

func (e *xorGenerationEvaluator) orgEvaluate(organism *genetics.Organism) (bool, error) {
	// The four possible input combinations to xor
	// The first number is for biasing
	in := [][]float64{
		{1.0, 0.0, 0.0},
		{1.0, 0.0, 1.0},
		{1.0, 1.0, 0.0},
		{1.0, 1.0, 1.0}}

	netDepth, err := organism.Phenotype.MaxActivationDepthFast(0) // The max depth of the network to be activated

	if netDepth == 0 {
		return false, nil
	}

	success := false          // Check for successful activation
	out := make([]float64, 4) // The four outputs

	// Load and activate the network on each input
	for count := 0; count < 4; count++ {
		if err = organism.Phenotype.LoadSensors(in[count]); err != nil {
			return false, err
		}

		// Use depth to ensure full relaxation
		if success, err = organism.Phenotype.ForwardSteps(netDepth); err != nil {
			return false, err
		}
		out[count] = organism.Phenotype.Outputs[0].Activation

		// Flush network for subsequent use
		if _, err = organism.Phenotype.Flush(); err != nil {
			return false, err
		}
	}

	if success {
		// Mean Squared Error
		errorSum := math.Abs(out[0]) + math.Abs(1.0-out[1]) + math.Abs(1.0-out[2]) + math.Abs(out[3]) // ideal == 0
		target := 4.0 - errorSum                                                                      // ideal == 4.0
		organism.Fitness = math.Pow(4.0-errorSum, 2.0)
		organism.Error = math.Pow(4.0-target, 2.0)
	} else {
		// The network is flawed (shouldn't happen) - flag as anomaly
		organism.Error = 1.0
		organism.Fitness = 0.0
	}

	if organism.Fitness > fitnessThreshold {
		organism.IsWinner = true

	} else {
		organism.IsWinner = false
	}

	return organism.IsWinner, nil
}
