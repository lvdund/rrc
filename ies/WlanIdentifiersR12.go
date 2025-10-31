package ies

import "rrc/utils"

// WLAN-Identifiers-r12 ::= SEQUENCE
// Extensible
type WlanIdentifiersR12 struct {
	SsidR12   *utils.OCTETSTRING `lb:1,ub:32`
	BssidR12  *utils.OCTETSTRING `lb:6,ub:6`
	HessidR12 *utils.OCTETSTRING `lb:6,ub:6`
}
