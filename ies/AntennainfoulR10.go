package ies

import "rrc/utils"

// AntennaInfoUL-r10 ::= SEQUENCE
type AntennainfoulR10 struct {
	TransmissionmodeulR10       *utils.ENUMERATED
	FourantennaportactivatedR10 *utils.ENUMERATED
}
