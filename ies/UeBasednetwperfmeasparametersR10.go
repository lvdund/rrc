package ies

import "rrc/utils"

// UE-BasedNetwPerfMeasParameters-r10 ::= SEQUENCE
type UeBasednetwperfmeasparametersR10 struct {
	LoggedmeasurementsidleR10 *utils.ENUMERATED
	StandalonegnssLocationR10 *utils.ENUMERATED
}
