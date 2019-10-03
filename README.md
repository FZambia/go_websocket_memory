Investigate Gorilla Websocket memory usage in various scenarios.

Stats showed here were reported by each server after all 10k connections from client script successfully established.

### server_01

Smaller buffers, no goroutine.

```
Alloc = 138 MiB	TotalAlloc = 182 MiB	Sys = 272 MiB
```

### server_02

Reuse buffers, no goroutine.

```
Alloc = 155 MiB	TotalAlloc = 160 MiB	Sys = 273 MiB
```

### server_03

Smaller buffers, with goroutine.

```
Alloc = 37 MiB	TotalAlloc = 182 MiB	Sys = 138 MiB
```

### server_04

Reuse buffers, with goroutine.

```
Alloc = 94 MiB	TotalAlloc = 160 MiB	Sys = 206 MiB
```

### server_05

Smaller read buffer, with goroutine, write buffer pool. 

```
Alloc = 26 MiB	TotalAlloc = 171 MiB	Sys = 138 MiB
```
