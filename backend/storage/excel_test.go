package storage

import (
	"os"
	"reflect"
	"testing"

	"github.com/Mr-Cheen1/go-reg-wails/backend/models"
)

func TestExcelStorage_WithFilename(t *testing.T) {
	storage := NewExcelStorage()
	filename := "test.xlsx"
	result := storage.WithFilename(filename)

	if result.filename != filename {
		t.Errorf("WithFilename() filename = %v, want %v", result.filename, filename)
	}

	if result != storage {
		t.Errorf("WithFilename() должен возвращать тот же инстанс хранилища")
	}
}

func TestExcelStorage_SaveAndLoad(t *testing.T) {
	// Используем временный файл
	tempFile := "test_temp.xlsx"
	defer os.Remove(tempFile) // Удаляем временный файл после тестов

	storage := NewExcelStorage().WithFilename(tempFile)

	// Создаем тестовые данные
	testProducts := models.Products{
		{ID: 1, Name: "Тестовый продукт 1", ProcessingTime: 1.5, TimeCalculation: "1.5"},
		{ID: 2, Name: "Тестовый продукт 2", ProcessingTime: 2.0, TimeCalculation: "2.0"},
	}

	// Сохраняем данные
	err := storage.Save(testProducts)
	if err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	// Загружаем данные
	loadedProducts, err := storage.Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	// Проверяем, что данные загружены корректно
	if !reflect.DeepEqual(loadedProducts, testProducts) {
		t.Errorf("Load() = %v, want %v", loadedProducts, testProducts)
	}

	// Тестируем Close
	err = storage.Close()
	if err != nil {
		t.Errorf("Close() error = %v", err)
	}
}

func TestExcelStorage_LoadNonExistentFile(t *testing.T) {
	// Используем несуществующий файл
	nonExistentFile := "non_existent.xlsx"
	defer os.Remove(nonExistentFile) // На всякий случай удаляем файл после тестов, если он будет создан

	storage := NewExcelStorage().WithFilename(nonExistentFile)

	// Загружаем данные из несуществующего файла
	products, err := storage.Load()
	if err != nil {
		t.Fatalf("Load() из несуществующего файла должно создать новый файл, error = %v", err)
	}

	// Проверяем, что возвращен пустой слайс продуктов
	if len(products) != 0 {
		t.Errorf("Load() из несуществующего файла = %v, хотим пустой слайс", products)
	}

	// Проверяем, что файл был создан
	_, err = os.Stat(nonExistentFile)
	if os.IsNotExist(err) {
		t.Errorf("Load() должен создать файл, если он не существует")
	}
}

func TestExcelStorage_SaveError(t *testing.T) {
	// Тестируем ошибку при сохранении в некорректную директорию
	invalidPath := "/invalid/directory/file.xlsx"
	storage := NewExcelStorage().WithFilename(invalidPath)

	err := storage.Save(models.Products{})
	if err == nil {
		t.Errorf("Save() должен вернуть ошибку при некорректном пути")
	}
}
