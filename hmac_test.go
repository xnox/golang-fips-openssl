//go:build linux && !android
// +build linux,!android

package openssl

import (
	"bytes"
	"hash"
	"testing"
)

func TestHMAC(t *testing.T) {
	var tests = []struct {
		name string
		fn   func() hash.Hash
	}{
		{"sha1", NewSHA1},
		{"sha224", NewSHA224},
		{"sha256", NewSHA256},
		{"sha384", NewSHA384},
		{"sha512", NewSHA512},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			h := NewHMAC(tt.fn, nil)
			h.Write([]byte("hello"))
			sumHello := h.Sum(nil)

			h = NewHMAC(tt.fn, nil)
			h.Write([]byte("hello world"))
			sumHelloWorld := h.Sum(nil)

			// Test that Sum has no effect on future Sum or Write operations.
			// This is a bit unusual as far as usage, but it's allowed
			// by the definition of Go hash.Hash, and some clients expect it to work.
			h = NewHMAC(tt.fn, nil)
			h.Write([]byte("hello"))
			if sum := h.Sum(nil); !bytes.Equal(sum, sumHello) {
				t.Fatalf("1st Sum after hello = %x, want %x", sum, sumHello)
			}
			if sum := h.Sum(nil); !bytes.Equal(sum, sumHello) {
				t.Fatalf("2nd Sum after hello = %x, want %x", sum, sumHello)
			}

			h.Write([]byte(" world"))
			if sum := h.Sum(nil); !bytes.Equal(sum, sumHelloWorld) {
				t.Fatalf("1st Sum after hello world = %x, want %x", sum, sumHelloWorld)
			}
			if sum := h.Sum(nil); !bytes.Equal(sum, sumHelloWorld) {
				t.Fatalf("2nd Sum after hello world = %x, want %x", sum, sumHelloWorld)
			}

			h.Reset()
			h.Write([]byte("hello"))
			if sum := h.Sum(nil); !bytes.Equal(sum, sumHello) {
				t.Fatalf("Sum after Reset + hello = %x, want %x", sum, sumHello)
			}
		})
	}
}

func BenchmarkHMACSHA256_32(b *testing.B) {
	b.StopTimer()
	key := make([]byte, 32)
	buf := make([]byte, 32)
	h := NewHMAC(NewSHA256, key)
	b.SetBytes(int64(len(buf)))
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		h.Write(buf)
		mac := h.Sum(nil)
		h.Reset()
		buf[0] = mac[0]
	}
}

func BenchmarkHMACNewWriteSum(b *testing.B) {
	b.StopTimer()
	buf := make([]byte, 32)
	b.SetBytes(int64(len(buf)))
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		h := NewHMAC(NewSHA256, make([]byte, 32))
		h.Write(buf)
		mac := h.Sum(nil)
		buf[0] = mac[0]
	}
}
