package ies

import "rrc/utils"

// Handover-targetRAT-Type ::= utils.ENUMERATED // Extensible
type HandoverTargetratType struct {
	Value utils.ENUMERATED
}

const (
	HandoverTargetratTypeEnumeratedNothing = iota
	HandoverTargetratTypeEnumeratedUtra
	HandoverTargetratTypeEnumeratedGeran
	HandoverTargetratTypeEnumeratedCdma2000_1xrtt
	HandoverTargetratTypeEnumeratedCdma2000_Hrpd
	HandoverTargetratTypeEnumeratedNr
	HandoverTargetratTypeEnumeratedEutra
	HandoverTargetratTypeEnumeratedSpare2
	HandoverTargetratTypeEnumeratedSpare1
)
