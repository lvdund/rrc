package ies

// MUSIM-GapConfig-r17 ::= SEQUENCE
// Extensible
type MusimGapconfigR17 struct {
	MusimGaptoreleaselistR17 *[]MusimGapidR17 `lb:1,ub:3`
	MusimGaptoaddmodlistR17  *[]MusimGapR17   `lb:1,ub:3`
	MusimAperiodicgapR17     *MusimGapinfoR17
}
