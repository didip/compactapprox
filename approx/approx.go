package approx


import (
    "crypto/sha256"
    "encoding/binary"
    "errors"
    "hash"
    "hash/fnv"
    "math"
)

// HashType defines the available hash algorithms.
type HashType int

const (
    FNV HashType = iota
    SHA256
)

// HashFunctionProvider returns a function that hashes bytes using the specified algorithm.
func HashFunctionProvider(hashType HashType) func([]byte, uint64) uint64 {
    switch hashType {
    case SHA256:
        return sha256HashWithSalt
    default:
        return fnvHashWithSalt
    }
}

type CompactApproximator struct {
    numBuckets        uint
    numHashFunctions  uint
    valueTable        []uint64
    hashFunctionSalts []uint64
    hashFunction      func([]byte, uint64) uint64
}

func NewCompactApproximator(numBuckets, numHashFunctions uint, hashType HashType) (*CompactApproximator, error) {
    if numBuckets == 0 || numHashFunctions == 0 {
        return nil, errors.New("numBuckets and numHashFunctions must be greater than 0")
    }

    salts := make([]uint64, numHashFunctions)
    for i := range salts {
        salts[i] = uint64(i + 1)
    }

    return &CompactApproximator{
        numBuckets:        numBuckets,
        numHashFunctions:  numHashFunctions,
        valueTable:        make([]uint64, numBuckets),
        hashFunctionSalts: salts,
        hashFunction:      HashFunctionProvider(hashType),
    }, nil
}

func fnvHashWithSalt(data []byte, salt uint64) uint64 {
    h := fnv.New64a()
    h.Write(data)
    var saltBytes [8]byte
    binary.LittleEndian.PutUint64(saltBytes[:], salt)
    h.Write(saltBytes[:])
    return h.Sum64()
}

func sha256HashWithSalt(data []byte, salt uint64) uint64 {
    h := sha256.New()
    h.Write(data)
    var saltBytes [8]byte
    binary.LittleEndian.PutUint64(saltBytes[:], salt)
    h.Write(saltBytes[:])
    sum := h.Sum(nil)
    return binary.LittleEndian.Uint64(sum[:8])
}

func (a *CompactApproximator) Insert(key []byte, value uint64) {
    for _, salt := range a.hashFunctionSalts {
        idx := a.hashFunction(key, salt) % uint64(a.numBuckets)
        if a.valueTable[idx] < value {
            a.valueTable[idx] = value
        }
    }
}

func (a *CompactApproximator) Get(key []byte) uint64 {
    minValue := uint64(math.MaxUint64)
    for _, salt := range a.hashFunctionSalts {
        idx := a.hashFunction(key, salt) % uint64(a.numBuckets)
        if a.valueTable[idx] < minValue {
            minValue = a.valueTable[idx]
        }
    }
    return minValue
}
