package ies

import "rrc/utils"

// BCCH-Config-v1310-modificationPeriodCoeff-v1310 ::= ENUMERATED
type BcchConfigV1310ModificationperiodcoeffV1310 struct {
	Value utils.ENUMERATED
}

const (
	BcchConfigV1310ModificationperiodcoeffV1310EnumeratedNothing = iota
	BcchConfigV1310ModificationperiodcoeffV1310EnumeratedN64
)
