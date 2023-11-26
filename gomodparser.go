package main

import (
	"regexp"
)

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

var goVersionRe = regexp.MustCompile(`go\s(1\.\d+\.?\d+)`)
var moduleName = regexp.MustCompile(`module\s(.+)`)
var directDependency = regexp.MustCompile(`require\s(.+)\s(.+)`)

func parseGoModFile(file []byte) GoModFile {
	gversion := goVersionRe.FindSubmatch(file)[1]
	pkgName := moduleName.FindSubmatch(file)[1]

	return GoModFile{
		packagename: string(pkgName),
		goversion:   string(gversion),
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
