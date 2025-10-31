package ies

import "rrc/utils"

// WLAN-Identifiers-r16 ::= SEQUENCE
// Extensible
type WlanIdentifiersR16 struct {
	SsidR16   *utils.OCTETSTRING `lb:1,ub:32`
	BssidR16  *utils.OCTETSTRING `lb:6,ub:6`
	HessidR16 *utils.OCTETSTRING `lb:6,ub:6`
}
