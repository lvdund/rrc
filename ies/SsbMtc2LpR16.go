package ies

// SSB-MTC2-LP-r16 ::= SEQUENCE
type SsbMtc2LpR16 struct {
	PciList     *[]Physcellid `lb:1,ub:maxNrofPCIsPerSMTC`
	Periodicity SsbMtc2LpR16Periodicity
}
