/*
	(c) 2019 Launix, Inh. Carl-Philip Hänsch
	Author: Tim Kluge

	Dual licensed with custom aggreements or GPLv3
*/

package packrat

import (
	"regexp"
	"strings"
)

type AtomParser struct {
	r      *regexp.Regexp
	skipWs bool
	atom   string
}

func NewAtomParser(str string, caseInsensitive bool, skipWs bool) *AtomParser {
	prefix := ""
	if caseInsensitive {
		prefix += "(?i)"
	}
	prefix += "^"
	r := regexp.MustCompile(prefix + regexp.QuoteMeta(str))
	p := &AtomParser{skipWs: skipWs, r: r, atom: str}
	return p
}

func (p *AtomParser) Description(stack map[Parser]bool) string {
	b := strings.Builder{}
	b.WriteString("Atom(")
	b.WriteString(p.atom)
	b.WriteString(")")
	return b.String()
}

// Match matches only the given string. If skipWs is set to true, leading whitespace according to the scanner's skip regexp is skipped, but not matched by the parser.
func (p *AtomParser) Match(s *Scanner) *Node {
	startPosition := s.position

	if p.skipWs {
		s.Skip()

		if !s.isAtBreak() {
			s.setPosition(startPosition)
			return nil
		}
	}

	matched := s.MatchRegexp(p.r)
	if matched == nil {
		s.setPosition(startPosition)
		return nil
	}

	if p.skipWs {
		if !s.isAtBreak() {
			s.setPosition(startPosition)
			return nil
		}
	}

	return &Node{Matched: *matched, Parser: p}
}
