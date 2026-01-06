package checks

import (
    "fmt"

    "github.com/example/terraform-doctor/internal/graph"
    "github.com/example/terraform-doctor/internal/report"
)

// DetectCircularDependencies finds simple two-node cycles as a placeholder.
func DetectCircularDependencies(g *graph.Graph) []report.Issue {
    var issues []report.Issue
    for a, neighbors := range g.Edges {
        for _, b := range neighbors {
            // if there's an edge b->a then report
            for _, bNeighbors := range g.Edges[b] {
                if bNeighbors == a {
                    issues = append(issues, report.Issue{Severity: report.SeverityWarning, Message: fmt.Sprintf("Circular dependency detected between modules %q and %q", a, b)})
                }
            }
        }
    }
    return issues
}
