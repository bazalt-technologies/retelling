package api

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func paramInt(r *http.Request, key string) int {
	params := paramsToLower(r.URL.Query())
	key = strings.ToLower(key)
	var id int
	if len(params[key]) > 0 {
		id, _ = strconv.Atoi(params[key][0])
	}
	return id
}

// Переводит имена параметров в нижний регистр.
func paramsToLower(params url.Values) map[string][]string {
	values := make(map[string][]string)
	for k, v := range params {
		values[strings.ToLower(k)] = v
	}
	return values
}
