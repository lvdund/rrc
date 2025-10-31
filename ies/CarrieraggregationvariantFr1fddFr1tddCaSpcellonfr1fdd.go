package ies

import "rrc/utils"

// CarrierAggregationVariant-fr1fdd-FR1TDD-CA-SpCellOnFR1FDD ::= ENUMERATED
type CarrieraggregationvariantFr1fddFr1tddCaSpcellonfr1fdd struct {
	Value utils.ENUMERATED
}

const (
	CarrieraggregationvariantFr1fddFr1tddCaSpcellonfr1fddEnumeratedNothing = iota
	CarrieraggregationvariantFr1fddFr1tddCaSpcellonfr1fddEnumeratedSupported
)
