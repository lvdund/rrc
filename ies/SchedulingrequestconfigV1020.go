package ies

import "rrc/utils"

// SchedulingRequestConfig-v1020 ::= SEQUENCE
type SchedulingrequestconfigV1020 struct {
	SrPucchResourceindexp1R10 *utils.INTEGER `lb:0,ub:2047`
}
