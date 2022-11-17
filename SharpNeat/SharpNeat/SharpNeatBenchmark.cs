using BenchmarkDotNet.Attributes;
using SharpNeatBenchmark.Xor;

namespace SharpNeatBenchmark;

[MemoryDiagnoser]
public class SharpNeatBenchmark
{
    [Benchmark]
    public void SharpNeat()
    {
        var experimentFactory = new XorExperimentFactory();
        var ea = Utils.Utils.CreateNeatEvolutionAlgorithm(experimentFactory);

        ea.Initialise();

        for (var i = 0; i < 100; i++)
        {
            ea.PerformOneGeneration();
        }
    }
}