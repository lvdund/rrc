package ies

import "rrc/utils"

// MUSIM-GapInfo-r17-musim-GapLength-r17 ::= ENUMERATED
type MusimGapinfoR17MusimGaplengthR17 struct {
	Value utils.ENUMERATED
}

const (
	MusimGapinfoR17MusimGaplengthR17EnumeratedNothing = iota
	MusimGapinfoR17MusimGaplengthR17EnumeratedMs3
	MusimGapinfoR17MusimGaplengthR17EnumeratedMs4
	MusimGapinfoR17MusimGaplengthR17EnumeratedMs6
	MusimGapinfoR17MusimGaplengthR17EnumeratedMs10
	MusimGapinfoR17MusimGaplengthR17EnumeratedMs20
)
