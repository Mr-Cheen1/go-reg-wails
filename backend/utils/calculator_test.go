package utils

import "testing"

func TestCalculateTime(t *testing.T) {
	tests := []struct {
		name     string
		timeStr  string
		expected float64
	}{
		{
			name:     "Пустая строка",
			timeStr:  "",
			expected: 0,
		},
		{
			name:     "Строка с пробелами",
			timeStr:  "   ",
			expected: 0,
		},
		{
			name:     "Одно число",
			timeStr:  "1.5",
			expected: 1.5,
		},
		{
			name:     "Сложение двух чисел",
			timeStr:  "1.5+2.5",
			expected: 4.0,
		},
		{
			name:     "Сложение нескольких чисел",
			timeStr:  "1.2+2.3+3.4",
			expected: 6.9,
		},
		{
			name:     "Числа с пробелами",
			timeStr:  " 1.5 + 2.5 ",
			expected: 4.0,
		},
		{
			name:     "Невалидные части строки игнорируются",
			timeStr:  "1.5+abc+2.5",
			expected: 4.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateTime(tt.timeStr)
			// Из-за особенностей работы с плавающей точкой, используем погрешность
			if result != tt.expected {
				t.Errorf("CalculateTime() = %v, want %v", result, tt.expected)
			}
		})
	}
}
