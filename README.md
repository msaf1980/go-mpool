# go-mpool

go-mpool is a simple memory pool library written in golang using stack.

## Requirement

- go (>= 1.16)

## Installation

```shell
go get github.com/msaf1980/go-mpool
```

## Example

MPool (allocation-free and lock-free memory pool)
Pool (Allocation-free memory pool, not thread-safe)

```go

// Create 10 pools of `*bytes.Buffer`
pools := Pool.NewMPool(10, func() interface{} {
    return &bytes.Buffer{}
})

pool1 := pools.Get()
pool2 := pools.Get()

buf1 := pool1.Value.(*bytes.Buffer)
buf1.WriteString("Hello world1")

buf2 := pool2.Value.(*bytes.Buffer)
buf2.WriteString("Hello world2")

pools.Put(pool1)    // Return pool1 to pools
pools.Put(pool2)    // Return pool2 to pools

```

## Benchmark
[Pool](https://github.com/msaf1980/go-mpool) vs [sync.Pool](https://github.com/golang/go/tree/master/src/sync)(Standard library)

```
goarch: amd64
pkg: github.com/msaf1980/go-mpool
cpu: Intel(R) Core(TM) i5-9400F CPU @ 2.90GHz
BenchmarkMPool-6      	201046626	         5.965 ns/op	       0 B/op	       0 allocs/op
BenchmarkPool-6       	196496100	         5.976 ns/op	       0 B/op	       0 allocs/op
BenchmarkSyncPool-6   	31270672	        38.59 ns/op	       0 B/op	       0 allocs/op
PASS
```

## Author
Based on [hlts2](https://github.com/hlts2)

## LICENSE
go-mpool released under MIT license, refer [LICENSE](https://github.com/hlts2/Pool/blob/master/LICENSE) file.
