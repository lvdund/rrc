package ies

import "rrc/utils"

// Phy-ParametersCommon-dynamicSwitchRA-Type0-1-PUSCH ::= ENUMERATED
type PhyParameterscommonDynamicswitchraType01Pusch struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonDynamicswitchraType01PuschEnumeratedNothing = iota
	PhyParameterscommonDynamicswitchraType01PuschEnumeratedSupported
)
