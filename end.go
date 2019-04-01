package packrat

type EndParser struct {
	skipWs bool
}

func NewEndParser(skipWs bool) *EndParser {
	return &EndParser{skipWs: skipWs}
}

// Match accepts only the end of the scanner's input and will not match any input.
func (p *EndParser) Match(os *Scanner) (*Scanner, Node) {
	s := os
	if p.skipWs {
		s = s.Copy()
		s.Skip()
	}

	if len(s.remainingInput) == 0 {
		return s, Node{}
	}

	return nil, Node{}
}
