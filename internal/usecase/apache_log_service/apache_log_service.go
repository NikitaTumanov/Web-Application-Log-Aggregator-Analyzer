package apachelogservice

import (
	"bufio"
	"fmt"
	"os"

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

}

func (a *ApacheLogService) Filter(logs []log.Log, input inputcommand.InputCommand) ([]log.Log, error) {

}

func (a *ApacheLogService) Statistic(logs []log.Log) string {

}
