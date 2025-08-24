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

	// Регистрируем все доступные решения
	factory.RegisterSolution("day1", NewDay1Solution())
	factory.RegisterSolution("day2", NewDay2Solution())
	factory.RegisterSolution("day3", NewDay3Solution())
	factory.RegisterSolution("day4", NewDay4Solution())
	factory.RegisterSolution("day5", NewDay5Solution())
	factory.RegisterSolution("day6", NewDay6Solution())
	factory.RegisterSolution("day7", NewDay7Solution())
	factory.RegisterSolution("day8", NewDay8Solution())
	factory.RegisterSolution("day9", NewDay9Solution())
	factory.RegisterSolution("day10", NewDay10Solution())
	factory.RegisterSolution("day11", NewDay11Solution())
	factory.RegisterSolution("day12", NewDay12Solution())
	factory.RegisterSolution("day13", NewDay13Solution())
	factory.RegisterSolution("day14", NewDay14Solution())
	factory.RegisterSolution("day15", NewDay15Solution())
	factory.RegisterSolution("day16", NewDay16Solution())
	factory.RegisterSolution("day17", NewDay17Solution())
	factory.RegisterSolution("day18", NewDay18Solution())
	factory.RegisterSolution("day19", NewDay19Solution())
	factory.RegisterSolution("day20", NewDay20Solution())
	factory.RegisterSolution("day21", NewDay21Solution())
	factory.RegisterSolution("day22", NewDay22Solution())
	factory.RegisterSolution("day23", NewDay23Solution())
	factory.RegisterSolution("day24", NewDay24Solution())
	factory.RegisterSolution("day25", NewDay25Solution())

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
