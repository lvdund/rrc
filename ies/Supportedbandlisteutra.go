package ies

import "rrc/utils"

// SupportedBandListEUTRA ::= SEQUENCE OF SupportedBandEUTRA
// SIZE (1..maxBands)
type Supportedbandlisteutra struct {
	Value []Supportedbandeutra
}
