package main

// Генератор случайных чисел в заданном диапазоне
func temperatureConvert(value float64, sign string) float64 {
	var result float64

	if sign == "c" {
		result = (value - 32.0) / 1.8
	}

	if sign == "f" {
		result = value*1.8 + 32.0
	}

	return result
}
