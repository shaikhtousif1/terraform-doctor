package cmd

import (
    "fmt"
    "path/filepath"

    "github.com/spf13/cobra"

    "github.com/example/terraform-doctor/internal/checks"
    "github.com/example/terraform-doctor/internal/graph"
    "github.com/example/terraform-doctor/internal/parser"
    "github.com/example/terraform-doctor/internal/report"
)

var scanCmd = &cobra.Command{
    Use:   "scan [path]",
    Short: "Scan a Terraform codebase for architectural issues",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        path := args[0]
        abs, err := filepath.Abs(path)
        if err == nil {
            path = abs
        }

        fmt.Printf("🩺 Terraform Doctor scanning: %s\n", path)

        modules, _ := parser.ParsePath(path)
        g := graph.BuildGraph(modules)

        var issues []report.Issue
        issues = append(issues, checks.DetectCircularDependencies(g)...)
        issues = append(issues, checks.DetectCoupling(g)...)
        issues = append(issues, checks.DetectStateMismatch(path)...)
        issues = append(issues, checks.DetectEnvironmentTagMismatch(path)...)

        report.PrintIssues(issues)
    },
}

func init() {
    rootCmd.AddCommand(scanCmd)
}
