/*
	(c) 2019 Launix, Inh. Carl-Philip Hänsch
	Author: Tim Kluge

	Dual licensed with custom aggreements or GPLv3
*/

package packrat

import "regexp"

type AtomParser struct {
	r    *regexp.Regexp
	skipWs bool
}

func NewAtomParser(str string, caseInsensitive bool, skipWs bool) *AtomParser {
	prefix := ""
	if caseInsensitive {
		prefix += "(?i)"
	}
	prefix += "^"
	r := regexp.MustCompile(prefix + regexp.QuoteMeta(str))
	p := &AtomParser{skipWs: skipWs, r: r}	
	return p
}

// Match matches only the given string. If skipWs is set to true, leading whitespace according to the scanner's skip regexp is skipped, but not matched by the parser.
func (p *AtomParser) Match(s *Scanner) (*Scanner, Node) {
	if p.skipWs {
		s.Skip()

		if !s.isAtBreak() {
			return nil, Node{}
		}
	}

	matched := s.MatchRegexp(p.r)
	if matched == nil {
		return nil, Node{}
	}

	if p.skipWs {
		if !s.isAtBreak() {
			return nil, Node{}
		}
	}

	r := scannerNode{Scanner: s, Node: Node{Matched: *matched, Parser: p}}
	return r.Scanner, r.Node
}
