package ies

// BandParameters-v1540-srs-CarrierSwitch-nr ::= SEQUENCE
type BandparametersV1540SrsCarrierswitchNr struct {
	SrsSwitchingtimeslistnr []SrsSwitchingtimenr `lb:1,ub:maxSimultaneousBands`
}
