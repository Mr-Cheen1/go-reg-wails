package storage

import "github.com/Mr-Cheen1/go-reg-wails/backend/models"

// Storage интерфейс для хранилища данных
type Storage interface {
	Load() (models.Products, error)
	Save(products models.Products) error
	Close() error
}
