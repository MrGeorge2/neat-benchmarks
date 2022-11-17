package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MrGeorge2/neat-benchmarks/goNeat/xor"
	"github.com/yaricom/goNEAT/v3/experiment"
	"github.com/yaricom/goNEAT/v3/neat"
	"github.com/yaricom/goNEAT/v3/neat/genetics"
)

func neatExecutor() {

	contextPath := "./data/xor.neat"
	genomePath := "./data/xorstartgenes"
	experimentName := "XOR"

	// Seed the random-number generator with current time so that
	// the numbers will be different every time we run.
	seed := time.Now().Unix()
	rand.Seed(seed)

	// Load NEAT options
	neatOptions, err := neat.ReadNeatOptionsFromFile(contextPath)
	if err != nil {
		log.Fatal("Failed to load NEAT options: ", err)
	}

	// Load Genome
	log.Printf("Loading start genome for %s experiment from file '%s'\n", experimentName, genomePath)
	reader, err := genetics.NewGenomeReaderFromFile(genomePath)
	if err != nil {
		log.Fatalf("Failed to open genome file, reason: '%s'", err)
	}
	startGenome, err := reader.Read()
	if err != nil {
		log.Fatalf("Failed to read start genome, reason: '%s'", err)
	}
	fmt.Println(startGenome)

	// create experiment
	expt := experiment.Experiment{
		Id:       0,
		Trials:   make(experiment.Trials, neatOptions.NumRuns),
		RandSeed: seed,
	}
	var generationEvaluator experiment.GenerationEvaluator
	expt.MaxFitnessScore = 16.0 // as given by fitness function definition
	generationEvaluator = xor.NewXorGenerationEvaluator("")

	// prepare to execute
	errChan := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())

	// run experiment in the separate GO routine
	go func() {
		if err = expt.Execute(neat.NewContext(ctx, neatOptions), startGenome, generationEvaluator, nil); err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	}()

	// register handler to wait for termination signals
	//
	go func(cancel context.CancelFunc) {
		fmt.Println("\nPress Ctrl+C to stop")

		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		select {
		case <-signals:
			// signal to stop test fixture
			cancel()
		case err = <-errChan:
			// stop waiting
		}
	}(cancel)

	// Wait for experiment completion
	//
	err = <-errChan
	if err != nil {
		// error during execution
		log.Fatalf("Experiment execution failed: %s", err)
	}

	// Print experiment results statistics
	//
	expt.PrintStatistics()
}

// The experiment runner boilerplate code
func main() {
	neatExecutor()
}
