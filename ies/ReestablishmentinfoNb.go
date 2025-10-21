package ies

import "rrc/utils"

// ReestablishmentInfo-NB ::= SEQUENCE
// Extensible
type ReestablishmentinfoNb struct {
	SourcephyscellidR13          Physcellid
	TargetcellshortmacIR13       ShortmacI
	AdditionalreestabinfolistR13 *Additionalreestabinfolist
}
