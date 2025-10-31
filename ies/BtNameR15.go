package ies

import "rrc/utils"

// BT-Name-r15 ::= OCTET STRING (SIZE (1..248))
type BtNameR15 struct {
	Value utils.OCTETSTRING `lb:1,ub:248`
}
