package bloom

import (
	"hash/crc32"
	"hash/fnv"
	"math"
)

type Bloom struct {
	bitset []byte
	m      uint // number of bits
	k      uint // number of hash functions
}

// @n: number of elements
// @p: false positive rate
// @return: a Bloom filter
func NewBloom(n uint, p float64) *Bloom {
	m := calM(n, p)
	k := calK(m, n)
	bitset := make([]byte, (m+7)/8)
	return &Bloom{
		bitset: bitset,
		m:      m,
		k:      k,
	}
}

// m:=-(n*ln(p))/(ln(2)^2)
func calM(n uint, p float64) uint {
	m := -float64(n) * math.Log(p) / (math.Ln2 * math.Ln2)
	return uint(math.Ceil(m))
}

// k:=m/n*ln(2)
func calK(m, n uint) uint {
	k := float64(m) / float64(n) * math.Ln2
	return uint(math.Ceil(k))
}

// @i:ith hash function
func (b *Bloom) hash(data []byte, i uint) uint {
	h1 := hash1(data)
	h2 := hash2(data)
	return (h1 + i*h2) % b.m
}

// 第一个hash:fnv
func hash1(data []byte) uint {
	h := fnv.New32()
	h.Write(data)
	return uint(h.Sum32())
}

// 第二个hash:crc32
func hash2(data []byte) uint {
	h := crc32.NewIEEE()
	h.Write(data)
	return uint(h.Sum32())
}

func (b *Bloom) Add(data []byte) {
	for i := uint(0); i < b.k; i++ {
		idx := b.hash(data, i)
		byteIdx := idx / 8
		bitIdx := idx % 8
		b.bitset[byteIdx] |= 1 << bitIdx
	}
}

func (b *Bloom) Exists(data []byte) bool {
	for i := uint(0); i < b.k; i++ {
		idx := b.hash(data, i)
		byteIdx := idx / 8
		bitIdx := idx % 8
		if b.bitset[byteIdx]&(1<<bitIdx) == 0 {
			return false
		}
	}
	return true
}
