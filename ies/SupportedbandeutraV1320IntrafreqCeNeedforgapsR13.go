package ies

import "rrc/utils"

// SupportedBandEUTRA-v1320-intraFreq-CE-NeedForGaps-r13 ::= ENUMERATED
type SupportedbandeutraV1320IntrafreqCeNeedforgapsR13 struct {
	Value utils.ENUMERATED
}

const (
	SupportedbandeutraV1320IntrafreqCeNeedforgapsR13EnumeratedNothing = iota
	SupportedbandeutraV1320IntrafreqCeNeedforgapsR13EnumeratedSupported
)
