package ies

import "rrc/utils"

// FrequencyInfoDL-SIB ::= SEQUENCE
type FrequencyinfodlSib struct {
	Frequencybandlist      MultifrequencybandlistnrSib
	Offsettopointa         utils.INTEGER        `lb:0,ub:2199`
	ScsSpecificcarrierlist []ScsSpecificcarrier `lb:1,ub:maxSCSs`
}
