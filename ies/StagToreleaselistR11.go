package ies

// STAG-ToReleaseList-r11 ::= SEQUENCE OF STAG-Id-r11
// SIZE (1..maxSTAG-r11)
type StagToreleaselistR11 struct {
	Value []StagIdR11 `lb:1,ub:maxSTAGR11`
}
