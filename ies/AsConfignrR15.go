package ies

import "rrc/utils"

// AS-ConfigNR-r15 ::= SEQUENCE
type AsConfignrR15 struct {
	SourcerbConfignrR15      *utils.OCTETSTRING
	SourcerbConfigsnNrR15    *utils.OCTETSTRING
	SourceotherconfigsnNrR15 *utils.OCTETSTRING
}
