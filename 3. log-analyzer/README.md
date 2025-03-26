# Log Analyzer

A Go-based application for analyzing HTTP logs and summarizing key metrics, such as requests per pod, HTTP response codes, bytes sent, and unique request paths.

# How It Works

1. **Input**:
   - The application takes a log file containing HTTP request logs as input.
   - The log file should be in a space-separated format and follow a structured log format.

2. **Log Parsing**:
   - The `logparser` package processes the log file and extracts key fields such as:
     - Pod name
     - Request path
     - HTTP response code
     - Bytes sent
     - Timestamp

3. **Analysis**:
   - The parsed log entries are analyzed to summarize key metrics:
     - Total requests per pod.
     - Distribution of HTTP response codes.
     - Total bytes sent by each pod.
     - List of unique request paths.

4. **Output**:
   - The results are displayed in a human-readable format, including tables and a numbered list of unique paths.

---

# Usage

1. **Run the Program**:
   - Use the following command to run the application:
     ```bash
     go run main.go -file=path/to/logfile.txt
     ```

2. **Required Parameter**:
   - `-file`: Provide the path to the log file for analysis.

3. **Example Command**:
   ```bash
   go run main.go -file=fp-sre-challenge.log
   ```
# Key Components
### Main Application (main.go)
Handles command-line arguments and validates the input.
Coordinates log parsing, analysis, and result display.

### Log Parser (logparser/parser.go)
Reads the log file and extracts structured log entries.
Handles parsing of fields like pod name, request path, HTTP response code, and bytes sent.
Skips invalid log lines with warnings.

### Models (models/log_entry.go)
Defines the data structure (LogEntry) for a parsed log entry.
Includes fields such as PodName, RequestPath, HTTPResponseCode, and BytesSent.

# Output
```
Analyzing log file: fp-sre-challenge.log

Successfully parsed 100 log entries

=== LOG ANALYSIS SUMMARY ===

1. TOTAL REQUESTS PER POD:
   Pod Name                               | Requests
   ---------------------------------------|---------
   pod-1                                  | 50
   pod-2                                  | 30
   pod-3                                  | 20

2. RESPONSE CODE SUMMARY:
   HTTP Status                            | Count
   ---------------------------------------|---------
   200                                    | 80
   404                                    | 15
   500                                    | 5

3. TOTAL BYTES SENT PER POD:
   Pod Name                               | Bytes Sent
   ---------------------------------------|---------
   pod-1                                  | 10240
   pod-2                                  | 7680
   pod-3                                  | 5120

4. UNIQUE REQUEST PATHS:
   1. /home
   2. /login
   3. /api/v1/resource
```