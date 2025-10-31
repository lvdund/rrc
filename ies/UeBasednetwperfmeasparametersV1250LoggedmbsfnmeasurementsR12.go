package ies

import "rrc/utils"

// UE-BasedNetwPerfMeasParameters-v1250-loggedMBSFNMeasurements-r12 ::= ENUMERATED
type UeBasednetwperfmeasparametersV1250LoggedmbsfnmeasurementsR12 struct {
	Value utils.ENUMERATED
}

const (
	UeBasednetwperfmeasparametersV1250LoggedmbsfnmeasurementsR12EnumeratedNothing = iota
	UeBasednetwperfmeasparametersV1250LoggedmbsfnmeasurementsR12EnumeratedSupported
)
