package ies

import "rrc/utils"

// MBMSSessionInfo-r13 ::= SEQUENCE
type MbmssessioninfoR13 struct {
	TmgiR13      TmgiR9
	SessionidR13 *utils.OCTETSTRING `lb:1,ub:1`
}
