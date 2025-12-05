package logrepository

import (
	"fmt"

	inputcommand "github.com/NikitaTumanov/Web-Application-Log-Aggregator-Analyzer/internal/entity/input_command"
	"github.com/NikitaTumanov/Web-Application-Log-Aggregator-Analyzer/internal/entity/log"
)

type LogRepository interface {
	Parse(filePath string) ([]log.Log, error)
	Filter(logs []log.Log, input inputcommand.InputCommand) ([]log.Log, error)
	Statistic(logs []log.Log)
}

type Analyzer struct {
	logRepository LogRepository
}

func (a *Analyzer) SetLogService(logService LogRepository) {
	a.logRepository = logService
}

func (a *Analyzer) ParseLogs(filePath string) ([]log.Log, error) {
	result, err := a.logRepository.Parse(filePath)
	if err != nil {
		return nil, fmt.Errorf("a.logRepository.Parse: %w", err)
	}
	return result, nil
}

func (a *Analyzer) FilterLogs(logs []log.Log, input inputcommand.InputCommand) ([]log.Log, error) {
	result, err := a.logRepository.Filter(logs, input)
	if err != nil {
		return nil, fmt.Errorf("a.logRepository.Filter: %w", err)
	}
	return result, nil
}

func (a *Analyzer) LogsStat(logs []log.Log) {
	a.logRepository.Statistic(logs)
}
