package ies

import "rrc/utils"

// FeatureSetUplink-v1720-interSubslotFreqHopping-PUCCH-r17 ::= ENUMERATED
type FeaturesetuplinkV1720IntersubslotfreqhoppingPucchR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1720IntersubslotfreqhoppingPucchR17EnumeratedNothing = iota
	FeaturesetuplinkV1720IntersubslotfreqhoppingPucchR17EnumeratedSupported
)
