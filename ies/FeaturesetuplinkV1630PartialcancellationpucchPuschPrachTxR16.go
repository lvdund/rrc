package ies

import "rrc/utils"

// FeatureSetUplink-v1630-partialCancellationPUCCH-PUSCH-PRACH-TX-r16 ::= ENUMERATED
type FeaturesetuplinkV1630PartialcancellationpucchPuschPrachTxR16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1630PartialcancellationpucchPuschPrachTxR16EnumeratedNothing = iota
	FeaturesetuplinkV1630PartialcancellationpucchPuschPrachTxR16EnumeratedSupported
)
