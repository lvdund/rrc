package ies

// MinTimeGap-r16 ::= SEQUENCE
type MintimegapR16 struct {
	Scs15khzR16  *MintimegapR16Scs15khzR16
	Scs30khzR16  *MintimegapR16Scs30khzR16
	Scs60khzR16  *MintimegapR16Scs60khzR16
	Scs120khzR16 *MintimegapR16Scs120khzR16
}
