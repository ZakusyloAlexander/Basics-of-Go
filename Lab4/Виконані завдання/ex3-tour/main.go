// Лабораторна робота №4. Калькулятор вартості туру (GUI: andlabs/ui).
package main

import (
	"fmt"
	"strconv"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

// Ціни за 1 день ($): pricesPerDay[країна][сезон].
// країна: 0 — Болгарія, 1 — Німеччина, 2 — Польща.
// сезон: 0 — літо, 1 — зима.
var pricesPerDay = [3][2]float64{
	{100, 150}, // Болгарія
	{160, 200}, // Німеччина
	{120, 180}, // Польща
}

const guidePricePerDay = 50 // гід: $ за день (на всю групу путівок за методичкою)
const luxuryMarkup = 0.2    // націнка 20% лише до вартості проживання при номері люкс

// CalcTourPrice: сума = проживання (дні × путівки × ціна/день) з опцією люкс + гід за дні.
func CalcTourPrice(country, season int, days int, vouchers int, hasGuide, hasLuxury bool) float64 {
	if country < 0 || country > 2 || season < 0 || season > 1 || days <= 0 || vouchers <= 0 {
		return 0
	}
	basePerDay := pricesPerDay[country][season]
	// "Люкс" збільшує лише частину проживання, не базову ціну гіда окремо.
	accommodation := float64(days*vouchers) * basePerDay
	if hasLuxury {
		accommodation *= (1 + luxuryMarkup)
	}
	guideCost := 0.0
	if hasGuide {
		guideCost = float64(days) * guidePricePerDay
	}
	return accommodation + guideCost
}

func main() {
	err := ui.Main(func() {
		window := ui.NewWindow("Калькулятор туру", 420, 420, false)
		window.SetMargined(true)

		box := ui.NewVerticalBox()
		box.SetPadded(true)

		box.Append(ui.NewLabel("Країна:"), false)
		countryCombo := ui.NewCombobox()
		countryCombo.Append("Болгарія")
		countryCombo.Append("Німеччина")
		countryCombo.Append("Польща")
		countryCombo.SetSelected(0)
		box.Append(countryCombo, false)

		box.Append(ui.NewLabel("Сезон:"), false)
		seasonCombo := ui.NewCombobox()
		seasonCombo.Append("Літо")
		seasonCombo.Append("Зима")
		seasonCombo.SetSelected(0)
		box.Append(seasonCombo, false)

		box.Append(ui.NewLabel("Кількість днів:"), false)
		daysEntry := ui.NewEntry()
		daysEntry.SetText("7")
		box.Append(daysEntry, false)

		box.Append(ui.NewLabel("Кількість путівок (туристів):"), false)
		vouchersEntry := ui.NewEntry()
		vouchersEntry.SetText("2")
		box.Append(vouchersEntry, false)

		guideCheck := ui.NewCheckbox("Індивідуальний гід (+$50/день)")
		box.Append(guideCheck, false)

		luxuryCheck := ui.NewCheckbox("Номер люкс (+20% до проживання)")
		box.Append(luxuryCheck, false)

		calcBtn := ui.NewButton("Розрахувати")
		box.Append(calcBtn, false)

		resultLabel := ui.NewLabel("Вартість: —")
		box.Append(resultLabel, false)

		calcBtn.OnClicked(func(*ui.Button) {
			days, err1 := strconv.Atoi(daysEntry.Text())
			vouchers, err2 := strconv.Atoi(vouchersEntry.Text())
			if err1 != nil || err2 != nil || days <= 0 || vouchers <= 0 {
				resultLabel.SetText("Помилка: введіть коректну кількість днів та путівок")
				return
			}
			country := countryCombo.Selected()
			season := seasonCombo.Selected()
			hasGuide := guideCheck.Checked()
			hasLuxury := luxuryCheck.Checked()
			total := CalcTourPrice(country, season, days, vouchers, hasGuide, hasLuxury)
			resultLabel.SetText(fmt.Sprintf("Вартість: $%.2f", total))
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
