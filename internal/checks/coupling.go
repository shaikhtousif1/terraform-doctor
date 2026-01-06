package checks

import (
    "fmt"

    "github.com/example/terraform-doctor/internal/graph"
    "github.com/example/terraform-doctor/internal/report"
)

// DetectCoupling returns warnings for modules that depend on many other modules.
// This is a heuristic placeholder — real coupling analysis requires deep inspection.
func DetectCoupling(g *graph.Graph) []report.Issue {
    var issues []report.Issue
    for from, tos := range g.Edges {
        if len(tos) > 3 {
            issues = append(issues, report.Issue{Severity: report.SeverityWarning, Message: fmt.Sprintf("Module %q has high outgoing coupling (%d dependencies)", from, len(tos))})
        }
    }
    return issues
}
