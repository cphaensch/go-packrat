/*
	(c) 2019 Launix, Inh. Carl-Philip Hänsch
	Author: Tim Kluge

	Dual licensed with custom aggreements or GPLv3
*/

package packrat

type EndParser struct {
	skipWs bool
}

func NewEndParser(skipWs bool) *EndParser {
	return &EndParser{skipWs: skipWs}
}

// Match accepts only the end of the scanner's input and will not match any input.
func (p *EndParser) Match(s *Scanner) (*Scanner, Node) {
	if p.skipWs {
		s.Skip()
	}

	if len(s.remainingInput) == 0 {
		return s, Node{}
	}

	return nil, Node{}
}
