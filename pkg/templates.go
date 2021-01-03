package pkg

import (
	"io/ioutil"
	"path/filepath"
	"text/template/parse"

	sprig "github.com/Masterminds/sprig/v3"
)

// TemplateFile refers to a template file
type TemplateFile struct {
	Path string
	Data []byte
}

// ReadTemplates reads template files from a directory
func ReadTemplates(templatesDirPath string) ([]TemplateFile, error) {
	fileInfos, err := ioutil.ReadDir(templatesDirPath)
	if err != nil {
		return nil, err
	}

	templateFiles := make([]TemplateFile, len(fileInfos))

	for _, fileInfo := range fileInfos {
		// TODO: Can template directories have templates within nested directories? Check!
		if fileInfo.IsDir() {
			continue
		}

		if fileInfo.Name() == "NOTES.txt" {
			continue
		}

		templatePath := filepath.Join(templatesDirPath, fileInfo.Name())

		templateFile, err := ReadTemplate(templatePath)
		if err != nil {
			return nil, err
		}
		templateFiles = append(templateFiles, templateFile)
	}

	return templateFiles, nil
}

// ReadTemplate reads a single template file
func ReadTemplate(templatePath string) (TemplateFile, error) {
	templateData, err := ioutil.ReadFile(templatePath)

	if err != nil {
		return TemplateFile{}, err
	}

	return TemplateFile{
		Path: templatePath,
		Data: templateData,
	}, nil
}

// ParseTemplateFile parses a template file into a parse tree
func ParseTemplateFile(templateFile TemplateFile) (map[string]*parse.Tree, error) {
	funcMap := getFuncMap()
	return parse.Parse(templateFile.Path, string(templateFile.Data), "", "", funcMap)
}

// ParseTemplateFiles parses template files into parse trees
func ParseTemplateFiles(templateFiles []TemplateFile) ([]map[string]*parse.Tree, error) {
	parseTrees := make([]map[string]*parse.Tree, len(templateFiles))

	for _, templateFile := range templateFiles {
		parseTree, err := ParseTemplateFile(templateFile)

		if err != nil {
			return nil, err
		}

		parseTrees = append(parseTrees, parseTree)
	}

	return parseTrees, nil
}

func getFuncMap() map[string]interface{} {
	funcMap := sprig.TxtFuncMap()

	funcMap["toToml"] = func() {}
	funcMap["toYaml"] = func() {}
	funcMap["fromYaml"] = func() {}
	funcMap["fromYamlArray"] = func() {}
	funcMap["toJson"] = func() {}
	funcMap["fromJson"] = func() {}
	funcMap["fromJsonArray"] = func() {}
	funcMap["include"] = func() {}
	funcMap["tpl"] = func() {}
	funcMap["required"] = func() {}
	funcMap["lookup"] = func() {}
	funcMap["and"] = func() {}
	funcMap["call"] = func() {}
	funcMap["html"] = func() {}
	funcMap["index"] = func() {}
	funcMap["slice"] = func() {}
	funcMap["js"] = func() {}
	funcMap["len"] = func() {}
	funcMap["not"] = func() {}
	funcMap["or"] = func() {}
	funcMap["print"] = func() {}
	funcMap["printf"] = func() {}
	funcMap["println"] = func() {}
	funcMap["urlquery"] = func() {}
	funcMap["eq"] = func() {}
	funcMap["ge"] = func() {}
	funcMap["gt"] = func() {}
	funcMap["le"] = func() {}
	funcMap["lt"] = func() {}
	funcMap["ne"] = func() {}

	return funcMap
}
