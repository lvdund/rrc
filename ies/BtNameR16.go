package ies

import "rrc/utils"

// BT-Name-r16 ::= OCTET STRING (SIZE (1..248))
type BtNameR16 struct {
	Value utils.OCTETSTRING `lb:1,ub:248`
}
