package ies

import "rrc/utils"

// AS-Context-v1320 ::= SEQUENCE
type AsContextV1320 struct {
	WlanconnectionstatusreportR13 *utils.OCTETSTRING
}
