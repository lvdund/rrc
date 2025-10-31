package ies

import "rrc/utils"

// CarrierAggregationVariant-fr1tdd-FR2TDD-CA-SpCellOnFR2TDD ::= ENUMERATED
type CarrieraggregationvariantFr1tddFr2tddCaSpcellonfr2tdd struct {
	Value utils.ENUMERATED
}

const (
	CarrieraggregationvariantFr1tddFr2tddCaSpcellonfr2tddEnumeratedNothing = iota
	CarrieraggregationvariantFr1tddFr2tddCaSpcellonfr2tddEnumeratedSupported
)
