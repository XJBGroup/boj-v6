package unittest

import "testing"

func TestLoad(t *testing.T) {
	load()
}

func BenchmarkLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Load("test.yaml", false)
	}
}

func BenchmarkLoadCached(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Load("test.yaml", true)
	}
}
