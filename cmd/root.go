package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"tmplexec/tools"

	"github.com/Masterminds/sprig"
	"github.com/spf13/cobra"
)

var (
	tmplDir   tools.TemplatesDir
	dataPath  string
	outDir    string
	funcMap   template.FuncMap
	goimports bool

	rootCmd = &cobra.Command{
		Use:   "tmplexec",
		Short: "tmplexec executes go templates from terminal",
		RunE: func(cmd *cobra.Command, args []string) error {
			dataSlice, err := tools.Unmarshal(dataPath)
			if err != nil {
				return err
			}

			for _, data := range dataSlice {
				err = tmplDir.Execute(data, outDir, funcMap)
				if err != nil {
					return err
				}
			}

			if goimports {
				for _, data := range dataSlice {
					path := filepath.Join(outDir, data.Filename)
					err = exec.Command("goimports", "-w", path).Run()
					if err != nil {
						return err
					}
				}
			}

			return nil
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&tmplDir.MainFile, "template-main", "m", "", "template to be executed")
	rootCmd.Flags().StringVarP(&tmplDir.Path, "template-path", "p", ".", "path to directory with templates to be executed")
	rootCmd.Flags().StringVarP(&dataPath, "data", "d", "", "data to execute template (json, yaml)")
	rootCmd.Flags().StringVarP(&outDir, "output", "o", "out", "output dir for executed templated")
	rootCmd.Flags().BoolVar(&goimports, "goimports", false, "use to run goimports -w on generated files")

	funcMap = sprig.TxtFuncMap()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
