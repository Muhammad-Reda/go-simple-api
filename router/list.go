package router

var ListPattern = map[string]bool{
	"/":                    true,
	"/users":               true,
	"/user/add/":           true,
	"/user/update/":        true,
	"/user/delete/":        true,
	"/products":            true,
	"/product/add/":        true,
	"/product/update/":     true,
	"/product/delete/":     true,
	"/transactions":        true,
	"/transaction/add/":    true,
	"/transaction/update/": true,
	"/transaction/delete/": true,
}
