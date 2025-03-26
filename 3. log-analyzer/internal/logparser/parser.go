package logparser

import (
	"bufio"
	"fmt"
	"log-analyzer/internal/models"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Regex for extracting HTTP request components from log lines
var requestRegex = regexp.MustCompile(`"(GET|POST|PUT|DELETE|HEAD|OPTIONS|PATCH|CONNECT|TRACE)\s+([^"]+)\s+(HTTP/[0-9.]+)"`)

// ParseLogFile reads a log file and returns structured log entries
func ParseLogFile(filePath string) ([]models.LogEntry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening log file: %w", err)
	}
	defer file.Close()

	var logEntries []models.LogEntry
	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		// Skip empty lines
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		entry, err := parseLine(line, lineNum)
		if err != nil {
			fmt.Printf("Warning: %v\n", err)
			continue
		}

		logEntries = append(logEntries, entry)
	}

	if err := scanner.Err(); err != nil {
		return logEntries, fmt.Errorf("error reading log file: %w", err)
	}

	return logEntries, nil
}

// parseLine extracts data from a log line into structured format
func parseLine(line string, lineNum int) (models.LogEntry, error) {
	parts := strings.Split(line, " ")
	if len(parts) < 20 {
		return models.LogEntry{}, fmt.Errorf("line %d: insufficient fields, found %d, expected at least 20",
			lineNum, len(parts))
	}

	// Find and extract HTTP request parts (method, path, protocol)
	var method, path, protocol string
	requestString := ""

	// Locate the request portion within the log line
	for i, part := range parts {
		if strings.HasPrefix(part, "\"GET") || strings.HasPrefix(part, "\"POST") ||
			strings.HasPrefix(part, "\"PUT") || strings.HasPrefix(part, "\"DELETE") {

			// Reconstruct the complete request string
			requestStart := i
			requestEnd := i
			for j := i; j < len(parts); j++ {
				if strings.HasSuffix(parts[j], "\"") && j > i {
					requestEnd = j
					break
				}
			}

			requestParts := parts[requestStart : requestEnd+1]
			requestString = strings.Join(requestParts, " ")
			requestString = strings.Trim(requestString, "\"")
			break
		}
	}

	// Extract request components
	matches := requestRegex.FindStringSubmatch(fmt.Sprintf("\"%s\"", requestString))
	if len(matches) >= 4 {
		method = matches[1]
		path = matches[2]
		protocol = matches[3]
	} else {
		method, path, protocol = "UNKNOWN", "UNKNOWN", "UNKNOWN"
	}

	// Find HTTP status code and response size
	var responseCode string
	var bytesSent int
	for i, part := range parts {
		if i > 0 && (part == "200" || part == "404" || part == "500" || part == "302" ||
			part == "301" || part == "400" || part == "401" || part == "403" ||
			part == "201" || part == "204") {
			responseCode = part
			if i+1 < len(parts) {
				bytesSentStr := parts[i+1]
				bytesSentVal, err := strconv.Atoi(bytesSentStr)
				if err == nil {
					bytesSent = bytesSentVal
				}
			}
			break
		}
	}

	// Extract timestamp (usually in brackets)
	var timestamp string
	for _, part := range parts {
		if strings.HasPrefix(part, "[") && strings.Contains(part, ":") {
			timestamp = part
			if !strings.HasSuffix(part, "]") && len(parts) > 1 {
				for j := range parts {
					if strings.HasSuffix(parts[j], "]") {
						timestamp = timestamp + " " + parts[j]
						break
					}
				}
			}
			timestamp = strings.Trim(timestamp, "[]")
			break
		}
	}

	// Extract user agent information
	var clientInfo string
	inClientInfo := false
	clientInfoParts := []string{}

	for _, part := range parts {
		if inClientInfo {
			clientInfoParts = append(clientInfoParts, part)
			if strings.HasSuffix(part, "\"") {
				inClientInfo = false
				break
			}
		} else if strings.HasPrefix(part, "\"Mozilla") || strings.HasPrefix(part, "\"AppleWebKit") {
			inClientInfo = true
			clientInfoParts = append(clientInfoParts, part)
		}
	}

	clientInfo = strings.Join(clientInfoParts, " ")
	clientInfo = strings.Trim(clientInfo, "\"")

	// Extract service IP address
	var serviceIP string
	for i := len(parts) - 10; i < len(parts); i++ {
		if i >= 0 && strings.Contains(parts[i], ":") && strings.Contains(parts[i], ".") {
			serviceIP = parts[i]
			break
		}
	}

	// Build the log entry structure
	entry := models.LogEntry{
		PodName:          parts[0],
		ContainerName:    parts[1],
		IPAddress:        parts[2],
		Timestamp:        timestamp,
		RequestMethod:    method,
		RequestPath:      path,
		HTTPProtocol:     protocol,
		HTTPResponseCode: responseCode,
		BytesSent:        bytesSent,
		ClientInfo:       clientInfo,
		ServiceIP:        serviceIP,
	}

	// Add trace ID if present
	if len(parts) > 0 && len(parts[len(parts)-1]) > 10 {
		entry.TraceID = parts[len(parts)-1]
	}

	return entry, nil
}

// AnalyzeLogEntries aggregates statistics from log entries
func AnalyzeLogEntries(entries []models.LogEntry) (map[string]int, map[string]int, map[string]int, []string) {
	requestsPerPod := make(map[string]int) // Counter for requests by pod
	responseCodes := make(map[string]int)  // Counter for HTTP response codes
	bytesPerPod := make(map[string]int)    // Total bytes sent by each pod
	uniquePaths := make(map[string]bool)   // Set of distinct request paths

	// Process each log entry
	for _, entry := range entries {
		requestsPerPod[entry.PodName]++
		responseCodes[entry.HTTPResponseCode]++
		bytesPerPod[entry.PodName] += entry.BytesSent
		uniquePaths[entry.RequestPath] = true
	}

	// Convert unique paths map to slice
	paths := make([]string, 0, len(uniquePaths))
	for path := range uniquePaths {
		paths = append(paths, path)
	}

	return requestsPerPod, responseCodes, bytesPerPod, paths
}
