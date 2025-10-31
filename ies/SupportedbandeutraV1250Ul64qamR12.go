package ies

import "rrc/utils"

// SupportedBandEUTRA-v1250-ul-64QAM-r12 ::= ENUMERATED
type SupportedbandeutraV1250Ul64qamR12 struct {
	Value utils.ENUMERATED
}

const (
	SupportedbandeutraV1250Ul64qamR12EnumeratedNothing = iota
	SupportedbandeutraV1250Ul64qamR12EnumeratedSupported
)
