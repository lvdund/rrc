package ies

import "rrc/utils"

// InterFreqRSTDMeasurementIndication-r10-IEs ::= SEQUENCE
type InterfreqrstdmeasurementindicationR10Ies struct {
	RstdInterfreqindicationR10 interface{}
	Latenoncriticalextension   *utils.OCTETSTRING
	Noncriticalextension       *interface{}
}
