package pkgo

import "strings"

const githubURL = "https://github.com/"

func TidyUpURL(pkg string, kind string) string {
	path := strings.Split(pkg, "/")

	if kind == "sub" {
		if 1 < len(path) {
			var pathStr string
			if kind == "sub" {
				for _, v := range path[1:] {
					pathStr += "/" + v
				}
				return githubURL + "golang/" + path[0] + "/tree/master" + pathStr
			}
		}

		return githubURL + "golang/" + pkg
	}

	if 1 < len(path) {
		var pathStr string
		for _, v := range path[2:] {
			pathStr += "/" + v
		}
		return githubURL + path[0] + "/" + path[1] + "/tree/master" + pathStr
	}

	return githubURL + pkg
}
