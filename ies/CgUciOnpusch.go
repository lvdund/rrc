package ies

// CG-UCI-OnPUSCH ::= CHOICE
const (
	CgUciOnpuschChoiceNothing = iota
	CgUciOnpuschChoiceDynamic
	CgUciOnpuschChoiceSemistatic
)

type CgUciOnpusch struct {
	Choice     uint64
	Dynamic    *[]Betaoffsets `lb:1,ub:4`
	Semistatic *Betaoffsets
}
