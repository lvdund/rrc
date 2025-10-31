package ies

import "rrc/utils"

// SL-QoS-Profile-r16 ::= SEQUENCE
// Extensible
type SlQosProfileR16 struct {
	SlPqiR16   *SlPqiR16
	SlGfbrR16  *utils.INTEGER `lb:0,ub:4000000000`
	SlMfbrR16  *utils.INTEGER `lb:0,ub:4000000000`
	SlRangeR16 *utils.INTEGER `lb:0,ub:1000`
}
