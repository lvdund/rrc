package ies

// SupportedBandListHRPD ::= SEQUENCE OF BandclassCDMA2000
// SIZE (1..maxCDMA-BandClass)
type Supportedbandlisthrpd struct {
	Value []Bandclasscdma2000 `lb:1,ub:maxCDMABandclass`
}
