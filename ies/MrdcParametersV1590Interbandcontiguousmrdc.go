package ies

import "rrc/utils"

// MRDC-Parameters-v1590-interBandContiguousMRDC ::= ENUMERATED
type MrdcParametersV1590Interbandcontiguousmrdc struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersV1590InterbandcontiguousmrdcEnumeratedNothing = iota
	MrdcParametersV1590InterbandcontiguousmrdcEnumeratedSupported
)
