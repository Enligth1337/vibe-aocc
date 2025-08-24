package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

// Day5Solution решение для 5-го дня
type Day5Solution struct{}

// NewDay5Solution создает новый экземпляр Day5Solution
func NewDay5Solution() *Day5Solution {
	return &Day5Solution{}
}

// Ch1 первая часть задачи 5-го дня - проверка порядка страниц
func (d *Day5Solution) Ch1(input []byte) (string, error) {
	lines := strings.Split(string(input), "\n")

	// Разделяем правила и обновления
	var rules []string
	var updates []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Если встречаем строку с запятыми, это обновление
		if strings.Contains(line, ",") {
			updates = append(updates, line)
		} else {
			// Иначе это правило (формат X|Y)
			rules = append(rules, line)
		}
	}

	// Строим граф зависимостей
	dependencies := d.buildDependencies(rules)

	// Проверяем каждое обновление
	var correctUpdates []string
	var middlePages []int

	for _, update := range updates {
		if d.isUpdateCorrect(update, dependencies) {
			correctUpdates = append(correctUpdates, update)
			middlePage := d.getMiddlePage(update)
			middlePages = append(middlePages, middlePage)
		}
	}

	// Суммируем средние страницы
	total := 0
	for _, page := range middlePages {
		total += page
	}

	return fmt.Sprintf("Correct updates: %d, Middle pages sum: %d", len(correctUpdates), total), nil
}

// Ch2 вторая часть задачи 5-го дня - исправление порядка страниц
func (d *Day5Solution) Ch2(input []byte) (string, error) {
	lines := strings.Split(string(input), "\n")

	// Разделяем правила и обновления
	var rules []string
	var updates []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Если встречаем строку с запятыми, это обновление
		if strings.Contains(line, ",") {
			updates = append(updates, line)
		} else {
			// Иначе это правило (формат X|Y)
			rules = append(rules, line)
		}
	}

	// Строим граф зависимостей
	dependencies := d.buildDependencies(rules)

	// Находим некорректные обновления и исправляем их
	var incorrectUpdates []string
	var correctedMiddlePages []int

	for _, update := range updates {
		if !d.isUpdateCorrect(update, dependencies) {
			incorrectUpdates = append(incorrectUpdates, update)
			correctedOrder := d.correctUpdateOrder(update, dependencies)
			middlePage := d.getMiddlePageFromPages(correctedOrder)
			correctedMiddlePages = append(correctedMiddlePages, middlePage)
		}
	}

	// Суммируем средние страницы исправленных обновлений
	total := 0
	for _, page := range correctedMiddlePages {
		total += page
	}

	return fmt.Sprintf("Incorrect updates: %d, Corrected middle pages sum: %d", len(incorrectUpdates), total), nil
}

// buildDependencies строит граф зависимостей из правил
func (d *Day5Solution) buildDependencies(rules []string) map[int][]int {
	deps := make(map[int][]int)

	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		if len(parts) != 2 {
			continue
		}

		before, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		after, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))

		if err1 != nil || err2 != nil {
			continue
		}

		deps[before] = append(deps[before], after)
	}

	return deps
}

// isUpdateCorrect проверяет, соответствует ли обновление правилам
func (d *Day5Solution) isUpdateCorrect(update string, dependencies map[int][]int) bool {
	pages := d.parseUpdate(update)

	// Проверяем каждую пару страниц
	for i := 0; i < len(pages); i++ {
		for j := i + 1; j < len(pages); j++ {
			before := pages[i]
			after := pages[j]

			// Если есть правило, что after должно быть перед before, то порядок неправильный
			// Но только если обе страницы присутствуют в обновлении
			if d.hasDirectDependency(after, before, dependencies) {
				return false
			}
		}
	}

	return true
}

// hasDirectDependency проверяет, есть ли прямая зависимость от from к to
func (d *Day5Solution) hasDirectDependency(from, to int, dependencies map[int][]int) bool {
	// Проверяем только прямые зависимости, без транзитивности
	for _, dep := range dependencies[from] {
		if dep == to {
			return true
		}
	}
	return false
}

// hasDependency проверяет, есть ли зависимость от from к to
func (d *Day5Solution) hasDependency(from, to int, dependencies map[int][]int) bool {
	visited := make(map[int]bool)
	return d.dfs(from, to, dependencies, visited)
}

// dfs выполняет поиск в глубину для проверки зависимости
func (d *Day5Solution) dfs(from, to int, dependencies map[int][]int, visited map[int]bool) bool {
	if from == to {
		return true
	}

	visited[from] = true

	for _, next := range dependencies[from] {
		if !visited[next] {
			if d.dfs(next, to, dependencies, visited) {
				return true
			}
		}
	}

	return false
}

// parseUpdate парсит строку обновления в массив страниц
func (d *Day5Solution) parseUpdate(update string) []int {
	parts := strings.Split(update, ",")
	var pages []int

	for _, part := range parts {
		page, err := strconv.Atoi(strings.TrimSpace(part))
		if err == nil {
			pages = append(pages, page)
		}
	}

	return pages
}

// getMiddlePage возвращает среднюю страницу из обновления
func (d *Day5Solution) getMiddlePage(update string) int {
	pages := d.parseUpdate(update)
	if len(pages) == 0 {
		return 0
	}

	// Для нечетного количества страниц - средняя, для четного - правая от центра
	middleIndex := len(pages) / 2
	return pages[middleIndex]
}

// min возвращает минимальное из двух чисел
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// correctUpdateOrder исправляет порядок страниц в обновлении согласно правилам
func (d *Day5Solution) correctUpdateOrder(update string, dependencies map[int][]int) []int {
	pages := d.parseUpdate(update)

	// Создаем граф зависимостей только для страниц в этом обновлении
	updateDeps := make(map[int][]int)

	// Добавляем зависимости между страницами в обновлении
	for _, page := range pages {
		for _, dep := range dependencies[page] {
			// Проверяем, что зависимость тоже присутствует в обновлении
			if d.contains(pages, dep) {
				updateDeps[page] = append(updateDeps[page], dep)
			}
		}
	}

	// Выполняем топологическую сортировку
	return d.topologicalSort(pages, updateDeps)
}

// contains проверяет, содержится ли элемент в слайсе
func (d *Day5Solution) contains(slice []int, item int) bool {
	for _, val := range slice {
		if val == item {
			return true
		}
	}
	return false
}

// topologicalSort выполняет топологическую сортировку
func (d *Day5Solution) topologicalSort(pages []int, dependencies map[int][]int) []int {
	// Считаем входящие степени для каждой страницы
	inDegree := make(map[int]int)
	for _, page := range pages {
		inDegree[page] = 0
	}

	for _, deps := range dependencies {
		for _, dep := range deps {
			inDegree[dep]++
		}
	}

	// Очередь для страниц с нулевой входящей степенью
	var queue []int
	for page, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}

	var result []int

	// Обрабатываем очередь
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		// Уменьшаем входящие степени для зависимостей
		for _, dep := range dependencies[current] {
			inDegree[dep]--
			if inDegree[dep] == 0 {
				queue = append(queue, dep)
			}
		}
	}

	// Если не все страницы обработаны, добавляем оставшиеся в конец
	for _, page := range pages {
		if !d.contains(result, page) {
			result = append(result, page)
		}
	}

	return result
}

// getMiddlePageFromPages возвращает среднюю страницу из слайса страниц
func (d *Day5Solution) getMiddlePageFromPages(pages []int) int {
	if len(pages) == 0 {
		return 0
	}

	// Для нечетного количества страниц - средняя, для четного - правая от центра
	middleIndex := len(pages) / 2
	return pages[middleIndex]
}
