package ies

import "rrc/utils"

// BCCH-Config-NB-r13-modificationPeriodCoeff-r13 ::= ENUMERATED
type BcchConfigNbR13ModificationperiodcoeffR13 struct {
	Value utils.ENUMERATED
}

const (
	BcchConfigNbR13ModificationperiodcoeffR13EnumeratedNothing = iota
	BcchConfigNbR13ModificationperiodcoeffR13EnumeratedN16
	BcchConfigNbR13ModificationperiodcoeffR13EnumeratedN32
	BcchConfigNbR13ModificationperiodcoeffR13EnumeratedN64
	BcchConfigNbR13ModificationperiodcoeffR13EnumeratedN128
)
