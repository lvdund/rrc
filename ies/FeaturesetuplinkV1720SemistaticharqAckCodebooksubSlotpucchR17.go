package ies

import "rrc/utils"

// FeatureSetUplink-v1720-semiStaticHARQ-ACK-CodebookSub-SlotPUCCH-r17 ::= ENUMERATED
type FeaturesetuplinkV1720SemistaticharqAckCodebooksubSlotpucchR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1720SemistaticharqAckCodebooksubSlotpucchR17EnumeratedNothing = iota
	FeaturesetuplinkV1720SemistaticharqAckCodebooksubSlotpucchR17EnumeratedSupported
)
