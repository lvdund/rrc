package ies

// MUSIM-GapPreferenceList-r17 ::= SEQUENCE OF MUSIM-GapInfo-r17
// SIZE (1..4)
type MusimGappreferencelistR17 struct {
	Value []MusimGapinfoR17 `lb:1,ub:4`
}
