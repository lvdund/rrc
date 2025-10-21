package ies

import "rrc/utils"

// PUCCH-ConfigDedicated ::= SEQUENCE
type PucchConfigdedicated struct {
	Acknackrepetition      interface{}
	TddAcknackfeedbackmode *utils.ENUMERATED
}
