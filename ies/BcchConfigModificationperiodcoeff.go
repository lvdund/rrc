package ies

import "rrc/utils"

// BCCH-Config-modificationPeriodCoeff ::= ENUMERATED
type BcchConfigModificationperiodcoeff struct {
	Value utils.ENUMERATED
}

const (
	BcchConfigModificationperiodcoeffEnumeratedNothing = iota
	BcchConfigModificationperiodcoeffEnumeratedN2
	BcchConfigModificationperiodcoeffEnumeratedN4
	BcchConfigModificationperiodcoeffEnumeratedN8
	BcchConfigModificationperiodcoeffEnumeratedN16
)
