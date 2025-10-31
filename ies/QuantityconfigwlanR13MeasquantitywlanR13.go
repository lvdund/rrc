package ies

import "rrc/utils"

// QuantityConfigWLAN-r13-measQuantityWLAN-r13 ::= ENUMERATED
type QuantityconfigwlanR13MeasquantitywlanR13 struct {
	Value utils.ENUMERATED
}

const (
	QuantityconfigwlanR13MeasquantitywlanR13EnumeratedNothing = iota
	QuantityconfigwlanR13MeasquantitywlanR13EnumeratedRssiwlan
)
