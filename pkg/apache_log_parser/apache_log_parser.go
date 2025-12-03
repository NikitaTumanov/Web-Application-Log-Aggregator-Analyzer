package apachelogparser

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"time"

	apachelog "github.com/NikitaTumanov/Web-Application-Log-Aggregator-Analyzer/internal/entity/log"
)

type ApacheLogParser struct {
	pattern *regexp.Regexp
}

func NewApacheLogParser() *ApacheLogParser {
	pattern := `^(\S+) (\S+) (\S+) \[([^\]]+)\] "(\S+) (\S+) (\S+)" (\d{3}) (\d+) "([^"]*)" "([^"]*)"`
	return &ApacheLogParser{
		pattern: regexp.MustCompile(pattern),
	}
}

func (p *ApacheLogParser) setFieldValue(fieldValue reflect.Value, value, parseType string) error {
	if value == "-" {
		value = ""
	}

	switch parseType {
	case "string":
		fieldValue.SetString(value)
	case "int":
		if value == "" {
			fieldValue.SetInt(0)
		} else {
			intVal, err := strconv.Atoi(value)
			if err != nil {
				return fmt.Errorf("strconv.Atoi(): %w", err)
			}
			fieldValue.SetInt(int64(intVal))
		}
	case "apache_time":
		t, err := time.Parse("02/Jan/2006:15:04:05 -0700", value)
		if err != nil {
			return fmt.Errorf("time.Parse(): %w", err)
		}
		fieldValue.Set(reflect.ValueOf(t))
	default:
		fieldValue.SetString(value)
	}
	return nil
}

func (p *ApacheLogParser) Parse(line string) (*apachelog.Log, error) {
	matches := p.pattern.FindStringSubmatch(line)
	if matches == nil {
		return nil, fmt.Errorf("log line doesn't match pattern")
	}

	logEntry := &apachelog.Log{}
	v := reflect.ValueOf(logEntry).Elem()
	vType := v.Type()

	for i := 1; i < len(matches) && i-1 < vType.NumField(); i++ {
		field := vType.Field(i - 1)
		parseType := field.Tag.Get("parse")
		value := matches[i]

		fieldValue := v.Field(i - 1)
		if err := p.setFieldValue(fieldValue, value, parseType); err != nil {
			return nil, fmt.Errorf("error parsing field %s: %w", field.Name, err)
		}
	}
	return logEntry, nil
}
