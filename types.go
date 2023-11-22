package main

import "fmt"

type Package struct {
	name    string
	version string
}

type GoModFile struct {
	packagename          string
	goversion            string
	directDependencies   []Package
	indirectDependencies []Package
}

func parseGoModFile(file []byte) GoModFile {
	fmt.Print(string(file))

	return GoModFile{
		packagename: "test",
		goversion:   "1.20",
		directDependencies: []Package{
			{
				name:    "pkg",
				version: "1.0",
			},
		},
		indirectDependencies: []Package{
			{
				name:    "pkg2",
				version: "2.0",
			},
		},
	}

}
