package ies

import "rrc/utils"

// RSS-Config-r15-duration-r15 ::= ENUMERATED
type RssConfigR15DurationR15 struct {
	Value utils.ENUMERATED
}

const (
	RssConfigR15DurationR15EnumeratedNothing = iota
	RssConfigR15DurationR15EnumeratedSf8
	RssConfigR15DurationR15EnumeratedSf16
	RssConfigR15DurationR15EnumeratedSf32
	RssConfigR15DurationR15EnumeratedSf40
)
