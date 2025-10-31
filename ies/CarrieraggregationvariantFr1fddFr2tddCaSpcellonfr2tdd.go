package ies

import "rrc/utils"

// CarrierAggregationVariant-fr1fdd-FR2TDD-CA-SpCellOnFR2TDD ::= ENUMERATED
type CarrieraggregationvariantFr1fddFr2tddCaSpcellonfr2tdd struct {
	Value utils.ENUMERATED
}

const (
	CarrieraggregationvariantFr1fddFr2tddCaSpcellonfr2tddEnumeratedNothing = iota
	CarrieraggregationvariantFr1fddFr2tddCaSpcellonfr2tddEnumeratedSupported
)
