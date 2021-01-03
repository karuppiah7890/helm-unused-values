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

	printSomeNodesInParseTrees(parseTrees)
}

func printSomeNodesInParseTrees(parseTrees []map[string]*parse.Tree) {
	for _, namedParseTree := range parseTrees {
		for _, parseTree := range namedParseTree {
			for _, node := range parseTree.Root.Nodes {
				if node.Type() == parse.NodeAction {
					pipe := node.(*parse.ActionNode).Pipe
					for _, cmd := range pipe.Cmds {
						for _, argNode := range cmd.Args {
							if argNode.Type() == parse.NodeField {
								identifiers := argNode.(*parse.FieldNode).Ident
								if len(identifiers) > 0 && identifiers[0] == "Values" {
									fmt.Println(argNode)
								}
							}
							// TODO: Handle parse.NodeChain node type of arg
							// node.
						}
					}
				}
			}
		}
	}
}
