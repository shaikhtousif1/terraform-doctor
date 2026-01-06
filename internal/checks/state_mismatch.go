package checks

import (
    "io/ioutil"
    "path/filepath"
    "strings"

    "github.com/example/terraform-doctor/internal/report"
)

// DetectStateMismatch looks for signs of mixed backends (local vs remote) across the codebase.
// This is heuristic-based and intentionally conservative.
func DetectStateMismatch(root string) []report.Issue {
    var issues []report.Issue
    foundLocal := false
    foundRemote := false

    _ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        package checks

        import (
            "io/ioutil"
            "os"
            "path/filepath"
            "strings"

            "github.com/example/terraform-doctor/internal/report"
        )

        // DetectStateMismatch looks for signs of mixed backends (local vs remote) across the codebase.
        // This is heuristic-based and intentionally conservative.
        func DetectStateMismatch(root string) []report.Issue {
            var issues []report.Issue
            foundLocal := false
            foundRemote := false

            _ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
                if err != nil {
                    return nil
                }
                if info.IsDir() || !strings.HasSuffix(info.Name(), ".tf") {
                    return nil
                }
                b, err := ioutil.ReadFile(path)
                if err != nil {
                    return nil
                }
                s := string(b)
                if strings.Contains(s, "backend \"local\"") || strings.Contains(s, "backend\"local\"") {
                    foundLocal = true
                }
                // detect common remote backends (S3, azurerm, gcs) as remote execution/state examples
                if strings.Contains(s, "backend \"s3\"") || strings.Contains(s, "backend \"azurerm\"") || strings.Contains(s, "backend \"gcs\"") || strings.Contains(s, "remote") {
                    foundRemote = true
                }
                return nil
            })

            if foundLocal && foundRemote {
                issues = append(issues, report.Issue{Severity: report.SeverityWarning, Message: "State mismatch detected between local and remote backends"})
            }

            return issues
        }
