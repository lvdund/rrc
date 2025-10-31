package ies

// SSB-MTC2 ::= SEQUENCE
type SsbMtc2 struct {
	PciList     *[]Physcellid `lb:1,ub:maxNrofPCIsPerSMTC`
	Periodicity SsbMtc2Periodicity
}
