package ies

import "rrc/utils"

// RSS-Config-r15-powerBoost-r15 ::= ENUMERATED
type RssConfigR15PowerboostR15 struct {
	Value utils.ENUMERATED
}

const (
	RssConfigR15PowerboostR15EnumeratedNothing = iota
	RssConfigR15PowerboostR15EnumeratedDb0
	RssConfigR15PowerboostR15EnumeratedDb3
	RssConfigR15PowerboostR15EnumeratedDb4dot8
	RssConfigR15PowerboostR15EnumeratedDb6
)
