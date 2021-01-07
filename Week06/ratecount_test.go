package Week06

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWindow_Count(t *testing.T) {
	w, _ := NewWindow(10, 3)
	w.Add(10)
	w.Add(20)
	w.Add(30)
	w.Add(40)
	require.Equal(t, float64(90), w.Count())
}

func BenchmarkWindow_Count(b *testing.B) {
	w, _ := NewWindow(10, 3)
	for i := 0; i < b.N; i++ {
		w.Add(float64(i))
		w.Count()
	}
}
