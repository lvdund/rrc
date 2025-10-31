package ies

// SRI-PUSCH-PowerControl ::= SEQUENCE
type SriPuschPowercontrol struct {
	SriPuschPowercontrolid        SriPuschPowercontrolid
	SriPuschPathlossreferencersId PuschPathlossreferencersId
	SriP0PuschAlphasetid          P0PuschAlphasetid
	SriPuschClosedloopindex       SriPuschPowercontrolSriPuschClosedloopindex
}
