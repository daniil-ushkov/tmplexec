package tools

import (
	"os"
	"path/filepath"
	"text/template"
)

type TemplatesDir struct {
	MainFile string
	Path     string
}

func (d TemplatesDir) Execute(data *TemplateData, out string, funcMap template.FuncMap) error {
	err := os.MkdirAll(out, 0775)
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(out, data.Filename))
	if err != nil {
		return err
	}

	files, err := templateFilesInDir(d.Path)
	if err != nil {
		return err
	}

	tmpl, err := template.New(d.MainFile).Funcs(funcMap).ParseFiles(files...)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, data.Data)
	if err != nil {
		return err
	}

	return file.Close()
}

func templateFilesInDir(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".tmpl" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
