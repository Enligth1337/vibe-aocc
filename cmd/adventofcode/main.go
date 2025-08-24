package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"new-go-project/internal/config"
	"new-go-project/internal/solutions"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/adventofcode/main.go <day> [test_file]")
		fmt.Println("Example: go run cmd/adventofcode/main.go day1")
		fmt.Println("Example: go run cmd/adventofcode/main.go day4 test_example.txt")
		os.Exit(1)
	}

	day := os.Args[1]

	var inputFilePath string

	if len(os.Args) >= 3 {
		// Если передан второй аргумент, используем его как тестовый файл
		inputFilePath = os.Args[2]
	} else {
		// Иначе загружаем конфигурацию
		cfg, err := config.Load("config.yml")
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}

		// Получаем задачу по дню
		task, err := cfg.GetTaskByDay(day)
		if err != nil {
			log.Fatalf("Failed to get task: %v", err)
		}

		// Получаем путь к файлу ввода
		inputFilePath = cfg.GetInputFilePath(task)
	}

	// Читаем файл ввода
	inputData, err := os.ReadFile(inputFilePath)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	// Создаем фабрику решений
	factory := solutions.NewSolutionFactory()

	// Получаем решение по ключу
	solution, err := factory.New(day)
	if err != nil {
		log.Fatalf("Failed to create solution: %v", err)
	}

	fmt.Printf("Solving %s with input file: %s\n", day, inputFilePath)
	fmt.Println("=" + strings.Repeat("=", 50))

	// Решаем первую часть
	result1, err := solution.Ch1(inputData)
	if err != nil {
		log.Printf("Error in Ch1: %v", err)
	} else {
		fmt.Printf("Ch1 result: %s\n", result1)
	}

	// Решаем вторую часть
	result2, err := solution.Ch2(inputData)
	if err != nil {
		log.Printf("Error in Ch2: %v", err)
	} else {
		fmt.Printf("Ch2 result: %s\n", result2)
	}
}
