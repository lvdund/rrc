package ies

// IDC-SubframePatternList-r11 ::= SEQUENCE OF IDC-SubframePattern-r11
// SIZE (1..maxSubframePatternIDC-r11)
type IdcSubframepatternlistR11 struct {
	Value []IdcSubframepatternR11 `lb:1,ub:maxSubframePatternIDCR11`
}
