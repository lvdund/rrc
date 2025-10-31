package ies

import "rrc/utils"

// Phy-ParametersCommon-dynamicSwitchRA-Type0-1-PDSCH ::= ENUMERATED
type PhyParameterscommonDynamicswitchraType01Pdsch struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonDynamicswitchraType01PdschEnumeratedNothing = iota
	PhyParameterscommonDynamicswitchraType01PdschEnumeratedSupported
)
