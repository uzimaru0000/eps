package main

import (
	"fmt"
	"os"

	pipeline "github.com/mattn/go-pipeline"
	"github.com/uzimaru0000/eps/packages"

	fuzzy "github.com/ktr0731/go-fuzzyfinder"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "EPS [Elm Package Supporter]"
	app.Usage = "Install supported for Elm packages"
	app.Version = "0.0.1"

	app.Action = action
	app.Run(os.Args)
}

func action(context *cli.Context) error {
	packages := packages.FetchPackageDatas()
	idx, err := fuzzy.FindMulti(
		packages,
		func(i int) string {
			return packages[i].Name
		},
		fuzzy.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf(
				"License: %s\nSummary: %s",
				packages[i].License,
				packages[i].Summary,
			)
		}))

	if err != nil {
		return err
	}

	out, err := pipeline.Output(
		[]string{"echo", "y"},
		[]string{"elm", "install", packages[idx[0]].Name},
	)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", string(out))

	return nil
}
