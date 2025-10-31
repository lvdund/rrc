package ies

// MeasObjectUTRA ::= SEQUENCE
// Extensible
type Measobjectutra struct {
	Carrierfreq             ArfcnValueutra
	Offsetfreq              QOffsetrangeinterrat
	Cellstoremovelist       *Cellindexlist
	Cellstoaddmodlist       *MeasobjectutraCellstoaddmodlist
	Cellforwhichtoreportcgi *MeasobjectutraCellforwhichtoreportcgi
}
