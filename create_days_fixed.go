package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Создаем дни с 3 по 25
	for day := 3; day <= 25; day++ {
		filename := fmt.Sprintf("internal/solutions/day%d.go", day)
		content := fmt.Sprintf(`package solutions

// Day%dSolution решение для %d-го дня
type Day%dSolution struct{}

// NewDay%dSolution создает новый экземпляр Day%dSolution
func NewDay%dSolution() *Day%dSolution {
	return &Day%dSolution{}
}

// Ch1 первая часть задачи %d-го дня
func (d *Day%dSolution) Ch1(input []byte) (string, error) {
	return "Part 1 not implemented yet", nil
}

// Ch2 вторая часть задачи %d-го дня
func (d *Day%dSolution) Ch2(input []byte) (string, error) {
	return "Part 2 not implemented yet", nil
}
`, day, day, day, day, day, day, day, day, day, day, day, day)

		// Создаем директорию если не существует
		dir := filepath.Dir(filename)
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			continue
		}

		// Записываем файл
		if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
			fmt.Printf("Error writing file %s: %v\n", filename, err)
		} else {
			fmt.Printf("Created %s\n", filename)
		}
	}

	fmt.Println("All solution files created successfully!")
}
