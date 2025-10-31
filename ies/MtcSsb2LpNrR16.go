package ies

// MTC-SSB2-LP-NR-r16 ::= SEQUENCE
type MtcSsb2LpNrR16 struct {
	PciListR16     *[]PhyscellidnrR15 `lb:1,ub:maxNrofPCIPersmtcR16`
	PeriodicityR16 MtcSsb2LpNrR16PeriodicityR16
}
