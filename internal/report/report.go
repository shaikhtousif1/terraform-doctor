package report

import (
    "fmt"
)

type Severity string

const (
    SeverityInfo    Severity = "INFO"
    SeverityWarning Severity = "WARN"
    SeverityError   Severity = "ERROR"
)

type Issue struct {
    Severity Severity
    Message  string
}

// PrintIssues prints human-friendly issues to stdout.
func PrintIssues(issues []Issue) {
    if len(issues) == 0 {
        fmt.Println("✅ No design issues found.")
        return
    }

    for _, it := range issues {
        switch it.Severity {
        case SeverityWarning:
            fmt.Printf("⚠️ %s\n", it.Message)
        case SeverityError:
            fmt.Printf("❌ %s\n", it.Message)
        default:
            fmt.Printf("ℹ️ %s\n", it.Message)
        }
    }
}
