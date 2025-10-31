package ies

import "rrc/utils"

// MeasConfigMN-gapPurpose ::= ENUMERATED
type MeasconfigmnGappurpose struct {
	Value utils.ENUMERATED
}

const (
	MeasconfigmnGappurposeEnumeratedNothing = iota
	MeasconfigmnGappurposeEnumeratedPerue
	MeasconfigmnGappurposeEnumeratedPerfr1
)
