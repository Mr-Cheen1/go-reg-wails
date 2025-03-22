package storage

import (
	"fmt"
	"strconv"

	"github.com/Mr-Cheen1/go-reg-wails/backend/models"
	"github.com/xuri/excelize/v2"
)

type ExcelStorage struct {
	file     *excelize.File
	filename string
}

func NewExcelStorage() *ExcelStorage {
	return &ExcelStorage{
		filename: "database.xlsx",
	}
}

// WithFilename позволяет указать имя файла
func (es *ExcelStorage) WithFilename(filename string) *ExcelStorage {
	es.filename = filename
	return es
}

// Load загружает данные из Excel файла
func (es *ExcelStorage) Load() (models.Products, error) {
	var products models.Products

	// Закрываем предыдущий файл если он был открыт
	if es.file != nil {
		es.file.Close()
	}

	// Попытка открыть существующий файл
	var err error
	es.file, err = excelize.OpenFile(es.filename)
	if err != nil {
		// Если файл не существует, создаем новый
		es.file = excelize.NewFile()
		// Создаем заголовки
		if err := es.file.SetCellValue("Sheet1", "A1", "ID"); err != nil {
			return products, fmt.Errorf("ошибка при установке заголовка ID: %w", err)
		}
		if err := es.file.SetCellValue("Sheet1", "B1", "Наименование"); err != nil {
			return products, fmt.Errorf("ошибка при установке заголовка Наименование: %w", err)
		}
		if err := es.file.SetCellValue("Sheet1", "C1", "Время обработки в часах"); err != nil {
			return products, fmt.Errorf("ошибка при установке заголовка Время обработки: %w", err)
		}
		if err := es.file.SetCellValue("Sheet1", "D1", "Расчет времени"); err != nil {
			return products, fmt.Errorf("ошибка при установке заголовка Расчет времени: %w", err)
		}
		return products, es.file.SaveAs(es.filename)
	}

	// Читаем данные
	rows, err := es.file.GetRows("Sheet1")
	if err != nil {
		return products, err
	}

	// Пропускаем заголовок
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		if len(row) < 4 {
			continue
		}

		id, _ := strconv.Atoi(row[0])
		processingTime, _ := strconv.ParseFloat(row[2], 64)

		product := models.Product{
			ID:              id,
			Name:            row[1],
			ProcessingTime:  processingTime,
			TimeCalculation: row[3],
		}
		products = append(products, product)
	}

	return products, nil
}

// Save сохраняет данные в Excel файл
func (es *ExcelStorage) Save(products models.Products) error {
	// Закрываем текущий файл если он открыт
	if es.file != nil {
		es.file.Close()
	}

	// Создаем новый файл
	es.file = excelize.NewFile()

	// Записываем заголовки
	if err := es.file.SetCellValue("Sheet1", "A1", "ID"); err != nil {
		return err
	}
	if err := es.file.SetCellValue("Sheet1", "B1", "Наименование"); err != nil {
		return err
	}
	if err := es.file.SetCellValue("Sheet1", "C1", "Время обработки в часах"); err != nil {
		return err
	}
	if err := es.file.SetCellValue("Sheet1", "D1", "Расчет времени"); err != nil {
		return err
	}

	// Записываем данные
	for i, product := range products {
		row := i + 2
		if err := es.file.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), product.ID); err != nil {
			return err
		}
		if err := es.file.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), product.Name); err != nil {
			return err
		}
		if err := es.file.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), product.ProcessingTime); err != nil {
			return err
		}
		if err := es.file.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), product.TimeCalculation); err != nil {
			return err
		}
	}

	// Сохраняем файл
	return es.file.SaveAs(es.filename)
}

// Close закрывает файл Excel
func (es *ExcelStorage) Close() error {
	if es.file != nil {
		return es.file.Close()
	}
	return nil
}
