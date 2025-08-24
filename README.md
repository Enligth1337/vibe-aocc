# Advent of Code Project

Проект для решения задач Advent of Code на языке Go с современной архитектурой.

## Структура проекта

```
new-go-project/
├── cmd/
│   └── adventofcode/
│       └── main.go              # Основная точка входа
├── internal/
│   ├── config/
│   │   └── config.go            # Управление конфигурацией
│   ├── solutions/
│   │   ├── interfaces.go        # Интерфейсы решений
│   │   ├── factory.go           # Фабрика решений
│   │   ├── day1.go             # Решение для первого дня
│   │   └── day2.go             # Решение для второго дня
│   └── parser/
│       └── parser.go            # Парсеры данных
├── pkg/
│   └── input/
│       └── parser.go            # Парсеры input файлов
├── resources/
│   ├── inputs/                  # Файлы ввода для задач
│   └── tests/                   # Тестовые данные
├── config.yml                   # Конфигурационный файл
├── go.mod                       # Зависимости Go
└── README.md                    # Документация
```

## Конфигурация

Проект использует YAML конфигурацию для управления задачами и путями к файлам.

### config.yml
```yaml
# Конфигурация для Advent of Code
input_resources_path: "resources/inputs"

# Задачи для решения
tasks:
  - day: "day1"
    input_file: "day1_input.txt"
  - day: "day2"
    input_file: "day2_example.txt"
```

## Использование

### Запуск программы

```bash
# Решение для первого дня
go run cmd/adventofcode/main.go day1

# Решение для второго дня
go run cmd/adventofcode/main.go day2
```

### Запуск тестов

```bash
# Запуск всех тестов
go test ./...

# Запуск тестов для конкретного пакета
go test ./internal/solutions
```

## Архитектура

### Пакеты

1. **cmd/adventofcode** - основная точка входа приложения
2. **internal/config** - управление конфигурацией
3. **internal/solutions** - решения задач по дням
4. **internal/parser** - парсеры данных
5. **pkg/input** - парсеры input файлов
6. **resources** - ресурсы (input файлы, тесты)

### Интерфейсы

- **Solution** - интерфейс для решения задач (Ch1, Ch2)
- **SolutionFactory** - фабрика для создания решений
- **Parser** - интерфейс для парсинга данных

### Реализации

- **Day1Solution** - решение для первого дня
- **Day2Solution** - решение для второго дня
- **SolutionFactoryImpl** - фабрика с регистрацией решений

## Добавление новых решений

1. Создайте новый файл в `internal/solutions/` (например, `day3.go`)
2. Реализуйте интерфейс `Solution`
3. Зарегистрируйте решение в `NewSolutionFactory()` в файле `factory.go`
4. Добавьте задачу в `config.yml`
5. Добавьте тесты

Пример:
```go
// internal/solutions/day3.go
type Day3Solution struct{}

func (d *Day3Solution) Ch1(input []byte) (string, error) {
    // Ваша логика здесь
    return "result", nil
}

func (d *Day3Solution) Ch2(input []byte) (string, error) {
    // Ваша логика здесь
    return "result", nil
}

// В factory.go добавьте:
factory.RegisterSolution("day3", NewDay3Solution())
```

## Требования

- Go 1.23.4 или выше
- Зависимость `gopkg.in/yaml.v3` для парсинга конфигурации

## Установка зависимостей

```bash
go mod tidy
```
