package ies

import "rrc/utils"

// LWA-Parameters-r13 ::= SEQUENCE
type LwaParametersR13 struct {
	LwaR13            *LwaParametersR13LwaR13
	LwaSplitbearerR13 *LwaParametersR13LwaSplitbearerR13
	WlanMacAddressR13 *utils.OCTETSTRING `lb:6,ub:6`
	LwaBuffersizeR13  *LwaParametersR13LwaBuffersizeR13
}
