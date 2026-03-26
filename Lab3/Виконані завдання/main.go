// Лабораторна робота №3. Керуючі конструкції, функції, псевдовипадкові числа.
// Реалізовано лінійний конгруентний генератор (ЛКГ) та статистичну обробку послідовностей.
// Перед обчисленнями користувач обирає варіант з таблиці 3.2 (1–20).
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Variant зберігає параметри одного рядка таблиці 3.2 для ЛКГ.
// Формула: X_{k+1} = (A*X_k + C) mod M.
type Variant struct {
	A    int64 // множник у рекурентній формулі
	C    int64 // приріст (additive constant)
	M    int64 // модуль; зазвичу велике 2^n або 2^n-1
	N    int   // верхня межа діапазону значень [0, N) для цілих / масштабу для дійсних
	K    int   // довжина генерованої послідовності
	MStr string // рядкове позначення m (для виводу в таблиці)
}

// variants — усі 20 варіантів з методички (параметри a, c, m, n, k).
var variants = []Variant{
	{1103515245, 12345, 1 << 31, 100, 10000, "2^31"},
	{1664525, 1013904223, 1 << 32, 150, 20000, "2^32"},
	{16807, 0, 1<<31 - 1, 200, 30000, "2^31-1"},
	{22695477, 1, 1 << 32, 250, 40000, "2^32"},
	{48271, 0, 1<<31 - 1, 300, 50000, "2^31-1"},
	{214013, 2531011, 1 << 32, 100, 50000, "2^32"},
	{2147483629, 2147483587, 1<<31 - 1, 150, 40000, "2^31-1"},
	{65539, 0, 1 << 32, 200, 30000, "2^32"},
	{1140671485, 12820163, 1 << 24, 250, 20000, "2^24"},
	{69069, 1, 1 << 32, 300, 10000, "2^32"},
	{1103515245, 12345, 1 << 31, 300, 50000, "2^31"},
	{1664525, 1013904223, 1 << 32, 250, 40000, "2^32"},
	{16807, 0, 1<<31 - 1, 200, 30000, "2^31-1"},
	{22695477, 1, 1 << 32, 150, 20000, "2^32"},
	{48271, 0, 1<<31 - 1, 100, 10000, "2^31-1"},
	{214013, 2531011, 1 << 32, 300, 10000, "2^32"},
	{2147483629, 2147483587, 1<<31 - 1, 250, 20000, "2^31-1"},
	{65539, 0, 1 << 32, 200, 30000, "2^32"},
	{1140671485, 12820163, 1 << 24, 150, 40000, "2^24"},
	{134775813, 1, 1 << 32, 100, 50000, "2^32"},
}

// nextInt — один крок ЛКГ: повертає наступне ціле в [0, m), оновлюючи *seed.
// У мові Go передача *int64 дозволяє змінювати змінну-зерно поза функцією.
func nextInt(seed *int64, a, c, m int64) int64 {
	next := (a*(*seed) + c) % m
	if next < 0 {
		// У Go % для від'ємних чисел може дати від'ємний залишок — зсуваємо в [0, m).
		next += m
	}
	*seed = next
	return next
}

// generateIntSequence будує K цілих чисел у діапазоні [0, N): беремо nextInt mod N.
func generateIntSequence(seed *int64, v Variant) []int {
	seq := make([]int, v.K)
	for i := 0; i < v.K; i++ {
		x := nextInt(seed, v.A, v.C, v.M)
		seq[i] = int(x % int64(v.N))
	}
	return seq
}

// generateRealSequence нормує nextInt до [0, N) у вигляді дійсного числа.
func generateRealSequence(seed *int64, v Variant) []float64 {
	seq := make([]float64, v.K)
	for i := 0; i < v.K; i++ {
		x := nextInt(seed, v.A, v.C, v.M)
		seq[i] = (float64(x) / float64(v.M)) * float64(v.N)
	}
	return seq
}

func main() {
	fmt.Println("=== Лабораторна робота №3 ===")
	fmt.Println("Лінійний конгруентний метод: X_{k+1} = (a*X_k + c) mod m")
	fmt.Println()
	fmt.Println("Таблиця 3.2 — варіанти (1–20):")
	fmt.Println("  Варіант |    a         |      c       |   m     | Діапазон  |    K")
	fmt.Println("  --------|--------------|--------------|---------|------------|-------")
	for i := range variants {
		v := &variants[i]
		fmt.Printf("  %7d | %12d | %12d | %-7s | [0, %-5d) | %5d\n",
			i+1, v.A, v.C, v.MStr, v.N, v.K)
	}
	fmt.Println()
	fmt.Print("Введіть номер варіанту (1–20): ")

	// Читання рядка з клавіатури (до символу нового рядка).
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Помилка читання:", err)
		return
	}
	line = strings.TrimSpace(line)
	num, err := strconv.Atoi(line)
	if err != nil || num < 1 || num > 20 {
		fmt.Println("Помилка: потрібно число від 1 до 20.")
		return
	}

	v := variants[num-1]
	// Початкове зерно X_0: беремо c, але приводимо до допустимого діапазону (0, m).
	x0 := v.C
	if x0 <= 0 {
		x0 = 1
	}
	if x0 >= v.M {
		x0 = v.M - 1
	}

	fmt.Printf("\nОбрано варіант %d: a=%d, c=%d, m=%s, діапазон [0, %d), K=%d\n\n",
		num, v.A, v.C, v.MStr, v.N, v.K)

	// ----- Завдання 1: цілочисельна послідовність і статистика -----
	seed := x0
	seq := generateIntSequence(&seed, v)
	n := v.N

	// freq[i] — скільки разів у послідовності зустрілось значення i.
	freq := make([]int, n)
	for _, val := range seq {
		if val >= 0 && val < n {
			freq[val]++
		}
	}

	fmt.Println("--- Завдання 1: обробка цілочисельної послідовності ---")
	show := 10
	if n < show {
		show = n
	}
	fmt.Printf("1) Частота інтервалів (перші %d значень 0..%d):\n", show, show-1)
	for i := 0; i < show; i++ {
		fmt.Printf("   значення %d: частота = %d\n", i, freq[i])
	}
	if n > show {
		fmt.Println("   ...")
	}

	// P_i = n_i / K — емпірична (статистична) ймовірність.
	prob := make([]float64, n)
	for i := 0; i < n; i++ {
		prob[i] = float64(freq[i]) / float64(v.K)
	}
	fmt.Printf("\n2) Статистична ймовірність P_i = частота_i / K (перші 5):\n")
	for i := 0; i < 5 && i < n; i++ {
		fmt.Printf("   P(%d) = %.6f\n", i, prob[i])
	}

	// M(X) = sum_i i * P(i) — математичне сподівання дискретної випадкової величини.
	var mx float64
	for i := 0; i < n; i++ {
		mx += float64(i) * prob[i]
	}
	fmt.Printf("\n3) Математичне сподівання M(X) = %.6f\n", mx)

	// D(X) = sum_i P(i) * (i - M)^2 — дисперсія.
	var dx float64
	for i := 0; i < n; i++ {
		diff := float64(i) - mx
		dx += prob[i] * (diff * diff)
	}
	fmt.Printf("4) Дисперсія D(X) = %.6f\n", dx)

	// sigma = sqrt(D(X)) — середньоквадратичне відхилення.
	sigma := math.Sqrt(dx)
	fmt.Printf("5) Середньоквадратичне відхилення sigma = %.6f\n", sigma)

	// ----- Завдання 2: дійсна послідовність -----
	fmt.Println("\n--- Завдання 2: послідовність дійсних псевдовипадкових значень ---")
	seed2 := x0
	realSeq := generateRealSequence(&seed2, v)
	showReal := 10
	if len(realSeq) < showReal {
		showReal = len(realSeq)
	}
	fmt.Printf("Перші %d значень у діапазоні [0, %d):\n", showReal, v.N)
	for i := 0; i < showReal; i++ {
		fmt.Printf("   [%d] = %.4f\n", i, realSeq[i])
	}

	// Окремий прохід генерації для обчислення вибіркового середнього та дисперсії дійсних значень.
	seed3 := x0
	realLong := generateRealSequence(&seed3, v)
	var sum float64
	for _, val := range realLong {
		sum += val
	}
	meanReal := sum / float64(len(realLong))
	var sumSq float64
	for _, val := range realLong {
		d := val - meanReal
		sumSq += d * d
	}
	// Вибіркова дисперсія: середнє квадратів відхилень від середнього (тут без Bessel correction).
	varReal := sumSq / float64(len(realLong))
	fmt.Printf("\nДля K=%d дійсних значень: середнє = %.4f, дисперсія = %.4f\n", v.K, meanReal, varReal)
}
