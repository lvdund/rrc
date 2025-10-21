package ies

import "rrc/utils"

// AS-Context-v1610 ::= SEQUENCE
type AsContextV1610 struct {
	SidelinkueinformationnrR16   *utils.OCTETSTRING
	UeassistanceinformationnrR16 *utils.OCTETSTRING
	ConfigrestrictinfodapsR16    *ConfigrestrictinfodapsR16
}
