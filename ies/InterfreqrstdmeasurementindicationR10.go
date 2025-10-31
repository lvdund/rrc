package ies

import "rrc/utils"

// InterFreqRSTDMeasurementIndication-r10-IEs ::= SEQUENCE
type InterfreqrstdmeasurementindicationR10 struct {
	RstdInterfreqindicationR10 InterfreqrstdmeasurementindicationR10IesRstdInterfreqindicationR10
	Latenoncriticalextension   *utils.OCTETSTRING
	Noncriticalextension       *InterfreqrstdmeasurementindicationR10IesNoncriticalextension
}
