package main

import (
	"fmt"
	"os"
	"path/filepath"

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

	namedParseTrees, err := pkg.ParseTemplateFiles(templateFiles)

	if err != nil {
		fmt.Printf("error while parsing templates: %v", err)
	}

	values := pkg.GetValues(namedParseTrees)
	printValues(values)
}

func printValues(values []string) {
	for _, value := range values {
		fmt.Println(value)
	}
}
