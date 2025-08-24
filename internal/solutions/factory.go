package solutions

import (
	"fmt"
)

// SolutionFactoryImpl реализация SolutionFactory
type SolutionFactoryImpl struct {
	solutions map[string]Solution
}

// NewSolutionFactory создает новый экземпляр SolutionFactoryImpl
func NewSolutionFactory() *SolutionFactoryImpl {
	factory := &SolutionFactoryImpl{
		solutions: make(map[string]Solution),
	}

	// Автоматически регистрируем все решения
	days := []string{"day1", "day2", "day3", "day4", "day5", "day6", "day7", "day8", "day9", "day10",
		"day11", "day12", "day13", "day14", "day15", "day16", "day17", "day18", "day19", "day20",
		"day21", "day22", "day23", "day24", "day25"}

	// Регистрируем все решения
	for _, day := range days {
		switch day {
		case "day1":
			factory.RegisterSolution(day, NewDay1Solution())
		case "day2":
			factory.RegisterSolution(day, NewDay2Solution())
		case "day3":
			factory.RegisterSolution(day, NewDay3Solution())
		case "day4":
			factory.RegisterSolution(day, NewDay4Solution())
		case "day5":
			factory.RegisterSolution(day, NewDay5Solution())
		case "day6":
			factory.RegisterSolution(day, NewDay6Solution())
		case "day7":
			factory.RegisterSolution(day, NewDay7Solution())
		case "day8":
			factory.RegisterSolution(day, NewDay8Solution())
		case "day9":
			factory.RegisterSolution(day, NewDay9Solution())
		case "day10":
			factory.RegisterSolution(day, NewDay10Solution())
		case "day11":
			factory.RegisterSolution(day, NewDay11Solution())
		case "day12":
			factory.RegisterSolution(day, NewDay12Solution())
		case "day13":
			factory.RegisterSolution(day, NewDay13Solution())
		case "day14":
			factory.RegisterSolution(day, NewDay14Solution())
		case "day15":
			factory.RegisterSolution(day, NewDay15Solution())
		case "day16":
			factory.RegisterSolution(day, NewDay16Solution())
		case "day17":
			factory.RegisterSolution(day, NewDay17Solution())
		case "day18":
			factory.RegisterSolution(day, NewDay18Solution())
		case "day19":
			factory.RegisterSolution(day, NewDay19Solution())
		case "day20":
			factory.RegisterSolution(day, NewDay20Solution())
		case "day21":
			factory.RegisterSolution(day, NewDay21Solution())
		case "day22":
			factory.RegisterSolution(day, NewDay22Solution())
		case "day23":
			factory.RegisterSolution(day, NewDay23Solution())
		case "day24":
			factory.RegisterSolution(day, NewDay24Solution())
		case "day25":
			factory.RegisterSolution(day, NewDay25Solution())
		}
	}

	return factory
}

// RegisterSolution регистрирует решение по ключу
func (f *SolutionFactoryImpl) RegisterSolution(key string, solution Solution) {
	f.solutions[key] = solution
}

// New создает решение по ключу
func (f *SolutionFactoryImpl) New(key string) (Solution, error) {
	solution, exists := f.solutions[key]
	if !exists {
		return nil, fmt.Errorf("solution not found for key: %s", key)
	}
	return solution, nil
}
