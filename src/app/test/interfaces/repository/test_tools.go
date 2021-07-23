package repository

import (
	"io/ioutil"
	"strings"
)

func ReadSqlFile(filepath string) ([]string, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	s := string(b)
	queries := strings.Split(s, "\n")
	return queries, nil
}
