package ies

import "rrc/utils"

// PhyLayerParameters-v1540 ::= SEQUENCE
type PhylayerparametersV1540 struct {
	SttiSptCapabilitiesV1540       *interface{}
	CrsImTm1Totm9OnerxPortV1540    *utils.ENUMERATED
	CchImRefrectypeaOnerxPortV1540 *utils.ENUMERATED
}
