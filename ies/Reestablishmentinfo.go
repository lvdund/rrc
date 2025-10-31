package ies

// ReestablishmentInfo ::= SEQUENCE
type Reestablishmentinfo struct {
	Sourcephyscellid          Physcellid
	TargetcellshortmacI       ShortmacI
	Additionalreestabinfolist *ReestabNCellInfoList
}
