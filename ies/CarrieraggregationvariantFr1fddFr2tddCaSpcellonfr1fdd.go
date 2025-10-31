package ies

import "rrc/utils"

// CarrierAggregationVariant-fr1fdd-FR2TDD-CA-SpCellOnFR1FDD ::= ENUMERATED
type CarrieraggregationvariantFr1fddFr2tddCaSpcellonfr1fdd struct {
	Value utils.ENUMERATED
}

const (
	CarrieraggregationvariantFr1fddFr2tddCaSpcellonfr1fddEnumeratedNothing = iota
	CarrieraggregationvariantFr1fddFr2tddCaSpcellonfr1fddEnumeratedSupported
)
