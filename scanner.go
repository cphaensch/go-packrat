package packrat

import (
	"regexp"
	"strings"
)

type scannerNode struct {
	Scanner *Scanner
	Node    Node
}

type Scanner struct {
	input          string
	remainingInput string
	position       int
	line           int
	memoization    []map[Parser]scannerNode

	skipRegex *regexp.Regexp
}

func (s *Scanner) Copy() *Scanner {
	ns := *s
	return &ns
}

var skipWhitespaceRegex = regexp.MustCompile("^[\r\n\t ]+")

func NewScanner(input string, skipWhitespace bool) *Scanner {
	s := &Scanner{input: input, position: 0, line: 1, memoization: make([]map[Parser]scannerNode, len(input))}
	for i := range s.input {
		s.memoization[i] = make(map[Parser]scannerNode)
	}
	s.remainingInput = s.input
	if skipWhitespace {
		s.skipRegex = skipWhitespaceRegex
	}
	return s
}

func (s *Scanner) updatePosition(reads string) {
	l := len(reads)
	if l > 0 {
		s.remainingInput = s.remainingInput[l:]
		s.position += l
		s.line += strings.Count(reads, "\n")
	}
}

func (s *Scanner) MatchRegexp(r *regexp.Regexp) *string {
	matched := r.FindStringSubmatch(s.remainingInput)
	if matched != nil {
		s.updatePosition(matched[0])
		return &matched[0]
	}

	return nil
}

func (s *Scanner) MatchString(str string) *string {
	if strings.HasPrefix(s.remainingInput, str) {
		s.updatePosition(str)
		return &str
	}

	return nil
}

func (s *Scanner) Skip() {
	s.MatchRegexp(s.skipRegex)
}
