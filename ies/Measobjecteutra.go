package ies

import "rrc/utils"

// MeasObjectEUTRA ::= SEQUENCE
// Extensible
type Measobjecteutra struct {
	Carrierfreq             ArfcnValueeutra
	Allowedmeasbandwidth    Allowedmeasbandwidth
	Presenceantennaport1    Presenceantennaport1
	Neighcellconfig         Neighcellconfig
	Offsetfreq              QOffsetrange
	Cellstoremovelist       *Cellindexlist
	Cellstoaddmodlist       *Cellstoaddmodlist
	Blackcellstoremovelist  *Cellindexlist
	Blackcellstoaddmodlist  *Blackcellstoaddmodlist
	Cellforwhichtoreportcgi *Physcellid
}
