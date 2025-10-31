package ies

import "rrc/utils"

// SupportedBandEUTRA-v1310-ue-PowerClass-5-r13 ::= ENUMERATED
type SupportedbandeutraV1310UePowerclass5R13 struct {
	Value utils.ENUMERATED
}

const (
	SupportedbandeutraV1310UePowerclass5R13EnumeratedNothing = iota
	SupportedbandeutraV1310UePowerclass5R13EnumeratedSupported
)
