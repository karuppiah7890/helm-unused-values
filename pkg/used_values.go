package pkg

import "text/template/parse"

// GetValues from the parse trees of templates
func GetValues(namedParseTrees []map[string]*parse.Tree) []string {
	values := make([]string, 0)

	for _, namedParseTree := range namedParseTrees {
		for _, parseTree := range namedParseTree {
			for _, node := range parseTree.Root.Nodes {
				if node.Type() == parse.NodeAction {
					pipe := node.(*parse.ActionNode).Pipe
					for _, cmd := range pipe.Cmds {
						for _, argNode := range cmd.Args {
							if argNode.Type() == parse.NodeField {
								identifiers := argNode.(*parse.FieldNode).Ident
								if len(identifiers) > 0 && identifiers[0] == "Values" {
									values = append(values, argNode.String())
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

	return values
}
