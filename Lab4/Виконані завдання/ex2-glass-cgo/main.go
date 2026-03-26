// Лабораторна робота №4. Калькулятор склопакета.
// Варіант 2: обчислення в C (calc.c), виклик через cgo; GUI — andlabs/ui.
package main

// Директиви cgo: CFLAGS додає поточну теку до include-шляху; calc.h підключає оголошення C-функції.
// Реалізація calc_glass_price збирається разом із проєктом (див. calc.c у цій папці).
/*
#cgo CFLAGS: -I.
#include "calc.h"
*/
import "C"

import (
	"fmt"
	"strconv"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// CalcGlassPrice — обгортка Go над C.calc_glass_price: перетворює типи Go → C і результат назад у float64.
// bool у C передаємо як int 0/1.
func CalcGlassPrice(width, height float64, glassType, frameType int, hasSill bool) float64 {
	sill := 0
	if hasSill {
		sill = 1
	}

	return float64(
		C.calc_glass_price(
			C.double(width),
			C.double(height),
			C.int(glassType),
			C.int(frameType),
			C.int(sill),
		),
	)
}

func main() {
	err := ui.Main(func() {
		window := ui.NewWindow("Калькулятор склопакета (C + Go)", 400, 350, false)
		window.SetMargined(true)

		box := ui.NewVerticalBox()
		box.SetPadded(true)

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
		glassCombo.SetSelected(0)
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

		calcBtn.OnClicked(func(*ui.Button) {
			width, err1 := strconv.ParseFloat(widthEntry.Text(), 64)
			height, err2 := strconv.ParseFloat(heightEntry.Text(), 64)

			if err1 != nil || err2 != nil || width <= 0 || height <= 0 {
				resultLabel.SetText("Помилка: введіть коректні значення")
				return
			}

			total := CalcGlassPrice(
				width,
				height,
				glassCombo.Selected(),
				frameCombo.Selected(),
				sillCheck.Checked(),
			)

			resultLabel.SetText(fmt.Sprintf("Вартість: %.2f грн", total))
		})

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.SetChild(box)
		window.Show()
	})

	if err != nil {
		panic(err)
	}
}
