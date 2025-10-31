package ies

import "rrc/utils"

// Phy-ParametersCommon-bwp-SwitchingDelay ::= ENUMERATED
type PhyParameterscommonBwpSwitchingdelay struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonBwpSwitchingdelayEnumeratedNothing = iota
	PhyParameterscommonBwpSwitchingdelayEnumeratedType1
	PhyParameterscommonBwpSwitchingdelayEnumeratedType2
)
