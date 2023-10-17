package interation

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// 함수이름이 Benchmark로 시작하고 매개변수고 testing.B를 받으면 벤츠마크 함수로 인식한다.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ { //benchmarRepaet이 실행되면 b.N번만큼 반복된다  그리고 얼마나 걸리는지 측정됨
		Repeat("a")
	}
}
