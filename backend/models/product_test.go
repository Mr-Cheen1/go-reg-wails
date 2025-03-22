package models

import (
	"reflect"
	"testing"
)

func TestProductsSearch(t *testing.T) {
	products := Products{
		{ID: 1, Name: "Продукт 1", ProcessingTime: 1.5, TimeCalculation: "1.5"},
		{ID: 2, Name: "Другой продукт", ProcessingTime: 2.0, TimeCalculation: "2.0"},
		{ID: 3, Name: "Продукт специальный", ProcessingTime: 3.0, TimeCalculation: "1.5+1.5"},
	}

	tests := []struct {
		name     string
		query    string
		expected int   // теперь ожидаем количество продуктов, а не точный список
		ids      []int // ожидаемые ID продуктов
	}{
		{
			name:     "Пустой запрос возвращает все продукты",
			query:    "",
			expected: 3,
			ids:      []int{1, 2, 3},
		},
		{
			name:     "Поиск по 'продукт' возвращает результаты с этим словом",
			query:    "продукт",
			expected: 3, // все продукты содержат слово "продукт"
			ids:      []int{1, 2, 3},
		},
		{
			name:     "Поиск по 'другой' возвращает один продукт",
			query:    "другой",
			expected: 1,
			ids:      []int{2},
		},
		{
			name:     "Поиск по 'специальный' возвращает один продукт",
			query:    "специальный",
			expected: 1,
			ids:      []int{3},
		},
		{
			name:     "Поиск по 'несуществующий' возвращает пустой результат",
			query:    "несуществующий",
			expected: 0,
			ids:      []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := products.Search(tt.query)

			// Проверяем количество найденных продуктов
			if len(result) != tt.expected {
				t.Errorf("Products.Search() количество элементов = %d, want %d", len(result), tt.expected)
				return
			}

			// Проверяем, что все нужные элементы присутствуют
			if len(tt.ids) > 0 {
				foundIDs := make(map[int]bool)
				for _, product := range result {
					foundIDs[product.ID] = true
				}

				for _, id := range tt.ids {
					if !foundIDs[id] {
						t.Errorf("Products.Search() не содержит продукт с ID %d", id)
					}
				}
			}
		})
	}
}

func TestProductsDelete(t *testing.T) {
	tests := []struct {
		name     string
		products Products
		id       int
		expected Products
	}{
		{
			name: "Удаление существующего продукта",
			products: Products{
				{ID: 1, Name: "Продукт 1"},
				{ID: 2, Name: "Продукт 2"},
				{ID: 3, Name: "Продукт 3"},
			},
			id: 2,
			expected: Products{
				{ID: 1, Name: "Продукт 1"},
				{ID: 3, Name: "Продукт 3"},
			},
		},
		{
			name: "Удаление несуществующего продукта",
			products: Products{
				{ID: 1, Name: "Продукт 1"},
				{ID: 2, Name: "Продукт 2"},
			},
			id: 3,
			expected: Products{
				{ID: 1, Name: "Продукт 1"},
				{ID: 2, Name: "Продукт 2"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			products := tt.products
			products.Delete(tt.id)
			if !reflect.DeepEqual(products, tt.expected) {
				t.Errorf("После Products.Delete() = %v, want %v", products, tt.expected)
			}
		})
	}
}

func TestProductsDeleteMultiple(t *testing.T) {
	tests := []struct {
		name     string
		products Products
		ids      []int
		expected Products
	}{
		{
			name: "Удаление нескольких существующих продуктов",
			products: Products{
				{ID: 1, Name: "Продукт 1"},
				{ID: 2, Name: "Продукт 2"},
				{ID: 3, Name: "Продукт 3"},
				{ID: 4, Name: "Продукт 4"},
			},
			ids: []int{2, 4},
			expected: Products{
				{ID: 1, Name: "Продукт 1"},
				{ID: 3, Name: "Продукт 3"},
			},
		},
		{
			name: "Удаление несуществующих продуктов",
			products: Products{
				{ID: 1, Name: "Продукт 1"},
				{ID: 2, Name: "Продукт 2"},
			},
			ids: []int{3, 4},
			expected: Products{
				{ID: 1, Name: "Продукт 1"},
				{ID: 2, Name: "Продукт 2"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			products := tt.products
			products.DeleteMultiple(tt.ids)
			if !reflect.DeepEqual(products, tt.expected) {
				t.Errorf("После Products.DeleteMultiple() = %v, want %v", products, tt.expected)
			}
		})
	}
}

func TestProductsUpdate(t *testing.T) {
	products := Products{
		{ID: 1, Name: "Продукт 1", ProcessingTime: 1.5, TimeCalculation: "1.5"},
		{ID: 2, Name: "Продукт 2", ProcessingTime: 2.0, TimeCalculation: "2.0"},
	}

	updatedProduct := Product{ID: 1, Name: "Обновленный продукт", ProcessingTime: 3.0, TimeCalculation: "3.0"}
	expected := Products{
		{ID: 1, Name: "Обновленный продукт", ProcessingTime: 3.0, TimeCalculation: "3.0"},
		{ID: 2, Name: "Продукт 2", ProcessingTime: 2.0, TimeCalculation: "2.0"},
	}

	products.Update(updatedProduct)
	if !reflect.DeepEqual(products, expected) {
		t.Errorf("После Products.Update() = %v, want %v", products, expected)
	}

	// Проверка обновления несуществующего продукта (не должно быть изменений)
	nonExistingProduct := Product{ID: 3, Name: "Несуществующий продукт"}
	products.Update(nonExistingProduct)
	if !reflect.DeepEqual(products, expected) {
		t.Errorf("После Products.Update() с несуществующим ID = %v, want %v", products, expected)
	}
}

func TestProductsGetNextID(t *testing.T) {
	tests := []struct {
		name     string
		products Products
		expected int
	}{
		{
			name:     "Пустой список продуктов",
			products: Products{},
			expected: 1,
		},
		{
			name: "Список с последовательными ID",
			products: Products{
				{ID: 1, Name: "Продукт 1"},
				{ID: 2, Name: "Продукт 2"},
				{ID: 3, Name: "Продукт 3"},
			},
			expected: 4,
		},
		{
			name: "Список с произвольными ID",
			products: Products{
				{ID: 5, Name: "Продукт 5"},
				{ID: 2, Name: "Продукт 2"},
				{ID: 10, Name: "Продукт 10"},
			},
			expected: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.products.GetNextID()
			if result != tt.expected {
				t.Errorf("Products.GetNextID() = %v, want %v", result, tt.expected)
			}
		})
	}
}
