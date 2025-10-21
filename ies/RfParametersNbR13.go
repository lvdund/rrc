package ies

import "rrc/utils"

// RF-Parameters-NB-r13 ::= SEQUENCE
type RfParametersNbR13 struct {
	SupportedbandlistR13 SupportedbandlistNbR13
	MultinsPmaxR13       *utils.ENUMERATED
}
