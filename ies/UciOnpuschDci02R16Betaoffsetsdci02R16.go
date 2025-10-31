package ies

// UCI-OnPUSCH-DCI-0-2-r16-betaOffsetsDCI-0-2-r16 ::= CHOICE
const (
	UciOnpuschDci02R16Betaoffsetsdci02R16ChoiceNothing = iota
	UciOnpuschDci02R16Betaoffsetsdci02R16ChoiceDynamicdci02R16
	UciOnpuschDci02R16Betaoffsetsdci02R16ChoiceSemistaticdci02R16
)

type UciOnpuschDci02R16Betaoffsetsdci02R16 struct {
	Choice             uint64
	Dynamicdci02R16    *UciOnpuschDci02R16Betaoffsetsdci02R16Dynamicdci02R16
	Semistaticdci02R16 *Betaoffsets
}
