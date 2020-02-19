package utils

import "net/http"

func IsExistPage(url string) (bool, error) {
	res, err := http.Get(url)
	if err != nil {
		return false, err
	}
	if res.StatusCode == 404 {
		return false, nil
	}
	return true, nil
}
