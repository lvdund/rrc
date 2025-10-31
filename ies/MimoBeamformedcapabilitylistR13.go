package ies

// MIMO-BeamformedCapabilityList-r13 ::= SEQUENCE OF MIMO-BeamformedCapabilities-r13
// SIZE (1..maxCSI-Proc-r11)
type MimoBeamformedcapabilitylistR13 struct {
	Value []MimoBeamformedcapabilitiesR13 `lb:1,ub:maxCSIProcR11`
}
