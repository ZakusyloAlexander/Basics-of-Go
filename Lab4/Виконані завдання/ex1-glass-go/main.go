// Лабораторна робота №4. Калькулятор склопакета.
// Варіант 1: розрахунок функціями мовою Go, інтерфейс — пакет github.com/andlabs/ui.
package main

import (
	"fmt"
	"strconv"

	"github.com/andlabs/ui"
	// Порожній імпорт winmanifest — ресурси Windows для коректного відображення GUI.
	_ "github.com/andlabs/ui/winmanifest"
)

// Ціни за 1 см² склопакета: prices[тип_скла][матеріал_рами].
// тип: 0 — однокамерний, 1 — двокамерний.
// матеріал: 0 — дерев'яний, 1 — металевий, 2 — металопластиковий.
var prices = [2][3]float64{
	{2.5, 0.5, 1.5}, // однокамерний
	{3, 1, 2},       // двокамерний
}

const sillPrice = 350 // вартість підвіконня (грн), додається за прапорцем

// CalcGlassPrice: площа (см²) × ціна за см² + підвіконня за потреби.
func CalcGlassPrice(width, height float64, glassType, frameType int, hasSill bool) float64 {
	if glassType < 0 || glassType > 1 || frameType < 0 || frameType > 2 {
		return 0
	}
	total := width * height * prices[glassType][frameType]
	if hasSill {
		total += sillPrice
	}
	return total
}

func main() {
	// ui.Main запускає цикл подій графічної бібліотеки; усі віджети створюємо всередині callback.
	err := ui.Main(func() {
		// NewWindow(заголовок, ширина, висота, чи є меню).
		window := ui.NewWindow("Калькулятор склопакета", 400, 350, false)
		window.SetMargined(true) // відступи від країв вікна до вмісту

		// Вертикальний Box — стовпчик елементів керування.
		box := ui.NewVerticalBox()
		box.SetPadded(true) // проміжки між дочірніми віджетами

		box.Append(ui.NewLabel("Ширина (см):"), false)
		widthEntry := ui.NewEntry()
		widthEntry.SetText("100")
		box.Append(widthEntry, false)

		box.Append(ui.NewLabel("Висота (см):"), false)
		heightEntry := ui.NewEntry()
		heightEntry.SetText("120")
		box.Append(heightEntry, false)

		box.Append(ui.NewLabel("Тип склопакета:"), false)
		glassCombo := ui.NewCombobox()
		glassCombo.Append("Однокамерний")
		glassCombo.Append("Двокамерний")
		glassCombo.SetSelected(0) // індекс обраного рядка (0 — перший)
		box.Append(glassCombo, false)

		box.Append(ui.NewLabel("Матеріал рами:"), false)
		frameCombo := ui.NewCombobox()
		frameCombo.Append("Дерев'яний")
		frameCombo.Append("Металевий")
		frameCombo.Append("Металопластиковий")
		frameCombo.SetSelected(0)
		box.Append(frameCombo, false)

		sillCheck := ui.NewCheckbox("Підвіконня (+350 грн)")
		box.Append(sillCheck, false)

		calcBtn := ui.NewButton("Розрахувати")
		box.Append(calcBtn, false)

		resultLabel := ui.NewLabel("Вартість: —")
		box.Append(resultLabel, false)

		// Обробник кліку: читаємо текст з Entry, індекси з ComboBox, стан з Checkbox.
		calcBtn.OnClicked(func(*ui.Button) {
			width, err1 := strconv.ParseFloat(widthEntry.Text(), 64)
			height, err2 := strconv.ParseFloat(heightEntry.Text(), 64)
			if err1 != nil || err2 != nil || width <= 0 || height <= 0 {
				resultLabel.SetText("Помилка: введіть коректні ширину та висоту")
				return
			}
			glassType := glassCombo.Selected()
			frameType := frameCombo.Selected()
			hasSill := sillCheck.Checked()
			total := CalcGlassPrice(width, height, glassType, frameType, hasSill)
			resultLabel.SetText(fmt.Sprintf("Вартість: %.2f грн", total))
		})

		// Закриття вікна — завершуємо цикл подій через ui.Quit().
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true // true — дозволити закрити вікно
		})

		window.SetChild(box)
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
