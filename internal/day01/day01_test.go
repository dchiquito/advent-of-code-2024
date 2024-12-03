package day01

import (
	"bytes"
	"os"
	"testing"
)

func BenchmarkParse(b *testing.B) {
	file, _ := os.ReadFile("data/01.txt")
	b.ResetTimer()
	for range b.N {
		reader := bytes.NewReader(file)
		parse(reader)
	}
}
func BenchmarkLevel1(b *testing.B) {
	file, _ := os.ReadFile("data/01.txt")
	b.ResetTimer()
	for range b.N {
		reader := bytes.NewReader(file)
		Level1(reader)
	}
}
func BenchmarkLevel2(b *testing.B) {
	file, _ := os.ReadFile("data/01.txt")
	b.ResetTimer()
	for range b.N {
		reader := bytes.NewReader(file)
		Level2(reader)
	}
}
