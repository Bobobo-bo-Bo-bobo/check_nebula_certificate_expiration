package main

import (
	"io/ioutil"
)

func readCertificateFile(f string) ([]byte, error) {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}
