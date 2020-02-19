package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pkg/browser"
	"github.com/urfave/cli/v2"
)

const pkgGoDevUrl = "https://pkg.go.dev/"

func main() {
	app := cli.NewApp()
	app.Name = "pkgo"
	app.Usage = "Open the document for Go package"

	app.Action = func(c *cli.Context) error {
		if 0 < c.NArg() {
			pkg := c.Args().Get(0)
			exist, err := isExistPage(pkgGoDevUrl + pkg)
			if err != nil {
				return err
			}
			if exist {
				if err := browser.OpenURL(pkgGoDevUrl + pkg); err != nil {
					return err
				}
				return nil
			}
			exist, err = isExistPage(pkgGoDevUrl + "golang.org/x/" + pkg)
			if err != nil {
				return err
			}
			if exist {
				if err := browser.OpenURL(pkgGoDevUrl + "golang.org/x/" + pkg); err != nil {
					return err
				}
				return nil
			}
			exist, err = isExistPage(pkgGoDevUrl + "github.com/" + pkg)
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

func isExistPage(url string) (bool, error) {
	res, err := http.Get(url)
	if err != nil {
		return false, err
	}
	if res.StatusCode == 404 {
		return false, nil
	}
	return true, nil
}
