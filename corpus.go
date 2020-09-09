package corpustool

import "strings"

func Clear(content string) string {
	for strings.Contains(content, "\n\n") {
		content = strings.ReplaceAll(content, "\n\n", "\n")
	}
	return content
}
