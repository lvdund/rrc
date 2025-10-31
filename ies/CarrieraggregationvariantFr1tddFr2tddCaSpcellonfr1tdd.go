package ies

import "rrc/utils"

// CarrierAggregationVariant-fr1tdd-FR2TDD-CA-SpCellOnFR1TDD ::= ENUMERATED
type CarrieraggregationvariantFr1tddFr2tddCaSpcellonfr1tdd struct {
	Value utils.ENUMERATED
}

const (
	CarrieraggregationvariantFr1tddFr2tddCaSpcellonfr1tddEnumeratedNothing = iota
	CarrieraggregationvariantFr1tddFr2tddCaSpcellonfr1tddEnumeratedSupported
)
