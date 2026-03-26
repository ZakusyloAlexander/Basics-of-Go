// Пункт 6: тести для пакета calc (файл *_test.go, package calc — той самий пакет).
// Запуск: з каталогу модуля виконати `go test ./...` або `go test ./calc -v`.
package calc

import "testing"

func TestMin3(t *testing.T) {
	// Табличне тестування: кілька вхідних наборів у слайсі анонімних структур.
	tests := []struct {
		a, b, c float64
		want    float64
	}{
		{5, 2, 8, 2},
		{-1, -5, -3, -5},
		{0, 0, 0, 0},
		{3, 3, 3, 3},
	}
	for _, tt := range tests {
		got := Min3(tt.a, tt.b, tt.c)
		if got != tt.want {
			// t.Errorf не зупиняє тест одразу — можна побачити інші помилки.
			t.Errorf("Min3(%v, %v, %v) = %v, потрібно %v", tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}

func TestAverage3(t *testing.T) {
	tests := []struct {
		a, b, c float64
		want    float64
	}{
		{1, 2, 3, 2},
		{0, 0, 0, 0},
		{3, 6, 9, 6},
	}
	for _, tt := range tests {
		got := Average3(tt.a, tt.b, tt.c)
		if got != tt.want {
			t.Errorf("Average3(%v, %v, %v) = %v, потрібно %v", tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}

func TestSolveLinear(t *testing.T) {
	// Нормальний випадок: 2x + 4 = 0 => x = -2.
	x, err := SolveLinear(2, 4)
	if err != nil {
		t.Fatalf("SolveLinear(2, 4) помилка: %v", err)
	}
	if x != -2 {
		t.Errorf("SolveLinear(2, 4) = %v, потрібно -2", x)
	}

	// -3x + 6 = 0 => x = 2.
	x2, err2 := SolveLinear(-3, 6)
	if err2 != nil {
		t.Fatalf("SolveLinear(-3, 6) помилка: %v", err2)
	}
	if x2 != 2 {
		t.Errorf("SolveLinear(-3, 6) = %v, потрібно 2", x2)
	}

	// a = 0, b ≠ 0 — очікуємо помилку.
	_, err3 := SolveLinear(0, 5)
	if err3 == nil {
		t.Error("SolveLinear(0, 5) повинен повертати помилку")
	}

	// a = 0, b = 0 — також помилка (нескінченно багато розв'язків).
	_, err4 := SolveLinear(0, 0)
	if err4 == nil {
		t.Error("SolveLinear(0, 0) повинен повертати помилку")
	}
}

// Benchmark* — вимірювання швидкості; b.N керує кількістю ітерацій.
func BenchmarkMin3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Min3(5, 2, 8)
	}
}

func BenchmarkAverage3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Average3(1, 2, 3)
	}
}

func BenchmarkSolveLinear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolveLinear(2, 4)
	}
}
