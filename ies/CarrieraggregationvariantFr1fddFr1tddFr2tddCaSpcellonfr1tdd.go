package ies

import "rrc/utils"

// CarrierAggregationVariant-fr1fdd-FR1TDD-FR2TDD-CA-SpCellOnFR1TDD ::= ENUMERATED
type CarrieraggregationvariantFr1fddFr1tddFr2tddCaSpcellonfr1tdd struct {
	Value utils.ENUMERATED
}

const (
	CarrieraggregationvariantFr1fddFr1tddFr2tddCaSpcellonfr1tddEnumeratedNothing = iota
	CarrieraggregationvariantFr1fddFr1tddFr2tddCaSpcellonfr1tddEnumeratedSupported
)
