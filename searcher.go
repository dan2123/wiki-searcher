package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type WikiResponse struct {
	QueryTerm          string
	ResultTitles       []string
	ResultDescriptions []string
	ResultLinks        []string
}

func SearchWiki(searchTerm string, limit int) (WikiResponse, error) {
	var wikiRes WikiResponse
	req, err := http.NewRequest("GET", "https://en.wikipedia.org/w/api.php", nil)
	if err != nil {
		return WikiResponse{}, err
	}

	q := req.URL.Query()
	q.Add("action", "opensearch")
	q.Add("search", searchTerm)
	q.Add("limit", strconv.Itoa(limit))
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return WikiResponse{}, err
	}
	defer res.Body.Close()

	var decodedRes []interface{}
	err = json.NewDecoder(res.Body).Decode(&decodedRes)
	if err != nil {
		return WikiResponse{}, err
	}

	wikiRes = WikiResponse{
		QueryTerm:          decodedRes[0].(string),
		ResultTitles:       convertInterfaceSliceToStringSlice(decodedRes[1].([]interface{})),
		ResultDescriptions: convertInterfaceSliceToStringSlice(decodedRes[2].([]interface{})),
		ResultLinks:        convertInterfaceSliceToStringSlice(decodedRes[3].([]interface{})),
	}
	return wikiRes, nil
}

func convertInterfaceSliceToStringSlice(s []interface{}) []string {
	out := make([]string, 0, len(s))
	for _, val := range s {
		out = append(out, val.(string))
	}
	return out
}
