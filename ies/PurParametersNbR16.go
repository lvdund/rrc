package ies

import "rrc/utils"

// PUR-Parameters-NB-r16 ::= SEQUENCE
type PurParametersNbR16 struct {
	PurCpEpcR16           *utils.ENUMERATED
	PurCp5gcR16           *utils.ENUMERATED
	PurUpEpcR16           *utils.ENUMERATED
	PurUp5gcR16           *utils.ENUMERATED
	PurNrsrpValidationR16 *utils.ENUMERATED
	PurCpL1ackR16         *utils.ENUMERATED
}
