// aez_test.go - AEZ tests.
//
// To the extent possible under law, Yawning Angel has waived all copyright
// and related or neighboring rights to aez, using the Creative
// Commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package aez

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func readJsonTestdata(t *testing.T, name string, destination interface{}) {
	var file *os.File
	file, err := os.Open(filepath.Join("testdata", name))
	if err != nil {
		t.Fatalf("Failed to read test vectors in %s", name)
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&destination); err != nil {
		t.Fatalf("Failed to parse test vectors in %s", name)
	}
}

// (a,b)  ==>  Extract(a) = b.
type ExtractVector struct {
	A string `json:"a"`
	B string `json:"b"`
}

func TestExtract(t *testing.T) {
	var extractVectors []ExtractVector

	readJsonTestdata(t, "extract.json", &extractVectors)

	for i, vec := range extractVectors {
		var extractedKey [extractedKeySize]byte

		vecA, err := hex.DecodeString(vec.A)
		if err != nil {
			t.Fatal(err)
		}
		vecB, err := hex.DecodeString(vec.B)
		if err != nil {
			t.Fatal(err)
		}

		extract(vecA, &extractedKey)
		assertEqual(t, i, vecB, extractedKey[:])
	}
}

// {K, tau, [N, A...], V) ==> AEZ-hash{K, {[tau]_128, N, A...)) = V
type HashVector struct {
	K    string   `json:"k"`
	Tau  int      `json:"tau"`
	Data []string `json:"data"`
	V    string   `json:"v"`
}

func TestHash(t *testing.T) {
	var e eState
	var hashVectors []HashVector

	readJsonTestdata(t, "hash.json", &hashVectors)

	for i, vec := range hashVectors {
		vecK, err := hex.DecodeString(vec.K)
		if err != nil {
			t.Fatal(err)
		}
		var data [][]byte
		for _, v := range vec.Data {
			d, err := hex.DecodeString(v)
			if err != nil {
				t.Fatal(err)
			}
			data = append(data, d)
		}
		vecV, err := hex.DecodeString(vec.V)
		if err != nil {
			t.Fatal(err)
		}

		var nonce []byte
		var ad [][]byte
		if len(data) > 0 {
			nonce = data[0]
			if len(data) > 1 {
				ad = data[1:]
			}
		}

		var result [blockSize]byte

		e.init(vecK)
		e.aezHash(nonce, ad, vec.Tau, result[:])
		assertEqual(t, i, vecV, result[:])
	}
}

// (K, delta, tau, R) ==> AEZ-prf(K, T, tau*8) = R where delta = AEZ-hash(K,T)
type PrfVector struct {
	K     string `json:"k"`
	Delta string `json:"delta"`
	Tau   int    `json:"tau"`
	R     string `json:"R"`
}

func TestPRF(t *testing.T) {
	var e eState
	var prfVectors []PrfVector

	readJsonTestdata(t, "prf.json", &prfVectors)

	for i, vec := range prfVectors {
		vecK, err := hex.DecodeString(vec.K)
		if err != nil {
			t.Fatal(err)
		}
		vecDelta, err := hex.DecodeString(vec.Delta)
		if err != nil {
			t.Fatal(err)
		}
		vecR, err := hex.DecodeString(vec.R)
		if err != nil {
			t.Fatal(err)
		}

		var vDelta [blockSize]byte
		copy(vDelta[:], vecDelta)

		result := make([]byte, len(vecR))
		e.init(vecK)
		e.aezPRF(&vDelta, vec.Tau, result)
		assertEqual(t, i, vecR, result)
	}
}

// (K, N, A, taubytes, M, C) ==> Encrypt(K,N,A,taubytes,M) = C
type EncryptVector struct {
	K     string   `json:"k"`
	Nonce string   `json:"nonce"`
	Data  []string `json:"data"`
	Tau   int      `json:"tau"`
	M     string   `json:"m"`
	C     string   `json:"c"`
}

func TestEncryptDecrypt(t *testing.T) {
	var encryptVectors []EncryptVector

	readJsonTestdata(t, "encrypt.json", &encryptVectors)
	assertEncrypt(t, encryptVectors)

	//
	// No AD test cases
	//
	readJsonTestdata(t, "encrypt_no_ad.json", &encryptVectors)
	assertEncrypt(t, encryptVectors)

	//
	// 33 bytes of AD test cases
	//
	readJsonTestdata(t, "encrypt_33_byte_ad.json", &encryptVectors)
	assertEncrypt(t, encryptVectors)

	//
	// 16 byte key test cases
	//
	readJsonTestdata(t, "encrypt_16_byte_key.json", &encryptVectors)
	assertEncrypt(t, encryptVectors)
}

func assertEncrypt(t *testing.T, vectors []EncryptVector) {
	var e eState

	for i, vec := range vectors {
		vecK, err := hex.DecodeString(vec.K)
		if err != nil {
			t.Fatal(err)
		}
		vecNonce, err := hex.DecodeString(vec.Nonce)
		if err != nil {
			t.Fatal(err)
		}
		var vecData [][]byte
		for _, s := range vec.Data {
			d, err := hex.DecodeString(s)
			if err != nil {
				t.Fatal(err)
			}
			vecData = append(vecData, d)
		}
		vecM, err := hex.DecodeString(vec.M)
		if err != nil {
			t.Fatal(err)
		}
		vecC, err := hex.DecodeString(vec.C)
		if err != nil {
			t.Fatal(err)
		}

		// Test the cipher.AEAD code as well, for applicable test vectors.
		var aead cipher.AEAD
		var ad []byte
		if len(vecNonce) == aeadNonceSize && vec.Tau == aeadOverhead && len(vecData) <= 1 {
			aead, err = New(vecK)
			if err != nil {
				t.Fatal(err)
			}
			if len(vecData) == 1 {
				ad = vecData[0]
			}
		}

		e.init(vecK)
		c := Encrypt(vecK, vecNonce, vecData, vec.Tau, vecM, nil)
		assertEqual(t, i, vecC, c)
		if aead != nil {
			ac := aead.Seal(nil, vecNonce, vecM, ad)
			assertEqual(t, i, vecC, ac)
		}

		m, ok := Decrypt(vecK, vecNonce, vecData, vec.Tau, vecC, nil)
		if !ok {
			t.Fatalf("decrypt failed: [%d]", i)
		}
		assertEqual(t, i, vecM, m)
		if aead != nil {
			am, err := aead.Open(nil, vecNonce, vecC, ad)
			if err != nil {
				t.Fatal(err)
			}
			assertEqual(t, i, vecM, am)
		}
	}
}

func assertEqual(t *testing.T, idx int, expected, actual []byte) {
	if !bytes.Equal(expected, actual) {
		for i, v := range actual {
			if expected[i] != v {
				t.Errorf("[%d] first mismatch at offset: %d (%02x != %02x)", idx, i, expected[i], v)
				break
			}
		}
		t.Errorf("expected: %s", hex.Dump(expected))
		t.Errorf("actual: %s", hex.Dump(actual))
		t.FailNow()
	}
}

var benchOutput []byte

func doBenchEncrypt(b *testing.B, n int) {
	var key [extractedKeySize]byte
	if _, err := rand.Read(key[:]); err != nil {
		b.Error(err)
		b.Fail()
	}

	const tau = 16

	var nonce [16]byte
	src := make([]byte, n)
	dst := make([]byte, n+tau)
	check := make([]byte, n+tau)

	b.SetBytes(int64(n))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		dst = Encrypt(key[:], nonce[:], nil, tau, src[:n], dst[:0])
		b.StopTimer()
		dec, ok := Decrypt(key[:], nonce[:], nil, tau, dst, check[:0])
		if !ok {
			b.Fatalf("decrypt failed")
		}
		if !bytes.Equal(dec, src) {
			b.Fatalf("decrypt produced invalid output")
		}
		copy(src, dst[:n])
	}

	benchOutput = src
}

func BenchmarkEncrypt(b *testing.B) {
	sizes := []int{1, 32, 512, 1024, 2048, 16384, 32768, 65536, 1024768}
	if testing.Short() {
		sizes = []int{1, 32, 512, 1024, 16384, 65536, 1024768}
	}

	b.SetParallelism(1) // AES-NI is a per-physical core thing.

	for _, sz := range sizes {
		n := fmt.Sprintf("%d", sz)
		b.Run(n, func(b *testing.B) { doBenchEncrypt(b, sz) })
	}
}
