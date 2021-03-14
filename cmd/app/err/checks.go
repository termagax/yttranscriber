package err

import "net/url"

// LinkIsValid ...
func LinkIsValid(uniresloc string) bool {
	_, err := url.ParseRequestURI(uniresloc)
	if err != nil {
		return false
	}

	u, err := url.Parse(uniresloc)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
