package ies

// UplinkTxDirectCurrentCell ::= SEQUENCE
// Extensible
type Uplinktxdirectcurrentcell struct {
	Servcellindex             Servcellindex
	Uplinkdirectcurrentbwp    []Uplinktxdirectcurrentbwp  `lb:1,ub:maxNrofBWPs`
	UplinkdirectcurrentbwpSul *[]Uplinktxdirectcurrentbwp `lb:1,ub:maxNrofBWPs,ext`
}
