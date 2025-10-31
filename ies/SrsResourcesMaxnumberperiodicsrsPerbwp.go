package ies

import "rrc/utils"

// SRS-Resources-maxNumberPeriodicSRS-PerBWP ::= ENUMERATED
type SrsResourcesMaxnumberperiodicsrsPerbwp struct {
	Value utils.ENUMERATED
}

const (
	SrsResourcesMaxnumberperiodicsrsPerbwpEnumeratedNothing = iota
	SrsResourcesMaxnumberperiodicsrsPerbwpEnumeratedN1
	SrsResourcesMaxnumberperiodicsrsPerbwpEnumeratedN2
	SrsResourcesMaxnumberperiodicsrsPerbwpEnumeratedN4
	SrsResourcesMaxnumberperiodicsrsPerbwpEnumeratedN8
	SrsResourcesMaxnumberperiodicsrsPerbwpEnumeratedN16
)
