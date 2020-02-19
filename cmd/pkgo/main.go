package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/browser"
	"github.com/urfave/cli/v2"
	"github.com/yuzuy/pkgo/utils"
)

const pkgGoDevUrl = "https://pkg.go.dev/"

func main() {
	pkgo()
}

func pkgo() {
	app := cli.NewApp()
	app.Name = "pkgo"
	app.Usage = "Open the document for Go package"

	app.Action = func(c *cli.Context) error {
		if 0 < c.NArg() {
			pkg := c.Args().Get(0)
			exist, err := utils.IsExistPage(pkgGoDevUrl + pkg)
			if err != nil {
				return err
			}
			if exist {
				if err := browser.OpenURL(pkgGoDevUrl + pkg); err != nil {
					return err
				}
				return nil
			}
			exist, err = utils.IsExistPage(pkgGoDevUrl + "golang.org/x/" + pkg)
			if err != nil {
				return err
			}
			if exist {
				if err := browser.OpenURL(pkgGoDevUrl + "golang.org/x/" + pkg); err != nil {
					return err
				}
				return nil
			}
			exist, err = utils.IsExistPage(pkgGoDevUrl + "github.com/" + pkg)
			if err != nil {
				return err
			}
			if exist {
				if err := browser.OpenURL(pkgGoDevUrl + "github.com/" + pkg); err != nil {
					return err
				}
				return nil
			}
			fmt.Println("Cannot find this package.")
			return nil
		}
		if err := browser.OpenURL(pkgGoDevUrl); err != nil {
			return err
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
