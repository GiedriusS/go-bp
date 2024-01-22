# go-bp

Binary packing in Go

## Features / TODO list

- Scalar implementation
- Fuzzing against the C++ implementation
- SIMD implementation
- Extend to AVX256 (ymm registers)
- Regular fuzz tests
- Benchmarks

## How this works?

- Integers are in 128 mini blocks
- Each block contains 16 mini blocks (16 * 128 = 2048)
- Alignment to 128 bit boundaries for SIMD to work its magic
- AVX256 needs alignment to 256 bit boundaries
- Sorted input of integers, no duplicates
