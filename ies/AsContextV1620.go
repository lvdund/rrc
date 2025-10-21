package ies

import "rrc/utils"

// AS-Context-v1620 ::= SEQUENCE
type AsContextV1620 struct {
	UeassistanceinformationnrScgR16 *utils.OCTETSTRING
}
