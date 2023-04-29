package sonolus

import (
	"strings"
)

func JoinUrl(base string, path string) (string, error) {
	if strings.HasPrefix(path, "http") {
		return path, nil
	}
	url := base + "/" + path
	url = strings.ReplaceAll(url, "//", "/")
  url = strings.Replace(url, ":/", "://", 1)
	return url, nil
}
