package solutions

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Day2Solution решение для 2-го дня
type Day2Solution struct{}

// NewDay2Solution создает новый экземпляр Day2Solution
func NewDay2Solution() *Day2Solution {
	return &Day2Solution{}
}

// Ch1 первая часть задачи 2-го дня - проверка безопасности отчетов
func (d *Day2Solution) Ch1(input []byte) (string, error) {
	safeReports := 0
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := regexp.MustCompile(`\s+`).Split(line, -1)
		var levels []int
		for _, part := range parts {
			if num, err := strconv.Atoi(strings.TrimSpace(part)); err == nil {
				levels = append(levels, num)
			}
		}
		if len(levels) >= 2 && isReportSafe(levels) {
			safeReports++
		}
	}
	return fmt.Sprintf("Safe reports: %d", safeReports), nil
}

// Ch2 вторая часть задачи 2-го дня - проверка безопасности с Problem Dampener
func (d *Day2Solution) Ch2(input []byte) (string, error) {
	safeReports := 0
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := regexp.MustCompile(`\s+`).Split(line, -1)
		var levels []int
		for _, part := range parts {
			if num, err := strconv.Atoi(strings.TrimSpace(part)); err == nil {
				levels = append(levels, num)
			}
		}
		if len(levels) >= 2 && isReportSafeWithDampener(levels) {
			safeReports++
		}
	}
	return fmt.Sprintf("Safe reports with dampener: %d", safeReports), nil
}

// isReportSafe проверяет, безопасен ли отчет
func isReportSafe(levels []int) bool {
	if len(levels) < 2 {
		return false
	}
	for i := 1; i < len(levels); i++ {
		if abs(levels[i]-levels[i-1]) > 10 {
			return false
		}
	}
	return true
}

// isReportSafeWithDampener проверяет, безопасен ли отчет с учетом Problem Dampener
func isReportSafeWithDampener(levels []int) bool {
	if isReportSafe(levels) {
		return true
	}
	for i := 0; i < len(levels); i++ {
		newLevels := make([]int, 0, len(levels)-1)
		newLevels = append(newLevels, levels[:i]...)
		newLevels = append(newLevels, levels[i+1:]...)
		if len(newLevels) >= 2 && isReportSafe(newLevels) {
			return true
		}
	}
	return false
}
