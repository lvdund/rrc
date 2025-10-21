package ies

import "rrc/utils"

// AntennaInfoDedicated ::= SEQUENCE
type Antennainfodedicated struct {
	Transmissionmode           utils.ENUMERATED
	Codebooksubsetrestriction  *interface{}
	UeTransmitantennaselection utils.ENUMERATED
}
