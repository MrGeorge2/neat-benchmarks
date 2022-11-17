using SharpNeatBenchmark.Utils;
using SharpNeatBenchmark.Xor;

var experimentFactory = new XorExperimentFactory();
var ea = Utils.CreateNeatEvolutionAlgorithm(experimentFactory);

ea.Initialise();

for (var i = 0; i < 100; i++)
{
    ea.PerformOneGeneration();
}

/*
BenchmarkRunner.Run<SharpNeatBenchmark.SharpNeatBenchmark>();
*/