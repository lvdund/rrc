package ies

// MIMO-UE-BeamformedCapabilities-r13 ::= SEQUENCE
type MimoUeBeamformedcapabilitiesR13 struct {
	AltcodebookR13                *MimoUeBeamformedcapabilitiesR13AltcodebookR13
	MimoBeamformedcapabilitiesR13 MimoBeamformedcapabilitylistR13
}
