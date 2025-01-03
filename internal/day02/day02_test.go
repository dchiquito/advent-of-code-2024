package day02

import (
	"bytes"
	"os"
	"testing"
)

func readData() []byte {
	file, _ := os.ReadFile("data/02.txt")
	return file
}

func BenchmarkParse(b *testing.B) {
	file := readData()
	b.ResetTimer()
	for range b.N {
		reader := bytes.NewReader(file)
		parse(reader)
	}
}
func BenchmarkLevel1(b *testing.B) {
	file := readData()
	b.ResetTimer()
	for range b.N {
		reader := bytes.NewReader(file)
		Level1(reader)
	}
}
func BenchmarkLevel2(b *testing.B) {
	file := readData()
	b.ResetTimer()
	for range b.N {
		reader := bytes.NewReader(file)
		Level2(reader)
	}
}
