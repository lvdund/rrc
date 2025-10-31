package ies

import "rrc/utils"

// MeasParameters-v1250-alternativeTimeToTrigger-r12 ::= ENUMERATED
type MeasparametersV1250AlternativetimetotriggerR12 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1250AlternativetimetotriggerR12EnumeratedNothing = iota
	MeasparametersV1250AlternativetimetotriggerR12EnumeratedSupported
)
