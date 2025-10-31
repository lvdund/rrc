package ies

import "rrc/utils"

// WLAN-Name-r16 ::= OCTET STRING (SIZE (1..32))
type WlanNameR16 struct {
	Value utils.OCTETSTRING `lb:1,ub:32`
}
