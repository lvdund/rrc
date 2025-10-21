package ies

import "rrc/utils"

// LWA-Parameters-r13 ::= SEQUENCE
type LwaParametersR13 struct {
	LwaR13            *utils.ENUMERATED
	LwaSplitbearerR13 *utils.ENUMERATED
	WlanMacAddressR13 *utils.OCTETSTRING
	LwaBuffersizeR13  *utils.ENUMERATED
}
