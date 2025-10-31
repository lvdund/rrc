package ies

import "rrc/utils"

// HRNN-r16 ::= SEQUENCE
type HrnnR16 struct {
	HrnnR16 *utils.OCTETSTRING `lb:1,ub:maxHRNNLenR16`
}
