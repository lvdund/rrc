package ies

import "rrc/utils"

// SL-PSSCH-Config-r16 ::= SEQUENCE
// Extensible
type SlPsschConfigR16 struct {
	SlPsschDmrsTimepatternlistR16 *[]utils.INTEGER    `lb:1,ub:3`
	SlBetaoffsets2ndsciR16        *[]SlBetaoffsetsR16 `lb:4,ub:4`
	SlScalingR16                  *SlPsschConfigR16SlScalingR16
}
