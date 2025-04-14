package bloom

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// Create a new Bloom filter with expected number of elements and false positive rate
	n := uint(100)       // expected number of elements
	p := 0.01            // false positive rate
	bloom := NewBloom(n, p)

	// Test data to add to the Bloom filter
	data := []byte("test-data")

	// Add the data to the Bloom filter
	bloom.Add(data)

	// Verify that the data exists in the Bloom filter
	if !bloom.Exists(data) {
		t.Errorf("Expected data to exist in the Bloom filter after adding, but it does not")
	}

	// Verify that unrelated data does not exist in the Bloom filter
	unrelatedData := []byte("unrelated-data")
	if bloom.Exists(unrelatedData) {
		t.Errorf("Expected unrelated data to not exist in the Bloom filter, but it does")
	}
}
