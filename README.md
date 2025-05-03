# Compact Approximator

A **Compact Approximator** is a space-efficient probabilistic data structure that maps keys to values using a generalization of Bloom filters. Instead of bits, it stores values from a lattice (e.g., integers, timestamps, counters), allowing efficient tracking of upper bounds across keys.

This Go library implements a customizable compact approximator with pluggable hash functions.

---

## ✨ Features

- ✅ Associate approximate values with keys
- ⚙️ Use `max` on insert, `min` on lookup (lattice semantics)
- 🔁 Customizable hash algorithms (FNV-1a, SHA-256)
- 🧪 Thoroughly tested
- 🪪 MIT licensed
- 📁 Includes real-world usage examples
- Library is written by AI and inspected by me.

---

## 📦 Installation

```bash
go get github.com/didip/compactapprox
```

---

## 🚀 Usage

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/didip/compactapprox/approx"
)

func main() {
    approx, err := approx.NewCompactApproximator(1024, 4, approx.FNV)
    if err != nil {
        panic(err)
    }

    key := []byte("user:567")
    approx.Insert(key, 123)
    approx.Insert(key, 110)
    value := approx.Get(key)

    fmt.Printf("Approximate value for key %s: %d\n", key, value)
}
```

---

## 🧠 How It Works

- Uses an array of `uint64` values (instead of bits).
- `Insert(key, value)` updates `k` positions in the array using:
  ```go
  table[i] = max(table[i], value)
  ```
- `Get(key)` reads the same `k` positions and returns:
  ```go
  min(table[i])
  ```

This provides an **upper bound** approximation of the value associated with a key.

---

## 🔧 Hash Algorithm Options

You can choose which hash function to use at creation:

```go
approx.NewCompactApproximator(numBuckets, numHashes, approx.FNV)
approx.NewCompactApproximator(numBuckets, numHashes, approx.SHA256)
```

- `FNV`: Fast and simple (default)
- `SHA256`: Cryptographic-quality hash (slower, more uniform)

---

## ✅ Testing

Run tests with:

```bash
go test ./approx
```

---

## 📁 Project Structure

```
compactapproximator/
├── LICENSE
├── go.mod
├── README.md
├── approx/
│   ├── compact_approximator.go
│   └── compact_approximator_test.go
└── examples/
    ├── basic_usage.go
    └── sensor_usage.go
```

---

## 📄 License

This project is licensed under the [MIT License](./LICENSE).

---

## 🤝 Contributing

Feel free to open issues or pull requests. Suggestions and improvements welcome!
