package ies

import "rrc/utils"

// MeasGapSharingScheme ::= ENUMERATED
type Measgapsharingscheme struct {
	Value utils.ENUMERATED
}

const (
	MeasgapsharingschemeEnumeratedNothing = iota
	MeasgapsharingschemeEnumeratedScheme00
	MeasgapsharingschemeEnumeratedScheme01
	MeasgapsharingschemeEnumeratedScheme10
	MeasgapsharingschemeEnumeratedScheme11
)
