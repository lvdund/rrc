package ies

// UCI-OnPUSCH-DCI-0-2-r16-betaOffsetsDCI-0-2-r16-dynamicDCI-0-2-r16 ::= CHOICE
const (
	UciOnpuschDci02R16Betaoffsetsdci02R16Dynamicdci02R16ChoiceNothing = iota
	UciOnpuschDci02R16Betaoffsetsdci02R16Dynamicdci02R16ChoiceOnebitR16
	UciOnpuschDci02R16Betaoffsetsdci02R16Dynamicdci02R16ChoiceTwobitsR16
)

type UciOnpuschDci02R16Betaoffsetsdci02R16Dynamicdci02R16 struct {
	Choice     uint64
	OnebitR16  *[]Betaoffsets `lb:2,ub:2`
	TwobitsR16 *[]Betaoffsets `lb:4,ub:4`
}
