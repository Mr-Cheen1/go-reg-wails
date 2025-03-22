<p align="center">
  <img src="https://github.com/Mr-Cheen1/go-reg-wails/raw/main/build/appicon.png" width="150">
</p>

<p align="center">
  <a href="https://github.com/Mr-Cheen1/go-reg-wails/actions/workflows/lint.yml"><img src="https://github.com/Mr-Cheen1/go-reg-wails/actions/workflows/lint.yml/badge.svg" alt="Lint Status"/></a>
  <a href="https://github.com/Mr-Cheen1/go-reg-wails/actions/workflows/test.yml"><img src="https://github.com/Mr-Cheen1/go-reg-wails/actions/workflows/test.yml/badge.svg" alt="Test Status"/></a>
  <a href="https://github.com/Mr-Cheen1/go-reg-wails/actions/workflows/build.yml"><img src="https://github.com/Mr-Cheen1/go-reg-wails/actions/workflows/build.yml/badge.svg" alt="Build Status"/></a>
</p>

# 📊 Desktop Go Reg (Wails + React)

Приложение для учета и расчета времени обработки деталей с современным веб-интерфейсом, написанное на Go (бэкенд) и React с TypeScript (фронтенд) с использованием фреймворка Wails.

## 🚀 Возможности

- ✨ Добавление, редактирование и удаление записей
- 🔍 Поиск по наименованию
- ⏱️ Расчет времени обработки с поддержкой формул (например: 8+2+5)
- ✅ Множественное выделение записей для удаления
- 💾 Автоматическое сохранение в Excel файл
- 🎨 Современный адаптивный интерфейс с темной темой
- 🖥️ Кроссплатформенность (Windows, macOS, Linux)
- 🧪 Полное покрытие тестами бэкенд-части

## 📥 Установка

### Требования

- Go 1.21 или выше
- Node.js 18 или выше
- Wails CLI

### Установка Wails CLI

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### Сборка из исходного кода

1. Клонируйте репозиторий:

```bash
git clone https://github.com/Mr-Cheen1/go-reg-wails.git
cd go-reg-wails
```

2. Установите зависимости и соберите приложение:

```bash
wails build
```

### Разработка

Для запуска в режиме разработки:

```bash
wails dev
```

### Тестирование

Для запуска всех тестов:

```bash
go test ./... -v
```

Для проверки покрытия кода тестами:

```bash
go test ./... -cover
```

## 🏗️ Структура проекта

```
go-reg-wails/
├── 📁 backend/              # Бэкенд на Go
│   ├── 📁 models/          # Модели данных
│   │   ├── product.go     # Структура продукта и методы работы с ним
│   │   └── product_test.go # Тесты для продуктов
│   ├── 📁 storage/         # Слой хранения данных
│   │   ├── excel.go       # Работа с Excel файлом
│   │   ├── storage.go     # Интерфейс хранилища
│   │   └── excel_test.go  # Тесты для хранилища
│   └── 📁 utils/           # Вспомогательные функции
│       ├── calculator.go  # Калькулятор времени
│       └── calculator_test.go # Тесты для калькулятора
├── 📁 frontend/             # Фронтенд на React
│   ├── 📁 src/             # Исходный код React
│   │   ├── 📁 components/  # React компоненты
│   │   ├── 📁 hooks/       # React хуки
│   │   ├── App.tsx        # Основной компонент приложения
│   │   └── main.tsx       # Точка входа
│   ├── package.json       # Зависимости NPM
│   └── tailwind.config.js # Конфигурация Tailwind CSS
├── 📝 app.go                # Логика приложения
├── 📝 app_test.go           # Тесты для логики приложения
├── 📝 main.go               # Точка входа в приложение
├── 📊 database.xlsx         # Файл базы данных
└── 📖 README.md             # Документация
```

## 🔧 Технологии

### Бэкенд

- **Go** - основной язык программирования
- **Wails** - фреймворк для создания десктопных приложений
- **excelize** - библиотека для работы с Excel файлами

### Фронтенд

- **React** - библиотека для создания пользовательского интерфейса
- **TypeScript** - типизированный JavaScript
- **Tailwind CSS** - утилитарный CSS-фреймворк
- **shadcn/ui** - компоненты пользовательского интерфейса
- **Vite** - инструмент сборки

## 🔄 Логика работы приложения

```mermaid
flowchart TD
    A[Запуск приложения] --> B[Загрузка данных из Excel]
    B --> C[Инициализация UI]
  
    %% События пользовательского интерфейса
    C --> D[Пользовательские действия]
  
    %% Ветви действий
    D --> E[Просмотр списка]
    D --> F[Поиск]
    D --> G[Добавление]
    D --> H[Редактирование]
    D --> I[Удаление]
  
    %% Логика поиска
    F --> F1[Ввод поискового запроса]
    F1 --> F2[Фильтрация по имени]
    F2 --> F3[Отображение результатов]
  
    %% Логика добавления
    G --> G1[Ввод данных]
    G1 --> G2[Расчет времени обработки]
    G2 --> G3[Добавление в список]
    G3 --> G4[Сохранение в Excel]
  
    %% Логика редактирования
    H --> H1[Выбор записи]
    H1 --> H2[Изменение данных]
    H2 --> H3[Расчет времени обработки]
    H3 --> H4[Обновление в списке]
    H4 --> H5[Сохранение в Excel]
  
    %% Логика удаления
    I --> I1[Выбор записей]
    I1 --> I2[Удаление из списка]
    I2 --> I3[Сохранение в Excel]
  
    %% Стили
    classDef process fill:#d4f1f9,stroke:#05a0c8,stroke-width:2px
    classDef action fill:#ffe6cc,stroke:#d79b00,stroke-width:2px
    classDef storage fill:#e1d5e7,stroke:#9673a6,stroke-width:2px
  
    class A,B,C process
    class D,E,F,G,H,I action
    class G4,H5,I3 storage
```

## 📊 Архитектура приложения

```mermaid
classDiagram
    class App {
        -context.Context ctx
        -storage.Storage storage
        -models.Products products
        +startup(ctx)
        +GetProducts() []Product
        +SearchProducts(query) []Product
        +AddProduct(name, timeCalculation) error
        +UpdateProduct(id, name, timeCalculation) error
        +DeleteProduct(id) error
        +DeleteProducts(ids) error
    }
  
    class Product {
        +int ID
        +string Name
        +float64 ProcessingTime
        +string TimeCalculation
    }
  
    class Products {
        +Search(query) Products
        +Delete(id)
        +DeleteMultiple(ids)
        +Update(product)
        +GetNextID() int
    }
  
    class Storage {
        <<interface>>
        +Load() (Products, error)
        +Save(products) error
        +Close() error
    }
  
    class ExcelStorage {
        -*excelize.File file
        -string filename
        +WithFilename(filename) *ExcelStorage
        +Load() (Products, error)
        +Save(products) error
        +Close() error
    }
  
    class Calculator {
        +CalculateTime(timeStr) float64
    }
  
    App --> Storage : использует
    App --> Products : хранит
    ExcelStorage ..|> Storage : реализует
    Products o-- Product : содержит
    App --> Calculator : использует
```

## 📝 Лицензия

Copyright © 2025

## 🔄 CI/CD

Проект использует GitHub Actions для автоматизации следующих процессов:

- 🔍 **Линтинг**: Проверка качества кода с помощью golangci-lint
- 🧪 **Тестирование**: Запуск всех тестов
- 📦 **Сборка**: Компиляция приложения для Windows

Статус процессов можно увидеть в бейджах вверху README.
