package main

import (
	"flag"
	"fmt"
	"log-analyzer/internal/logparser"
	"os"
	"sort"
	"strings"
)

// main is the entry point of the log analysis application.
// It processes command-line arguments, parses the log file,
// analyzes log data, and displays a summary of results.
func main() {
	logFile := flag.String("file", "", "Path to the log file")
	flag.Parse()

	if *logFile == "" {
		fmt.Println("Error: Log file path is required")
		fmt.Println("Usage: go run main.go -file=path/to/logfile.txt")
		os.Exit(1)
	}

	fmt.Printf("Analyzing log file: %s\n", *logFile)

	// Parse the log file using the logparser package
	// This converts raw log data into structured logEntries
	logEntries, err := logparser.ParseLogFile(*logFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully parsed %d log entries\n\n", len(logEntries))

	// Analyze the log entries to extract meaningful metrics
	// The function returns several analysis results:
	// - requestsPerPod: Count of requests handled by each pod
	// - responseCodes: Distribution of HTTP response codes
	// - bytesPerPod: Total bytes sent by each pod
	// - uniquePaths: List of distinct request paths
	requestsPerPod, responseCodes, bytesPerPod, uniquePaths := logparser.AnalyzeLogEntries(logEntries)
	fmt.Println("=== LOG ANALYSIS SUMMARY ===")

	// 1. Display total requests per pod in descending order
	fmt.Println("\n1. TOTAL REQUESTS PER POD:")
	printSortedMap(requestsPerPod, "Pod", "Requests")

	// 2. Display HTTP response code distribution
	fmt.Println("\n2. RESPONSE CODE SUMMARY:")
	printSortedMap(responseCodes, "HTTP Status", "Count")

	// 3. Display total bytes sent by each pod
	fmt.Println("\n3. TOTAL BYTES SENT PER POD:")
	printSortedMap(bytesPerPod, "Pod", "Bytes Sent")

	// 4. Display all unique request paths in alphabetical order
	fmt.Println("\n4. UNIQUE REQUEST PATHS:")
	sort.Strings(uniquePaths)
	for i, path := range uniquePaths {
		fmt.Printf("   %d. %s\n", i+1, path)
	}
}

// printSortedMap displays the contents of a map with keys sorted alphabetically.
// It formats the output as a two-column table with custom headers.
//
// Parameters:
//   - m: The map to display (string keys, integer values)
//   - keyLabel: The header label for the keys column
//   - valueLabel: The header label for the values column
func printSortedMap(m map[string]int, keyLabel, valueLabel string) {
	// Extract all keys from the map for sorting
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// Sort keys alphabetically for consistent output
	sort.Strings(keys)

	// Print table header with column labels
	fmt.Printf("   %-40s | %s\n", keyLabel, valueLabel)
	fmt.Printf("   %s\n", strings.Repeat("-", 55)) // Separator line

	// Print each key-value pair in the sorted order
	for _, k := range keys {
		fmt.Printf("   %-40s | %d\n", k, m[k])
	}
}
