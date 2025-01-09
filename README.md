# Catena (WIP)

**Catena** is a custom serialization library that provides high-performance memory management using an arena-based allocator. The library is designed to optimize memory usage during serialization and deserialization processes by allocating memory in a single, pre-allocated chunk (arena). This can be beneficial for applications where frequent object serialization occurs and memory management needs to be efficient.

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
BenchmarkJSONSerialization-12                    7212676               152.3 ns/op            48 B/op          1 allocs/op
BenchmarkJSONDeserialization-12                  2033766               649.0 ns/op           280 B/op          7 allocs/op
BenchmarkCatenaJSONSerialization-12             22543629                50.20 ns/op           72 B/op          2 allocs/op
BenchmarkCatenaJSONDeserialization-12            7464746               162.5 ns/op           280 B/op          7 allocs/op
BenchmarkGoogleProtoSerialization-12            14805024                80.87 ns/op           32 B/op          1 allocs/op
BenchmarkGoogleProtoDeserialization-12          10505408               113.2 ns/op            29 B/op          2 allocs/op
BenchmarkCatenaProtoSerialization-12             9760736               128.3 ns/op            56 B/op          2 allocs/op
BenchmarkCatenaProtoDeserialization-12           7727215               148.3 ns/op           109 B/op          3 allocs/op
PASS
ok      github.com/ezrantn/catena       11.026s
```

### Takeaways

1. **JSON Serialization and Deserialization:**

   - **Catena JSON Serialization** is significantly faster than the standard JSON serialization (50.20 ns/op vs. 152.3 ns/op), delivering a 3.04x improvement in speed. However, this comes with a trade-off in **memory usage** (72 bytes vs. 48 bytes) and **allocations** (2 allocations vs. 1). While memory usage is slightly higher, the improved speed could be valuable for performance-sensitive applications.
   - **Catena JSON Deserialization** also outperforms the standard JSON deserialization (162.5 ns/op vs. 649.0 ns/op), achieving a 4.00x speedup. The trade-offs here are **no change in memory usage** (280 bytes) and 7 **allocations** (vs. 7 allocations in the standard library). The extra speed is especially beneficial in workloads requiring rapid deserialization, though the increased allocations might slightly impact memory usage in highly concurrent systems.

2. **Proto Serialization and Deserialization:**

   - **Catena Proto Serialization** shows a trade-off: while it’s slower than the Google Proto library (128.3 ns/op vs. 81.82 ns/op), the custom implementation provides more flexibility by utilizing a **memory arena** and **object pooling**. This flexibility may result in higher **memory usage** (56 bytes vs. 32 bytes) and **allocations** (2 vs. 1) but gives you the option to manage memory more efficiently in custom scenarios. The reduced speed is a reasonable compromise for systems that need precise control over memory management.
   - **Catena Proto Deserialization** is slower than Google’s implementation (148.3 ns/op vs. 106.6 ns/op), resulting in a 1.39x slower operation. The trade-offs here are a **higher memory usage** (109 bytes vs. 29 bytes) and more **allocations** (3 vs. 2). These overheads might be acceptable when the primary goal is memory control and when deserialization time is less critical than efficient memory reuse in a high-concurrency environment.

### Conclusion

The custom Catena implementations offer better control over memory allocation and scalability in exchange for slightly higher memory usage and additional allocations, especially noticeable in serialization. These trade-offs are particularly beneficial when optimizing for performance in scenarios involving large datasets and high concurrency. However, if raw speed and minimal memory footprint are the key goals, the standard libraries (JSON or Proto) will outperform Catena, especially in simpler use cases where advanced memory management isn’t required.

## License

This tool is open-source and available under the [MIT License](https://github.com/ezrantn/catena/blob/main/LICENSE).

## Contributing

Contributions are welcome! Please feel free to submit a pull request. For major changes, please open an issue first to discuss what you would like to change.
