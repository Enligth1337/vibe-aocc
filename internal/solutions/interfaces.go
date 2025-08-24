package solutions

// Solution интерфейс для решения задач
type Solution interface {
	Ch1(input []byte) (string, error)
	Ch2(input []byte) (string, error)
}

// SolutionFactory интерфейс для создания решений
type SolutionFactory interface {
	New(key string) (Solution, error)
}
