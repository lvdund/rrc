package ies

// UCI-OnPUSCH-betaOffsets ::= CHOICE
const (
	UciOnpuschBetaoffsetsChoiceNothing = iota
	UciOnpuschBetaoffsetsChoiceDynamic
	UciOnpuschBetaoffsetsChoiceSemistatic
)

type UciOnpuschBetaoffsets struct {
	Choice     uint64
	Dynamic    *[]Betaoffsets `lb:4,ub:4`
	Semistatic *Betaoffsets
}
