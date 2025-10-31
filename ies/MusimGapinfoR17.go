package ies

// MUSIM-GapInfo-r17 ::= SEQUENCE
// Extensible
type MusimGapinfoR17 struct {
	MusimStartingSfnAndsubframeR17 *MusimStartingSfnAndsubframeR17
	MusimGaplengthR17              *MusimGapinfoR17MusimGaplengthR17
	MusimGaprepetitionandoffsetR17 *MusimGapinfoR17MusimGaprepetitionandoffsetR17
}
