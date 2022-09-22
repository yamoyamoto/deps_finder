package cmd

import (
	"depsfinder/lib"
	"github.com/spf13/cobra"
	"log"
)

type TmpFinder struct{}

func (finder TmpFinder) Find() (*lib.Dependencies, error) {
	return nil, nil
}

var findDepsCmd = &cobra.Command{
	Use: "find-deps",
	Run: func(cmd *cobra.Command, args []string) {
		tmpFinder := TmpFinder{}

		err := lib.FindDeps(tmpFinder)
		if err != nil {
			log.Fatalf("failed to find deps. err: %s", err)
		}
	},
}
