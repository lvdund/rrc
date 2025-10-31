package ies

import "rrc/utils"

// CarrierAggregationVariant-fr1fdd-FR1TDD-FR2TDD-CA-SpCellOnFR2TDD ::= ENUMERATED
type CarrieraggregationvariantFr1fddFr1tddFr2tddCaSpcellonfr2tdd struct {
	Value utils.ENUMERATED
}

const (
	CarrieraggregationvariantFr1fddFr1tddFr2tddCaSpcellonfr2tddEnumeratedNothing = iota
	CarrieraggregationvariantFr1fddFr1tddFr2tddCaSpcellonfr2tddEnumeratedSupported
)
