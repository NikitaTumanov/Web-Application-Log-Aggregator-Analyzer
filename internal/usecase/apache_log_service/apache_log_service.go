package apachelogservice

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"time"

	inputcommand "github.com/NikitaTumanov/Web-Application-Log-Aggregator-Analyzer/internal/entity/input_command"
	"github.com/NikitaTumanov/Web-Application-Log-Aggregator-Analyzer/internal/entity/log"
	apachelogparser "github.com/NikitaTumanov/Web-Application-Log-Aggregator-Analyzer/pkg/apache_log_parser"
)

type ApacheLogService struct{}

func (a *ApacheLogService) Parse(filePath string) ([]log.Log, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	parser := apachelogparser.NewApacheLogParser()
	logs := make([]log.Log, 0)

	for scanner.Scan() {
		line := scanner.Text()

		log, err := parser.Parse(line)
		if err != nil {
			return nil, fmt.Errorf("parser.Parse: %w", err)
		}

		logs = append(logs, *log)
	}

	return logs, nil
}

func filterByField[T comparable](logs []log.Log, fieldValue T, fieldName string) ([]log.Log, error) {
	result := make([]log.Log, 0)

	for _, log := range logs {
		v := reflect.ValueOf(log)
		field := v.FieldByName(fieldName)
		if !field.IsValid() {
			return nil, fmt.Errorf("field '%s' does not exist", fieldName)
		}
		if !field.CanInterface() {
			return nil, fmt.Errorf("field '%s' is not accessible", fieldName)
		}

		fieldInterface := field.Interface()

		fieldVal, ok := fieldInterface.(T)
		if !ok {
			if reflect.DeepEqual(fieldInterface, fieldValue) {
				result = append(result, log)
			}
			continue
		}

		if fieldVal == fieldValue {
			result = append(result, log)
		}
	}

	return result, nil
}

func (a *ApacheLogService) Filter(logs []log.Log, input inputcommand.InputCommand) ([]log.Log, error) {
	v := reflect.ValueOf(input)
	t := reflect.TypeOf(input)

	for i := 3; i < v.NumField(); i++ {
		fieldValue := v.Field(i).Interface()
		fieldName := t.Field(i).Name

		if fieldValue == "" || fieldValue == -1 || fieldValue == time.Date(0001, 01, 01, 00, 00, 00, 00, time.UTC) {
			continue
		}

		var err error
		logs, err = filterByField(logs, fieldValue, fieldName)
		if err != nil {
			return nil, fmt.Errorf("filterByField: %w", err)
		}
	}

	return logs, nil
}

func (a *ApacheLogService) Statistic(logs []log.Log) string {
	return ""
}
