package ies

// BandParameters-v1540-srs-CarrierSwitch-eutra ::= SEQUENCE
type BandparametersV1540SrsCarrierswitchEutra struct {
	SrsSwitchingtimeslisteutra []SrsSwitchingtimeeutra `lb:1,ub:maxSimultaneousBands`
}
