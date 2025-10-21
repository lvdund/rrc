package ies

import "rrc/utils"

// SL-Parameters-v1610 ::= SEQUENCE
type SlParametersV1610 struct {
	SlParameternrR16 *utils.OCTETSTRING
	Dummy            *V2xSupportedbandcombinationeutraNrR16
}
