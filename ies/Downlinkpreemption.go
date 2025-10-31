package ies

import "rrc/utils"

// DownlinkPreemption ::= SEQUENCE
// Extensible
type Downlinkpreemption struct {
	IntRnti                        RntiValue
	Timefrequencyset               DownlinkpreemptionTimefrequencyset
	DciPayloadsize                 utils.INTEGER                    `lb:0,ub:maxINTDciPayloadsize`
	IntConfigurationperservingcell []IntConfigurationperservingcell `lb:1,ub:maxNrofServingCells`
}
