package ies

import "rrc/utils"

// IRAT-ParametersNR-r15 ::= SEQUENCE
type IratParametersnrR15 struct {
	EnDcR15                  *utils.ENUMERATED
	Eventb2R15               *utils.ENUMERATED
	SupportedbandlistenDcR15 *SupportedbandlistnrR15
}
