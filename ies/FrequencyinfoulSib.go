package ies

import "rrc/utils"

// FrequencyInfoUL-SIB ::= SEQUENCE
// Extensible
type FrequencyinfoulSib struct {
	Frequencybandlist       *MultifrequencybandlistnrSib
	Absolutefrequencypointa *ArfcnValuenr
	ScsSpecificcarrierlist  []ScsSpecificcarrier `lb:1,ub:maxSCSs`
	PMax                    *PMax
	Frequencyshift7p5khz    *utils.BOOLEAN
}
