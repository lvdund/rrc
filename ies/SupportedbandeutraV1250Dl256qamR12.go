package ies

import "rrc/utils"

// SupportedBandEUTRA-v1250-dl-256QAM-r12 ::= ENUMERATED
type SupportedbandeutraV1250Dl256qamR12 struct {
	Value utils.ENUMERATED
}

const (
	SupportedbandeutraV1250Dl256qamR12EnumeratedNothing = iota
	SupportedbandeutraV1250Dl256qamR12EnumeratedSupported
)
