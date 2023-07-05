# Benchmarking

- Run go tests with benchmarks:

```bash
go test -bench .
go test -bench=FibInt
```

- Run go tests with benchmarks and create statistics in binary files for later analysis:

```bash
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
go test -cpuprofile cpu.prof -memprofile mem.prof -bench=FibInt
go test -cpuprofile cpu.prof -memprofile mem.prof -bench=FibBig
```

- Analyze the benchmark data with `pprof` command, and type in `top` in pprof CLI:

```bash
go tool pprof cpu.prof
go tool pprof mem.prof
```
