// Лабораторна робота №5: структури, методи, вбудовані типи, інтерфейс fmt.Stringer.
// Програма вводить з клавіатури список товарів (різні валюти), виводить їх і знаходить
// найдешевший та найдорожчий за загальною вартістю в гривнях.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Глобальний читач рядків з stdin — один раз створюємо, далі використовуємо в readLine.
var reader = bufio.NewReader(os.Stdin)

// ---------------------- Допоміжні функції вводу ----------------------

func readLine(prompt string) string {
	fmt.Print(prompt)
	s, _ := reader.ReadString('\n') // на кінці рядка може бути \n; помилку ігноруємо для простоти
	s = strings.TrimSpace(s)
	// Порожній рядок для текстових полів — допустимий варіант.
	return s
}

func parseFloatFlexible(s string) (float64, error) {
	// Користувач може ввести десяткову кому — замінюємо на крапку для strconv.
	s = strings.ReplaceAll(strings.TrimSpace(s), ",", ".")
	return strconv.ParseFloat(s, 64)
}

func readFloat(prompt string, nonNegative bool, positive bool) float64 {
	for {
		s := readLine(prompt)
		v, err := parseFloatFlexible(s)
		if err != nil {
			fmt.Println("Некоректне число. Спробуйте ще раз.")
			continue
		}
		if nonNegative && v < 0 {
			fmt.Println("Число не може бути від'ємним. Спробуйте ще раз.")
			continue
		}
		if positive && v <= 0 {
			fmt.Println("Число має бути додатним. Спробуйте ще раз.")
			continue
		}
		return v
	}
}

func readUint(prompt string) uint64 {
	for {
		s := readLine(prompt)
		v, err := strconv.ParseUint(strings.TrimSpace(s), 10, 64)
		if err != nil {
			fmt.Println("Некоректне ціле додатне значення. Спробуйте ще раз.")
			continue
		}
		return v
	}
}

// ---------------------- Структури ----------------------

// Currency описує валюту товару та курс до гривні.
type Currency struct {
	Name   string  // коротка назва (USD, EUR, …)
	ExRate float64 // скільки грн за 1 одиницю валюти
}

// NewCurrency — конструктор (фабрична функція замість обов'язкового New у інших мовах).
func NewCurrency(name string, exRate float64) Currency {
	return Currency{Name: name, ExRate: exRate}
}

// Сеттери з приймачем *Currency змінюють поля оригінальної структури.
// Геттери з приймачем Currency (копія) безпечні для читання невеликих структур.
func (c *Currency) SetName(name string) { c.Name = name }
func (c Currency) GetName() string      { return c.Name }

func (c *Currency) SetExRate(exRate float64) { c.ExRate = exRate }
func (c Currency) GetExRate() float64         { return c.ExRate }

// ToUAH перекладає суму з валюти в гривні за поточним курсом.
func (c Currency) ToUAH(amount float64) float64 {
	return amount * c.ExRate
}

// String реалізує інтерфейс fmt.Stringer — fmt.Println автоматично викличе цей метод.
func (c Currency) String() string {
	return fmt.Sprintf("%s (курс: %.4f грн за 1)", c.Name, c.ExRate)
}

// Product — товар на складі; поле Cost типу Currency — композиція (вкладена структура).
type Product struct {
	Name     string
	Price    float64  // ціна одиниці в одиницях валюти Cost
	Cost     Currency // валюта, в якій задана Price
	Quantity uint64
	Producer string
	Weight   float64 // вага однієї одиниці товару
}

func NewProduct(name string, price float64, cost Currency, quantity uint64, producer string, weight float64) Product {
	return Product{
		Name:     name,
		Price:    price,
		Cost:     cost,
		Quantity: quantity,
		Producer: producer,
		Weight:   weight,
	}
}

// Методи доступу до полів Product (аналогічно Currency).
func (p *Product) SetName(name string)        { p.Name = name }
func (p Product) GetName() string             { return p.Name }
func (p *Product) SetPrice(price float64)      { p.Price = price }
func (p Product) GetPrice() float64            { return p.Price }
func (p *Product) SetCost(cost Currency)       { p.Cost = cost }
func (p Product) GetCost() Currency            { return p.Cost }
func (p *Product) SetQuantity(q uint64)     { p.Quantity = q }
func (p Product) GetQuantity() uint64      { return p.Quantity }
func (p *Product) SetProducer(s string)     { p.Producer = s }
func (p Product) GetProducer() string       { return p.Producer }
func (p *Product) SetWeight(w float64)     { p.Weight = w }
func (p Product) GetWeight() float64       { return p.Weight }

// ---------------------- Методи обчислення для Product ----------------------

// GetPriceIn — ціна однієї одиниці товару в гривнях (Price у валюті × курс).
func (p Product) GetPriceIn() float64 {
	return p.Cost.ToUAH(p.Price)
}

// GetTotalPrice — загальна вартість усіх одиниць на складі (грн).
func (p Product) GetTotalPrice() float64 {
	return p.GetPriceIn() * float64(p.Quantity)
}

// GetTotalWeight — сумарна вага всіх одиниць (вага одиниці × кількість).
func (p Product) GetTotalWeight() float64 {
	return p.Weight * float64(p.Quantity)
}

// String — детальний рядок для виводу (знову fmt.Stringer).
func (p Product) String() string {
	return fmt.Sprintf(
		"%s | Виробник: %s | Ціна: %.2f %s (%.2f грн) | К-сть: %d | Вага одиниці: %.2f | Разом: %.2f грн",
		p.Name,
		p.Producer,
		p.Price,
		p.Cost.Name,
		p.GetPriceIn(),
		p.Quantity,
		p.Weight,
		p.GetTotalPrice(),
	)
}

// ---------------------- Функції за завданням ----------------------

// ReadProductsArray зчитує n товарів; для кожного можна задати свою валюту та курс.
func ReadProductsArray() []Product {
	n := int(readUint("Введіть кількість товарів n: "))
	products := make([]Product, 0, n) // len=0, cap=n — додаємо через append

	for i := 0; i < n; i++ {
		fmt.Printf("\nТовар #%d\n", i+1)
		name := readLine("  Назва товару: ")
		price := readFloat("  Вартість одиниці (в валюті Cost): ", true, false)

		curName := readLine("  Назва валюти (Cost): ")
		exRate := readFloat("  Курс валюти (грн за 1 Cost): ", true, true) // курс > 0

		quantity := readUint("  Кількість на складі: ")
		producer := readLine("  Виробник: ")
		weight := readFloat("  Вага одиниці: ", true, false) // >= 0

		currency := NewCurrency(curName, exRate)
		products = append(products, NewProduct(name, price, currency, quantity, producer, weight))
	}

	return products
}

func PrintProduct(p Product) {
	fmt.Println(p.String())
	fmt.Printf("  Загальна вага: %.2f\n", p.GetTotalWeight())
}

func PrintProducts(products []Product) {
	fmt.Println("\nСписок товарів:")
	for _, p := range products {
		PrintProduct(p)
		fmt.Println()
	}
}

// GetProductsInfo записує найдешевший і найдорожчий товар за GetTotalPrice() (грн).
// Параметри cheapest та mostExpensive — вказівники, щоб повернути два результати зі зміни зовнішніх змінних.
func GetProductsInfo(products []Product, cheapest *Product, mostExpensive *Product) {
	if len(products) == 0 || cheapest == nil || mostExpensive == nil {
		return
	}

	min := products[0]
	max := products[0]

	for i := 1; i < len(products); i++ {
		if products[i].GetTotalPrice() < min.GetTotalPrice() {
			min = products[i]
		}
		if products[i].GetTotalPrice() > max.GetTotalPrice() {
			max = products[i]
		}
	}

	*cheapest = min
	*mostExpensive = max
}

func main() {
	products := ReadProductsArray()
	PrintProducts(products)

	var cheapest Product
	var mostExpensive Product
	GetProductsInfo(products, &cheapest, &mostExpensive) // передаємо адреси, щоб функція заповнила змінні

	fmt.Println("\nНайдешевший товар (за загальною вартістю в грн):")
	PrintProduct(cheapest)

	fmt.Println("\nНайдорожчий товар (за загальною вартістю в грн):")
	PrintProduct(mostExpensive)
}
