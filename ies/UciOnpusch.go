package ies

// UCI-OnPUSCH ::= SEQUENCE
type UciOnpusch struct {
	Betaoffsets *UciOnpuschBetaoffsets
	Scaling     UciOnpuschScaling
}
