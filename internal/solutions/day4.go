package solutions

import (
	"fmt"
	"strings"
)

// Day4Solution решение для 4-го дня
type Day4Solution struct{}

// NewDay4Solution создает новый экземпляр Day4Solution
func NewDay4Solution() *Day4Solution {
	return &Day4Solution{}
}

// Ch1 первая часть задачи 4-го дня - поиск всех вхождений XMAS
func (d *Day4Solution) Ch1(input []byte) (string, error) {
	lines := strings.Split(string(input), "\n")

	// Убираем пустые строки
	var grid []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			grid = append(grid, line)
		}
	}

	if len(grid) == 0 {
		return "Empty grid", nil
	}

	rows := len(grid)
	cols := len(grid[0])

	// Направления для поиска: горизонтально, вертикально, диагонально
	directions := [][]int{
		{0, 1},   // вправо
		{0, -1},  // влево
		{1, 0},   // вниз
		{-1, 0},  // вверх
		{1, 1},   // вправо-вниз
		{1, -1},  // влево-вниз
		{-1, 1},  // вправо-вверх
		{-1, -1}, // влево-вверх
	}

	totalOccurrences := 0

	// Проходим по каждой позиции в сетке
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Проверяем каждое направление
			for _, dir := range directions {
				if d.checkXMAS(grid, row, col, dir[0], dir[1], rows, cols) {
					totalOccurrences++
				}
			}
		}
	}

	return fmt.Sprintf("Total XMAS occurrences: %d", totalOccurrences), nil
}

// Ch2 вторая часть задачи 4-го дня - поиск X-MAS (две MAS в форме X)
func (d *Day4Solution) Ch2(input []byte) (string, error) {
	lines := strings.Split(string(input), "\n")

	// Убираем пустые строки
	var grid []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			grid = append(grid, line)
		}
	}

	if len(grid) == 0 {
		return "Empty grid", nil
	}

	rows := len(grid)
	cols := len(grid[0])

	totalXMAS := 0

	// Проходим по каждой позиции в сетке
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Проверяем, может ли здесь быть центр X-MAS (должен быть символ 'A')
			if grid[row][col] == 'A' {
				// Проверяем X-MAS паттерн с центром в этой позиции
				if d.checkXMASPattern(grid, row, col, rows, cols) {
					totalXMAS++
				}
			}
		}
	}

	return fmt.Sprintf("Total X-MAS occurrences: %d", totalXMAS), nil
}

// checkXMAS проверяет, начинается ли слово XMAS с позиции (row, col) в направлении (dRow, dCol)
func (d *Day4Solution) checkXMAS(grid []string, row, col, dRow, dCol, rows, cols int) bool {
	word := "XMAS"

	// Проверяем, что можем пройти всю длину слова в данном направлении
	for i := 0; i < len(word); i++ {
		newRow := row + i*dRow
		newCol := col + i*dCol

		// Проверяем границы
		if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
			return false
		}

		// Проверяем, что символ совпадает
		if grid[newRow][newCol] != word[i] {
			return false
		}
	}

	return true
}

// checkXMASPattern проверяет, есть ли X-MAS с центром в позиции (row, col)
func (d *Day4Solution) checkXMASPattern(grid []string, row, col, rows, cols int) bool {
	// Проверяем, что можем проверить строки выше и ниже
	if row == 0 || row == rows-1 {
		return false
	}

	// Проверяем, что можем проверить столбцы слева и справа
	if col == 0 || col == cols-1 {
		return false
	}

	// Проверяем диагональные позиции для формирования X
	// Левая диагональ: (row-1, col-1) и (row+1, col-1)
	// Правая диагональ: (row-1, col+1) и (row+1, col+1)

	leftAbove := grid[row-1][col-1]
	leftBelow := grid[row+1][col-1]
	rightAbove := grid[row-1][col+1]
	rightBelow := grid[row+1][col+1]

	// Если leftAbove == rightBelow, то по диагонали будет SAS или MAM - невалидно
	// if leftAbove == rightBelow {
	// 	return false
	// }

	// Проверяем, что левая диагональ содержит M и S (разность ASCII == 6)
	// Диагональ валидна если содержит M и S (|leftAbove - leftBelow| == 6)
	leftValid := abs(int(leftAbove)-int(rightBelow)) == 6

	// Проверяем, что правая диагональ содержит M и S (разность ASCII == 6)
	// Диагональ валидна если содержит M и S (|rightAbove - rightBelow| == 6)
	rightValid := abs(int(leftBelow)-int(rightAbove)) == 6

	// Обе диагонали должны быть валидными для формирования X-MAS
	return leftValid && rightValid
}
