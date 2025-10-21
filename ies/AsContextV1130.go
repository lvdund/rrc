package ies

import "rrc/utils"

// AS-Context-v1130 ::= SEQUENCE
// Extensible
type AsContextV1130 struct {
	IdcIndicationR11           *utils.OCTETSTRING
	MbmsinterestindicationR11  *utils.OCTETSTRING
	UeassistanceinformationR11 *utils.OCTETSTRING
}
