package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/browser"
	"github.com/urfave/cli/v2"
	"github.com/yuzuy/pkgo/utils"
)

const (
	pkgGoDevURL = "https://pkg.go.dev/"
	githubURL   = "https://github.com/"
)

func main() {
	pkgo()
}

func pkgo() {
	app := cli.NewApp()
	app.Name = "pkgo"
	app.Usage = "Open the document for Go package"

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "Open the GitHub repository",
		},
	}

	app.Action = func(c *cli.Context) error {
		if 0 < c.NArg() {
			pkg := c.Args().Get(0)
			exist, err := utils.IsExistPage(pkgGoDevURL + pkg)
			if err != nil {
				return err
			}
			if exist {
				url := pkgGoDevURL + pkg
				if c.IsSet("repo") {
					url = githubURL + "golang/go/tree/master/src/" + pkg
				}
				if err := browser.OpenURL(url); err != nil {
					return err
				}
				return nil
			}
			exist, err = utils.IsExistPage(pkgGoDevURL + "golang.org/x/" + pkg)
			if err != nil {
				return err
			}
			if exist {
				url := pkgGoDevURL + "golang.org/x/" + pkg
				if c.IsSet("repo") {
					path := strings.Split(pkg, "/")
					if 1 < len(path) {
						var pathStr string
						for _, v := range path[1:] {
							pathStr += "/" + v
						}
						url = fmt.Sprintf("%sgolang/%s/tree/master%s", githubURL, path[0], pathStr)
					} else {
						url = githubURL + "golang/" + pkg
					}
				}
				if err := browser.OpenURL(url); err != nil {
					return err
				}
				return nil
			}
			exist, err = utils.IsExistPage(pkgGoDevURL + "github.com/" + pkg)
			if err != nil {
				return err
			}
			if exist {
				url := pkgGoDevURL + "github.com/" + pkg
				if c.IsSet("repo") {
					url = githubURL + pkg
				}
				if err := browser.OpenURL(url); err != nil {
					return err
				}
				return nil
			}
			fmt.Println("Cannot find this package.")
			return nil
		}
		if err := browser.OpenURL(pkgGoDevURL); err != nil {
			return err
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
