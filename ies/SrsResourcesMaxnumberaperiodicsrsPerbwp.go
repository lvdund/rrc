package ies

import "rrc/utils"

// SRS-Resources-maxNumberAperiodicSRS-PerBWP ::= ENUMERATED
type SrsResourcesMaxnumberaperiodicsrsPerbwp struct {
	Value utils.ENUMERATED
}

const (
	SrsResourcesMaxnumberaperiodicsrsPerbwpEnumeratedNothing = iota
	SrsResourcesMaxnumberaperiodicsrsPerbwpEnumeratedN1
	SrsResourcesMaxnumberaperiodicsrsPerbwpEnumeratedN2
	SrsResourcesMaxnumberaperiodicsrsPerbwpEnumeratedN4
	SrsResourcesMaxnumberaperiodicsrsPerbwpEnumeratedN8
	SrsResourcesMaxnumberaperiodicsrsPerbwpEnumeratedN16
)
