package ies

import "rrc/utils"

// MeasObjectCDMA2000 ::= SEQUENCE
// Extensible
type Measobjectcdma2000 struct {
	Cdma2000Type            Cdma2000Type
	Carrierfreq             Carrierfreqcdma2000
	Searchwindowsize        *utils.INTEGER `lb:0,ub:15`
	Offsetfreq              QOffsetrangeinterrat
	Cellstoremovelist       *Cellindexlist
	Cellstoaddmodlist       *Cellstoaddmodlistcdma2000
	Cellforwhichtoreportcgi *Physcellidcdma2000
}
