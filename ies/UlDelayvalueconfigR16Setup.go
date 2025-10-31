package ies

// UL-DelayValueConfig-r16-setup ::= SEQUENCE
type UlDelayvalueconfigR16Setup struct {
	DelayDrblistR16 []DrbIdentity `lb:1,ub:maxDRB`
}
