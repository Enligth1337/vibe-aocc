package input

import (
	"bufio"
	"strings"
)

// Parser интерфейс для парсинга input файлов
type Parser interface {
	ParseInput(data []byte) ([]string, error)
}

// FileParser реализация Parser для чтения файлов
type FileParser struct{}

// NewFileParser создает новый экземпляр FileParser
func NewFileParser() *FileParser {
	return &FileParser{}
}

// ParseInput читает содержимое файла и возвращает срез строк
func (p *FileParser) ParseInput(data []byte) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	return lines, scanner.Err()
}
