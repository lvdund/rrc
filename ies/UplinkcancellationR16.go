package ies

import "rrc/utils"

// UplinkCancellation-r16 ::= SEQUENCE
// Extensible
type UplinkcancellationR16 struct {
	CiRntiR16                        RntiValue
	DciPayloadsizeforciR16           utils.INTEGER                      `lb:0,ub:maxCIDciPayloadsizeR16`
	CiConfigurationperservingcellR16 []CiConfigurationperservingcellR16 `lb:1,ub:maxNrofServingCells`
}
