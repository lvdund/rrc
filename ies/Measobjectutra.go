package ies

import "rrc/utils"

// MeasObjectUTRA ::= SEQUENCE
// Extensible
type Measobjectutra struct {
	Carrierfreq             ArfcnValueutra
	Offsetfreq              QOffsetrangeinterrat
	Cellstoremovelist       *Cellindexlist
	Cellstoaddmodlist       *interface{}
	Cellforwhichtoreportcgi *interface{}
}
