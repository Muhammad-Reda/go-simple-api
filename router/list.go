package router

var ListPattern = map[string]bool{
	"/":                    true,
	"/users":               true,
	"/user/":               true,
	"/user/add":            true,
	"/user/update/":        true,
	"/user/delete/":        true,
	"/products":            true,
	"/product/":            true,
	"/product/add":         true,
	"/product/update/":     true,
	"/product/delete/":     true,
	"/transactions":        true,
	"/transaction/":        true,
	"/transaction/add":     true,
}
