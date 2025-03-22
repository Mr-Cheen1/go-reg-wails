package main

import (
	"context"
	"log"

	"github.com/Mr-Cheen1/go-reg-wails/backend/models"
	"github.com/Mr-Cheen1/go-reg-wails/backend/storage"
	"github.com/Mr-Cheen1/go-reg-wails/backend/utils"
)

// App структура приложения
type App struct {
	ctx      context.Context
	storage  storage.Storage
	products models.Products
}

// NewApp создает новый экземпляр приложения
func NewApp(storage storage.Storage) *App {
	return &App{
		storage: storage,
	}
}

// startup вызывается при запуске приложения
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	var err error
	a.products, err = a.storage.Load()
	if err != nil {
		log.Printf("Ошибка загрузки данных: %v\n", err)
	}
}

// GetProducts возвращает все продукты
func (a *App) GetProducts() []models.Product {
	return a.products
}

// SearchProducts ищет продукты по запросу
func (a *App) SearchProducts(query string) []models.Product {
	return a.products.Search(query)
}

// AddProduct добавляет новый продукт
func (a *App) AddProduct(name, timeCalculation string) error {
	processingTime := utils.CalculateTime(timeCalculation)
	product := models.Product{
		ID:              a.products.GetNextID(),
		Name:            name,
		ProcessingTime:  processingTime,
		TimeCalculation: timeCalculation,
	}
	a.products = append(a.products, product)
	return a.storage.Save(a.products)
}

// UpdateProduct обновляет существующий продукт
func (a *App) UpdateProduct(id int, name, timeCalculation string) error {
	processingTime := utils.CalculateTime(timeCalculation)
	product := models.Product{
		ID:              id,
		Name:            name,
		ProcessingTime:  processingTime,
		TimeCalculation: timeCalculation,
	}
	a.products.Update(product)
	return a.storage.Save(a.products)
}

// DeleteProduct удаляет продукт по ID
func (a *App) DeleteProduct(id int) error {
	a.products.Delete(id)
	return a.storage.Save(a.products)
}

// DeleteProducts удаляет несколько продуктов по ID
func (a *App) DeleteProducts(ids []int) error {
	a.products.DeleteMultiple(ids)
	return a.storage.Save(a.products)
}
