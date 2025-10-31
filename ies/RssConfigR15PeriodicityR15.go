package ies

import "rrc/utils"

// RSS-Config-r15-periodicity-r15 ::= ENUMERATED
type RssConfigR15PeriodicityR15 struct {
	Value utils.ENUMERATED
}

const (
	RssConfigR15PeriodicityR15EnumeratedNothing = iota
	RssConfigR15PeriodicityR15EnumeratedMs160
	RssConfigR15PeriodicityR15EnumeratedMs320
	RssConfigR15PeriodicityR15EnumeratedMs640
	RssConfigR15PeriodicityR15EnumeratedMs1280
)
