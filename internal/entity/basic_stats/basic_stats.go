package basicstats

import (
	"fmt"
	"time"

	"github.com/NikitaTumanov/Web-Application-Log-Aggregator-Analyzer/internal/entity/log"
)

type Statistics struct {
	TotalRequests int
	UniqueIPs     int
	TotalBytes    int64
	StartTime     time.Time
	EndTime       time.Time
	Methods       map[string]int
	StatusCodes   map[int]int
	TopPaths      map[string]int
	TopIPs        map[string]int
	UserAgents    map[string]int
}

func GetBasicStats(logs []log.Log) Statistics {
	if len(logs) == 0 {
		return Statistics{}
	}

	stats := Statistics{
		Methods:     make(map[string]int),
		StatusCodes: make(map[int]int),
		TopPaths:    make(map[string]int),
		TopIPs:      make(map[string]int),
		UserAgents:  make(map[string]int),
		StartTime:   logs[0].Timestamp,
		EndTime:     logs[0].Timestamp,
	}

	ipSet := make(map[string]bool)

	for _, log := range logs {
		stats.TotalRequests++

		ipSet[log.IP] = true

		stats.TotalBytes += int64(log.Size)

		stats.Methods[log.Method]++

		stats.StatusCodes[log.Status]++

		stats.TopPaths[log.Path]++

		stats.TopIPs[log.IP]++

		stats.UserAgents[log.UserAgent]++

		if log.Timestamp.Before(stats.StartTime) {
			stats.StartTime = log.Timestamp
		}
		if log.Timestamp.After(stats.EndTime) {
			stats.EndTime = log.Timestamp
		}
	}

	stats.UniqueIPs = len(ipSet)
	return stats
}

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func DisplayStats(stats Statistics) {
	fmt.Printf("Total Requests: %d\n", stats.TotalRequests)
	fmt.Printf("Unique IPs: %d\n", stats.UniqueIPs)
	fmt.Printf("Total Data Transferred: %s\n", formatBytes(stats.TotalBytes))
	fmt.Printf("Time Range: %s - %s\n",
		stats.StartTime.Format("2006-01-02 15:04:05"),
		stats.EndTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("Duration: %v\n", stats.EndTime.Sub(stats.StartTime))

	fmt.Println("\nHTTP Methods:")
	for method, count := range stats.Methods {
		percent := float64(count) / float64(stats.TotalRequests) * 100
		fmt.Printf("  %-6s: %d (%.1f%%)\n", method, count, percent)
	}

	fmt.Println("\nStatus Codes:")
	for code, count := range stats.StatusCodes {
		percent := float64(count) / float64(stats.TotalRequests) * 100
		fmt.Printf("  %d: %d (%.1f%%)\n", code, count, percent)
	}
}
