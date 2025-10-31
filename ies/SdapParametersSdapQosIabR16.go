package ies

import "rrc/utils"

// SDAP-Parameters-sdap-QOS-IAB-r16 ::= ENUMERATED
type SdapParametersSdapQosIabR16 struct {
	Value utils.ENUMERATED
}

const (
	SdapParametersSdapQosIabR16EnumeratedNothing = iota
	SdapParametersSdapQosIabR16EnumeratedSupported
)
