package main

import (
	"testing"
)

func Benchmark_concurrentArchivator(b *testing.B) {
	for i:=0; i< b.N; i++{
		concurrent([]string{
			"one.txt",
			"two.txt",
			"three.txt",
			"one.txt",
			"two.txt",
			"three.txt",
			"one.txt",
			"two.txt",
			"three.txt",
			"one.txt",
			"two.txt",
			"three.txt",
			"one.txt",
			"two.txt",
			"three.txt",
			"one.txt",
			"two.txt",
			"three.txt",
		})
	}
}

func Benchmark_sequencedArchivator(b *testing.B) {
	for i:=0; i< b.N; i++{
		sequenced([]string{
			"one.txt",
			"two.txt",
			"three.txt",
			"one.txt",
			"two.txt",
			"three.txt",
			"one.txt",
			"two.txt",
			"three.txt",
			"one.txt",
			"two.txt",
			"three.txt",
			"one.txt",
			"two.txt",
			"three.txt",
		})
	}
}