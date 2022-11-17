``` ini

BenchmarkDotNet=v0.13.2, OS=ubuntu 22.04
12th Gen Intel Core i7-1255U, 1 CPU, 12 logical and 10 physical cores
.NET SDK=7.0.100
  [Host]     : .NET 7.0.0 (7.0.22.51805), X64 RyuJIT AVX2
  DefaultJob : .NET 7.0.0 (7.0.22.51805), X64 RyuJIT AVX2


```
|    Method |     Mean |    Error |   StdDev |      Gen0 |      Gen1 |     Gen2 | Allocated |
|---------- |---------:|---------:|---------:|----------:|----------:|---------:|----------:|
| SharpNeat | 41.77 ms | 0.818 ms | 1.321 ms | 8000.0000 | 3312.5000 | 812.5000 |  46.38 MB |
