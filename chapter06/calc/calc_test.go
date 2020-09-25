package calc

import "testing"

func TestAdd(t *testing.T) {
	if 3 != add(1, 2) {
		t.Error("1 plus 2 != 3")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(1, 2)
	}
}
