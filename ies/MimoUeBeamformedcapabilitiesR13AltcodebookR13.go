package ies

import "rrc/utils"

// MIMO-UE-BeamformedCapabilities-r13-altCodebook-r13 ::= ENUMERATED
type MimoUeBeamformedcapabilitiesR13AltcodebookR13 struct {
	Value utils.ENUMERATED
}

const (
	MimoUeBeamformedcapabilitiesR13AltcodebookR13EnumeratedNothing = iota
	MimoUeBeamformedcapabilitiesR13AltcodebookR13EnumeratedSupported
)
