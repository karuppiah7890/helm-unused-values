package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template/parse"

	"github.com/karuppiah7890/helm-unused-values/pkg"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("pass exactly one argument")
		fmt.Println("usage: helm-unused-values <helm-chart-path>")
		os.Exit(1)
		return
	}

	helmChartPath := os.Args[1]

	templatesDirPath := filepath.Join(helmChartPath, "templates")

	templateFiles, err := pkg.ReadTemplates(templatesDirPath)

	if err != nil {
		fmt.Printf("error while reading templates: %v", err)
	}

	parseTrees, err := pkg.ParseTemplateFiles(templateFiles)

	if err != nil {
		fmt.Printf("error while parsing templates: %v", err)
	}

	printParseTrees(parseTrees)
}

func printParseTrees(parseTrees []map[string]*parse.Tree) {
	for _, namedParseTree := range parseTrees {
		for name, parseTree := range namedParseTree {
			fmt.Printf("Name: %v\n", name)
			fmt.Printf("Parse Tree: %v\n", parseTree)
			fmt.Println()
		}
	}
}
