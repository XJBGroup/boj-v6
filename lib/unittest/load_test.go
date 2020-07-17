package unittest

import "testing"

func TestLoad(t *testing.T) {
	load(V1Opt)
}

func BenchmarkLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Load("test.yaml", false, V1Opt)
	}
}

func BenchmarkLoadCached(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Load("test.yaml", true, V1Opt)
	}
}
