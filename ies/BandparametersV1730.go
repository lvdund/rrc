package ies

// BandParameters-v1730 ::= SEQUENCE
type BandparametersV1730 struct {
	SrsSwitchingaffectedbandslistnrR17 []SrsSwitchingaffectedbandsnrR17 `lb:1,ub:maxSimultaneousBands`
}
