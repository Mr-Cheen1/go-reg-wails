package main

import (
	"context"
	"testing"

	"github.com/Mr-Cheen1/go-reg-wails/backend/models"
)

// MockStorage - мок для Storage
type MockStorage struct {
	products models.Products
	saveFunc func(models.Products) error
}

func NewMockStorage(products models.Products) *MockStorage {
	return &MockStorage{
		products: products,
		saveFunc: func(p models.Products) error {
			return nil
		},
	}
}

func (ms *MockStorage) Load() (models.Products, error) {
	return ms.products, nil
}

func (ms *MockStorage) Save(products models.Products) error {
	return ms.saveFunc(products)
}

func (ms *MockStorage) Close() error {
	return nil
}

// Вспомогательная функция для проверки равенства продуктов по важным полям
func checkProductsEqual(t *testing.T, got, want models.Products, testName string) {
	if len(got) != len(want) {
		t.Errorf("%s: количество продуктов = %d, want %d", testName, len(got), len(want))
		return
	}

	// Создаем карту ожидаемых ID
	wantIDs := make(map[int]models.Product)
	for _, product := range want {
		wantIDs[product.ID] = product
	}

	// Проверяем каждый продукт
	for _, gotProduct := range got {
		if wantProduct, ok := wantIDs[gotProduct.ID]; ok {
			if gotProduct.Name != wantProduct.Name {
				t.Errorf("%s: продукт с ID=%d имеет Name=%s, want %s", testName, gotProduct.ID, gotProduct.Name, wantProduct.Name)
			}
			// Можно добавить проверки других полей при необходимости
		} else {
			t.Errorf("%s: неожиданный продукт с ID=%d", testName, gotProduct.ID)
		}
	}
}

func TestApp_GetProducts(t *testing.T) {
	// Создаем тестовые данные
	testProducts := models.Products{
		{ID: 1, Name: "Продукт 1", ProcessingTime: 1.5, TimeCalculation: "1.5"},
		{ID: 2, Name: "Продукт 2", ProcessingTime: 2.0, TimeCalculation: "2.0"},
	}

	// Создаем мок хранилища и приложение
	mockStorage := NewMockStorage(testProducts)
	app := NewApp(mockStorage)
	app.startup(context.Background())

	// Проверяем, что GetProducts возвращает все продукты
	result := app.GetProducts()
	checkProductsEqual(t, result, testProducts, "GetProducts()")
}

func TestApp_SearchProducts(t *testing.T) {
	// Создаем тестовые данные
	testProducts := models.Products{
		{ID: 1, Name: "Продукт A", ProcessingTime: 1.5, TimeCalculation: "1.5"},
		{ID: 2, Name: "Товар B", ProcessingTime: 2.0, TimeCalculation: "2.0"},
		{ID: 3, Name: "Продукт C", ProcessingTime: 3.0, TimeCalculation: "3.0"},
	}

	// Создаем мок хранилища и приложение
	mockStorage := NewMockStorage(testProducts)
	app := NewApp(mockStorage)
	app.startup(context.Background())

	// Тест 1: Поиск по слову "Продукт"
	result := app.SearchProducts("Продукт")
	expected := models.Products{
		{ID: 1, Name: "Продукт A", ProcessingTime: 1.5, TimeCalculation: "1.5"},
		{ID: 3, Name: "Продукт C", ProcessingTime: 3.0, TimeCalculation: "3.0"},
	}
	checkProductsEqual(t, result, expected, "SearchProducts('Продукт')")

	// Тест 2: Поиск по слову "Товар"
	result = app.SearchProducts("Товар")
	expected = models.Products{
		{ID: 2, Name: "Товар B", ProcessingTime: 2.0, TimeCalculation: "2.0"},
	}
	checkProductsEqual(t, result, expected, "SearchProducts('Товар')")

	// Тест 3: Поиск по несуществующему слову
	result = app.SearchProducts("Несуществующий")
	if len(result) != 0 {
		t.Errorf("SearchProducts('Несуществующий') = %v, want пустой слайс", result)
	}
}

func TestApp_AddProduct(t *testing.T) {
	// Создаем тестовые данные
	initialProducts := models.Products{
		{ID: 1, Name: "Продукт 1", ProcessingTime: 1.5, TimeCalculation: "1.5"},
	}

	// Создаем мок хранилища с функцией, проверяющей сохраняемые данные
	mockStorage := NewMockStorage(initialProducts)
	savedProducts := models.Products{}
	mockStorage.saveFunc = func(products models.Products) error {
		savedProducts = products
		return nil
	}

	// Создаем приложение
	app := NewApp(mockStorage)
	app.startup(context.Background())

	// Добавляем новый продукт
	err := app.AddProduct("Новый продукт", "2.5")
	if err != nil {
		t.Fatalf("AddProduct() error = %v", err)
	}

	// Проверяем, что продукт добавлен
	if len(savedProducts) != 2 {
		t.Fatalf("После AddProduct() количество продуктов = %d, want %d", len(savedProducts), 2)
	}

	newProduct := savedProducts[1]
	if newProduct.Name != "Новый продукт" {
		t.Errorf("Имя добавленного продукта = %s, want %s", newProduct.Name, "Новый продукт")
	}
	if newProduct.ProcessingTime != 2.5 {
		t.Errorf("Время обработки = %f, want %f", newProduct.ProcessingTime, 2.5)
	}
	if newProduct.TimeCalculation != "2.5" {
		t.Errorf("Расчет времени = %s, want %s", newProduct.TimeCalculation, "2.5")
	}
	if newProduct.ID != 2 {
		t.Errorf("ID добавленного продукта = %d, want %d", newProduct.ID, 2)
	}
}

func TestApp_UpdateProduct(t *testing.T) {
	// Создаем тестовые данные
	initialProducts := models.Products{
		{ID: 1, Name: "Продукт 1", ProcessingTime: 1.5, TimeCalculation: "1.5"},
		{ID: 2, Name: "Продукт 2", ProcessingTime: 2.0, TimeCalculation: "2.0"},
	}

	// Создаем мок хранилища с функцией, проверяющей сохраняемые данные
	mockStorage := NewMockStorage(initialProducts)
	savedProducts := models.Products{}
	mockStorage.saveFunc = func(products models.Products) error {
		savedProducts = products
		return nil
	}

	// Создаем приложение
	app := NewApp(mockStorage)
	app.startup(context.Background())

	// Обновляем существующий продукт
	err := app.UpdateProduct(1, "Обновленный продукт", "3.5")
	if err != nil {
		t.Fatalf("UpdateProduct() error = %v", err)
	}

	// Проверяем, что продукт обновлен
	if len(savedProducts) != 2 {
		t.Fatalf("После UpdateProduct() количество продуктов = %d, want %d", len(savedProducts), 2)
	}

	updatedProduct := savedProducts[0]
	if updatedProduct.Name != "Обновленный продукт" {
		t.Errorf("Имя обновленного продукта = %s, want %s", updatedProduct.Name, "Обновленный продукт")
	}
	if updatedProduct.ProcessingTime != 3.5 {
		t.Errorf("Время обработки = %f, want %f", updatedProduct.ProcessingTime, 3.5)
	}
	if updatedProduct.TimeCalculation != "3.5" {
		t.Errorf("Расчет времени = %s, want %s", updatedProduct.TimeCalculation, "3.5")
	}
	if updatedProduct.ID != 1 {
		t.Errorf("ID обновленного продукта = %d, want %d", updatedProduct.ID, 1)
	}
}

func TestApp_DeleteProduct(t *testing.T) {
	// Создаем тестовые данные
	initialProducts := models.Products{
		{ID: 1, Name: "Продукт 1", ProcessingTime: 1.5, TimeCalculation: "1.5"},
		{ID: 2, Name: "Продукт 2", ProcessingTime: 2.0, TimeCalculation: "2.0"},
		{ID: 3, Name: "Продукт 3", ProcessingTime: 3.0, TimeCalculation: "3.0"},
	}

	// Создаем мок хранилища с функцией, проверяющей сохраняемые данные
	mockStorage := NewMockStorage(initialProducts)
	savedProducts := models.Products{}
	mockStorage.saveFunc = func(products models.Products) error {
		savedProducts = products
		return nil
	}

	// Создаем приложение
	app := NewApp(mockStorage)
	app.startup(context.Background())

	// Удаляем продукт с ID=2
	err := app.DeleteProduct(2)
	if err != nil {
		t.Fatalf("DeleteProduct() error = %v", err)
	}

	// Проверяем, что продукт удален
	if len(savedProducts) != 2 {
		t.Fatalf("После DeleteProduct() количество продуктов = %d, want %d", len(savedProducts), 2)
	}

	// Проверяем, что остались правильные продукты
	expectedIDs := []int{1, 3}
	for i, product := range savedProducts {
		if product.ID != expectedIDs[i] {
			t.Errorf("ID продукта после удаления = %d, want %d", product.ID, expectedIDs[i])
		}
	}
}

func TestApp_DeleteProducts(t *testing.T) {
	// Создаем тестовые данные
	initialProducts := models.Products{
		{ID: 1, Name: "Продукт 1", ProcessingTime: 1.5, TimeCalculation: "1.5"},
		{ID: 2, Name: "Продукт 2", ProcessingTime: 2.0, TimeCalculation: "2.0"},
		{ID: 3, Name: "Продукт 3", ProcessingTime: 3.0, TimeCalculation: "3.0"},
		{ID: 4, Name: "Продукт 4", ProcessingTime: 4.0, TimeCalculation: "4.0"},
	}

	// Создаем мок хранилища с функцией, проверяющей сохраняемые данные
	mockStorage := NewMockStorage(initialProducts)
	savedProducts := models.Products{}
	mockStorage.saveFunc = func(products models.Products) error {
		savedProducts = products
		return nil
	}

	// Создаем приложение
	app := NewApp(mockStorage)
	app.startup(context.Background())

	// Удаляем несколько продуктов
	err := app.DeleteProducts([]int{2, 4})
	if err != nil {
		t.Fatalf("DeleteProducts() error = %v", err)
	}

	// Проверяем, что продукты удалены
	if len(savedProducts) != 2 {
		t.Fatalf("После DeleteProducts() количество продуктов = %d, want %d", len(savedProducts), 2)
	}

	// Проверяем, что остались правильные продукты
	expectedIDs := []int{1, 3}
	for i, product := range savedProducts {
		if product.ID != expectedIDs[i] {
			t.Errorf("ID продукта после удаления = %d, want %d", product.ID, expectedIDs[i])
		}
	}
}
