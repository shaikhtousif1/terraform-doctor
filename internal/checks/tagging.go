package checks

import (
    "io/ioutil"
    "path/filepath"
    "strings"

    "github.com/example/terraform-doctor/internal/report"
)

// DetectEnvironmentTagMismatch looks for obvious copy-paste tag drift like dev tags in prod modules.
// Placeholder: scans for `tags = {` blocks containing `env`/`environment` keys with dev/prod values.
func DetectEnvironmentTagMismatch(root string) []report.Issue {
    var issues []report.Issue

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
        // crude checks
        if strings.Contains(s, "environment = \"dev\"") || strings.Contains(s, "env = \"dev\"") {
            // if path contains prod, flag it
            if strings.Contains(strings.ToLower(path), "prod") {
                issues = append(issues, report.Issue{Severity: report.SeverityWarning, Message: "Environment tag mismatch detected (dev tag in prod)"})
            }
        }
        return nil
    })

    return issues
}
