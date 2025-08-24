package solutions

import (
	"fmt"
	"regexp"
	"strconv"
)

// Day3Solution решение для 3-го дня
type Day3Solution struct{}

// NewDay3Solution создает новый экземпляр Day3Solution
func NewDay3Solution() *Day3Solution {
	return &Day3Solution{}
}

// Ch1 первая часть задачи 3-го дня - поиск корректных mul инструкций
func (d *Day3Solution) Ch1(input []byte) (string, error) {
	// Регулярное выражение для поиска корректных mul инструкций
	// mul(число,число) где числа от 1 до 3 цифр
	mulPattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := mulPattern.FindAllStringSubmatch(string(input), -1)

	totalSum := 0
	for _, match := range matches {
		if len(match) == 3 {
			// match[1] и match[2] содержат числа из скобок
			if x, err := strconv.Atoi(match[1]); err == nil {
				if y, err := strconv.Atoi(match[2]); err == nil {
					result := x * y
					totalSum += result
				}
			}
		}
	}

	return fmt.Sprintf("Total sum of mul results: %d", totalSum), nil
}

// Ch2 вторая часть задачи 3-го дня - поддержка do() и don't() инструкций
func (d *Day3Solution) Ch2(input []byte) (string, error) {
	inputStr := string(input)

	// Регулярные выражения для поиска инструкций
	mulPattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doPattern := regexp.MustCompile(`do\(\)`)
	dontPattern := regexp.MustCompile(`don't\(\)`)

	// Находим все позиции инструкций
	mulMatches := mulPattern.FindAllStringIndex(inputStr, -1)
	doMatches := doPattern.FindAllStringIndex(inputStr, -1)
	dontMatches := dontPattern.FindAllStringIndex(inputStr, -1)

	// Сортируем все инструкции по позиции
	type Instruction struct {
		pos  int
		kind string   // "mul", "do", "dont"
		data []string // для mul: [число1, число2]
	}

	var instructions []Instruction

	// Добавляем mul инструкции
	for _, match := range mulMatches {
		start, end := match[0], match[1]
		fullMatch := inputStr[start:end]
		// Извлекаем числа из mul(число,число)
		subMatches := mulPattern.FindStringSubmatch(fullMatch)
		if len(subMatches) == 3 {
			instructions = append(instructions, Instruction{
				pos:  start,
				kind: "mul",
				data: []string{subMatches[1], subMatches[2]},
			})
		}
	}

	// Добавляем do() инструкции
	for _, match := range doMatches {
		start := match[0]
		instructions = append(instructions, Instruction{
			pos:  start,
			kind: "do",
		})
	}

	// Добавляем don't() инструкции
	for _, match := range dontMatches {
		start := match[0]
		instructions = append(instructions, Instruction{
			pos:  start,
			kind: "dont",
		})
	}

	// Сортируем по позиции в тексте
	for i := 0; i < len(instructions)-1; i++ {
		for j := i + 1; j < len(instructions); j++ {
			if instructions[i].pos > instructions[j].pos {
				instructions[i], instructions[j] = instructions[j], instructions[i]
			}
		}
	}

	// Обрабатываем инструкции по порядку
	mulEnabled := true // по умолчанию включено
	totalSum := 0

	for _, instr := range instructions {
		switch instr.kind {
		case "do":
			mulEnabled = true
		case "dont":
			mulEnabled = false
		case "mul":
			if mulEnabled {
				// Выполняем умножение только если включено
				if x, err := strconv.Atoi(instr.data[0]); err == nil {
					if y, err := strconv.Atoi(instr.data[1]); err == nil {
						result := x * y
						totalSum += result
					}
				}
			}
		}
	}

	return fmt.Sprintf("Total sum of enabled mul results: %d", totalSum), nil
}
