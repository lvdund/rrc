package ies

import "rrc/utils"

// SSB-MTC3-r16 ::= SEQUENCE
type SsbMtc3R16 struct {
	PeriodicityandoffsetR16 SsbMtc3R16PeriodicityandoffsetR16
	DurationR16             SsbMtc3R16DurationR16
	PciListR16              *[]Physcellid `lb:1,ub:maxNrofPCIsPerSMTC`
	SsbTomeasureR16         *utils.Setuprelease[SsbTomeasure]
}
