package ies

import "rrc/utils"

// SDAP-Parameters ::= SEQUENCE
// Extensible
type SdapParameters struct {
	AsReflectiveqos  *utils.BOOLEAN
	SdapQosIabR16    *SdapParametersSdapQosIabR16    `ext`
	SdapheaderiabR16 *SdapParametersSdapheaderiabR16 `ext`
}
