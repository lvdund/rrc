package ies

import "rrc/utils"

// SDAP-Parameters-sdapHeaderIAB-r16 ::= ENUMERATED
type SdapParametersSdapheaderiabR16 struct {
	Value utils.ENUMERATED
}

const (
	SdapParametersSdapheaderiabR16EnumeratedNothing = iota
	SdapParametersSdapheaderiabR16EnumeratedSupported
)
