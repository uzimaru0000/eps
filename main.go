package main

import (
	"fmt"
	"log"
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
	var datas []byte
	var err error
	path := fmt.Sprintf("%s/0.19.0/search.json", os.Getenv("ELM_HOME"))
	log.Printf(path)

	if packages.PackagesFileExist(path) {
		datas, err = packages.ReadPackagesFile(path)
		if err != nil {
			return err
		}
	} else {
		datas, err = packages.FetchPackagesFile()
		if err != nil {
			return err
		}
		if err = packages.SavePackagesFile(path, datas); err != nil {
			return err
		}
	}

	packages, err := packages.ConverteJSON(datas)
	if err != nil {
		return err
	}

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
				"- Summary -\n%s\n\nLicense: %s",
				packages[i].Summary,
				packages[i].License,
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
