package tempaltes

import (
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/russross/blackfriday"
)

const Indentation = `  `

type normalizer struct {
	string
}

// LongDesc normalizes a command's long description to follow the conventions.
func LongDesc(s string) string {
	if len(s) == 0 {
		return s
	}
	return normalizer{s}.heredoc().markdown().trim().string
}

func (s normalizer) heredoc() normalizer {
	s.string = heredoc.Doc(s.string)
	return s
}

func (s normalizer) markdown() normalizer {
	bytes := []byte(s.string)
	formatted := blackfriday.Markdown(
		bytes,
		&ASCIIRenderer{Indentation: Indentation},
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS,
	)
	s.string = string(formatted)
	return s
}

func (s normalizer) trim() normalizer {
	s.string = strings.TrimSpace(s.string)
	return s
}
