package ies

import "rrc/utils"

// MBS-FSAI-r17 ::= OCTET STRING (SIZE (3))
type MbsFsaiR17 struct {
	Value utils.OCTETSTRING `lb:3,ub:3`
}
