package cors

import (
	"net"
	"net/http"
	"simple-api/router"
	"strings"
)

func IPFilter(next http.HandlerFunc, pattern map[string]bool) http.HandlerFunc {
	host := map[string]bool{
		"localhost": true,
		"127.0.0.1": true,
		"::1":       true,
	}

	return func(res http.ResponseWriter, req *http.Request) {

		ip, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			router.ErrorJson(res, "Forbidder", http.StatusForbidden)
			return
		}

		if !host[ip] {
			router.ErrorJson(res, "Forbidden", http.StatusForbidden)
			return
		}

		// First check for exact path match
		if pattern[req.URL.Path] {
			next(res, req)
			return
		}

		// Then check for prefix matches with trailing slash
		for patternPath := range pattern {
			if patternPath == "/" {
				continue
			}

			if strings.HasPrefix(req.URL.Path, patternPath) && strings.HasSuffix(patternPath, "/") {
				next(res, req)
				return
			}
		}

		router.ErrorJson(res, "Path not allowed", http.StatusForbidden)

	}

}
