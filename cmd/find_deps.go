package cmd

import (
	"depsfinder/lib"
	"depsfinder/lib/finders"
	"depsfinder/lib/services"
	"encoding/json"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
)

var dirPath string
var outputDirDirPath string
var language string

var allFinders = map[string]lib.DepsFinder{
	"java": finders.NewJavaDepsFinder(finders.NewJavaParser(), services.NewDirWalker()),
}

var findDepsCmd = &cobra.Command{
	Use: "find-deps",
	Run: func(cmd *cobra.Command, args []string) {
		finder, ok := allFinders[language]
		if !ok {
			log.Fatalf("language '%s' was not found", language)
		}

		dependencies, err := lib.FindDeps(finder, dirPath)
		if err != nil {
			log.Fatalf("failed to find deps. err: %s", err)
		}

		b, err := json.Marshal(dependencies)
		if err != nil {
			log.Fatalf("failed to find deps. err: %s", err)
		}

		f, err := os.Create(path.Join(outputDirDirPath, language+".json"))
		defer f.Close()

		if err != nil {
			log.Fatalf("failed to find deps. err: %s", err)
		}

		if _, err := f.Write(b); err != nil {
			log.Fatalf("failed to find deps. err: %s", err)
		}

		log.Printf("complete to save dependencies to '%s'", outputDirDirPath)
	},
}
