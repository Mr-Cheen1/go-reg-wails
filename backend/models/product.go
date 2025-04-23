package models

import "strings"

// Product представляет собой структуру продукта
type Product struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	ProcessingTime  float64 `json:"processingTime"`
	TimeCalculation string  `json:"timeCalculation"`
}

// Products представляет собой срез продуктов с методами для работы
type Products []Product

// Search ищет продукты по запросу
func (p Products) Search(query string) Products {
	if query == "" {
		return p
	}

	// Безопасная обработка запроса
	query = strings.TrimSpace(query)
	if len(query) == 0 {
		return p
	}

	var result Products
	query = strings.ToLower(query)

	// Безопасный поиск с ограничением длины запроса
	if len(query) > 50 {
		query = query[:50]
	}

	for _, product := range p {
		if strings.Contains(strings.ToLower(product.Name), query) {
			result = append(result, product)
		}
	}

	// Возвращаем результаты поиска, даже если они пустые
	return result
}

// Delete удаляет продукт по ID
func (p *Products) Delete(id int) {
	for i, product := range *p {
		if product.ID == id {
			*p = append((*p)[:i], (*p)[i+1:]...)
			break
		}
	}
}

// DeleteMultiple удаляет несколько продуктов по ID
func (p *Products) DeleteMultiple(ids []int) {
	idMap := make(map[int]bool)
	for _, id := range ids {
		idMap[id] = true
	}

	var result Products
	for _, product := range *p {
		if !idMap[product.ID] {
			result = append(result, product)
		}
	}
	*p = result
}

// Update обновляет продукт
func (p *Products) Update(product Product) {
	for i, prod := range *p {
		if prod.ID == product.ID {
			(*p)[i] = product
			break
		}
	}
}

// GetNextID возвращает следующий доступный ID
func (p Products) GetNextID() int {
	maxID := 0
	for _, product := range p {
		if product.ID > maxID {
			maxID = product.ID
		}
	}
	return maxID + 1
}
