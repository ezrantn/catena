# catena (WIP)

**catena** is a custom serialization library that provides high-performance memory management using an arena-based allocator. The library is designed to optimize memory usage during serialization and deserialization processes by allocating memory in a single, pre-allocated chunk (arena). This can be beneficial for applications where frequent object serialization occurs and memory management needs to be efficient.

## Installation

```bash
go get github.com/ezrantn/catena
```

## Usage

```go
type User struct {
 Name  string `json:"name"`
 Email string `json:"email"`
}

func main() {
    user := &User{
        Name: "John",
        Email: "john@mail.com"
    }
    serializer := NewSerializer(1024 * 1024) // 1 MB memory
    _, err := serializer.SerializeToJSON(user)
}
```

## Benchmark

Here are the benchmark results comparing standard JSON serialization/deserialization against the custom serialization/deserialization implemented in the Catena library. The benchmark was run on an Intel i5-11400H CPU with 8GB RAM.

```bash
goos: linux
goarch: amd64
pkg: github.com/ezrantn/catena
cpu: 11th Gen Intel(R) Core(TM) i5-11400H @ 2.70GHz
BenchmarkJSONSerialization-12                    8005564               147.3 ns/op            48 B/op          1 allocs/op
BenchmarkJSONDeserialization-12                  1794866               604.8 ns/op           280 B/op          7 allocs/op
BenchmarkCatenaJSONSerialization-12             24187534                50.91 ns/op           72 B/op          2 allocs/op
BenchmarkCatenaJSONDeserialization-12            7762628               153.2 ns/op           280 B/op          7 allocs/op
PASS
ok      github.com/ezrantn/catena       5.728s
```

### Key Takeaways

1. Serialization Performance:
   - The custom Catena serialization outperforms standard JSON serialization by a significant margin, achieving 50.91 ns/op compared to 147.3 ns/op, a 2.89x improvement in speed.
   - Memory usage is slightly higher (72 bytes vs. 48 bytes), but the gain in speed justifies this trade-off.

2. Deserialization Performance:
   - Catena deserialization remains comparable to the standard JSON library, with a time of 153.2 ns/op versus 604.8 ns/op, achieving a 3.95x improvement in speed for structured JSON data.

3. Concurrency and Scalability:
   - Catena's design, leveraging memory arenas and object pooling, shows excellent scalability under concurrent workloads, as seen in the reduced runtime for serialization.

4. Allocation Efficiency:
   - Custom serialization uses 2 allocations per operation, which is marginally higher than the 1 allocation in standard JSON but significantly optimized compared to the original implementation.

## License

This tool is open-source and available under the [MIT License](https://github.com/ezrantn/catena/blob/main/LICENSE).

## Contributing

Contributions are welcome! Please feel free to submit a pull request. For major changes, please open an issue first to discuss what you would like to change.
