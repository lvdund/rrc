package ies

import "rrc/utils"

// FrequencyInfoUL ::= SEQUENCE
// Extensible
type Frequencyinfoul struct {
	Frequencybandlist          *Multifrequencybandlistnr
	Absolutefrequencypointa    *ArfcnValuenr
	ScsSpecificcarrierlist     []ScsSpecificcarrier `lb:1,ub:maxSCSs`
	Additionalspectrumemission *Additionalspectrumemission
	PMax                       *PMax
	Frequencyshift7p5khz       *utils.BOOLEAN
}
