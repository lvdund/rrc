package ies

import "rrc/utils"

// MIMO-UE-BeamformedCapabilities-r13 ::= SEQUENCE
type MimoUeBeamformedcapabilitiesR13 struct {
	AltcodebookR13                *utils.ENUMERATED
	MimoBeamformedcapabilitiesR13 MimoBeamformedcapabilitylistR13
}
