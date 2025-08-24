package solutions

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Day1Solution решение для 1-го дня
type Day1Solution struct{}

// NewDay1Solution создает новый экземпляр Day1Solution
func NewDay1Solution() *Day1Solution {
	return &Day1Solution{}
}

// Ch1 первая часть задачи 1-го дня - вычисление расстояния между списками
func (d *Day1Solution) Ch1(input []byte) (string, error) {
	var leftList, rightList []int
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := regexp.MustCompile(`\s+`).Split(line, -1)
		if len(parts) < 2 {
			continue
		}
		if leftNum, err := strconv.Atoi(strings.TrimSpace(parts[0])); err == nil {
			leftList = append(leftList, leftNum)
		}
		if rightNum, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
			rightList = append(rightList, rightNum)
		}
	}
	sort.Ints(leftList)
	sort.Ints(rightList)
	totalDistance := 0
	for i := 0; i < len(leftList) && i < len(rightList); i++ {
		distance := abs(leftList[i] - rightList[i])
		totalDistance += distance
	}
	return fmt.Sprintf("Total distance: %d", totalDistance), nil
}

// Ch2 вторая часть задачи 1-го дня - вычисление similarity score
func (d *Day1Solution) Ch2(input []byte) (string, error) {
	var leftList, rightList []int
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := regexp.MustCompile(`\s+`).Split(line, -1)
		if len(parts) < 2 {
			continue
		}
		if leftNum, err := strconv.Atoi(strings.TrimSpace(parts[0])); err == nil {
			leftList = append(leftList, leftNum)
		}
		if rightNum, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
			rightList = append(rightList, rightNum)
		}
	}

	// Создаем карту для подсчета частоты каждого числа в левом списке
	leftFreq := make(map[int]int)
	for _, num := range leftList {
		leftFreq[num]++
	}

	// Создаем карту для подсчета частоты каждого числа в правом списке
	rightFreq := make(map[int]int)
	for _, num := range rightList {
		rightFreq[num]++
	}

	// Вычисляем similarity score
	similarityScore := 0
	for num, leftCount := range leftFreq {
		if rightCount, exists := rightFreq[num]; exists {
			similarityScore += leftCount * rightCount
		}
	}

	return fmt.Sprintf("Similarity score: %d", similarityScore), nil
}
