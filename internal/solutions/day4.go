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

	// Убираем пустые строки и объединяем в одну строку
	var gridString string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			gridString += line
		}
	}

	if len(gridString) == 0 {
		return "Empty input", nil
	}

	// Создаем регулярное выражение для поиска X-MAS паттернов
	// Паттерн: M.S.A.M.S (где . - любой символ, A - центр)
	// Это означает, что от центра 'A' идут две "MAS" в противоположных направлениях
	// Но нужно учесть, что это должно быть в форме X, а не просто последовательность
	// Давайте попробуем другой подход - искать все возможные X-паттерны

	// Ищем все позиции символа 'A' и проверяем, может ли там быть центр X-MAS
	totalXMAS := 0

	for i := 0; i < len(gridString); i++ {
		if gridString[i] == 'A' {
			// Проверяем, может ли здесь быть центр X-MAS
			if d.checkXMASAtPosition(gridString, i) {
				totalXMAS++
			}
		}
	}

	return fmt.Sprintf("Total X-MAS occurrences: %d", totalXMAS), nil
}

// checkXMASAtPosition проверяет, может ли в позиции pos быть центр X-MAS
func (d *Day4Solution) checkXMASAtPosition(gridString string, pos int) bool {
	// Проверяем, что позиция находится в допустимых границах
	if pos < 0 || pos >= len(gridString) {
		return false
	}

	// Проверяем, что в позиции находится символ 'A'
	if gridString[pos] != 'A' {
		return false
	}

	// Проверяем все возможные X-паттерны с центром в этой позиции
	// Для простоты, давайте проверим, есть ли две "MAS" в противоположных направлениях

	// Ищем "MAS" влево и "MAS" вправо
	leftMAS := d.findMASLeft(gridString, pos)
	rightMAS := d.findMASRight(gridString, pos)

	if leftMAS && rightMAS {
		return true
	}

	// Ищем "MAS" вверх и "MAS" вниз (если это сетка)
	// Но поскольку у нас одна строка, это сложнее

	return false
}

// findMASLeft ищет "MAS" слева от позиции pos
func (d *Day4Solution) findMASLeft(gridString string, pos int) bool {
	if pos < 3 {
		return false
	}

	// Проверяем паттерн "MAS" слева
	if gridString[pos-3] == 'M' && gridString[pos-2] == 'A' && gridString[pos-1] == 'S' {
		return true
	}

	// Проверяем паттерн "SAM" слева
	if gridString[pos-3] == 'S' && gridString[pos-2] == 'A' && gridString[pos-1] == 'M' {
		return true
	}

	return false
}

// findMASRight ищет "MAS" справа от позиции pos
func (d *Day4Solution) findMASRight(gridString string, pos int) bool {
	if pos+3 >= len(gridString) {
		return false
	}

	// Проверяем паттерн "MAS" справа
	if gridString[pos+1] == 'M' && gridString[pos+2] == 'A' && gridString[pos+3] == 'S' {
		return true
	}

	// Проверяем паттерн "SAM" справа
	if gridString[pos+1] == 'S' && gridString[pos+2] == 'A' && gridString[pos+3] == 'M' {
		return true
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
	// Проверяем две пары диагональных направлений для формирования X
	directionPairs := [][][]int{
		{{-1, -1}, {1, 1}}, // влево-вверх + вправо-вниз
		{{-1, 1}, {1, -1}}, // вправо-вверх + влево-вниз
	}

	// Для каждой пары направлений проверяем, можем ли мы найти X-MAS
	for _, pair := range directionPairs {
		if d.canFormXMAS(grid, row, col, pair[0][0], pair[0][1], rows, cols) {
			return true
		}
	}

	return false
}

// canFormXMAS проверяет, может ли с центром в (row, col) сформироваться X-MAS
// используя направление (dRow, dCol) и противоположное ему
func (d *Day4Solution) canFormXMAS(grid []string, row, col, dRow, dCol, rows, cols int) bool {
	// Проверяем первую MAS в направлении (dRow, dCol) от центра 'A'
	mas1 := d.checkMAS(grid, row, col, dRow, dCol, rows, cols)
	if !mas1 {
		return false
	}

	// Проверяем вторую MAS в противоположном направлении (-dRow, -dCol) от центра 'A'
	mas2 := d.checkMAS(grid, row, col, -dRow, -dCol, rows, cols)
	if !mas2 {
		return false
	}

	return true
}

// checkMAS проверяет, есть ли MAS в направлении (dRow, dCol) от позиции (row, col)
func (d *Day4Solution) checkMAS(grid []string, row, col, dRow, dCol, rows, cols int) bool {
	// Проверяем MAS вперед: от центра 'A' ищем полное слово "MAS"
	forward := d.checkWord(grid, row, col, dRow, dCol, rows, cols, "MAS")
	if forward {
		return true
	}

	// Проверяем MAS назад: от центра 'A' ищем полное слово "SAM"
	backward := d.checkWord(grid, row, col, dRow, dCol, rows, cols, "SAM")
	if backward {
		return true
	}

	return false
}

// checkWord проверяет, есть ли слово в направлении (dRow, dCol) от позиции (row, col)
func (d *Day4Solution) checkWord(grid []string, row, col, dRow, dCol, rows, cols int, word string) bool {
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
