package must

import "testing"

func BenchmarkRetrunWithVal(b *testing.B){
	for i := 0; i < b.N; i++ {
		_=retrunWithVal()
	}
}
func BenchmarkRetrunWithPoint(b *testing.B){
	for i := 0; i < b.N; i++ {
		_=retrunWithVal()
	}
}