package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("pass exactly one argument")
		fmt.Println("usage: helm-unused-values <helm-chart-path>")
		os.Exit(1)
		return
	}

	helmChartPath := os.Args[1]

	fmt.Println(helmChartPath)
}
