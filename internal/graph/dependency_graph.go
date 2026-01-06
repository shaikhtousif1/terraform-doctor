package graph

import "github.com/example/terraform-doctor/internal/parser"

type Graph struct {
    Nodes map[string]parser.Module
    Edges map[string][]string // from -> to
}

// BuildGraph builds a very small dependency graph from modules.
// This is a placeholder; future versions should analyze module source for real dependencies.
func BuildGraph(mods []parser.Module) *Graph {
    g := &Graph{
        Nodes: make(map[string]parser.Module),
        Edges: make(map[string][]string),
    }
    for _, m := range mods {
        g.Nodes[m.Name] = m
    }

    // Placeholder: simulate a circular dependency if both "eks" and "rds" exist.
    if _, eks := g.Nodes["eks"]; eks {
        if _, rds := g.Nodes["rds"]; rds {
            g.Edges["eks"] = append(g.Edges["eks"], "rds")
            g.Edges["rds"] = append(g.Edges["rds"], "eks")
        }
    }

    return g
}
