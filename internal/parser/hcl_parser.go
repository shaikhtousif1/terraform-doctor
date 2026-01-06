package parser

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

type Module struct {
    Name  string
    Path  string
    Files []string
}

// ParsePath walks the provided path and returns folders that contain .tf files as modules.
// This is intentionally lightweight and designed to be replaced by a proper HCL parser.
func ParsePath(root string) ([]Module, error) {
    var modules []Module

    // Walk filesystem and collect .tf files per directory
    _ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return nil
        }
        if info.IsDir() {
            return nil
        }
        if strings.HasSuffix(info.Name(), ".tf") {
            dir := filepath.Dir(path)
            // find existing module entry
            found := -1
            for i, m := range modules {
                if m.Path == dir {
                    found = i
                    break
                }
            }
            if found == -1 {
                modules = append(modules, Module{Name: filepath.Base(dir), Path: dir, Files: []string{path}})
            } else {
                modules[found].Files = append(modules[found].Files, path)
            }
        }
        return nil
    })

    // If no modules found but root is a directory, try to treat root as a module if it has .tf files
    if len(modules) == 0 {
        files, _ := ioutil.ReadDir(root)
        for _, f := range files {
            if !f.IsDir() && strings.HasSuffix(f.Name(), ".tf") {
                modules = append(modules, Module{Name: filepath.Base(root), Path: root, Files: []string{}})
                break
            }
        }
    }

    return modules, nil
}
