package ies

import "rrc/utils"

// MAC-MainConfigSL-r16 ::= SEQUENCE
// Extensible
type MacMainconfigslR16 struct {
	SlBsrConfigR16           *BsrConfig
	UlPrioritizationthresR16 *utils.INTEGER `lb:0,ub:16`
	SlPrioritizationthresR16 *utils.INTEGER `lb:0,ub:8`
}
