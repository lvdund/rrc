package ies

import "rrc/utils"

// Phy-ParametersCommon-directionalCollisionDC-IAB-r17 ::= ENUMERATED
type PhyParameterscommonDirectionalcollisiondcIabR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonDirectionalcollisiondcIabR17EnumeratedNothing = iota
	PhyParameterscommonDirectionalcollisiondcIabR17EnumeratedSupported
)
