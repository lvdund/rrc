package ies

import "rrc/utils"

// WLAN-Name-r15 ::= OCTET STRING (SIZE (1..32))
type WlanNameR15 struct {
	Value utils.OCTETSTRING `lb:1,ub:32`
}
