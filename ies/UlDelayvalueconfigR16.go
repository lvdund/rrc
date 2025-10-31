package ies

// UL-DelayValueConfig-r16 ::= SEQUENCE
type UlDelayvalueconfigR16 struct {
	DelayDrblistR16 []DrbIdentity `lb:1,ub:maxDRB`
}
