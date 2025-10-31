package ies

import "rrc/utils"

// Phy-ParametersCommon-pathlossEstimation2PortCSI-RS-r16 ::= ENUMERATED
type PhyParameterscommonPathlossestimation2portcsiRsR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonPathlossestimation2portcsiRsR16EnumeratedNothing = iota
	PhyParameterscommonPathlossestimation2portcsiRsR16EnumeratedSupported
)
