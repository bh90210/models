package elektronmodels

import "testing"

var x map[voice]*track

func BenchmarkEmptyMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = make(map[voice]*track)
	}
}

// func BenchmarkNewProject(b *testing.B) {
// 	b.StopTimer()

// 	b.StartTimer()
// 	p := NewProject(CYCLES)
// 	p.PatternInit(0)
// 	b.StopTimer()
// }
