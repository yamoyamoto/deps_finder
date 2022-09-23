package cmd

import (
	"depsfinder/lib"
	"depsfinder/lib/finders"
	"depsfinder/lib/services"
	"github.com/spf13/cobra"
	"log"
)

var dirPath string

var findDepsCmd = &cobra.Command{
	Use: "find-deps",
	Run: func(cmd *cobra.Command, args []string) {
		err := lib.FindDeps(finders.NewJavaDepsFinder(finders.NewJavaParser(), services.NewDirWalker()), dirPath)
		if err != nil {
			log.Fatalf("failed to find deps. err: %s", err)
		}
	},
}
