package ies

import "rrc/utils"

// MeasParameters-NB-r16 ::= SEQUENCE
type MeasparametersNbR16 struct {
	DlChannelqualityreportingR16 *utils.ENUMERATED
}
