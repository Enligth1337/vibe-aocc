package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Task представляет задачу для решения
type Task struct {
	Day       string `yaml:"day"`
	InputFile string `yaml:"input_file"`
}

// Config представляет основную конфигурацию
type Config struct {
	InputResourcesPath string `yaml:"input_resources_path"`
	Tasks              []Task `yaml:"tasks"`
}

// Load загружает конфигурацию из файла
func Load(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

// GetTaskByDay возвращает задачу по номеру дня
func (c *Config) GetTaskByDay(day string) (*Task, error) {
	for _, task := range c.Tasks {
		if task.Day == day {
			return &task, nil
		}
	}
	return nil, fmt.Errorf("task not found for day: %s", day)
}

// GetInputFilePath возвращает полный путь к файлу ввода для задачи
func (c *Config) GetInputFilePath(task *Task) string {
	return fmt.Sprintf("%s/%s", c.InputResourcesPath, task.InputFile)
}
