package ies

import "rrc/utils"

// ReestablishmentInfo ::= SEQUENCE
// Extensible
type Reestablishmentinfo struct {
	Sourcephyscellid          Physcellid
	TargetcellshortmacI       ShortmacI
	Additionalreestabinfolist *Additionalreestabinfolist
}
