package utils

import (
	"crypto/rand"
	"io"
	"strconv"
	"strings"
)

var (
	//table for code generator
	table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
)

//QueryParams ...
type QueryParams struct {
	Filters  map[string]string
	Page     int64
	Limit    int64
	Ordering []string
	Search   string
	Author   string
	Category []string
}

//ParseQueryParams ...
func ParseQueryParams(queryParams map[string][]string) (*QueryParams, []string) {
	params := QueryParams{
		Filters:  make(map[string]string),
		Page:     1,
		Limit:    10,
		Ordering: []string{},
		Search:   "",
	}
	var errStr []string
	var err error

	for key, value := range queryParams {
		if key == "page" {
			params.Page, err = strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				errStr = append(errStr, "Invalid `page` param")
			}
			continue
		}

		if key == "limit" {
			params.Limit, err = strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				errStr = append(errStr, "Invalid `limit` param")
			}
			continue
		}

		if key == "search" {
			params.Search = value[0]
			continue
		}

		if key == "ordering" {
			params.Ordering = strings.Split(value[0], ",")
			continue
		}

		if key == "author" {
			params.Author = value[0]
		}

		if key == "category" {
			params.Category = value
		}
		params.Filters[key] = value[0]
	}

	return &params, errStr
}

// GenerateCode is function generating n-digit random code
func GenerateCode(max int, isDev ...bool) string {
	if len(isDev) == 1 {
		return "7777"
	}

	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
