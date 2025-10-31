package ies

// ReestablishmentInfo ::= SEQUENCE
// Extensible
type Reestablishmentinfo struct {
	Sourcephyscellid          Physcellid
	TargetcellshortmacI       ShortmacI
	Additionalreestabinfolist *Additionalreestabinfolist
}
