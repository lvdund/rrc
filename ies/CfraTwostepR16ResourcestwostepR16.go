package ies

import "rrc/utils"

// CFRA-TwoStep-r16-resourcesTwoStep-r16 ::= SEQUENCE
type CfraTwostepR16ResourcestwostepR16 struct {
	SsbResourcelist        []CfraSsbResource `lb:1,ub:maxRASsbResources`
	RaSsbOccasionmaskindex utils.INTEGER     `lb:0,ub:15`
}
