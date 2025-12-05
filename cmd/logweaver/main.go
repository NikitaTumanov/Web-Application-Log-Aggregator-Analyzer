package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	inputcommand "github.com/NikitaTumanov/Web-Application-Log-Aggregator-Analyzer/internal/entity/input_command"
	logrepository "github.com/NikitaTumanov/Web-Application-Log-Aggregator-Analyzer/internal/repository/log_repository"
	apachelogservice "github.com/NikitaTumanov/Web-Application-Log-Aggregator-Analyzer/internal/usecase/apache_log_service"
)

func main() {
	var (
		command  string
		filePath = flag.String("file_path", "", "file_path")
		logType  = flag.String("log_type", "", "log_type")

		ip        = flag.String("ip", "", "ip")
		identd    = flag.String("identd", "", "identd")
		user      = flag.String("user", "", "user")
		timestamp = flag.String("timestamp", "", "timestamp")
		method    = flag.String("method", "", "method")
		path      = flag.String("path", "", "path")
		protocol  = flag.String("protocol", "", "protocol")
		status    = flag.Int("status", -1, "status")
		size      = flag.Int("size", -1, "size")
		referer   = flag.String("referer", "", "referer")
		userAgent = flag.String("user_agent", "", "user_agent")
	)

	flag.StringVar(&command, "c", "", "command")
	flag.Parse()

	var t time.Time
	var err error
	if *timestamp != "" {
		t, err = time.Parse("02/Jan/2006:15:04:05 -0700", *timestamp)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	input := &inputcommand.InputCommand{
		Command:   command,
		FilePath:  *filePath,
		LogType:   *logType,
		IP:        *ip,
		Identd:    *identd,
		User:      *user,
		Timestamp: t,
		Method:    *method,
		Path:      *path,
		Protocol:  *protocol,
		Status:    *status,
		Size:      *size,
		Referer:   *referer,
		UserAgent: *userAgent,
	}

	logService := &logrepository.Analyzer{}

	switch strings.ToLower(*logType) {
	case "apache":
		// надо сделать интерфейс хэндлер, который будет запускать метод хэндл
		// внутри него будет вызываться парсер файла и делать список логов отфильтрованных
		// с помощью потоков (каналов) построчно фильтровать и записывать логи в список
		// после чего будет вызываться визуализатор статистики логов
		logService.SetLogService(&apachelogservice.ApacheLogService{})
	default:
		fmt.Println("incorrect log type")
		return
	}

	logs, err := logService.ParseLogs(input.FilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	logs, err = logService.FilterLogs(logs, *input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(logs)
	// logService.LogsStat()
}
