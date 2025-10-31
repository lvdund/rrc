package ies

import "rrc/utils"

// UE-BasedNetwPerfMeasParameters-v1430-locationReport-r14 ::= ENUMERATED
type UeBasednetwperfmeasparametersV1430LocationreportR14 struct {
	Value utils.ENUMERATED
}

const (
	UeBasednetwperfmeasparametersV1430LocationreportR14EnumeratedNothing = iota
	UeBasednetwperfmeasparametersV1430LocationreportR14EnumeratedSupported
)
