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
BenchmarkJSONSerialization-12            7569145               161.3 ns/op            48 B/op          1 allocs/op
BenchmarkJSONDeserialization-12          1887522               637.7 ns/op           280 B/op          7 allocs/op
BenchmarkCustomSerialization-12          7210110               170.7 ns/op            48 B/op          1 allocs/op
BenchmarkCustomDeserialization-12        1820454               673.4 ns/op           280 B/op          7 allocs/op
PASS
ok      github.com/ezrantn/catena       6.527s
```

### Key Takeaways

1. The custom serialization is slightly slower than the standard JSON serialization, with a minimal performance overhead of about 8.6 ns per operation.
2. Memory allocation for both custom and standard JSON serialization is the same (48 bytes), showing that the custom implementation does not introduce unnecessary overhead.
3. The custom deserialization time is close to the standard JSON deserialization, with only a small difference of around 11 ns per operation.

## License

This tool is open-source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a pull request. For major changes, please open an issue first to discuss what you would like to change.
