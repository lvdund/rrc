package ies

// STAG-ToAddModList-r11 ::= SEQUENCE OF STAG-ToAddMod-r11
// SIZE (1..maxSTAG-r11)
type StagToaddmodlistR11 struct {
	Value []StagToaddmodR11 `lb:1,ub:maxSTAGR11`
}
