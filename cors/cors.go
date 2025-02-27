package cors

import (
	"net"
	"net/http"
)

func IPFilter(next http.HandlerFunc, pattern map[string]bool) http.HandlerFunc {
	host := map[string]bool{
		"localhost": true,
		"127.0.0.1": true,
		"::1":       true,
	}

	return func(res http.ResponseWriter, req *http.Request) {

		if !pattern[req.URL.Path] {
			http.Error(res, "Pattern not found", http.StatusNotFound)
			return
		}

		ip, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			http.Error(res, "Forbidder", http.StatusForbidden)
			return
		}

		if !host[ip] {
			http.Error(res, "Forbidden", http.StatusForbidden)
			return
		}

		next(res, req)
	}

}
