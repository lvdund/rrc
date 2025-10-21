package ies

import "rrc/utils"

// CSG-ProximityIndicationParameters-r9 ::= SEQUENCE
type CsgProximityindicationparametersR9 struct {
	IntrafreqproximityindicationR9 *utils.ENUMERATED
	InterfreqproximityindicationR9 *utils.ENUMERATED
	UtranProximityindicationR9     *utils.ENUMERATED
}
