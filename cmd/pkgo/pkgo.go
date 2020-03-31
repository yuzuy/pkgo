package main

import (
	"fmt"
	"os"

	"github.com/pkg/browser"
	"github.com/urfave/cli/v2"
	"github.com/yuzuy/pkgo"
)

const docURL = "https://pkg.go.dev/"

func main() {
	if err := run(); err != nil {
		fmt.Println("pkgo: " + err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	app := cli.NewApp()
	app.Name = "pkgo"
	app.Usage = "pkgo [package name]"

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "Open the GitHub repository",
		},
		&cli.BoolFlag{
			Name:    "official",
			Aliases: []string{"o"},
			Usage:   "Open the document in golang.org",
		},
	}

	app.Action = func(c *cli.Context) error {
		if 0 < c.NArg() {
			pkg := c.Args().Get(0)

			exist, err := pkgo.IsExistPage(docURL + pkg)
			if err != nil {
				return err
			}
			if exist {
				url := docURL + pkg
				if c.IsSet("repo") {
					url = "https://github.com/golang/go/tree/master/src/" + pkg
				}
				if c.IsSet("official") {
					url = "https://golang.org/pkg/" + pkg
				}
				if err := browser.OpenURL(url); err != nil {
					return err
				}
				return nil
			}
			if c.IsSet("official") {
				fmt.Println("Cannot find this package")
				return nil
			}

			exist, err = pkgo.IsExistPage(docURL + "golang.org/x/" + pkg)
			if err != nil {
				return err
			}
			if exist {
				url := docURL + "golang.org/x/" + pkg
				if c.IsSet("repo") {
					url = pkgo.TidyUpURL(pkg, "sub")
				}
				if err := browser.OpenURL(url); err != nil {
					return err
				}
				return nil
			}

			exist, err = pkgo.IsExistPage(docURL + "github.com/" + pkg)
			if err != nil {
				return err
			}
			if exist {
				url := docURL + "github.com/" + pkg
				if c.IsSet("repo") {
					url = pkgo.TidyUpURL(pkg, "external")
				}
				if err := browser.OpenURL(url); err != nil {
					return err
				}
				return nil
			}
			fmt.Println("Cannot find this package.")
			return nil
		}
		if err := browser.OpenURL(docURL); err != nil {
			return err
		}

		return nil
	}

	return app.Run(os.Args)
}
