// bloom_test.go - Bloom filter tests.
// Written in 2017 by Yawning Angel
//
// To the extent possible under law, the author(s) have dedicated all copyright
// and related and neighboring rights to this software to the public domain
// worldwide. This software is distributed without any warranty.
//
// You should have received a copy of the CC0 Public Domain Dedication along
// with this software. If not, see <http://creativecommons.org/publicdomain/zero/1.0/>.

package bloom

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	const (
		entryLength       = 32
		falsePositiveRate = 0.01
		filterSize        = 15 // 2^15 bits = 4 KiB
	)

	assert := assert.New(t)
	require := require.New(t)

	// 4 KiB filter, 0.01 false positive rate.
	f, err := New(rand.Reader, filterSize, falsePositiveRate)
	require.NoError(err, "New()")
	assert.Equal(0, f.Entries(), "Entries(), empty filter")
	assert.Equal(4096, len(f.b), "Backing store size")

	// I could assert on these since the values calculated by New() are
	// supposed to be optimal, but I won't for now.
	t.Logf("Hashes: %v", f.nrHashes)         // 7 hashes is "ideal".
	t.Logf("MaxEntries: %v", f.MaxEntries()) // 3418 entries with these params.

	// Generate enough entries to fully saturate the filter.
	max := f.MaxEntries()
	entries := make(map[[entryLength]byte]bool)
	for count := 0; count < max; {
		var ent [entryLength]byte
		rand.Read(ent[:])

		// This needs to ignore false positives.
		if !f.TestAndSet(ent[:]) {
			entries[ent] = true
			count++
		}
	}
	assert.Equal(max, f.Entries(), "After populating")

	// Ensure that all the entries are present in the filter.
	idx := 0
	for ent := range entries {
		assert.True(f.Test(ent[:]), "Test(ent #: %v)", idx)
		assert.True(f.TestAndSet(ent[:]), "TestAndSet(ent #: %v)", idx)
		idx++
	}

	// Test the false positive rate, by generating another set of entries
	// NOT in the filter, and counting the false positives.
	//
	// This may have suprious failures once in a blue moon because the
	// algorithm is probabalistic, but that's *exceedingly* unlikely with
	// the chosen delta.
	randomEntries := make(map[[entryLength]byte]bool)
	for count := 0; count < max; {
		var ent [entryLength]byte
		rand.Read(ent[:])
		if !entries[ent] && !randomEntries[ent] {
			randomEntries[ent] = true
			count++
		}
	}
	falsePositives := 0
	for ent := range randomEntries {
		if f.Test(ent[:]) {
			falsePositives++
		}
	}
	observedP := float64(falsePositives) / float64(max)
	t.Logf("Observed False Positive Rate: %v", observedP)
	assert.InDelta(falsePositiveRate, observedP, 0.02, "False positive rate")

	assert.Equal(max, f.Entries(), "After tests") // Should still be = max.
}
