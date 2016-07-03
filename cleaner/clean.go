package cleaner

import "github.com/BenLubar/htmlcleaner"

func Clean(content string) string {
	if s, err := htmlcleaner.Preprocess(Config, content); err == nil {
		content = s
	}
	return htmlcleaner.Clean(Config, content)
}
