package ies

import "rrc/utils"

// BandListEUTRA ::= SEQUENCE OF BandInfoEUTRA
// SIZE (1..maxBands)
type Bandlisteutra struct {
	Value []Bandinfoeutra
}
