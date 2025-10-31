package ies

import "rrc/utils"

// UE-BasedNetwPerfMeasParameters-r10-loggedMeasurementsIdle-r10 ::= ENUMERATED
type UeBasednetwperfmeasparametersR10LoggedmeasurementsidleR10 struct {
	Value utils.ENUMERATED
}

const (
	UeBasednetwperfmeasparametersR10LoggedmeasurementsidleR10EnumeratedNothing = iota
	UeBasednetwperfmeasparametersR10LoggedmeasurementsidleR10EnumeratedSupported
)
