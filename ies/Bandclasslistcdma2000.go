package ies

// BandClassListCDMA2000 ::= SEQUENCE OF BandClassInfoCDMA2000
// SIZE (1..maxCDMA-BandClass)
type Bandclasslistcdma2000 struct {
	Value []Bandclassinfocdma2000 `lb:1,ub:maxCDMABandclass`
}
