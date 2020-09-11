package corpustool

import (
	"github.com/kevin-zx/corpustool/clear"
	"strings"
)

func ClearDoubleBreak(content string) string {
	for strings.Contains(content, "\n\n") {
		content = strings.ReplaceAll(content, "\n\n", "\n")
	}
	for strings.Contains(content,"  "){
		content = strings.ReplaceAll(content, "  ", " ")
	}
	return content
}

func ClearDuplicatePart(contents []string, keepMinLen int) []string {
	smPart := clear.DiscriminateSiteTextSamePart(contents, 10, 0.3)
	var rContents []string
	tmpContent := ""
	for i := range contents {
		tmpContent = contents[i]
		for _, part := range smPart {
			tmpContent = strings.ReplaceAll(tmpContent, part.Text, "")
		}
		if len(strings.Split(tmpContent, "")) < keepMinLen {
			continue
		}
		rContents = append(rContents, tmpContent)
	}
	return rContents
}

func Clear(contents []string, minContentLen int) []string {
	for i := range contents {
		contents[i] = ClearDoubleBreak(contents[i])
	}
	if len(contents) > 1 {
		contents = ClearDuplicatePart(contents, minContentLen)
	}

	return contents
}
