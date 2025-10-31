package ies

import "rrc/utils"

// RadioLinkMonitoringRS-purpose ::= ENUMERATED
type RadiolinkmonitoringrsPurpose struct {
	Value utils.ENUMERATED
}

const (
	RadiolinkmonitoringrsPurposeEnumeratedNothing = iota
	RadiolinkmonitoringrsPurposeEnumeratedBeamfailure
	RadiolinkmonitoringrsPurposeEnumeratedRlf
	RadiolinkmonitoringrsPurposeEnumeratedBoth
)
